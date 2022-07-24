package main

import (
	"log"

	"github.com/codemaestro64/stardust/cli"
)

func mainCore() error {
	if command := cli.GetCommandStub(); command != nil {
		return command.Execute()
	}

	return nil
}

func main() {
	if err := mainCore(); err != nil {
		log.Fatal(err)
	}
}
