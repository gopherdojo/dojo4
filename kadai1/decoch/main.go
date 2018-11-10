package main

import (
	"log"

	"github.com/decoch/dojo4/kadai1/decoch/cmd"
	"github.com/decoch/dojo4/kadai1/decoch/reader"
)

func main() {
	command, err := cmd.Parse()
	if err != nil {
		log.Fatal(err)
	}

	fileNames := reader.Files(command.Dir, *command.InputType.Extensions())
	for _, fileName := range fileNames {
		if err := command.OutputType.Converter()(fileName); err != nil {
			log.Fatal(err)
		}
	}
}
