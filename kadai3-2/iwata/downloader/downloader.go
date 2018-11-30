package downloader

import (
	"context"
	"io"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/sync/errgroup"
)

type Downloader struct {
	Procs uint
}

type ChunkRange struct {
	start uint
	end   uint
}

type ChunkFiles []string

func NewDownloader(procs uint) *Downloader {
	return &Downloader{Procs: procs}
}

func (d *Downloader) Do(url string, timeout time.Duration) (ChunkFiles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	size, err := GetContentLength(ctx, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get ContentType")
	}

	cr := makeChunkRages(size, d.Procs)

	eg, ctxEg := errgroup.WithContext(ctx)
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

func makeChunkRages(size, p uint) []ChunkRange {
	chunkSize := size / p
	buf := make([]ChunkRange, p)
	for i := 0; i < size; i + chunkSize {
	}

	return buf
}

func (c ChunkFiles) Save(dist string) error {
	d, err := os.Create(dist)
	if err != nil {
		return errors.Wrapf(err, "failed to create %s", dist)
	}

	sort.Strings(c)
	for _, f := range c {
		chunk, err := os.Open(f)
		if err != nil {
			return errors.Wrapf(err, "failed to open %s", chunk)
		}
		defer chunk.Close()
		io.Copy(d, chunk)
		// cleanup temp files
		if err := os.Remove(f); err != nil {
			return errors.Wrapf(err, "failed to remove %s", chunk)
		}
	}

	return nil
}
