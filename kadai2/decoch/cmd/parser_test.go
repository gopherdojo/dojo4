package cmd

import (
	"testing"
)

func TestParseCommandLineArgs(t *testing.T) {
	validArgs := []string{"dummy"}
	dirName, err := parseCommandLineArgs(validArgs)
	if err != nil {
		t.Fatal(validArgs)
	}

	if validArgs[0] != dirName {
		t.Fatal(validArgs, dirName)
	}
}

func TestParseCommandLineArgsIncorrectNumber(t *testing.T) {
	zeroArgs := []string{"dummy", "dummy"}
	tooManyArgs := []string{"dummy", "dummy"}
	if _, err := parseCommandLineArgs(zeroArgs); err == nil {
		t.Fatal(zeroArgs)
	}
	if _, err := parseCommandLineArgs(tooManyArgs); err == nil {
		t.Fatal(zeroArgs)
	}
}

func TestParseImageType(t *testing.T) {
	testParseImage(t, "jpeg")
	testParseImage(t, "png")
	testParseImage(t, "gif")
}

func testParseImage(t *testing.T, str string) {
	t.Helper()
	if _, err := parseImageType(str); err != nil {
		t.Errorf("invalid extension: %s", str)
	}
}

func TestParseNotImageType(t *testing.T) {
	if _, err := parseImageType("dummy"); err == nil {
		t.Fatal(err)
	}
}
