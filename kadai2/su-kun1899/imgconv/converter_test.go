package imgconv_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/gopherdojo/dojo4/kadai2/su-kun1899/imgconv"
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
		{
			name: "gif file is not jpeg",
			args: args{
				path:        filepath.Join("testdata", "Gif.gif"),
				imageFormat: imgconv.JpegFormat,
			},
			want: false,
		},
		{
			name: "directory is not jpeg",
			args: args{
				path:        filepath.Join("testdata"),
				imageFormat: imgconv.JpegFormat,
			},
			want: false,
		},
		{
			name: "not exist file is not jpeg",
			args: args{
				path:        filepath.Join("testdata", "dummy.jpg"),
				imageFormat: imgconv.JpegFormat,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imgconv.Is(tt.args.path, tt.args.imageFormat); got != tt.want {
				t.Errorf("imgConv.Is() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTo(t *testing.T) {
	type args struct {
		srcFile     string
		targetFile  string
		imageFormat string
	}
	tests := []struct {
		name     string
		args     args
		wantFile string
		wantErr  bool
	}{
		{
			name: "jpg to png",
			args: args{
				srcFile:     filepath.Join("testdata", "Jpeg.jpg"),
				targetFile:  filepath.Join("testdata", "work_Jpeg.jpg"),
				imageFormat: imgconv.PngFormat,
			},
			wantFile: filepath.Join("testdata", "work_Jpeg.png"),
			wantErr:  false,
		},
		{
			name: "jpg to gif",
			args: args{
				srcFile:     filepath.Join("testdata", "Jpeg.jpg"),
				targetFile:  filepath.Join("testdata", "work_Jpeg.jpg"),
				imageFormat: imgconv.GifFormat,
			},
			wantFile: filepath.Join("testdata", "work_Jpeg.gif"),
			wantErr:  false,
		},
		{
			name: "png to jpeg",
			args: args{
				srcFile:     filepath.Join("testdata", "Png.png"),
				targetFile:  filepath.Join("testdata", "work_Png.png"),
				imageFormat: imgconv.JpegFormat,
			},
			wantFile: filepath.Join("testdata", "work_Png.jpg"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup working file
			copyFile(t, tt.args.srcFile, tt.args.targetFile)
			defer os.Remove(tt.wantFile)

			if err := imgconv.ConvertTo(tt.args.targetFile, tt.args.imageFormat); (err != nil) != tt.wantErr {
				t.Errorf("imgConv.ConvertTo() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got := imgconv.Is(tt.wantFile, tt.args.imageFormat); !got {
				t.Errorf("imgConv.Is() = %v, want %v", got, true)
			}
		})
	}
}

func copyFile(t *testing.T, src, dest string) {
	t.Helper()

	bytes, err := ioutil.ReadFile(src)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	err = ioutil.WriteFile(dest, bytes, 0655)
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
}
