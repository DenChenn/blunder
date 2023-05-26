package util

import (
	"fmt"
	"os"
	"text/template"
)

func Generate(path string, templatePath string, data any) error {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	mkErr := os.MkdirAll(GetFileDirPath(path), 0o755)
	if mkErr != nil {
		fmt.Println(mkErr)
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if err := os.Chmod(path, 0o644); err != nil {
		return err
	}

	if err := t.Execute(f, data); err != nil {
		return err
	}
	return nil
}
