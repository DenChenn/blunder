package util

import (
	"fmt"
	"github.com/DenChenn/blunder/internal/constant"
	"io/fs"
	"os"
	"path/filepath"
)

func GetCWD() string {
	wd, _ := os.Getwd()
	return wd
}

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

func GetFileDirPath(path string) string {
	return filepath.Dir(path)
}
