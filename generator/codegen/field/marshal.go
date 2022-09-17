package field

import (
	"go/ast"
	"go/token"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
)

// b = strconv.AppendUint(buf[:0], uint64(s.Height), 10)
// if _, err = result.Write(b); err != nil {
// 	return nil, err
// }
func (f fld) typeMarshal(src ast.Expr, v, t string) []ast.Stmt {
	switch t {

	case "int", "int8", "int16", "int32":
		return intMarshal(src)

	case "int64":
		return int64Marshal(src)

	case "uint", "uint8", "uint16", "uint32":
		return uintMarshal(src)

	case "uint64":
		return uint64Marshal(src)

	case "float32", "float64":
		return floatMarshal(src)

	case "bool":
		return boolMarshal(src)

	case "string":
		return stringMarshal(src)

	default:
		// todo @menshenin return nestedMarshal(name, v, f.t.JsonName())
		return nil

	}
}

// if s.HeightRef != nil {
// 	result.WriteString("\"heightRef\":")
// 	b = strconv.AppendUint(buf[:0], uint64(*s.HeightRef), 10)
// 	if _, err = result.Write(b); err != nil {
// 		return nil, err
// 	}
// } else {
// 	result.WriteString("\"heightRef\":null")
// }
func (f fld) typeRefMarshal(src ast.Expr, v, t string) []ast.Stmt {
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
func (f fld) ifNil() []ast.Stmt {
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
