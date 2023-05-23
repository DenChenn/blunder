package cmd

import (
	"errors"
	"github.com/DenChenn/blunder/internal/codegen/gpt"
	"github.com/DenChenn/blunder/internal/codegen/model"
	"github.com/DenChenn/blunder/internal/codegen/util"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

const (
	GeneratedDirName                    = "generated"
	ErrorFileName                       = "error.go"
	ErrorFileTemplateFileName           = "error.go.tmpl"
	GenerateBlunderYamlTemplateFileName = "gen_blunder.yaml.tmpl"
)

var Gen = &cli.Command{
	Name:  "gen",
	Usage: "generate all errors",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "complete",
			Value:   false,
			Usage:   "auto complete error detail with gpt3",
			Aliases: []string{"c"},
		},
	},
	Action: func(cCtx *cli.Context) error {
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

		// check all error code is provided
		if !checkAllErrorCodeIsProvided(&blunderConfig) {
			return errors.New("in blunder.yaml, all error code must be provided")
		}

		// complete error detail with gpt3
		complete := cCtx.Bool("complete")
		if complete {
			// check if user provide gpt3 api token
			_, exist := os.LookupEnv("OPENAI_API_TOKEN")
			if !exist {
				return errors.New("OPENAI_API_TOKEN not found, please export it with your own openai api token")
			}

			hasSomethingToComplete, indexMap, errorCodes := determineWhichToComplete(&blunderConfig)
			if hasSomethingToComplete {
				completed, err := gpt.CompleteErrorDetail(errorCodes)
				if err != nil {
					return err
				}

				for _, ec := range completed {
					blunderConfig.
						Details[indexMap[ec.Code].DetailIndex].
						Errors[indexMap[ec.Code].ErrorIndex].
						HttpStatusCode = ec.HttpStatusCode
					blunderConfig.
						Details[indexMap[ec.Code].DetailIndex].
						Errors[indexMap[ec.Code].ErrorIndex].
						GrpcStatusCode = ec.GrpcStatusCode
					blunderConfig.
						Details[indexMap[ec.Code].DetailIndex].
						Errors[indexMap[ec.Code].ErrorIndex].
						Message = ec.Message
				}
			}

			// generate blunder.yaml again to record the completion
			templatePath, err := util.GetTemplatePath(GenerateBlunderYamlTemplateFileName)
			if err != nil {
				return err
			}

			if err := util.Generate(blunderPath, templatePath, &blunderConfig); err != nil {
				return err
			}
		}

		generatedRootPath := filepath.Join(blunderRootPath, GeneratedDirName)
		clearGeneratedFolder(generatedRootPath)

		for _, detail := range blunderConfig.Details {
			errorFilePath := filepath.Join(generatedRootPath, detail.Package, ErrorFileName)
			templateFilePath, err := util.GetTemplatePath(ErrorFileTemplateFileName)
			if err != nil {
				return err
			}

			// generate id for this error
			for i := range detail.Errors {
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

func clearGeneratedFolder(generatedRootPath string) {
	// remove old generated folder
	_ = os.RemoveAll(generatedRootPath)
	// create generated folder
	_ = os.MkdirAll(generatedRootPath, os.ModePerm)
}

func checkAllErrorCodeIsProvided(b *model.Blunder) bool {
	for _, detail := range b.Details {
		for _, e := range detail.Errors {
			if e.Code == "" {
				return false
			}
		}
	}
	return true
}

func determineWhichToComplete(b *model.Blunder) (bool, map[string]model.Index, []string) {
	which := make(map[string]model.Index)
	whichErrorCodes := make([]string, 0)
	for detailIndex, detail := range b.Details {
		for eIndex, e := range detail.Errors {
			// grpc code = 0 means OK
			if e.HttpStatusCode == 0 || e.Message == "" {
				which[e.Code] = model.Index{
					DetailIndex: detailIndex,
					ErrorIndex:  eIndex,
				}
				whichErrorCodes = append(whichErrorCodes, e.Code)
			}
		}
	}

	if len(which) == 0 {
		return false, nil, nil
	} else {
		return true, which, whichErrorCodes
	}
}
