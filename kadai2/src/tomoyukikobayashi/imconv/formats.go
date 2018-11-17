package imconv

import "strings"

// HACK 本当はimmutableにしたいけどどうもできなそう
// ポインタだししょうがないか。。。厳密にやるなら直接触らせないのがいいかな
var (
	jpgExts     = []string{"jpg", "jpeg"}
	pngExts     = []string{"png"}
	supportExts = [][]string{
		jpgExts,
		pngExts,
	}
	allSupportExts = append(jpgExts, pngExts...)
)

// Supported returns if the package can handle speficied file format
func Supported(ext string) bool {
	for _, spext := range allSupportExts {
		if strings.ToLower(ext) == spext {
			return true
		}
	}
	return false
}

// SupportedExtensions returns image formats the package can handle
func SupportedExtensions() []string {
	clone := make([]string, len(allSupportExts))
	copy(clone, allSupportExts)
	return clone
}

// GetFormatThesaurus returns image formats that are the same with arg ext
func GetFormatThesaurus(ext string) []string {
	for _, spexts := range supportExts {
		for _, spext := range spexts {
			if strings.ToLower(ext) == spext {
				return spexts
			}
		}
	}
	return nil
}

func isSameFormat(from string, to string) bool {
	th := GetFormatThesaurus(from)
	for _, v := range th {
		if strings.ToLower(to) == v {
			return true
		}
	}
	return false
}
