package converter

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

// ConvertFiles converts a list of file to other formats
func ConvertFiles(files []string, c ConvertOption) error {
	if len(files) == 0 {
		return nil
	}

	eg := errgroup.Group{}
	for _, f := range files {
		f := f
		eg.Go(func() error {
			return convert(f, c)
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("Failed to convert a list of files: %s", err)
	}

	return nil
}

func convert(f string, c ConvertOption) error {
	src, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", f, err)
	}
	defer src.Close()

	cv, err := resolveConverter(src, c)
	if err != nil {
		return fmt.Errorf("Failed to resolve converter about %s: %s", f, err)
	}

	dist := distName(f, c.ToFormat())
	df, err := os.Create(dist)
	if err != nil {
		return fmt.Errorf("Failed to create a %s: %s", dist, err)
	}
	defer df.Close()

	if err := cv.Convert(df); err != nil {
		return fmt.Errorf("Failed to convert %s to %s: %s", f, dist, err)
	}
	fmt.Printf("Converted %s to %s\n", f, dist)

	return nil
}

func distName(f, toExt string) string {
	ext := filepath.Ext(f)
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(f, ext), toExt)
}
