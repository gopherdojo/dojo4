package main

import (
	"dojo4/kadai1/yasu/converter"
	"io/ioutil"
	"os"
	"sync"
)

func main() {
	directories := os.Args[1:]
	wg := sync.WaitGroup{}
	for _, directory := range directories {
		wg.Add(1)
		go func(directory string) {
			analyzeFiles(directory)
			wg.Done()
		}(directory)
	}
	wg.Wait()
}

func analyzeFiles(directory string) {
	files, err := ioutil.ReadDir(directory) // Get all file information in directory
	if err != nil {
		panic(err)
	}
	for _, fileInfo := range files {
		file, err := os.Open(directory + "/" + fileInfo.Name()) // Read file
		if err != nil {
			panic(err)
		}
		converter.ConvertImg(file, directory)
	}
}
