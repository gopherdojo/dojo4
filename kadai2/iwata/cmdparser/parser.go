package cmdparser

import (
	"flag"
	"fmt"
	"io"
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

// Cmd is a struct to parse arguments
type Cmd struct {
	stde io.Writer
}

func NewCmd(stde io.Writer) *Cmd {
	return &Cmd{stde}
}

// Parse is method to parse from command args
func (c *Cmd) Parse(args []string) (*CmdConfig, error) {
	conf := &CmdConfig{}
	flags := flag.NewFlagSet("imgconv", flag.ContinueOnError)
	flags.SetOutput(c.stde)
	flags.StringVar(&conf.fromExt, "from", "jpg", "Convert from image format")
	flags.StringVar(&conf.toExt, "to", "png", "Convert to image format")

	flags.Usage = func() {
		fmt.Fprintf(c.stde, `
Usage of %s:
   %s [OPTIONS] DIR
OPTIONS
`, args[0], args[0])
		flags.PrintDefaults()
	}

	if err := flags.Parse(args[1:]); err != nil {
		return nil, fmt.Errorf("Failed to paser args: %s", err)
	}

	if flags.NArg() != 1 {
		return nil, fmt.Errorf("Not support %d args, only one arg", flag.NArg())
	}
	if conf.fromExt == conf.toExt {
		return nil, fmt.Errorf("Cannot set the same format %s between -from and -to", conf.fromExt)
	}

	conf.dir = flags.Arg(0)
	return conf, nil
}
