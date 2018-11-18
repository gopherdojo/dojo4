package converter

// Argument is a command line options
type Argument struct {
	Dir          string
	InputFormat  string
	OutputFormat string
}

// InputExtensions retrun input extentions
func (argument *Argument) InputExtensions() []string {
	switch argument.InputFormat {
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
func (argument *Argument) OutputExtension() string {
	switch argument.OutputFormat {
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
func (argument *Argument) IsValid() bool {
	if len(argument.InputExtensions()) == 0 {
		return false
	}
	if argument.OutputExtension() == "" {
		return false
	}
	return true
}
