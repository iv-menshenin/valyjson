package field

import (
	"go/ast"
	"go/token"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

// offset := v.Get("offset")
func (f fld) extract(v string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(v)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{f.getValue()},
	}
}

// v.Get("offset")
func (f fld) getValue() ast.Expr {
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent(names.VarNameJsonValue),
			Sel: ast.NewIdent("Get"),
		},
		Args: []ast.Expr{&ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"" + f.t.JsonName() + "\"",
		}},
	}
}

//	s.Offset, err = offset.Int()
//	if err != nil {
//		return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//	}
func (f fld) fillFrom(name, v string) []ast.Stmt {
	var dstType = f.f.Type
	star, isStar := dstType.(*ast.StarExpr)
	if isStar {
		dstType = star.X
		f.f.Type = dstType
	}
	var bufVariable = ast.NewIdent("val" + name)
	var result []ast.Stmt
	result = append(result, f.typedValue(bufVariable, v)...)
	result = append(result, f.checkErr()...)

	var fldExpr = &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}
	result = append(result, f.fillField(bufVariable, fldExpr, v)...)
	return result
}

//  var val{name} {type}
//	val{name}, err = {v}.(Int|Int64|String|Bool)()
func (f fld) typedValue(dst *ast.Ident, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.f.Type.(type) {

	case *ast.Ident:
		result = append(result, f.typeExtraction(dst, v, t.Name)...)

	case *ast.StructType:
		panic("unsupported field type 'struct'")

	case *ast.SelectorExpr:
		result = append(result, nestedExtraction(dst, f.f.Type, v, f.t.JsonName())...)

	case *ast.ArrayType:
		result = append(result, arrayExtraction(dst, v, f.t.JsonName())...)
		return result

	default:
		panic("unsupported field type")
	}
	return result
}

func (f fld) typeExtraction(dst *ast.Ident, v, t string) []ast.Stmt {
	switch t {

	case "int", "int8", "int16", "int32":
		return intExtraction(dst, v)

	case "int64":
		return int64Extraction(dst, v)

	case "uint", "uint8", "uint16", "uint32":
		return uintExtraction(dst, v)

	case "uint64":
		return uint64Extraction(dst, v)

	case "float32", "float64":
		return floatExtraction(dst, v)

	case "bool":
		return boolExtraction(dst, v)

	case "string":
		return stringExtraction(dst, v, f.t.JsonName())

	default:
		return nestedExtraction(dst, f.f.Type, v, f.t.JsonName())

	}
}

//	if err != nil {
//		return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//	}
func (f fld) checkErr() []ast.Stmt {
	if t, ok := f.f.Type.(*ast.Ident); ok && t.Name == "string" {
		// no error checking for string
		return nil
	}
	format := "error parsing '%s" + f.t.JsonName() + "' value: %w"
	return []ast.Stmt{
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(names.VarNameError),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{helpers.FmtError(format, ast.NewIdent(names.VarNameObjPath), ast.NewIdent(names.VarNameError))},
				},
			}},
		},
	}
}

func (f fld) fillField(rhs, dst ast.Expr, t string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.f.Type.(type) {

	case *ast.Ident:
		return f.typedFillIn(rhs, dst, t.Name)

	case *ast.StructType:
		return result

	case *ast.SelectorExpr:
		return result

	case *ast.ArrayType:
		return result

	case *ast.StarExpr:
		var tName = "nested"
		if ident, ok := t.X.(*ast.Ident); ok {
			tName = ident.Name
		}
		return f.typedRefFillIn(rhs, dst, tName)

	}
	return nil
}

func (f fld) typedFillIn(rhs, dst ast.Expr, t string) []ast.Stmt {
	switch t {
	case "string", "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{rhs},
			},
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent(t),
					Args: []ast.Expr{rhs},
				}},
			},
		}

	default:
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{rhs},
			},
		}
	}
}

func (f fld) typedRefFillIn(rhs, dst ast.Expr, t string) []ast.Stmt {
	switch t {
	case "string", "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.UnaryExpr{X: rhs, Op: token.AND}},
			},
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		var result []ast.Stmt
		result = append(
			result,
			// s.HeightRef = new(uint32)
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent("new"),
					Args: []ast.Expr{ast.NewIdent(t)},
				}},
			},
			// *s.HeightRef = uint32(xHeightRef)
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.StarExpr{X: dst}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent(t),
					Args: []ast.Expr{rhs},
				}},
			},
		)
		return result

	default:
		return nil
	}
}

//	} else {
//		s.{name} = 100
//	}
func (f fld) ifDefault(name string) []ast.Stmt {
	if f.t.DefaultValue() == "" {
		if f.t.JsonTags().Has("required") {
			// return fmt.Errorf("required element '%s{json}' is missing", objPath)
			return []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent("fmt"), Sel: ast.NewIdent("Errorf")},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"required element '%s" + f.t.JsonName() + "' is missing\""},
							ast.NewIdent(names.VarNameObjPath),
						},
					}},
				},
			}
		}
		return nil
	}
	if star, ok := f.f.Type.(*ast.StarExpr); ok {
		return []ast.Stmt{
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names:  []*ast.Ident{ast.NewIdent("x" + name)},
							Type:   star.X,
							Values: []ast.Expr{helpers.BasicLiteralFromType(star.X, f.t.DefaultValue())},
						},
					},
				},
			},
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.UnaryExpr{Op: token.AND, X: ast.NewIdent("x" + name)}},
			},
		}
	}
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{helpers.BasicLiteralFromType(f.f.Type, f.t.DefaultValue())},
		},
	}
}
