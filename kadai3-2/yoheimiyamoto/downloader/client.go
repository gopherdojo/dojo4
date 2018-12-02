package downloader

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/errgroup"
)

// Client ...
type Client struct {
	*http.Client
	url string
}

// NewClient ...
func NewClient(url string) *Client {
	return &Client{new(http.Client), url}
}

// Download ...
func (c *Client) Download(ctx context.Context) error {
	l, err := c.contentLength()
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "contentLength: %d\n", l)

	eg, ctx := errgroup.WithContext(ctx)

	ps := newRangeProperties(l)
	fmt.Fprintf(os.Stdout, "ps: %d\n", len(ps))

	for _, p := range ps {
		p := p
		eg.Go(func() error {
			return c.rangeDownload(ctx, p)
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	defer removeFiles(ps)

	src := make([]string, len(ps))
	for i, p := range ps {
		src[i] = p.path
	}
	err = mergeFiles(src, "output.jpg")
	if err != nil {
		return nil
	}
	return nil
}

// contentLengthを取得
func (c *Client) contentLength() (int, error) {
	res, err := c.Head(c.url)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res.Header.Get("Content-Length"))
}

// ファイルの分割ダウンロード
// start -> 開始のbyte
// end -> 終了のbyte
// dst -> ダウンロード先のファイル名
func (c *Client) rangeDownload(ctx context.Context, r *rangeProperty) error {
	errCh := make(chan error, 1)

	go func() {
		req, _ := http.NewRequest("GET", c.url, nil)
		req.Header.Set("Authorization", "Bearer access-token")
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.start, r.end))

		res, err := c.Do(req)
		if err != nil {
			errCh <- err
			return
		}
		defer res.Body.Close()

		f, err := os.Create(r.path)
		if err != nil {
			errCh <- err
			return
		}
		defer f.Close()

		_, err = io.Copy(f, res.Body)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		return ctx.Err()
	}

	return nil
}

// ダウンロードした分割ファイルのマージ
// src -> 元の分割ファイル名
// dst -> マージ先のファイル名
func mergeFiles(src []string, dst string) error {
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, s := range src {
		_f, err := os.Open(s)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, _f)
		if err != nil {
			return err
		}
		_f.Close()
	}
	return nil
}

// 一時保存ファイルの全削除
func removeFiles(ps []*rangeProperty) error {
	for _, p := range ps {
		err := os.Remove(p.path)
		if err != nil {
			return err
		}
	}
	return nil
}
