package generator

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

type (
	Gen struct {
		fileName string
		parsed   ast.Node
		result   ast.File
	}
)

func (g *Gen) Parse() (err error) {
	g.parsed, err = parseGo(g.fileName)
	if err == nil {
		g.result.Name = g.parsed.(*ast.File).Name
	}
	return
}

func (g *Gen) Print(name string) {
	if err := printer.Fprint(os.Stdout, token.NewFileSet(), &g.result); err != nil {
		panic(err)
	}
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
