package cmd

import (
	"errors"
	"github.com/DenChenn/blunder/internal/codegen/model"
	"github.com/DenChenn/blunder/internal/codegen/util"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	GeneratedDirName          = "generated"
	ErrorFileName             = "error.go"
	ErrorFileTemplateFileName = "error.go.tmpl"
)

var Gen = &cli.Command{
	Name:  "gen",
	Usage: "generate all errors",
	Action: func(context *cli.Context) error {
		blunderPath := util.LocateBlunderYamlPath()
		if blunderPath == "" {
			return errors.New("blunder.yaml not found, please init the project first")
		}
		blunderRootPath := util.GetFileDirPath(blunderPath)

		f, err := os.ReadFile(blunderPath)
		if err != nil {
			return err
		}

		var blunderConfig model.Blunder
		if err := yaml.Unmarshal(f, &blunderConfig); err != nil {
			return err
		}

		generatedRootPath := filepath.Join(blunderRootPath, GeneratedDirName)
		// remove old generated folder
		_ = os.RemoveAll(generatedRootPath)

		// create generated folder
		_ = os.MkdirAll(generatedRootPath, os.ModePerm)

		for _, detail := range blunderConfig.Details {
			errorFilePath := filepath.Join(generatedRootPath, detail.Package, ErrorFileName)
			templateFilePath, err := util.GetTemplatePath(ErrorFileTemplateFileName)
			if err != nil {
				return err
			}

			// generate id for this error
			for i, _ := range detail.Errors {
				id := util.GetId(errorFilePath + detail.Errors[i].Code)
				detail.Errors[i].Id = id
			}

			if err := util.Generate(errorFilePath, templateFilePath, &detail); err != nil {
				return err
			}
		}

		return nil
	},
}
