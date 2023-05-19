package discoverer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
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

func (d *Discoverer) GetPackagePath(packageName string) (string, error) {
	packagePath, ok := d.packages[packageName]
	if !ok {
		return "", fmt.Errorf("can't find package '%s'", packageName)
	}
	return packagePath, nil
}

func (d *Discoverer) GetPackageFullName(packageName string) (string, error) {
	packagePath, ok := d.packages[packageName]
	if !ok {
		return "", fmt.Errorf("can't find package '%s'", packageName)
	}
	if strings.HasPrefix(packagePath, d.modPath) {
		packagePath = d.module + packagePath[len(d.modPath):]
	}
	return packagePath, nil
}

func (d *Discoverer) GetModuleName() string {
	return d.module
}
