package imgconv_test

import (
	"github.com/gopherdojo/dojo4/kadai2/su-kun1899/imgconv"
	"path/filepath"
	"testing"
)

func TestFilePath_Is(t *testing.T) {
	type args struct {
		path        string
		imageFormat string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "jpeg file is jpeg",
			args: args{
				path:        filepath.Join("testdata", "Jpeg.jpg"),
				imageFormat: imgconv.JpegFormat,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imgconv.Is(tt.args.path, tt.args.imageFormat); got != tt.want {
				t.Errorf("FilePath.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTo(t *testing.T) {
	t.Skip()

	type args struct {
		imageFormat string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imgconv.ConvertTo(tt.args.imageFormat); got != tt.want {
				t.Errorf("ConvertTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
