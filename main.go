package main

import (
	"github.com/DenChenn/blunder/internal/codegen/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "blunder",
		Usage: "a modern golang module for error handling",
	}

	app.Commands = []*cli.Command{
		cmd.Gen,
		cmd.Init,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
