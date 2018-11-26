package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

// Client ...
type Client struct {
	*http.Client
	url string
}

type rangeProperty struct {
	dst   string // 一時保存先のファイル名
	start int    // 開始のバイト数
	end   int    // 終了のバイト数
}

// NewClient ...
func NewClient(url string) *Client {
	return &Client{new(http.Client), url}
}

//
func newRangeProperties(contentLength int) []*rangeProperty {
	// const range = 10000
	var out []*rangeProperty
	f := func(i int) (int, int) {
		return 0, 0
	}
	start := 1
	end := start + 10000
	for {
		r := &rangeProperty{
			dst:   fmt.Sprintf("file%d.jpg", i),
			start: start,
			end:   end,
		}
		out = append(out, r)
		start += 10000
		end := start + 10000
		if end > contentLength {
			end = contentLength
		}
	}
	return out
}

func Download() {

}

// 一時保存ファイルの全削除
func removeFiles(ps []*rangeProperty) error {
	for _, p := range ps {
		err := os.Remove(p.dst)
		if err != nil {
			return err
		}
	}
	return nil
}

// ファイルの分割ダウンロード
// start -> 開始のbyte
// end -> 終了のbyte
// dst -> ダウンロード先のファイル名
func (c *Client) rangeDownload(r *rangeProperty) error {
	req, _ := http.NewRequest("GET", c.url, nil)
	req.Header.Set("Authorization", "Bearer access-token")
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.start, r.end))

	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	f, err := os.Create(r.dst)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, res.Body)
	if err != nil {
		return err
	}
	return nil
}

// ダウンロードした分割ファイルのマージ
// src -> 元の分割ファイル名
// dst -> マージ先のファイル名
func (c *Client) mergeFiles(src []string, dst string) error {
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

// contentLengthを取得
func (c *Client) contentLength() (int, error) {
	res, err := c.Head(c.url)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(res.Header.Get("Content-Length"))
}
