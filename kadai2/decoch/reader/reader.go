package reader

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Files get files filter by extensions.
func Files(dir string, extensions []string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			files, err := Files(filepath.Join(dir, file.Name()), extensions)
			if err != nil {
				return []string{}, err
			}
			paths = append(paths, files...)
			continue
		}
		if hasExtension(file, extensions) {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths, nil
}

func hasExtension(file os.FileInfo, extensions []string) bool {
	for _, ex := range extensions {
		if strings.HasSuffix(file.Name(), ex) {
			return true
		}
	}
	return false
}
