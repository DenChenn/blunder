package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func GetRootPath() string {
	wd, _ := os.Getwd()
	return wd
}

const BlunderYamlFileName = "blunder.yaml"

func LocateBlunderYamlPath() string {
	root := GetRootPath()
	fmt.Println(root)
	location := ""
	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == BlunderYamlFileName {
			location = path
			return nil
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	return location
}

func GetFileDirPath(path string) string {
	return filepath.Dir(path)
}

func GetTemplatePath(templateFileName string) (string, error) {
	goPath, isExist := os.LookupEnv("GOPATH")
	if !isExist {
		return "", errors.New("GOPATH not found")
	}

	// find package file path
	walkPath := filepath.Join(goPath, "pkg/mod/github.com/!den!chenn")
	pathPattern := filepath.Join(goPath, "pkg/mod/github.com/!den!chenn/blunder@(.*?)/internal/codegen/template")

	re := regexp.MustCompile(pathPattern)

	templateRootPath := ""
	if err := filepath.Walk(walkPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		isMatched := re.MatchString(path)
		if isMatched && info.IsDir() {
			templateRootPath = path
			return nil
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	templatePath := filepath.Join(templateRootPath, templateFileName)
	return templatePath, nil
}
