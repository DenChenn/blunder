package cmd

import (
	"errors"
	"github.com/DenChenn/blunder/internal/codegen/template"
	"github.com/DenChenn/blunder/internal/codegen/util"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

const (
	BlunderYamlTemplateFileName = "blunder.yaml.tmpl"
)

var Init = &cli.Command{
	Name:  "init",
	Usage: "generate error",
	Action: func(cCtx *cli.Context) error {
		initPath := cCtx.Args().Get(0)
		if initPath == "" {
			return errors.New("you should specify the path to init")
		}

		blunderYamlPath := filepath.Join(initPath, "errors", util.BlunderYamlFileName)
		if err := template.Generate(blunderYamlPath, BlunderYamlTemplateFileName, nil); err != nil {
			return err
		}

		return nil
	},
}
