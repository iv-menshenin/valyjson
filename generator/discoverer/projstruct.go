package discoverer

import (
	"go/parser"
	"go/token"
	"io"
	"os"
	"path"

	"golang.org/x/mod/modfile"
)

type (
	Discoverer struct {
		modData  []byte
		modPath  string
		module   string
		packages map[string][]string
	}
)

func New(dir string) *Discoverer {
	var d = Discoverer{
		packages: make(map[string][]string),
	}
	err := d.discoveryModFile(dir)
	if err != nil {
		panic(err)
	}
	return &d
}

func (d *Discoverer) discoveryModFile(dir string) error {
	filePath := path.Join(dir, "go.mod")
	f, err := os.Open(filePath)
	if err == nil {
		defer f.Close()
		d.modData, err = io.ReadAll(f)
		d.modPath = dir
		return err
	}
	if dir == "/" || dir == "." || dir == "" {
		return err
	}
	if os.IsNotExist(err) {
		return d.discoveryModFile(path.Dir(dir))
	}
	return err
}

func (d *Discoverer) Discover() error {
	f, err := modfile.ParseLax("", d.modData, nil)
	if err != nil {
		return err
	}
	d.module = f.Module.Mod.Path
	return d.discoverPackages(d.modPath)
}

func (d *Discoverer) discoverPackages(dir string) error {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	var discovered bool
	for _, dirEntity := range dirs {
		entityPath := path.Join(dir, dirEntity.Name())
		if dirEntity.IsDir() {
			err = d.discoverPackages(entityPath)
			if err != nil {
				return err
			}
			continue
		}
		if discovered {
			continue
		}
		if path.Ext(entityPath) == ".go" {
			err = d.discoverPackage(entityPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Discoverer) discoverPackage(goFilePath string) error {
	f, err := parser.ParseFile(token.NewFileSet(), goFilePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return err
	}
	goFileDir := path.Dir(goFilePath)
	if packPaths, exists := d.packages[f.Name.Name]; exists {
		for _, packPath := range packPaths {
			if packPath == goFileDir {
				return nil
			}
		}
	}
	d.packages[f.Name.Name] = append(d.packages[f.Name.Name], goFileDir)
	return nil
}
