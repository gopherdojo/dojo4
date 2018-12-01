package downloader

import (
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

type Downloader struct {
	procs uint
	dir   string
}

type ChunkRequest struct {
	Start uint
	End   uint
}

type ChunkFiles []string

func New(procs uint, tempDir string) *Downloader {
	return &Downloader{procs: procs, dir: tempDir}
}

func (d *Downloader) Do(url string, timeout time.Duration) (ChunkFiles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	length, err := GetContentLength(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ContentType")
	}

	cr := makeChunkRequests(length, d.procs)
	cf := ChunkFiles(make([]string, d.procs))

	eg, ctxEg := errgroup.WithContext(ctx)
	for i, req := range cr {
		i := i
		req := req
		eg.Go(func() error {
			res, err := req.Do(ctxEg, url)
			if err != nil {
				return errors.Wrapf(err, "failed to do chunk request: %s", url)
			}
			defer res.Body.Close()

			f := filepath.Join(d.dir, fmt.Sprintf("%d", req.Start))
			w, err := os.Create(f)
			if err != nil {
				return errors.Wrapf(err, "failed to create a chunk file: %s", f)
			}
			defer w.Close()

			_, err = io.Copy(w, res.Body)
			if err != nil {
				return errors.Wrapf(err, "failed to dump a chunk body to %s", f)
			}

			cf[i] = f

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		err = errors.Wrap(err, "failed to download in parallel")
		if errR := os.RemoveAll(d.dir); errR != nil {
			err = errors.Wrapf(errR, "failed to clear temp files in %s", d.dir)
		}
		return nil, err
	}

	return cf, nil
}

func GetContentLength(ctx context.Context, url string) (uint, error) {
	res, err := ctxhttp.Head(ctx, http.DefaultClient, url)
	if err != nil {
		return 0, errors.Wrapf(err, "failed to HEAD request: %s", url)
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return 0, errors.Errorf("not supported range access: %s", url)
	}

	if res.ContentLength <= 0 {
		return 0, errors.Errorf("invalid content length about %s", url)
	}

	return uint(res.ContentLength), nil
}

func makeChunkRequests(length, p uint) []*ChunkRequest {
	chunkSize := uint(math.Ceil(float64(length) / float64(p)))
	buf := make([]*ChunkRequest, p)
	// Content-Rangesは0番目からはじまり, Content-Length-1番目まで
	for i := uint(0); i < p; i++ {
		s := i * chunkSize
		e := s + chunkSize - 1
		buf[i] = &ChunkRequest{Start: s, End: e}
	}

	return buf
}

func (r *ChunkRequest) Do(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to GET: %s", url)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.Start, r.End))

	return http.DefaultClient.Do(req)
}

func (c ChunkFiles) Save(dist string) error {
	d, err := os.Create(dist)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", dist)
	}

	for _, f := range c {
		chunk, err := os.Open(f)
		if err != nil {
			return errors.Wrapf(err, "failed to open %s", f)
		}
		defer chunk.Close()

		_, err = io.Copy(d, chunk)
		if err != nil {
			return errors.Wrapf(err, "failed to copy chunk data to %s", f)
		}
	}

	if err := os.RemoveAll(filepath.Dir(c[0])); err != nil {
		return errors.Wrap(err, "failed to clean up temp files")
	}

	return nil
}
