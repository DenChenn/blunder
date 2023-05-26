package util

import (
	"fmt"
	"github.com/DenChenn/blunder/internal/constant"
	"io/fs"
	"os"
	"path/filepath"
)

// GetCWD returns the current working directory
func GetCWD() string {
	wd, _ := os.Getwd()
	return wd
}

// LocateBlunderYamlPath returns the path of blunder.yaml
func LocateBlunderYamlPath() string {
	root := GetCWD()
	location := ""
	if err := filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == constant.BlunderYamlFileName {
			location = path
			return nil
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}

	return location
}

// GetFileDirPath returns the directory of the file
func GetFileDirPath(path string) string {
	return filepath.Dir(path)
}
