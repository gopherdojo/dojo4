package reader

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Files get files filter by extensions.
func Files(dir string, extensions []string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Files(filepath.Join(dir, file.Name()), extensions)...)
			continue
		}
		for _, ex := range extensions {
			if strings.HasSuffix(file.Name(), ex) {
				paths = append(paths, filepath.Join(dir, file.Name()))
			}
		}
	}

	return paths
}
