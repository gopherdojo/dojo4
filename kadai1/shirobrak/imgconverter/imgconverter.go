package imgconverter

import (
	"flag"
	"fmt"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

// ImageConverter is Image Format Converter.
type ImageConverter struct {
	srcName string // Convert Directory Name
	srcExt  string // File format before conversion
	toExt   string // File format after conversion
	mode    string // Convert Mode : File or Directory(Default)
}

// ImageConverterError has params for error message related to ImageConverter
type ImageConverterError struct {
	msg  string
	code int
}

// Init Funtion sets params inputted in cmd line.
func (ic *ImageConverter) Init() {

	// define option command
	flag.StringVar(&ic.srcExt, "src", "jpg", "変換元のファイルフォーマット [ jpeg, jpg, png ]")
	flag.StringVar(&ic.toExt, "to", "png", "変換後のファイルフォーマット [ jpeg, jpg, png ]")
	flag.StringVar(&ic.mode, "mode", "d", "変換モード [ f(file) , d(directory) ]")

	// set param entered from the cmd line
	flag.Parse()

	// set param source name
	ic.srcName = flag.Arg(0)

	return
}

// Run converts the Image Files.
func (ic *ImageConverter) Run() error {

	switch mode := ic.mode; mode {
	case "d":
		if err := ic.convertDir(); err != nil {
			return err
		}
	case "f":
		if err := ic.convertFile(); err != nil {
			return err
		}
	}
	return nil
}

// Error Function outputs error message related to ImageConverter
func (err *ImageConverterError) Error() string {
	return fmt.Sprintf("err %s [code=%d]", err.msg, err.code)
}

func (ic *ImageConverter) convertFile() error {

	if (ic.srcExt == "jpeg" || ic.srcExt == "jpg") && ic.toExt == "png" {
		if err := ic.convertJPEG2PNG(ic.srcName); err != nil {
			return err
		}
	} else if ic.srcExt == "png" && (ic.toExt == "jpeg" || ic.toExt == "jpg") {
		if err := ic.convertPNG2JPEG(ic.srcName); err != nil {
			return err
		}
	} else {
		return &ImageConverterError{msg: "Invalid Flags Set", code: 101}
	}
	return nil
}

// ConvertDir converts the Files in the specified Directory.
func (ic *ImageConverter) convertDir() error {

	srcFileList, err := ic.getFilePathList(ic.srcName)
	if err != nil {
		return err
	}

	if (ic.srcExt == "jpeg" || ic.srcExt == "jpg") && ic.toExt == "png" {
		for _, filePath := range srcFileList {
			if err := ic.convertJPEG2PNG(filePath); err != nil {
				return err
			}
		}
	} else if ic.srcExt == "png" && (ic.toExt == "jpeg" || ic.toExt == "jpg") {
		for _, filePath := range srcFileList {
			if err := ic.convertPNG2JPEG(filePath); err != nil {
				return err
			}
		}
	} else {
		return &ImageConverterError{msg: "Invalid Flags Set", code: 101}
	}

	return nil
}

// getFilePathList returns file`s path list existing in directory.
func (ic *ImageConverter) getFilePathList(dirName string) ([]string, error) {
	var filePathList []string
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		fileName := fileInfo.Name()
		if !fileInfo.IsDir() {
			if filepath.Ext(fileName) == "."+ic.srcExt {
				filePath := path.Join(dirName, fileName)
				filePathList = append(filePathList, filePath)
			}
		} else {
			dirPath := path.Join(dirName, fileName)
			res, err := ic.getFilePathList(dirPath)
			if err != nil {
				return nil, err
			}
			filePathList = append(filePathList, res...)
		}
	}
	return filePathList, nil
}

// newFileReader returns Reader.
func (ic *ImageConverter) newFileReader(filePath string) (io.Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// convertJPEG2PNG converts the PNG Format File into the JPEG Format File.
func (ic *ImageConverter) convertJPEG2PNG(filePath string) error {

	reader, err := ic.newFileReader(filePath)
	if err != nil {
		return err
	}

	Image, err := jpeg.Decode(reader)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(ic.srcExt + "$")
	newFilePath := re.ReplaceAllString(filePath, ic.toExt)
	newImageFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer newImageFile.Close()

	if err := png.Encode(newImageFile, Image); err != nil {
		return err
	}

	return nil
}

// convertPNG2JPEG converts the JPEG/JPG Format File into the PNG Format File.
func (ic *ImageConverter) convertPNG2JPEG(filePath string) error {

	reader, err := ic.newFileReader(filePath)
	if err != nil {
		return err
	}

	Image, err := png.Decode(reader)
	if err != nil {
		return err
	}

	re := regexp.MustCompile(ic.srcExt + "$")
	newFilePath := re.ReplaceAllString(filePath, ic.toExt)
	newImageFile, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer newImageFile.Close()

	if err := jpeg.Encode(newImageFile, Image, nil); err != nil {
		return err
	}

	return nil
}
