package field

import (
	"go/ast"
	"go/token"
)

// b := marshalString(s.Field, buf[:0])
func stringMarshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.DEFINE,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: ast.NewIdent("marshalString"),
					Args: []ast.Expr{
						src,
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
					},
				},
			},
		},
	}
}

// b = strconv.AppendInt(buf[:0], int64({src}), 10)
func intMarshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendInt")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						&ast.CallExpr{
							Fun:  ast.NewIdent("int64"),
							Args: []ast.Expr{src},
						},
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendInt(buf[:0], {src}, 10)
func int64Marshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendInt")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						src,
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendUint(buf[:0], uint64({src}), 10)
func uintMarshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendUint")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						&ast.CallExpr{
							Fun:  ast.NewIdent("uint64"),
							Args: []ast.Expr{src},
						},
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendUint(buf[:0], {src}, 10)
func uint64Marshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendUint")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						src,
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendFloat(buf[:0], float64({src}), 'f', 10, 64)
func floatMarshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendFloat")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						&ast.CallExpr{
							Fun:  ast.NewIdent("float64"),
							Args: []ast.Expr{src},
						},
						&ast.BasicLit{Kind: token.CHAR, Value: "'f'"},
						&ast.BasicLit{Kind: token.INT, Value: "-1"}, // todo @menshenin pass precision through structTags
						&ast.BasicLit{Kind: token.INT, Value: "64"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendBool(buf[:0], {src})
func boolMarshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendBool")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						src,
					},
				},
			},
		},
	}
}
