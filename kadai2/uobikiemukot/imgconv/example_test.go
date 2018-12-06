package imgconv_test

import (
	"fmt"
	"image/gif"
	"path/filepath"

	"github.com/gopherdojo/dojo4/kadai2/uobikiemukot/imgconv"
)

func ExampleConverter_Run() {
	c := imgconv.Converter{
		InputFormat:  "png",
		OutputFormat: "gif",
		GIFOptions:   gif.Options{NumColors: 256},
	}

	root, err := filepath.Abs("./testdata/subdir")
	if err != nil {
		return
	}

	err = c.Run(root)
	if err != nil {
		return
	}

	fmt.Println("succeeded!")
	// Output: succeeded!
}
