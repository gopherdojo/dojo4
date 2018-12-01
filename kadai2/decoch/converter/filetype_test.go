package converter

import (
	"reflect"
	"testing"
)

func TestNewImageType(t *testing.T) {
	cases := []struct {
		str       string
		imageType ImageType
	}{
		{str: "png", imageType: Png},
		{str: "jpeg", imageType: Jpeg},
		{str: "gif", imageType: Gif},
	}

	for _, c := range cases {
		actual := NewImageType(c.str)
		if actual != c.imageType {
			t.Errorf("expected %v, actual %v.", c.imageType, actual)
		}
	}
}

func TestExtensions(t *testing.T) {
	cases := []struct {
		imageType  ImageType
		extensions []string
	}{
		{imageType: Png, extensions: pngExtensions},
		{imageType: Jpeg, extensions: jpegExtensions},
		{imageType: Gif, extensions: gifExtensions},
	}

	for _, c := range cases {
		actual := c.imageType.Extensions()
		if reflect.DeepEqual(c.extensions, actual) {
			t.Errorf("expected %v, actual %v.", c.extensions, actual)
		}
	}
}
