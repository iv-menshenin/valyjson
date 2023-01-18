package generator

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"

	"github.com/iv-menshenin/valyjson/generator/codegen/impcovery"
	"github.com/iv-menshenin/valyjson/generator/static"
)

type (
	Gen struct {
		fileName string
		parsed   ast.Node
		result   ast.File
		packageN string
	}
)

func (g *Gen) Parse() (err error) {
	g.parsed, err = parseGo(g.fileName)
	g.packageN = g.parsed.(*ast.File).Name.Name
	if err == nil {
		g.result.Name = g.parsed.(*ast.File).Name
	}
	return
}

func (g *Gen) FixImports() {
	// discovery used imports and build their declaration
	discovery := impcovery.New()
	discovery.Explore(&g.result)
	var decls = make([]ast.Decl, 0, len(g.result.Decls)+1)

	decls = append(decls, &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: discovery.Spec(),
	})
	decls = append(decls, g.result.Decls...)
	g.result.Decls = decls
}

func (g *Gen) Print(name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fmtGOFile(name)
	defer f.Close()
	_, err = fmt.Fprint(f, "// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.\n")
	if err != nil {
		panic(err)
	}

	if err := printer.Fprint(f, token.NewFileSet(), &g.result); err != nil {
		panic(err)
	}
	err = static.Print(static.Context{
		Package: g.packageN,
	}, path.Dir(name))
	if err != nil {
		panic(err)
	}
}

func fmtGOFile(fileName string) (err error) {
	var fileSet = token.NewFileSet()
	f, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if e := out.Close(); e != nil && err == nil {
			err = e
		}
	}()
	return format.Node(out, fileSet, f)
}

func parseGo(file string) (ast.Node, error) {
	f := token.NewFileSet()
	return parser.ParseFile(f, file, nil, parser.ParseComments|parser.AllErrors)
}

func New(file string) *Gen {
	return &Gen{
		fileName: file,
	}
}
