package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCLI_Run(t *testing.T) {
	cases := []struct {
		command           string
		expectedOutStream string
		expectedErrStream string
		expectedExitCode  int
	}{
		{
			command:           "converter-cli",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up converter-cli: invalid argument\nPlease specify the exact one path to a directly or a file\n\n",
			expectedExitCode:  ExitCodeBadArgs,
		},
		{
			command:           "converter-cli testdata testdata2",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up converter-cli: invalid argument\nPlease specify the exact one path to a directly or a file\n\n",
			expectedExitCode:  ExitCodeBadArgs,
		},
		{
			command:           "converter-cli --from .svg",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up converter-cli: invalid extension `.svg` is given for --from flag\nPlease choose an extension from one of those: [.gif .jpeg .jpg .png]\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "converter-cli --to .svg",
			expectedOutStream: "",
			expectedErrStream: "Failed to set up converter-cli: invalid extension `.svg` is given for --to flag\nPlease choose an extension from one of those: [.gif .jpeg .jpg .png]\n\n",
			expectedExitCode:  ExitCodeInvalidFlagError,
		},
		{
			command:           "converter-cli testdata",
			expectedOutStream: "",
			expectedErrStream: "Failed to execute converter-cli\ncould not find files with the specified extension. path: testdata, extension: .jpg\n\n",
			expectedExitCode:  ExitCodeExpectedError,
		},
		{
			command:           "converter-cli testdata/unknown.jpg",
			expectedOutStream: "",
			expectedErrStream: "Failed to execute converter-cli\nlstat testdata/unknown.jpg: no such file or directory\n\n",
			expectedExitCode:  ExitCodeExpectedError,
		},
		{
			command:           "converter-cli --from .jpeg testdata",
			expectedOutStream: "converter-cli successfully converted following files to `.png`.\n[testdata/jpeg-image.jpeg]\n\n",
			expectedErrStream: "",
			expectedExitCode:  ExitCodeOK,
		},
	}

	for i, tc := range cases {
		outStream := new(bytes.Buffer)
		errStream := new(bytes.Buffer)

		cli := CLI{outStream: outStream, errStream: errStream}
		args := strings.Split(tc.command, " ")

		if got := cli.Run(args); got != tc.expectedExitCode {
			t.Errorf("#%d %q exits with %d, want %d", i, tc.command, got, tc.expectedExitCode)
		}

		if got := outStream.String(); got != tc.expectedOutStream {
			t.Errorf("#%d Unexpected outStream has returned: want: %s, got: %s", i, tc.expectedOutStream, got)
		}

		if got := errStream.String(); got != tc.expectedErrStream {
			t.Errorf("#%d Unexpected errStream has returned: want: %s, got: %s", i, tc.expectedErrStream, got)
		}

		cleanup(t)
	}
}

func cleanup(t *testing.T) {
	t.Helper()

	err := filepath.Walk("testdata", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path != "testdata/jpeg-image.jpeg" && path != "testdata" {
			return os.Remove(path)
		}

		return nil
	})

	if err != nil {
		t.Errorf("failed to cleanup testdata: %s", err)
	}
}
