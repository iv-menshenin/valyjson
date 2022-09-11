package field

import (
	"go/ast"
	"go/token"

	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

func (f fld) typeMarshal(name, v, t string) []ast.Stmt {
	switch t {

	case "int", "int8", "int16", "int32":
		return intMarshal(name, v)

	case "int64":
		return int64Marshal(name, v)

	case "uint", "uint8", "uint16", "uint32":
		return uintMarshal(name, v)

	case "uint64":
		return uint64Marshal(name, v)

	case "float32", "float64":
		return floatMarshal(name, v)

	case "bool":
		return boolMarshal(name, v)

	case "string":
		return stringMarshal(name, v, f.t.JsonName())

	default:
		// todo @menshenin return nestedMarshal(name, v, f.t.JsonName())
		return nil

	}
}

// b, err = marshalString(s.Field, buf[:0])
func stringMarshal(name, v, json string) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b"), ast.NewIdent("err")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: ast.NewIdent("marshalString"),
					Args: []ast.Expr{
						&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
					},
				},
			},
		},
	}
}

// b = strconv.AppendInt(buf[:0], int64(s.Height), 10)
func intMarshal(name, v string) []ast.Stmt {
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
							Fun: ast.NewIdent("int64"),
							Args: []ast.Expr{
								&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
							},
						},
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendInt(buf[:0], s.Height, 10)
func int64Marshal(name, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendInt")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendUint(buf[:0], uint64(s.Height), 10)
func uintMarshal(name, v string) []ast.Stmt {
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
							Fun: ast.NewIdent("uint64"),
							Args: []ast.Expr{
								&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
							},
						},
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendUint(buf[:0], s.Height, 10)
func uint64Marshal(name, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendUint")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
						&ast.BasicLit{Kind: token.INT, Value: "10"},
					},
				},
			},
		},
	}
}

// b = strconv.AppendFloat(buf[:0], float64(s.Height), 'f', 10, 64)
func floatMarshal(name, v string) []ast.Stmt {
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
							Fun: ast.NewIdent("float64"),
							Args: []ast.Expr{
								&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
							},
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

// b = strconv.AppendBool(buf[:0], s.Admin)
func boolMarshal(name, v string) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Tok: token.ASSIGN,
			Lhs: []ast.Expr{ast.NewIdent("b")},
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("strconv"), Sel: ast.NewIdent("AppendBool")},
					Args: []ast.Expr{
						&ast.SliceExpr{X: ast.NewIdent("buf"), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
					},
				},
			},
		},
	}
}
