package imgconv_test

import (
	"github.com/gopherdojo/dojo4/kadai1/su-kun1899/imgconv"
	"os"
	"path/filepath"
	"testing"
)

func TestPngConv_Convert(t *testing.T) {
	type args struct {
		src  string
		dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "png to jpeg",
			args: args{
				src:  filepath.Join("testdata", "syokuji_computer.jpg"),
				dest: filepath.Join("testdata", "syokuji_computer.png"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			j := imgconv.PngConv{}
			defer func() {
				if err := os.Remove(tt.args.dest); err != nil {
					t.Error("unexpected error:", err)
					return
				}
			}()

			// when
			err := j.Convert(tt.args.src, tt.args.dest)

			// then
			if (err != nil) != tt.wantErr {
				t.Errorf("PngConv.Convert() error = %v, wantErr %v", err, tt.wantErr)
			}

			// TODO 変換後のファイルを確認する
		})
	}
}
