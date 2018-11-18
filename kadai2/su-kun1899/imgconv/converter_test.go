package imgconv

import (
	"path/filepath"
	"testing"
)

func TestFilePath_Is(t *testing.T) {
	type args struct {
		imageFormat string
	}
	tests := []struct {
		name string
		path FilePath
		args args
		want bool
	}{
		{name: "jpeg file is jpeg", path: FilePath(filepath.Join("testdata", "Jpeg.jpg")), args: args{imageFormat: JpegFormat}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.path.Is(tt.args.imageFormat); got != tt.want {
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
			if got := ConvertTo(tt.args.imageFormat); got != tt.want {
				t.Errorf("ConvertTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
