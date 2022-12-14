package field

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

// if list := v.Get("list"); list != nil {
// 	var listA []*fastjson.Value
// 	listA, err = list.Array()
// 	if err != nil {
// 		return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
// 	}
// 	s.List = make([]int32, 0, len(listA))
// 	for listElemNum, listElem := range listA {
// 		var listElemVal int64
// 		listElemVal, err = listElem.Int64()
// 		if err != nil {
// 			return fmt.Errorf("error parsing '%slist[%d]' value: %w", objPath, listElemNum, err)
// 		}
// 		s.List = append(s.List, int32(listElemVal))
// 	}
// }
func arrayExtraction(dst ast.Expr, v, json string, t ast.Expr, stmtExtr []ast.Stmt) []ast.Stmt {
	var body = []ast.Stmt{
		// var listA []*fastjson.Value
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.TypeSpec{
						Name: ast.NewIdent("listA"),
						Type: &ast.ArrayType{Elt: &ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")}}},
					},
				},
			},
		},
		// listA, err = list.Array()
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("listA"), ast.NewIdent("err")},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Array")},
			}},
		},
		// 	if err != nil {
		// 		return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
		// 	}
		checkErrAnd(json),
		// valList := make([]int32, 0, len(listA))
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: ast.NewIdent("make"),
					Args: []ast.Expr{
						&ast.ArrayType{Elt: t},
						&ast.BasicLit{Kind: token.INT, Value: "0"},
						&ast.CallExpr{Fun: ast.NewIdent("len"), Args: []ast.Expr{ast.NewIdent("listA")}},
					},
				},
			},
		},
		// for _, listElem := range listA {
		&ast.RangeStmt{
			Key:   ast.NewIdent("_"),
			Value: ast.NewIdent("listElem"),
			Tok:   token.DEFINE,
			X:     ast.NewIdent("listA"),
			Body:  &ast.BlockStmt{List: stmtExtr},
		},
	}
	return body
}

// checkErrAnd generates a decoding error check
// 	if err != nil {
// 	    return fmt.Errorf("error parsing '%sname' value: %w", objPath, err)
// 	}
func checkErrAnd(jsonFieldName string) ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.BinaryExpr{X: ast.NewIdent("err"), Op: token.NEQ, Y: ast.NewIdent("nil")},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent("fmt"), Sel: ast.NewIdent("Errorf")},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: `"error parsing '%s` + jsonFieldName + `' value: %w"`},
							ast.NewIdent("objPath"),
							ast.NewIdent("err"),
						},
					},
				},
			},
		}},
	}
}

//	if {v}.Type() != fastjson.TypeString {
//		err = fmt.Errorf("value doesn't contain string; it contains %s", {v}.Type())
//		return fmt.Errorf("error parsing '%s{json}' value: %w", objPath, err)
//	}
//	{dst} = {v}.String()
func stringExtraction(dst ast.Expr, v, json string) []ast.Stmt {
	var result []ast.Stmt
	var valueType = &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Type")},
	}
	result = append(result, &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  valueType,
			Op: token.NEQ,
			Y:  &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("TypeString")},
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent(names.VarNameError)},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{helpers.FmtError("value doesn't contain string; it contains %s", valueType)},
			},
			&ast.ReturnStmt{Results: []ast.Expr{
				helpers.FmtError("error parsing '%s"+json+"' value: %w", ast.NewIdent(names.VarNameObjPath), ast.NewIdent(names.VarNameError)),
			}},
		}},
	})
	result = append(result, &ast.AssignStmt{
		Lhs: []ast.Expr{dst},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("String")},
		}},
	})
	return result
}

// var {dst} int
// {dst}, err = {v}.Int()
func intExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  ast.NewIdent("int"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Int")},
			}},
		},
	}
}

// var {dst} uint
// {dst}, err = {v}.Uint()
func uintExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  ast.NewIdent("uint"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Uint")},
			}},
		},
	}
}

// var {dst} int64
// {dst}, err = {v}.Int64()
func int64Extraction(dst *ast.Ident, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  ast.NewIdent("int64"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Int64")},
			}},
		},
	}
}

// var {dst} uint64
// {dst}, err = {v}.Uint64()
func uint64Extraction(dst *ast.Ident, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  ast.NewIdent("uint64"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Uint64")},
			}},
		},
	}
}

// var {dst} float64
// {dst}, err = {v}.Float64()
func floatExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  ast.NewIdent("float64"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Float64")},
			}},
		},
	}
}

// var {dst} bool
// {dst}, err = {v}.Bool()
func boolExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  ast.NewIdent("bool"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Bool")},
			}},
		},
	}
}

// err = {dst}.fill({v}, objPath+"{json}.")
func nestedExtraction(dst *ast.Ident, t ast.Expr, v, json string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  t,
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(names.VarNameError)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   dst,
						Sel: ast.NewIdent(names.FuncNameFill),
					},
					Args: []ast.Expr{
						ast.NewIdent(v),
						&ast.BinaryExpr{
							X:  ast.NewIdent(names.VarNameObjPath),
							Op: token.ADD,
							Y: &ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"" + strings.Replace(json, "\"", "\\\"", -1) + ".\"",
							},
						},
					},
				},
			},
		},
	}
}

// uuidExtraction generates the code of the standard conversion process from string to UUID format
//  var valfield uuid.UUID
//  b, err := field.StringBytes()
//  if err != nil {
//      return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
//  }
//  err = valfield.UnmarshalText(b)
func uuidExtraction(dst *ast.Ident, t ast.Expr, v, name string) []ast.Stmt {
	var stmt = []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  t,
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("b"), ast.NewIdent("err")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("StringBytes")},
				},
			},
		},
		checkErrAnd(name),
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("err")},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun:  &ast.SelectorExpr{X: dst, Sel: ast.NewIdent("UnmarshalText")},
					Args: []ast.Expr{ast.NewIdent("b")},
				},
			},
		},
	}
	return stmt
}

// {dst}, err := time.Parse({layout}, {v}.String())
func timeExtraction(dst *ast.Ident, v, layout string) []ast.Stmt {
	const (
		defLayout = "time.RFC3339"
	)
	if layout == "" {
		layout = defLayout
	}
	var layoutExpr ast.Expr
	if l := strings.Split(layout, "."); len(l) == 2 && l[0] == "time" {
		layoutExpr = &ast.SelectorExpr{
			X:   ast.NewIdent(l[0]),
			Sel: ast.NewIdent(l[1]),
		}
	} else {
		layoutExpr = &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%s"`, strings.Replace(layout, "\"", "\\\"", -1))}
	}
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent("err")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X:   ast.NewIdent("time"),
						Sel: ast.NewIdent("Parse"),
					},
					Args: []ast.Expr{
						layoutExpr,
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								Sel: ast.NewIdent("String"),
								X:   ast.NewIdent(v),
							},
						},
					},
				},
			},
		},
	}
}
