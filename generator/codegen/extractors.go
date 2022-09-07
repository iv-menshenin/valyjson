package codegen

import (
	"go/ast"
	"go/token"
	"strings"
)

// 	if list := v.Get("list"); list != nil {
//		var listA []*fastjson.Value
//		listA, err = list.Array()
//		if err != nil {
//			return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
//		}
//		s.List = make([]int64, 0, len(listA))
//		for listElemNum, listElem := range listA {
//			var listElemVal int64
//			listElemVal, err = listElem.Int64()
//			if err != nil {
//				return fmt.Errorf("error parsing '%slist[%d]' value: %w", objPath, listElemNum, err)
//			}
//			s.List = append(s.List, listElemVal)
//		}
//	}
func arrayExtraction(name, v, json string) []ast.Stmt {
	return nil
}

//	if {v}.Type() != fastjson.TypeString {
//		err = fmt.Errorf("value doesn't contain string; it contains %s", {v}.Type())
//		return fmt.Errorf("error parsing '%s{json}' value: %w", objPath, err)
//	}
//	s.{name} = {v}.String()
func stringExtraction(name, v, json string) []ast.Stmt {
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
				Lhs: []ast.Expr{ast.NewIdent(errVarName)},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{fmtError("value doesn't contain string; it contains %s", valueType)},
			},
			&ast.ReturnStmt{Results: []ast.Expr{
				fmtError("error parsing '%s"+json+"' value: %w", ast.NewIdent(objPathVarName), ast.NewIdent(errVarName)),
			}},
		}},
	})
	result = append(result, &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(name)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("String")},
		}},
	})
	return result
}

// x{name}, err = {v}.Int()
func intExtraction(name, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent(name)},
						Type:  ast.NewIdent("int"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(name), ast.NewIdent(errVarName)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Int")},
			}},
		},
	}
}

// s.{name}, err = {v}.Int64()
func int64Extraction(name, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent(name)},
						Type:  ast.NewIdent("int64"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(name), ast.NewIdent(errVarName)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Int64")},
			}},
		},
	}
}

// s.{name}, err = {v}.Bool()
func boolExtraction(name, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent(name)},
						Type:  ast.NewIdent("bool"),
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(name), ast.NewIdent(errVarName)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Bool")},
			}},
		},
	}
}

// err = s.{name}.fill({v}, objPath+"{json}.")
func nestedExtraction(name, v, json string) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(errVarName)},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{
						X: &ast.SelectorExpr{
							X:   ast.NewIdent(recvVarName),
							Sel: ast.NewIdent(name),
						},
						Sel: ast.NewIdent(fillerFuncName),
					},
					Args: []ast.Expr{
						ast.NewIdent(v),
						&ast.BinaryExpr{
							X:  ast.NewIdent(objPathVarName),
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
