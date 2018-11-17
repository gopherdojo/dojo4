package converter

type ImageType int

const (
	Unknown ImageType = iota
	Png
	Jpeg
	Gif
)

var (
	jpegExtensions = []string{"jpeg", "jpg"}
	pngExtensions  = []string{"png"}
	gifExtensions  = []string{"gif"}
)

// NewImageType create a ImageType instance.
func NewImageType(str string) *ImageType {
	for _, imageType := range []ImageType{Png, Jpeg, Gif} {
		for _, ex := range *imageType.Extensions() {
			if ex == str {
				return &imageType
			}
		}
	}
	return nil
}

// Extensions get extensions of type.
func (i *ImageType) Extensions() *[]string {
	switch *i {
	case Png:
		return &pngExtensions
	case Jpeg:
		return &jpegExtensions
	case Gif:
		return &gifExtensions
	default:
		return nil
	}
}

// Converter get a image convert function.
func (i ImageType) Converter() func(src string) error {
	switch i {
	case Png:
		return convertToPng
	case Jpeg:
		return convertToJpg
	case Gif:
		return convertToGif
	default:
		return nil
	}
}
