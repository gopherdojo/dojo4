package downloader

import (
	"fmt"
	"math"
)

type rangeProperty struct {
	path  string // 一時保存先のファイル名
	start int    // 開始のバイト数
	end   int    // 終了のバイト数
}

// rangeDownloadの引数として必要なrangePropertyを生成
func newRangeProperties(contentLength int) []*rangeProperty {
	num := 1000000

	maxIndex := int(math.Ceil(float64(contentLength) / float64(num)))

	f := func(i int) *rangeProperty {
		start := 0
		if i != 0 {
			start = i*num + 1
		}
		end := (i + 1) * num
		if end > contentLength {
			end = contentLength
		}
		return &rangeProperty{
			path:  fmt.Sprintf("file%d.jpg", i),
			start: start,
			end:   end,
		}
	}

	var out []*rangeProperty

	for i := 0; i < maxIndex; i++ {
		start := i + num
		end := start + num
		if end > contentLength {
			end = contentLength
		}
		out = append(out, f(i))
	}

	return out
}
