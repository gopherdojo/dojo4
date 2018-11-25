package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/sync/errgroup"
)

const processNum = 4
const tmpDir = "./tmp"

type RangeClients []RangeClient

type RangeClient struct {
	Index      int
	FilePath   string
	StartBytes int64
	EndBytes   int64
}

func NewRangeClients(fileBytes int64, fileName string) RangeClients {
	rc := make(RangeClients, 0, processNum)
	b := int64(fileBytes / processNum)
	sb := int64(0)
	for i := 0; i < processNum; i++ {
		r := RangeClient{
			Index:      i,
			FilePath:   fmt.Sprintf("%s/%s_%d", tmpDir, fileName, i),
			StartBytes: sb,
			EndBytes:   sb + b,
		}
		if i == processNum-1 {
			r.EndBytes = fileBytes
		}
		rc = append(rc, r)
		sb += b + 1
	}
	return rc
}

func main() {
	os.Exit(Run())
}

func Run() int {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./main TARGET_URL")
		return 1
	}
	targetUrl := os.Args[1]

	res, err := http.Head(targetUrl)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	if res.Header.Get("Accept-Ranges") != "bytes" {
		fmt.Fprintf(os.Stderr, "not support range header, url:%s", targetUrl)
		return 1
	}
	rcs := NewRangeClients(res.ContentLength, filepath.Base(targetUrl))

	// create tmp directory
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		fmt.Fprintln(os.Stderr, "could not create tmp directory")
	}
	// remove tmp directory with defer
	defer os.RemoveAll(tmpDir)

	eg, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for _, rc := range rcs {
		rc := rc
		eg.Go(func() error {
			res, err := rc.download(ctx, targetUrl)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			return rc.outputFile(ctx, res)
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fh, err := os.Create(filepath.Base(targetUrl))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer fh.Close()

	// bind files
	for _, rc := range rcs {
		fp, err := os.Open(rc.FilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		defer fp.Close()
		io.Copy(fh, fp)
	}

	return 0
}

func (rc *RangeClient) download(ctx context.Context, t string) (*http.Response, error) {
	type requestResult struct {
		res *http.Response
		err error
	}

	req, err := http.NewRequest("GET", t, nil)
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	if err != nil {
		return nil, err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", rc.StartBytes, rc.EndBytes))

	resCh := make(chan requestResult)
	go func() {
		time.Sleep(5 * time.Second)
		res, err := client.Do(req)
		resCh <- requestResult{
			res: res,
			err: err,
		}
	}()

	select {
	case <-ctx.Done():
		tr.CancelRequest(req)
		return nil, ctx.Err()
	case result := <-resCh:
		return result.res, result.err
	}
}

func (rc *RangeClient) outputFile(ctx context.Context, res *http.Response) error {
	fp, err := os.OpenFile(rc.FilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()

	errCh := make(chan error)
	go func() {
		_, err := io.Copy(fp, res.Body)
		errCh <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errCh:
		if err != nil {
			return err
		}
	}
	return nil
}
