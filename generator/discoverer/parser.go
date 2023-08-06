package discoverer

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func (d *Discoverer) GetPackage(packageName string, hints ...string) (ast.Node, error) {
	packagePaths, ok := d.packages[packageName]
	if !ok {
		return nil, fmt.Errorf("can't find package '%s'", packageName)
	}
	for _, packagePath := range packagePaths {
		pkgs, err := parser.ParseDir(token.NewFileSet(), packagePath, nil, parser.AllErrors)
		if err != nil {
			return nil, err
		}
		if len(packagePaths) == 1 || len(hints) == 0 || lookupNames(pkgs, hints...) {
			return pkgs[packageName], nil
		}
	}
	return nil, fmt.Errorf("can't find package '%s'", packageName)
}

func (d *Discoverer) GetPackagePath(packageName string, hints ...string) (string, error) {
	packagePaths, ok := d.packages[packageName]
	if !ok {
		return "", fmt.Errorf("can't find package '%s'", packageName)
	}
	for _, packagePath := range packagePaths {
		pkgs, err := parser.ParseDir(token.NewFileSet(), packagePath, nil, parser.AllErrors)
		if err != nil {
			return "", err
		}
		if len(packagePaths) == 1 || len(hints) == 0 || lookupNames(pkgs, hints...) {
			return packagePath, nil
		}
	}
	return "", fmt.Errorf("can't find package '%s'", packageName)
}

func (d *Discoverer) GetPackageFullName(packageName string, hints ...string) (string, error) {
	packagePaths, ok := d.packages[packageName]
	if !ok {
		return "", fmt.Errorf("can't find package '%s'", packageName)
	}
	for _, packagePath := range packagePaths {
		pkgs, err := parser.ParseDir(token.NewFileSet(), packagePath, nil, parser.AllErrors)
		if err != nil {
			return "", err
		}
		if len(packagePaths) == 1 || len(hints) == 0 || lookupNames(pkgs, hints...) {
			if strings.HasPrefix(packagePath, d.modPath) {
				packagePath = d.module + packagePath[len(d.modPath):]
			}
			return packagePath, nil
		}
	}
	return "", fmt.Errorf("can't find package '%s'", packageName)
}

func lookupNames(pkgs map[string]*ast.Package, hints ...string) bool {
	for name, pkg := range pkgs {
		if strings.HasSuffix(name, "_test") {
			continue
		}
		for _, ident := range hints {
			var v = lookupVisitor{
				found:  new(ast.Node),
				lookUp: ident,
			}
			ast.Walk(&v, pkg)
			if *v.found != nil {
				return true
			}
		}
	}
	return false
}

type lookupVisitor struct {
	found  *ast.Node
	lookUp string
}

func (v *lookupVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if v.found == nil {
		return v
	}

	switch currNode := node.(type) {

	case *ast.GenDecl:
		for _, declSpec := range currNode.Specs {
			switch spec := declSpec.(type) {
			case *ast.TypeSpec:
				if spec.Name.Name == v.lookUp {
					*v.found = node
				}

			case *ast.ValueSpec:
				for _, name := range spec.Names {
					if name.Name == v.lookUp {
						*v.found = node
					}
				}

			default:
				v1 := *v
				v1.found = nil
				return &v1
			}
		}

	case *ast.TypeSpec:
		if currNode.Name.Name == v.lookUp {
			*v.found = node
		}

	case *ast.ValueSpec:
		for _, name := range currNode.Names {
			if name.Name == v.lookUp {
				*v.found = node
			}
		}
	}
	return v
}

func (d *Discoverer) GetModuleName() string {
	return d.module
}
