package convert

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testNewStruct struct {
	name           string
	inputDir       string
	inputSrc       string
	inputDst       string
	wantErr        error
	wantConversion *Conversion
}

func TestNew(t *testing.T) {
	tests := []testNewStruct{
		{
			name:           "srcが想定外",
			inputDir:       "./testdata",
			inputSrc:       "txt",
			inputDst:       "png",
			wantErr:        ErrWrongInputSrcExtention,
			wantConversion: nil,
		},
		{
			name:           "dstが想定外",
			inputDir:       "./testdata",
			inputSrc:       "png",
			inputDst:       "txt",
			wantErr:        ErrWrongInputDstExtention,
			wantConversion: nil,
		},
		{
			name:     "png->jpg",
			inputDir: "./testdata",
			inputSrc: "png",
			inputDst: "jpg",
			wantErr:  nil,
			wantConversion: &Conversion{
				Files: []string{"testdata/talks.png", "testdata/subdir/ref.png"},
				Src:   "png",
				Dst:   "jpg",
			},
		},
	}

	// sort Files
	// https://godoc.org/github.com/google/go-cmp/cmp#example-Option--SortedSlice
	trans := cmp.Transformer("Sort", func(in []string) []string {
		out := append([]string(nil), in...)
		sort.Strings(out)
		return out
	})

	for _, tt := range tests {
		testNew(t, tt, trans)
	}
}

func testNew(t *testing.T, tt testNewStruct, tr cmp.Option) {
	t.Helper()
	ret, err := New(tt.inputDir, tt.inputSrc, tt.inputDst)
	if err != tt.wantErr {
		t.Errorf(`%s: New("%s", "%s", "%s") => error:%v, want error: %v`, tt.name, tt.inputDir, tt.inputSrc, tt.inputDst, err, tt.wantErr)
	}
	diff := cmp.Diff(ret, tt.wantConversion, tr)
	if diff != "" {
		t.Errorf(`%s: New("%s", "%s", "%s"): Conversion diff:%s`, tt.name, tt.inputDir, tt.inputSrc, tt.inputDst, diff)
	}
}

func TestValidateExtension(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{input: "jpg", want: true},
		{input: "png", want: true},
		{input: "gif", want: true},
		{input: "txt", want: false},
	}

	for _, tt := range tests {
		ret := validateExtension(tt.input)
		if ret != tt.want {
			t.Errorf(`validateExtension("%s") => %v, want %v`, tt.input, ret, tt.want)
		}
	}
}

func TestFilename(t *testing.T) {
	tests := []struct {
		inputName string
		inputSrc  string
		inputDst  string
		want      string
	}{
		{inputName: "testdata/talks.png", inputSrc: "png", inputDst: "jpg", want: "testdata/talks.jpg"},
		{inputName: "testdata/subdir/ref.png", inputSrc: "png", inputDst: "gif", want: "testdata/subdir/ref.gif"},
	}

	for _, tt := range tests {
		ret := filename(tt.inputName, tt.inputSrc, tt.inputDst)
		if ret != tt.want {
			t.Errorf(`filename("%s", "%s", "%s") => %s, want %s`, tt.inputName, tt.inputSrc, tt.inputDst, ret, tt.want)
		}
	}
}
