package impcovery

import (
	"fmt"
	"go/ast"
	"go/token"
	"sort"
)

type (
	Discoverer struct {
		imports map[string]Package
	}
	Package struct {
		Path string
		Kind PkgKind
	}
	PkgKind int8
)

const (
	PkgKindSystem PkgKind = iota
	PkgKindExternal
	PkgKindInternal
)

var (
	knownPackages = map[string]Package{
		"bytes":    {Path: "bytes", Kind: PkgKindSystem},
		"context":  {Path: "context", Kind: PkgKindSystem},
		"errors":   {Path: "errors", Kind: PkgKindSystem},
		"http":     {Path: "net/http", Kind: PkgKindSystem},
		"json":     {Path: "encoding/json", Kind: PkgKindSystem},
		"reflect":  {Path: "reflect", Kind: PkgKindSystem},
		"time":     {Path: "time", Kind: PkgKindSystem},
		"fmt":      {Path: "fmt", Kind: PkgKindSystem},
		"strconv":  {Path: "strconv", Kind: PkgKindSystem},
		"net":      {Path: "net", Kind: PkgKindSystem},
		"math":     {Path: "math", Kind: PkgKindSystem},
		"url":      {Path: "net/url", Kind: PkgKindSystem},
		"fasthttp": {Path: "github.com/valyala/fasthttp", Kind: PkgKindExternal},
		"fastjson": {Path: "github.com/valyala/fastjson", Kind: PkgKindExternal},
		"router":   {Path: "github.com/fasthttp/router", Kind: PkgKindExternal},
		"uuid":     {Path: "github.com/google/uuid", Kind: PkgKindExternal},
	}
)

func RegisterPackage(packName string, pkg Package) {
	knownPackages[packName] = pkg
}

func New() *Discoverer {
	return &Discoverer{
		imports: make(map[string]Package),
	}
}

func (i *Discoverer) Explore(node ast.Node) {
	ast.Walk(i, node)
}

func (i *Discoverer) Visit(node ast.Node) (w ast.Visitor) {
	sel, ok := node.(*ast.SelectorExpr)
	if !ok {
		return i
	}
	x, ok := sel.X.(*ast.Ident)
	if !ok {
		return i
	}
	pack, ok := knownPackages[x.String()]
	if ok {
		i.imports[pack.Path] = pack
	}
	return i
}

func (i *Discoverer) Spec() []ast.Spec {
	var imports []Package
	for _, pkg := range i.imports {
		imports = append(imports, pkg)
	}
	sort.Slice(imports, func(i, j int) bool {
		if imports[i].Kind < imports[j].Kind {
			return true
		}
		return imports[i].Path < imports[j].Path
	})

	var (
		currT PkgKind = -1
		specs []ast.Spec
	)
	for _, imp := range imports {
		var addLine string
		if currT != imp.Kind {
			currT = imp.Kind
			addLine = "\n\t"
		}
		specs = append(specs, &ast.ImportSpec{Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf("%s\"%s\"", addLine, imp.Path),
		}})
	}
	return specs
}
