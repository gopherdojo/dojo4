package downloader

import "testing"

const URL = "https://images.pexels.com/photos/248304/pexels-photo-248304.jpeg"

func TestContentLength(t *testing.T) {
	c := NewClient(URL)
	l, err := c.contentLength()
	if err != nil {
		t.Error(err)
	}
	expect := 6480509
	if l != expect {
		t.Errorf("expect: %d, actual: %d", expect, l)
	}
}

func TestNewRangeProperties(t *testing.T) {
	ps := newRangeProperties(40000)
	t.Logf("%v", ps)
}

func TestRangeDownload(t *testing.T) {
	// c := NewClient(URL)
	// c.rangeDownload("1.jpg", 0, 5000)
	// c.rangeDownload("2.jpg", 5001, 6480509)
}

func TestMergeFiles(t *testing.T) {
	c := NewClient(URL)
	src := []string{"1.jpg", "2.jpg"}
	err := c.mergeFiles(src, "image.jpg")
	if err != nil {
		t.Error(err)
	}
}
