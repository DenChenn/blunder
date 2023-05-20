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
		Usage: "generate errors",
	}

	app.Commands = []*cli.Command{
		cmd.Gen,
		cmd.Init,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
