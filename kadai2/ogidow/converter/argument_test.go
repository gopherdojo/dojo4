package converter

import (
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestInputExtensions(t *testing.T) {
	cases := []struct {
		argument   Argument
		extensions []string
	}{
		{Argument{"", "jpg", ""}, []string{".jpg", ".jpeg"}},
		{Argument{"", "jpeg", ""}, []string{".jpg", ".jpeg"}},
		{Argument{"", "png", ""}, []string{".png"}},
		{Argument{"", "gif", ""}, []string{".gif"}},
		{Argument{"", "hoge", ""}, []string{}},
	}

	for _, c := range cases {
		result := c.argument.InputExtensions()
		if diff := cmp.Diff(result, c.extensions); diff != "" {
			t.Errorf("InputExtensions differs: (-got +want)\n%s", diff)
		}
	}
}

func TestOutputExtension(t *testing.T) {
	cases := []struct {
		argument   Argument
		extensions string
	}{
		{Argument{"", "", "jpg"}, ".jpg"},
		{Argument{"", "", "jpeg"}, ".jpg"},
		{Argument{"", "", "png"}, ".png"},
		{Argument{"", "", "gif"}, ".gif"},
		{Argument{"", "", "hoge"}, ""},
	}

	for _, c := range cases {
		result := c.argument.OutputExtension()
		if diff := cmp.Diff(result, c.extensions); diff != "" {
			t.Errorf("OutputExtension differs: (-got +want)\n%s", diff)
		}
	}
}

func TestIsValid(t *testing.T) {
	rand.Seed(time.Now().Unix())
	validFormats := [...]string{"jpg", "jpeg", "png", "gif"}
	cases := []struct {
		argument Argument
		valid    bool
	}{
		{Argument{"", validFormats[rand.Intn(len(validFormats))], validFormats[rand.Intn(len(validFormats))]}, true},
		{Argument{"", validFormats[rand.Intn(len(validFormats))], "hoge"}, false},
		{Argument{"", "hoge", validFormats[rand.Intn(len(validFormats))]}, false},
	}

	for _, c := range cases {
		result := c.argument.IsValid()
		if result != c.valid {
			t.Errorf(`IsValid expect="%T" actual="%T"`, c.valid, result)
		}
	}
}
