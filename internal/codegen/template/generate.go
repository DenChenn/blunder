package template

import (
	"embed"
	"fmt"
	"github.com/DenChenn/blunder/internal/codegen/util"
	"os"
	"text/template"
)

//go:embed *.tmpl
var codegenTemplates embed.FS

func Generate(path string, templateName string, data any) error {
	t, err := template.ParseFS(codegenTemplates, templateName)
	if err != nil {
		return err
	}

	mkErr := os.MkdirAll(util.GetFileDirPath(path), 0o755)
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
