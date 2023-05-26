package cmd

import (
	"github.com/DenChenn/blunder/internal/codegen/template"
	"github.com/DenChenn/blunder/internal/constant"
	"github.com/DenChenn/blunder/internal/util"
	"github.com/urfave/cli/v2"
	"path/filepath"
)

var Init = &cli.Command{
	Name:  "init",
	Usage: "initiate blunder configuration file and module structure",
	Action: func(cCtx *cli.Context) error {
		initPath := cCtx.Args().Get(0)
		if initPath == "" {
			return util.PrintErrAndReturn("you should specify the path to init")
		}

		s := util.PrintLoading("generating blunder.yaml ...", "blunder.yaml is generated successfully")
		blunderYamlPath := filepath.Join(initPath, "errors", constant.BlunderYamlFileName)
		if err := template.Generate(blunderYamlPath, constant.BlunderYamlTemplateFileName, nil); err != nil {
			return util.PrintErrAndReturn(err.Error())
		}
		s.Stop()

		return nil
	},
}
