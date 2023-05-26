package util

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func GetRootPath() string {
	wd, _ := os.Getwd()
	return wd
}

const BlunderYamlFileName = "blunder.yaml"

func LocateBlunderYamlPath() string {
	root := GetRootPath()
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
