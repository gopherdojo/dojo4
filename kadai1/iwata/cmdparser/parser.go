package cmdparser

import (
	"flag"
	"fmt"
)

// CmdConfig is a configuration
type CmdConfig struct {
	dir     string
	fromExt string
	toExt   string
}

// SrcDir returns a directory path
func (c *CmdConfig) SrcDir() string {
	return c.dir
}

// FromFormat returns a format about source images
func (c *CmdConfig) FromFormat() string {
	return c.fromExt
}

// ToFormat returns a format for target images
func (c *CmdConfig) ToFormat() string {
	return c.toExt
}

// Parse is method to parse from command args
func Parse() (*CmdConfig, error) {
	c := &CmdConfig{}
	flag.StringVar(&c.fromExt, "from", "jpg", "Convert from image extention")
	flag.StringVar(&c.toExt, "to", "png", "Convert to image extention")
	flag.Parse()

	if flag.NArg() != 1 {
		return nil, fmt.Errorf("Not support %d args, only one arg", flag.NArg())
	}

	c.dir = flag.Arg(0)
	return c, nil
}
