package main

import (
	"log"

	"github.com/decoch/dojo4/kadai2/decoch/cmd"
	"github.com/decoch/dojo4/kadai2/decoch/reader"
)

func main() {
	command, err := cmd.Parse()
	if err != nil {
		log.Fatal(err)
	}

	fileNames, err := reader.Files(command.Dir, *command.InputType.Extensions())
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range fileNames {
		if err := command.OutputType.Converter()(fileName); err != nil {
			log.Fatal(err)
		}
	}
}
