package downloader

import (
	"fmt"
	"io"
	"math"
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
	path  string // 一時保存先のファイル名
	start int    // 開始のバイト数
	end   int    // 終了のバイト数
}

// NewClient ...
func NewClient(url string) *Client {
	return &Client{new(http.Client), url}
}

// Download ...
func (c *Client) Download() error {
	l, err := c.contentLength()
	if err != nil {
		return err
	}
	ps := newRangeProperties(l)
	for _, p := range ps {
		err = c.rangeDownload(p)
		if err != nil {
			return err
		}
	}
	for _, p := range ps {
		fmt.Printf("%v", p)
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

// rangeDownloadの引数として必要なrangePropertyを生成
func newRangeProperties(contentLength int) []*rangeProperty {
	num := 1000000

	maxIndex := int(math.Ceil(float64(contentLength) / float64(num)))

	f := func(i int) *rangeProperty {
		start := 0
		if i != 0 {
			start = i*num + 1
		}
		end := (i + 1) * num
		if end > contentLength {
			end = contentLength
		}
		return &rangeProperty{
			path:  fmt.Sprintf("file%d.jpg", i),
			start: start,
			end:   end,
		}
	}

	var out []*rangeProperty

	for i := 0; i < maxIndex; i++ {
		start := i + num
		end := start + num
		if end > contentLength {
			end = contentLength
		}
		out = append(out, f(i))
	}

	return out
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
	f, err := os.Create(r.path)
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
