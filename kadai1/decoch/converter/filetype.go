package converter

type ImageType int

const (
	Png ImageType = iota + 1
	Jpeg
	Gif
)

var (
	jpegExtensions = []string{"jpeg", "jpg"}
	pngExtensions  = []string{"png"}
	gifExtensions  = []string{"gif"}
)

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

func (i ImageType) Extensions() *[]string {
	switch i {
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

func (i ImageType) Convert() func(src string) error {
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
