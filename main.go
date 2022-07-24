package main

import (
	"fmt"
	"os"

	"github.com/codemaestro64/stardust_cli/cli"
)

func main() {
	if err := cli.New().Run(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
