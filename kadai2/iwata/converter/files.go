package converter

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/sync/errgroup"
)

// Converter is a struct to convert images
type Converter struct {
	out io.Writer
	opt Option
	mu  sync.Mutex
}

// Option is an interface for converter
type Option interface {
	ToFormat() string
}

func NewConverter(out io.Writer, opt Option) *Converter {
	return &Converter{out: out, opt: opt}
}

// ConvertFiles converts a list of file to other formats
func (c *Converter) ConvertFiles(files []string) error {
	if len(files) == 0 {
		return nil
	}

	eg := errgroup.Group{}
	for _, f := range files {
		f := f
		eg.Go(func() error {
			return c.convert(f)
		})
	}

	if err := eg.Wait(); err != nil {
		return fmt.Errorf("Failed to convert a list of files: %s", err)
	}

	return nil
}

func (c *Converter) convert(f string) (err error) {
	src, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("Failed to open %s: %s", f, err)
	}
	defer func() {
		if cerr := src.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	cv, err := c.resolveConverter(src)
	if err != nil {
		return fmt.Errorf("Failed to resolve converter about %s: %s", f, err)
	}

	dist := distName(f, c.opt.ToFormat())
	df, err := os.Create(dist)
	if err != nil {
		return fmt.Errorf("Failed to create a %s: %s", dist, err)
	}
	defer func() {
		if cerr := df.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if err := cv.convert(df); err != nil {
		return fmt.Errorf("Failed to convert %s to %s: %s", f, dist, err)
	}
	c.mu.Lock()
	fmt.Fprintf(c.out, "Converted %s to %s\n", f, dist)
	c.mu.Unlock()

	return nil
}

func distName(f, toExt string) string {
	ext := filepath.Ext(f)
	return fmt.Sprintf("%s.%s", strings.TrimSuffix(f, ext), toExt)
}
