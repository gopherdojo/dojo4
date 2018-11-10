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
	default:
		return []string{}
	}
}

// OutputExtension retrun output extentions
func (argument *Argument) OutputExtension() string {
	switch argument.OutputFormat {
	case "png":
		return ".png"
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
