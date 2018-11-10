package converter

type Type int

const (
	JpgToPng Type = iota + 1
	PngToJpg
)

func (c Type) TargetEx() *[]string {
	switch c {
	case JpgToPng:
		return &[]string{".jpeg", ".jpg"}
	case PngToJpg:
		return &[]string{".png"}
	default:
		return nil
	}
}

func (c Type) Convert() func(src string) error {
	switch c {
	case JpgToPng:
		return ToPng
	case PngToJpg:
		return ToJpg
	default:
		return nil
	}
}
