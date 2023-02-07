package discoverer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"strings"
)

func (d *Discoverer) GetPackage(packageName string) (ast.Node, error) {
	packagePath, ok := d.packages[packageName]
	if !ok {
		return nil, fmt.Errorf("can't find package '%s'", packageName)
	}
	pkgs, err := parser.ParseDir(token.NewFileSet(), packagePath, func(info fs.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return pkgs[packageName], nil
}
