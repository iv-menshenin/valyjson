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
		// result.WriteString("\"field\":")
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":"`, f.t.JsonName())}},
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
		// todo @menshenin return nestedMarshal(name, v, f.t.JsonName())
		return nil
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
	if f.t.DefaultValue() == "" {
		if f.t.JsonTags().Has("omitempty") {
			return nil
		}
		// result.WriteString("\"{name}\":null")
		return []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
					Args: []ast.Expr{
						&ast.BasicLit{Kind: token.STRING, Value: `"\"` + f.t.JsonName() + `\":null"`},
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
					&ast.BasicLit{Kind: token.STRING, Value: `"\"` + f.t.JsonName() + `\":` + helpers.StringFromType(f.f.Type, f.t.DefaultValue()) + `"`},
				},
			},
		},
	}
}
