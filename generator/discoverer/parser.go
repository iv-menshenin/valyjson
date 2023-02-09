package discoverer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func (d *Discoverer) GetPackage(packageName string) (ast.Node, error) {
	packagePath, ok := d.packages[packageName]
	if !ok {
		return nil, fmt.Errorf("can't find package '%s'", packageName)
	}
	pkgs, err := parser.ParseDir(token.NewFileSet(), packagePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	return pkgs[packageName], nil
}
