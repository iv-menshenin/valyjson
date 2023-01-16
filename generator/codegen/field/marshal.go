package field

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
)

// result.WriteString("\"{json}\":")
// b = strconv.AppendUint(buf[:0], uint64({src}), 10)
// result.Write(b)
func (f *fld) typeMarshal(src ast.Expr, v, t string) []ast.Stmt {
	var result = []ast.Stmt{
		&ast.IfStmt{
			// if result.Len() > 1 {
			Cond: &ast.BinaryExpr{
				X:  &ast.CallExpr{Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("Len")}},
				Op: token.GTR,
				Y:  &ast.BasicLit{Kind: token.INT, Value: "1"},
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				// result.WriteRune(',')
				&ast.ExprStmt{X: &ast.CallExpr{
					Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteRune")},
					Args: []ast.Expr{&ast.BasicLit{Kind: token.CHAR, Value: `','`}},
				}},
			}},
		},
		// result.WriteString("\"field\":")
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":"`, f.tags.JsonName())}},
		}},
	}
	switch t {

	case "int", "int8", "int16", "int32":
		result = append(result, intMarshal(src)...)

	case "int64":
		result = append(result, int64Marshal(src)...)

	case "uint", "uint8", "uint16", "uint32":
		result = append(result, uintMarshal(src)...)

	case "uint64":
		result = append(result, uint64Marshal(src)...)

	case "float32", "float64":
		result = append(result, floatMarshal(src)...)

	case "bool":
		result = append(result, boolMarshal(src)...)

	case "string":
		result = append(result, stringMarshal(src)...)

	default:
		result = append(result, nestedMarshal(src)...)
	}
	// result.Write(b)
	result = append(
		result,
		&ast.ExprStmt{
			X: &ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("Write")},
				Args: []ast.Expr{ast.NewIdent("b")},
			},
		},
	)
	return result
}

// result.WriteString("\"field\":\"\"")
func (f *fld) typeMarshalDefault(src ast.Expr, v, t string) []ast.Stmt {
	var result []ast.Stmt
	var args []ast.Expr
	switch t {

	case "int", "int8", "int16", "int32", "int64":
		args = append(args, &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":0"`, v)})

	case "uint", "uint8", "uint16", "uint32", "uint64":
		args = append(args, &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":0"`, v)})

	case "float32", "float64":
		args = append(args, &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":0.0"`, v)})

	case "bool":
		args = append(args, &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":false"`, v)})

	case "string":
		args = append(args, &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":\"\""`, v)})

	default:
		panic("not implemented")
	}
	result = append(result, &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
			Args: args,
		},
	})
	return result
}

// b, err = {src}.MarshalAppend(buf[:0])
// if err != nil {
// 	return nil, err
// }
func nestedMarshal(src ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("b"), ast.NewIdent("err")},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: src, Sel: ast.NewIdent("MarshalAppend")},
					Args: []ast.Expr{&ast.SliceExpr{
						X:    ast.NewIdent("buf"),
						High: &ast.BasicLit{Kind: token.INT, Value: "0"},
					}},
				},
			},
		},
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{X: ast.NewIdent("err"), Op: token.NEQ, Y: ast.NewIdent("nil")},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent("nil"), ast.NewIdent("err")}},
			}},
		},
	}
}

// if s.HeightRef != nil {
//     {v} := *{src}
//     result.WriteString("\"{json}\":")
//     b = strconv.AppendUint(buf[:0], uint64({v}), 10)
//     result.Write(b)
// } else {
//     result.WriteString("\"{json}\":{default}")
// }
func (f *fld) typeRefMarshal(src ast.Expr, v, t string) []ast.Stmt {
	var els ast.Stmt
	if stmt := f.ifNil(); len(stmt) > 0 {
		els = &ast.BlockStmt{List: stmt}
	}
	var result = []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(v)},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.StarExpr{X: src}},
		},
	}
	result = append(
		result,
		f.typeMarshal(ast.NewIdent(v), v, t)...,
	)

	return []ast.Stmt{
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  src,
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: result},
			Else: els,
		},
	}
}

// result.WriteString("\"{name}\":{default}")
func (f *fld) ifNil() []ast.Stmt {
	if f.tags.DefaultValue() == "" {
		if f.tags.JsonTags().Has("omitempty") {
			return nil
		}
		// result.WriteString("\"{name}\":null")
		return []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
					Args: []ast.Expr{
						&ast.BasicLit{Kind: token.STRING, Value: `"\"` + f.tags.JsonName() + `\":null"`},
					},
				},
			},
		}
	}
	return []ast.Stmt{
		&ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: `"\"` + f.tags.JsonName() + `\":` + helpers.StringFromType(f.expr, f.tags.DefaultValue()) + `"`},
				},
			},
		},
	}
}
