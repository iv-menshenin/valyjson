package generator

import (
	"fmt"
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
		// "bytes"
		//	"fmt"
		//	"github.com/valyala/fastjson"
		g.result.Decls = []ast.Decl{
			&ast.GenDecl{
				Tok: token.IMPORT,
				Specs: []ast.Spec{
					&ast.ImportSpec{
						Path: &ast.BasicLit{Kind: token.IMPORT, Value: "\"bytes\""},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{Kind: token.IMPORT, Value: "\"fmt\""},
					},
					&ast.ImportSpec{
						Path: &ast.BasicLit{Kind: token.IMPORT, Value: "\"github.com/valyala/fastjson\""},
					},
				},
			},
		}
	}
	return
}

func (g *Gen) Print(name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = fmt.Fprint(f, "// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.\n")
	if err != nil {
		panic(err)
	}
	if err := printer.Fprint(f, token.NewFileSet(), &g.result); err != nil {
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
