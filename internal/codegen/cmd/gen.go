package cmd

import (
	"github.com/DenChenn/blunder/internal/codegen/gpt"
	"github.com/DenChenn/blunder/internal/codegen/model"
	"github.com/DenChenn/blunder/internal/codegen/template"
	"github.com/DenChenn/blunder/internal/constant"
	"github.com/DenChenn/blunder/internal/util"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var Gen = &cli.Command{
	Name:  "gen",
	Usage: "generate all errors according to blunder.yaml",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "complete",
			Value:   false,
			Usage:   "auto complete error detail with gpt3",
			Aliases: []string{"c"},
		},
	},
	Action: func(cCtx *cli.Context) error {
		sGen := util.PrintLoading("generating all errors...", "all errors are generated successfully")

		blunderPath := util.LocateBlunderYamlPath()
		if blunderPath == "" {
			return util.PrintErrAndReturn("blunder.yaml not found, please init the project first")
		}
		blunderRootPath := util.GetFileDirPath(blunderPath)

		f, err := os.ReadFile(blunderPath)
		if err != nil {
			return util.PrintErrAndReturn(err.Error())
		}

		var blunderConfig model.Blunder
		if err := yaml.Unmarshal(f, &blunderConfig); err != nil {
			return util.PrintErrAndReturn(err.Error())
		}

		// check all error code is provided
		if !checkAllErrorCodeIsProvided(&blunderConfig) {
			return util.PrintErrAndReturn("in blunder.yaml, all error code must be provided")
		}

		// complete error detail with gpt3
		complete := cCtx.Bool("complete")
		if complete {
			// check if user provide gpt3 api token
			_, exist := os.LookupEnv("OPENAI_API_TOKEN")
			if !exist {
				return util.PrintErrAndReturn("OPENAI_API_TOKEN not found, please export with your own openai api token")
			}

			sComplete := util.PrintLoading("auto-completing error detail with gpt3...", "error detail is completed successfully")
			hasSomethingToComplete, indexMap, errorCodes := determineWhichToComplete(&blunderConfig)
			if hasSomethingToComplete {
				completed, err := gpt.CompleteErrorDetail(errorCodes)
				if err != nil {
					return util.PrintErrAndReturn(err.Error())
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
			if err := template.Generate(blunderPath, constant.GenerateBlunderYamlTemplateFileName, &blunderConfig); err != nil {
				return util.PrintErrAndReturn(err.Error())
			}
			sComplete.Stop()
		}

		generatedRootPath := filepath.Join(blunderRootPath, constant.GeneratedDirName)
		clearGeneratedFolder(generatedRootPath)

		for _, detail := range blunderConfig.Details {
			errorFilePath := filepath.Join(generatedRootPath, detail.Package, constant.ErrorFileName)

			// generate id for this error according to file path + code
			for i := range detail.Errors {
				id := util.GetId(errorFilePath + detail.Errors[i].Code)
				detail.Errors[i].Id = id
			}

			if err := template.Generate(errorFilePath, constant.ErrorFileTemplateFileName, &detail); err != nil {
				return util.PrintErrAndReturn(err.Error())
			}
		}

		sGen.Stop()
		return nil
	},
}

// clearGeneratedFolder remove old generated folder and create a new one
func clearGeneratedFolder(generatedRootPath string) {
	// remove old generated folder
	_ = os.RemoveAll(generatedRootPath)
	// create generated folder
	_ = os.MkdirAll(generatedRootPath, os.ModePerm)
}

// checkAllErrorCodeIsProvided check if all error code is provided
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

// determineWhichToComplete determine which error detail to complete
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
