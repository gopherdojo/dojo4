package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

type Client struct {
	URL           url.URL
	HTTPClient    *http.Client
	Logger        *log.Logger
	ContentLength int
}

func (c *Client) RequestHeader() bool {
	resp, err := c.HTTPClient.Head(c.URL.String())
	if err != nil {
		c.Logger.Println(err)
		return false
	}
	if resp.Header.Get("Accept-Ranges") == "" {
		return false
	}
	// set content length
	contentLengthStr := resp.Header.Get("Content-Length")
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		c.Logger.Println(err)
	}
	c.ContentLength = contentLength
	return true
}

// Make Goroutine
func (c *Client) GetContent() error {
	var n int
	downloadBytesPerGoroutine := c.ContentLength / 4
	// make Request
	eg := errgroup.Group{}
	var tmpFileNames []string
	for n < c.ContentLength {
		startRange := n
		endRange := n + downloadBytesPerGoroutine
		if startRange != 0 {
			startRange = startRange + 1
		}
		if endRange > c.ContentLength {
			endRange = c.ContentLength
		}
		n = endRange
		eg.Go(func() error {
			tmpFileName, err := c.download(startRange, endRange)
			if err != nil {
				return err
			}
			tmpFileNames = append(tmpFileNames, tmpFileName)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
	sort.Strings(tmpFileNames)
	return c.merge(tmpFileNames)
}
func (c *Client) merge(tmpFileNames []string) error {
	file, err := os.Create(c.getFileName())
	if err != nil {
		return err
	}
	for _, tmpFileName := range tmpFileNames {
		tmpFile, err := os.Open(tmpFileName)
		if err != nil {
			return err
		}
		io.Copy(file, tmpFile)
		tmpFile.Close()
		err = os.Remove(tmpFileName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) getFileName() string {
	downloadPathSplit := strings.Split(c.URL.Path, "/")
	downloadFileName := downloadPathSplit[len(downloadPathSplit)-1]
	return downloadFileName
}

func (c *Client) download(startRange int, endRange int) (string, error) {
	req, _ := http.NewRequest("GET", c.URL.String(), nil)
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", startRange, endRange))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	downloadFileName := c.getFileName()
	downloadTmpFileName := fmt.Sprintf(downloadFileName+"_"+"%d", startRange)
	file, err := os.Create(downloadTmpFileName)
	defer file.Close()
	if err != nil {
		return "", err
	}
	io.Copy(file, resp.Body)
	return downloadTmpFileName, nil
}

func NewGpcClient(url url.URL, logger *log.Logger) (GpcClient, error) {
	client := &http.Client{Timeout: time.Duration(5) * time.Second}
	return &Client{URL: url, HTTPClient: client, Logger: logger}, nil
}

type GpcClient interface {
	RequestHeader() bool
	GetContent() error
}
