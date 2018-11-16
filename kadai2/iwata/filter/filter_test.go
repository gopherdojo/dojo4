package filter_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/iwata/filter"
)

type mockOption struct {
	dir    string
	format string
}

func (o *mockOption) SrcDir() string {
	return o.dir
}

func (o *mockOption) FromFormat() string {
	return o.format
}

func newMockOption(dir, format string) filter.Option {
	return &mockOption{dir, format}
}

func TestFiles(t *testing.T) {
	var goLogoDir string = "../test/fixtures/images/Go-Logo"

	type args struct {
		c filter.Option
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{"invalid path", args{newMockOption("%$ !!", "png")}, nil, true},
		{"upper case format", args{newMockOption(goLogoDir, "PNG")}, nil, false},
		{"png", args{newMockOption(goLogoDir, "png")}, []string{
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_Aqua.png"),
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_Black.png"),
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_Blue.png"),
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_Fuchsia.png"),
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_LightBlue.png"),
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_White.png"),
			fmt.Sprintf("%s/%s", goLogoDir, "PNG/Go-Logo_Yellow.png"),
		}, false},
		{"jpg", args{newMockOption(goLogoDir, "jpg")}, []string{
			fmt.Sprintf("%s/%s", goLogoDir, "JPG/Go-Logo_Aqua.jpg"),
			fmt.Sprintf("%s/%s", goLogoDir, "JPG/Go-Logo_Black.jpg"),
			fmt.Sprintf("%s/%s", goLogoDir, "JPG/Go-Logo_Blue.jpg"),
			fmt.Sprintf("%s/%s", goLogoDir, "JPG/Go-Logo_Fuchsia.jpg"),
			fmt.Sprintf("%s/%s", goLogoDir, "JPG/Go-Logo_LightBlue.jpg"),
			fmt.Sprintf("%s/%s", goLogoDir, "JPG/Go-Logo_Yellow.jpg"),
		}, false},
		{"svg", args{newMockOption(goLogoDir, "svg")}, []string{
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_Aqua.svg"),
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_Black.svg"),
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_Blue.svg"),
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_Fuchsia.svg"),
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_LightBlue.svg"),
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_White.svg"),
			fmt.Sprintf("%s/%s", goLogoDir, "SVG/Go-Logo_Yellow.svg"),
		}, false},
		{"eps", args{newMockOption(goLogoDir, "eps")}, []string{
			fmt.Sprintf("%s/%s", goLogoDir, "EPS/Go-Logo_Versions.eps"),
		}, false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := filter.Files(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("filter.Files() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter.Files() = %v, want %v", got, tt.want)
			}
		})
	}
}
