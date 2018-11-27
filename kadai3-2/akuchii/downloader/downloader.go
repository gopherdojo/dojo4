package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

// Downloader downloads data in parallel
type Downloader struct {
	url      string
	fileSize uint
	splitNum uint
	fileName string
}

// Range defines range of data of start and end
type Range struct {
	start uint
	end   uint
	idx   uint
}

// NewDownloader generate new Downloader instance
func NewDownloader(url string, splitNum uint) *Downloader {
	return &Downloader{
		url:      url,
		splitNum: splitNum,
	}
}

// Prepare prepares for download in parallel
func (d *Downloader) Prepare() error {
	res, err := http.Head(d.url)
	if err != nil {
		return err
	}
	d.fileSize = uint(res.ContentLength)

	elems := strings.Split(d.url, "/")
	d.fileName = elems[len(elems)-1]

	return nil
}

// Download downloads in parallel
func (d Downloader) Download() error {
	var eg errgroup.Group

	for i := uint(0); i < d.splitNum; i++ {
		r := d.makeRange(i)
		eg.Go(func() error {
			return d.request(r)
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

// GetFileName gets name of download file
func (d Downloader) GetFileName() string {
	return d.fileName
}

// MergeFiles merges files downloads in parallel
func (d Downloader) MergeFiles() error {
	f, err := os.Create(d.fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := uint(0); i < d.splitNum; i++ {
		tmpFileName := fmt.Sprintf("%s.%d", d.fileName, i)
		tmpFile, err := os.Open(tmpFileName)
		if err != nil {
			return err
		}
		io.Copy(f, tmpFile)
		tmpFile.Close()
		if err := os.Remove(tmpFileName); err != nil {
			return err
		}
	}
	return nil
}

func (d Downloader) request(r Range) error {
	req, err := http.NewRequest(http.MethodGet, d.url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.start, r.end))
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmpFileName := fmt.Sprintf("%s.%d", d.fileName, r.idx)
	file, err := os.Create(tmpFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}

	return nil
}

func (d Downloader) makeRange(i uint) Range {
	splitSize := d.fileSize / d.splitNum
	start := i * splitSize
	end := (i+1)*splitSize - 1

	if i == d.splitNum {
		end = d.fileSize
	}

	return Range{
		start: start,
		end:   end,
		idx:   i,
	}
}
