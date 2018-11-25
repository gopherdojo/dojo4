package downloader

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type Downloader int

type downloadInfo struct {
	startByte int64
	endByte int64
}

// Download
// description:Download specified URL.
// parameter
// url :download target url.
// saveFilePath:save file path.
// div:Number of divided downloads.
func (d *Downloader) Download(url string,saveFilePath string,div int64) error{
	res,err := http.Head(url)
	if err != nil {
		fmt.Println("http Head error")
		return err
	}
	if(res.StatusCode != 200) {
		fmt.Println("bad status code")
		fmt.Println("status code:", res.StatusCode)
		return err
	}

	if !canRangeDownload(res.Header) {
		fmt.Println("range download not support.")
		return err
	}

	di := splitDownloadLength(res.ContentLength,div)
	filenames := make([]string,div)

	var eg errgroup.Group

	i:=1
	for _,d := range di{
		j:=i
		eg.Go (func() error {
			err := rangeDownload(j, url, d.startByte, d.endByte)
			return err
		})
		filenames[i-1] = strconv.Itoa(j)+".temp.download"
		i++
	}
	if err := eg.Wait();err != nil{
		fmt.Println("download error")
		return err
	}

	if err:= joinFiles(filenames,saveFilePath);err != nil{
		fmt.Println("join file error")
		return err
	}

	if err:= deleteFiles(filenames);err != nil{
		fmt.Println("join file error")
		return err
	}
	fmt.Println("download finish.")
	return nil
}

func canRangeDownload(h http.Header) bool{
	f := false
	for k, v := range h {
		if(k == "Accept-Ranges" && len(v) > 0 && v[0] == "bytes"){
			f=true
			break
		}
	}
	return f
}


func splitDownloadLength(length int64,div int64)[]downloadInfo {

	divLength := length / div

	a := make([]downloadInfo,div)


	var i int64
	for i=0;i<div;i++ {
		s := i * divLength
		e := (i+1)	 * divLength - 1

		a[i] = downloadInfo{s,e}
	}
	return a
}

func rangeDownload(index int,url string,s int64,e int64) error{
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", s, e))

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	if err != nil {
		return err
	}


	out_path := strconv.Itoa(index)+".temp.download"
	out, err := os.Create(out_path)
	defer out.Close()
	if err != nil {
		return err
	}

	io.Copy(out, res.Body)
	return nil
}

func joinFiles(filenames []string,saveFilePath string) error{
	files := make([]io.Reader, len(filenames))

	for i, filename := range filenames {
		files[i], _ = os.Open(filename)
	}

	reader := io.MultiReader(files...)
	b, _ := ioutil.ReadAll(reader)

	file, err := os.Create(saveFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	file.Write(([]byte)(b))
	return nil
}

func deleteFiles(filenames []string) error{
	for _,f := range  filenames {
		err := os.Remove(f)
		if err != nil {
			return err
		}
	}
	return nil
}