package main

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

type ConvImgMock struct {
	FakeConvert func() error
}
func (c *ConvImgMock) Convert() error {
	return c.FakeConvert()
}

func TestCLI_Run(t *testing.T) {
	t.Helper()
	cases := []struct {
		convimg *ConvImgMock
		output string
		status int
	}{
		{
			convimg: &ConvImgMock{
				FakeConvert: func() error {
					return nil
				},
			},
			output: "",
			status: 0,
		},
		{
			convimg: &ConvImgMock{
				FakeConvert: func() error {
					return errors.New("Error: convert img")
				},
			},
			output: "Error: convert img",
			status: 1,
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%#v, %s, %d", c.convimg, c.output, c.status), func(t *testing.T) {
			writer := new(bytes.Buffer)
			cli := CLI{outStream:writer, errStream: writer, convert: c.convimg}
			if cli.Run() != c.status {
				t.Errorf("unexpected status want %d, get %d", c.status, cli.Run())
			}
		})
	}
}
