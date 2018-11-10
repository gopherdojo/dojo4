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

	fileNames := reader.Files(command.Dir, *command.ConvertType.TargetEx())
	for _, fileName := range fileNames {
		if err := command.ConvertType.Convert()(fileName); err != nil {
			log.Fatal(err)
		}
	}
}
