package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

// RangeRequest represents a request with a range access
type RangeRequest struct {
	URL    string
	FName  string
	Ranges []*Range
}

// NonRangeRequest represents a request without a range access
type NonRangeRequest struct {
	URL   string
	FName string
}

// Request represents a request
type Request interface {
	Do() error
}

// Range tells tha range of file to download
type Range struct {
	start int64
	end   int64
}

// NewRequest initializes Request object
func NewRequest(rawURL string, parallel int) (Request, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	ss := strings.Split(u.Path, "/")
	fname := ss[len(ss)-1]

	res, err := http.Head(rawURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return &NonRangeRequest{URL: rawURL, FName: fname}, nil
	}

	total := res.ContentLength
	unit := total / int64(parallel)
	ranges := make([]*Range, parallel)

	for i := 0; i < parallel; i++ {
		var start int64
		if i == 0 {
			start = 0
		} else {
			start = int64(i)*unit + 1
		}

		var end int64
		if i == parallel-1 {
			end = total
		} else {
			end = int64(i+1) * unit
		}

		ranges[i] = &Range{start: start, end: end}
	}

	return &RangeRequest{URL: rawURL, FName: fname, Ranges: ranges}, nil
}

// Do sends a real HTTP requests in parallel
func (r *NonRangeRequest) Do() error {
	req, err := http.NewRequest(http.MethodGet, r.URL, nil)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return saveResponseBody(r.FName, res)
}

// Do sends a real HTTP requests in parallel
func (r *RangeRequest) Do() error {
	eg, _ := errgroup.WithContext(context.TODO())

	for idx := range r.Ranges {
		// DO NOT refer to idx directly since function below
		// is a closure and idx changes for each iterations
		i := idx
		eg.Go(func() error {
			return r.do(i)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return r.mergeFiles()
}

func (r *RangeRequest) do(idx int) error {
	req, err := http.NewRequest(http.MethodGet, r.URL, nil)
	if err != nil {
		return err
	}

	ran := r.Ranges[idx]
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", ran.start, ran.end))

	client := http.DefaultClient

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	tmpFName := fmt.Sprintf("%s.%d", r.FName, idx)
	return saveResponseBody(tmpFName, res)
}

func (r *RangeRequest) mergeFiles() error {
	f, err := os.Create(r.FName)
	if err != nil {
		return err
	}
	defer f.Close()

	for idx := range r.Ranges {
		tmpFName := fmt.Sprintf("%s.%d", r.FName, idx)
		tmpFile, err := os.Open(tmpFName)
		if err != nil {
			return err
		}

		io.Copy(f, tmpFile)
		tmpFile.Close()
		if err := os.Remove(tmpFName); err != nil {
			return err
		}
	}

	return nil
}

func saveResponseBody(fname string, response *http.Response) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return err
	}

	return nil
}
