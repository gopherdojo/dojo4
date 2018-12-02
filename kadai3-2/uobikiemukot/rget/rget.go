package rget

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync"
)

// payload store data acquired by rget
type payload struct {
	buf    *bytes.Buffer // buffer
	offset int64         // data offset
	err    error
}

// fanIn merge sereval channels into one channel
func fanIn(chs ...<-chan payload) <-chan payload {
	var wg sync.WaitGroup
	merged := make(chan payload)

	wg.Add(len(chs))
	for _, ch := range chs {
		go func(ch <-chan payload) {
			defer wg.Done()
			p := <-ch
			merged <- p
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

// Head HTTP HEAD and return Content-Length, Accept-Ranges
func Head(url string) (int64, string, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, "", fmt.Errorf("unexpected StatusCode: %s", resp.Status)
	}

	size, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 0, 64)
	if err != nil {
		return 0, "", err
	}

	rng := resp.Header.Get("Accept-Ranges")

	return size, rng, nil
}

// do HTTP GET with Range header
func do(ctx context.Context, url string, from, to int64) <-chan payload {
	ch := make(chan payload)

	go func() {
		defer close(ch)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			ch <- payload{buf: nil, offset: 0, err: err}
			return
		}
		req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", from, to))
		req = req.WithContext(ctx)

		var client http.Client
		resp, err := client.Do(req)
		if err != nil {
			ch <- payload{buf: nil, offset: 0, err: err}
			return
		}
		defer resp.Body.Close()

		var buf bytes.Buffer
		_, err = io.Copy(&buf, resp.Body)
		if err != nil {
			ch <- payload{buf: nil, offset: 0, err: err}
			return
		}


		ch <- payload{buf: &buf, offset: from, err: nil}
	}()

	return ch
}

// Serial HTTP Get in serial
func Serial(ctx context.Context, fp *os.File, url string, size int64) error {
	p := <-do(ctx, url, 0, size)
	if p.err != nil {
		return p.err
	}

	_, err := io.Copy(fp, p.buf)
	if err != nil {
		return err
	}

	_, err = fp.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	return nil
}

// Parallel HTTP Get in parallel
func Parallel(ctx context.Context, fp *os.File, url string, size int64) error {
	ncpu := runtime.NumCPU()
	chunkSize := size / int64(ncpu)

	// fan-out
	finders := make([]<-chan payload, ncpu)
	for i := 0; i < ncpu; i++ {
		var from, to int64

		from = chunkSize * int64(i)

		if i == (ncpu - 1) {
			to = size - 1
		} else {
			to = (from + chunkSize) - 1
		}

		finders[i] = do(ctx, url, from, to)
	}

	for p := range fanIn(finders...) {
		if p.err != nil {
			return p.err
		}

		fp.WriteAt(p.buf.Bytes(), p.offset)
	}

	_, err := fp.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	return nil
}
