package filter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Option is an interface to filter files in some directories
type Option interface {
	SrcDir() string
	FromFormat() string
}

// Files filter files by an extension
func Files(c Option) ([]string, error) {
	var files []string
	ext := fmt.Sprintf(".%s", c.FromFormat())

	err := filepath.Walk(c.SrcDir(), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ext) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
