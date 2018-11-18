package converter

// Argument is a command line options
type Argument struct {
	Dir          string
	InputFormat  string
	OutputFormat string
}

// InputExtensions retrun input extentions
func (a *Argument) InputExtensions() []string {
	switch a.InputFormat {
	case "jpg", "jpeg":
		return []string{".jpg", ".jpeg"}
	case "png":
		return []string{".png"}
	case "gif":
		return []string{".gif"}
	default:
		return []string{}
	}
}

// OutputExtension retrun output extentions
func (a *Argument) OutputExtension() string {
	switch a.OutputFormat {
	case "jpg", "jpeg":
		return ".jpg"
	case "png":
		return ".png"
	case "gif":
		return ".gif"
	default:
		return ""
	}
}

// IsValid is Verify options
func (a *Argument) IsValid() bool {
	if len(a.InputExtensions()) == 0 {
		return false
	}
	if a.OutputExtension() == "" {
		return false
	}
	return true
}
