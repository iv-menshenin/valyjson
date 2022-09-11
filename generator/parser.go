package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"

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
		//	"bytes"
		//	"fmt"
		//	"strconv"
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
						Path: &ast.BasicLit{Kind: token.IMPORT, Value: "\"strconv\""},
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
	g.result.Decls = append(g.result.Decls, valueIsNotNullDecl())
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

func parseGo(file string) (ast.Node, error) {
	f := token.NewFileSet()
	return parser.ParseFile(f, file, nil, parser.ParseComments|parser.AllErrors)
}

// func valueIsNotNull(v *fastjson.Value) bool {
//   return v != nil && v.Type() != fastjson.TypeNull
// }
func valueIsNotNullDecl() ast.Decl {
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// valueIsNotNull allows you to determine if the value is contained in a Json structure."},
			{Text: "// Checks if the structure itself or the value is Null."},
		}},
		Name: ast.NewIdent("valueIsNotNull"),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{
					Names: []*ast.Ident{ast.NewIdent("v")},
					Type:  &ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")}},
				},
			}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Type: ast.NewIdent("bool")},
			}},
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.BinaryExpr{
						X: &ast.BinaryExpr{
							X:  ast.NewIdent("v"),
							Op: token.NEQ,
							Y:  ast.NewIdent("nil"),
						},
						Op: token.LAND,
						Y: &ast.BinaryExpr{
							X: &ast.CallExpr{
								Fun: &ast.SelectorExpr{X: ast.NewIdent("v"), Sel: ast.NewIdent("Type")},
							},
							Op: token.NEQ,
							Y:  &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("TypeNull")},
						},
					},
				},
			},
		}},
	}
}

func New(file string) *Gen {
	return &Gen{
		fileName: file,
	}
}
