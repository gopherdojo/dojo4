package downloader

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/sync/errgroup"
)

// Downloader divide support donwloader
type Downloader int

type downloadInfo struct {
	startByte int64
	endByte   int64
}

type deleteFileError struct {
	filename string
	err      error
}

// Download download action
// description:Download specified URL.
// parameter
// url :download target url.
// saveFilePath:save file path.
// div:Number of divided downloads.
func (d *Downloader) Download(url, saveFilePath string, div int64) error {
	res, err := http.Head(url)
	if err != nil {
		fmt.Println("http Head error")
		return err
	}
	if res.StatusCode != 200 {
		fmt.Println("bad status code")
		fmt.Println("status code:", res.StatusCode)
		return err
	}

	if !canRangeDownload(res.Header) {
		fmt.Println("[warn]range download not support.")
		div = 1
	}

	di := splitDownloadLength(res.ContentLength, div)
	filenames := make([]string, div)

	var eg errgroup.Group

	i := 1
	for _, d := range di {
		j := i
		eg.Go(func() error {
			err := rangeDownload(j, url, d.startByte, d.endByte)
			return err
		})
		filenames[i-1] = strconv.Itoa(j) + ".temp.download"
		i++
	}
	if err := eg.Wait(); err != nil {
		fmt.Println("download error")
		return err
	}

	if err := joinFiles(filenames, saveFilePath); err != nil {
		fmt.Println("join file error")
		return err
	}

	deleteFileError := deleteFiles(filenames)

	if len(deleteFileError) != 0 {
		for _, e := range deleteFileError {
			fmt.Printf("[delete file error.%s %s", e.filename, e.err)
		}
		return errors.New("download fail")
	}

	fmt.Println("download finish.")
	return nil
}

func canRangeDownload(h http.Header) bool {
	accept := h.Get("Accept-Ranges")

	if accept == "bytes" {
		return true
	} else {
		return false
	}
}

func splitDownloadLength(length, div int64) []downloadInfo {

	divLength := length / div

	a := make([]downloadInfo, div)

	var i int64
	for i = 0; i < div; i++ {
		s := i * divLength
		e := (i+1)*divLength - 1

		a[i] = downloadInfo{s, e}
	}
	return a
}

func rangeDownload(index int, url string, s, e int64) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", s, e))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	outPath := strconv.Itoa(index) + ".temp.download"
	out, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, copyerr := io.Copy(out, res.Body)
	if copyerr != nil {
		return err
	}
	return nil
}

func joinFiles(filenames []string, saveFilePath string) error {
	files := make([]io.Reader, len(filenames))

	for i, filename := range filenames {
		files[i], _ = os.Open(filename)
	}

	src := io.MultiReader(files...)

	dst, err := os.Create(saveFilePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func deleteFiles(filenames []string) []deleteFileError {
	result := []deleteFileError{}

	for _, f := range filenames {
		err := os.Remove(f)
		if err != nil {
			result = append(result, deleteFileError{f, err})
		}
	}
	return result
}
