package static

import (
	"embed"
	"os"
	"path"
	"strings"
	"text/template"
)

//go:embed tpl/*.go.tpl
var tpls embed.FS

type Context struct {
	Package string
}

func Print(ctx Context, dir string) error {
	d, err := tpls.ReadDir(".")
	if err != nil {
		return err
	}
	for _, ent := range d {
		if !ent.IsDir() {
			continue
		}
		if err = recursivePrint(ctx, dir, ent.Name()); err != nil {
			return err
		}
	}
	return nil
}

func recursivePrint(ctx Context, basePath, dirPath string) error {
	d, err := tpls.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, ent := range d {
		fullName := path.Join(dirPath, ent.Name())
		basePathEx := path.Join(basePath, ent.Name())
		if ent.IsDir() {
			if err = recursivePrint(ctx, basePathEx, fullName); err != nil {
				return err
			}
		}
		fName := strings.TrimSuffix(ent.Name(), ".tpl")
		tpl, err := newTemplate(fullName, fName)
		if err != nil {
			return err
		}
		if err = tpl.Execute(ctx); err != nil {
			return err
		}
		if err = tpl.Print(path.Join(basePath, fName)); err != nil {
			return err
		}
	}
	return nil
}

type TPL struct {
	str strings.Builder
	tpl *template.Template
}

func (t *TPL) Execute(ctx Context) error {
	return t.tpl.Execute(&t.str, ctx)
}

func (t *TPL) Print(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(t.str.String())
	return err
}

func newTemplate(tplPath, fName string) (*TPL, error) {
	tpl, err := template.New(fName).ParseFS(tpls, tplPath)
	if err != nil {
		return nil, err
	}
	return &TPL{tpl: tpl.Templates()[0]}, nil
}
