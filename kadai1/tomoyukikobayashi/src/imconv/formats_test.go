package imconv

import (
	"reflect"
	"testing"
)

func Test_Supported(t *testing.T) {
	tests := []struct {
		ext  string
		want bool
	}{
		{"jpg", true},
		{"jpeg", true},
		{"JPG", true},
		{"png", true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			if got := Supported(tt.ext); got != tt.want {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}
}

func Test_GetFormatThesaurus(t *testing.T) {
	tests := []struct {
		ext  string
		want []string
	}{
		{"jpg", jpgExts},
		{"JPG", jpgExts},
		{"jpeg", jpgExts},
		{"png", pngExts},
		{"invalid", nil},
		{"", nil},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			if got := GetFormatThesaurus(tt.ext); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}
}

func Test_isSameFormat(t *testing.T) {
	tests := []struct {
		from string
		to   string
		want bool
	}{
		{"jpg", "jpeg", true},
		{"JPG", "jpeg", true},
		{"jpeg", "jpeg", true},
		{"JPG", "png", false},
		{"png", "invalid", false},
		{"invalid", "invalid", false},
	}

	for _, tt := range tests {
		t.Run(tt.from+"_"+tt.to, func(t *testing.T) {
			if got := isSameFormat(tt.from, tt.to); got != tt.want {
				t.Fatalf("want = %#v, got = %#v", tt.want, got)
			}
		})
	}
}
