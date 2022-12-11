package field

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

// offset := v.Get("offset")
func (f *fld) extract(v string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(v)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{f.getValue()},
	}
}

// v.Get("offset")
func (f *fld) getValue() ast.Expr {
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

func (f *fld) prepareRef() {
	var dstType = f.x
	star, isStar := dstType.(*ast.StarExpr)
	if isStar {
		f.x = star.X
		f.isStar = true
	}
	f.fillDenotedType()
}

func (f *fld) fillDenotedType() {
	if i, ok := f.x.(*ast.Ident); ok {
		f.d = denotedType(i)
	} else {
		f.d = f.x
	}
}

func denotedType(t *ast.Ident) ast.Expr {
	if t.Obj != nil {
		ts, ok := t.Obj.Decl.(*ast.TypeSpec)
		if ok {
			return ts.Type
		}
	}
	return t
}

// fillFrom makes statements to fill some field according to its type
//	s.Offset, err = offset.Int()
//	if err != nil {
//	    return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//	}
func (f *fld) fillFrom(name, v string) []ast.Stmt {
	var bufVariable = ast.NewIdent("val" + name)
	var result []ast.Stmt
	result = append(result, f.typedValue(bufVariable, v)...)
	result = append(result, f.checkErr()...)

	var fldExpr = &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}
	if f.isStar {
		result = append(result, f.fillRefField(bufVariable, fldExpr, v)...)
	} else {
		result = append(result, f.fillField(bufVariable, fldExpr, v)...)
	}
	return result
}

// var elem int
// if elem, err = listElem.Int(); err != nil {
// 	break
// }
// valList = append(valList, int32(elem))
func (f *fld) fillElem(dst ast.Expr, v string) []ast.Stmt {
	var bufVariable = ast.NewIdent("elem")
	var result []ast.Stmt
	result = append(result, f.typedValue(bufVariable, v)...)
	result = append(result, f.breakErr()...)

	// valList = append(valList, int32(elem))
	result = append(result, &ast.AssignStmt{
		Lhs: []ast.Expr{dst},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: ast.NewIdent("append"),
			Args: []ast.Expr{
				dst,
				&ast.CallExpr{
					Fun:  f.x,
					Args: []ast.Expr{ast.NewIdent("elem")},
				},
			},
		}},
	})
	return result
}

//  var val{name} {type}
//	val{name}, err = {v}.(Int|Int64|String|Bool)()
func (f *fld) typedValue(dst *ast.Ident, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.d.(type) {

	case *ast.Ident:
		result = append(result, f.typeExtraction(dst, v, t.Name)...)

	case *ast.StructType:
		result = append(result, f.typeExtraction(dst, v, "?")...)

	case *ast.SelectorExpr:
		switch t.Sel.Name {

		case "Time":
			result = append(result, timeExtraction(dst, v, f.t.Layout())...)

		case "UUID":
			result = append(result, uuidExtraction(dst, f.d, v, f.t.JsonName())...)

		default:
			result = append(result, nestedExtraction(dst, f.x, v, f.t.JsonName())...)
		}

	case *ast.ArrayType:
		intF := fld{
			x: t.Elt,
			t: tags.Parse(fmt.Sprintf(`json:"%s"`, f.t.JsonName())),
		}
		intF.prepareRef()
		result = append(result, arrayExtraction(dst, v, f.t.JsonName(), t.Elt, intF.fillElem(dst, "listElem"))...)
		return result

	default:
		panic("unsupported field type")
	}
	return result
}

func (f *fld) typeExtraction(dst *ast.Ident, v, t string) []ast.Stmt {
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
		return nestedExtraction(dst, f.x, v, f.t.JsonName())

	}
}

//	if err != nil {
//		return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//	}
func (f *fld) checkErr() []ast.Stmt {
	if t, ok := f.d.(*ast.Ident); ok && t.Name == "string" {
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

//	if err != nil {
//		break
//	}
func (f *fld) breakErr() []ast.Stmt {
	if t, ok := f.x.(*ast.Ident); ok && t.Name == "string" {
		// no error checking for string
		return nil
	}
	return []ast.Stmt{
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(names.VarNameError),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.BranchStmt{Tok: token.BREAK},
			}},
		},
	}
}

func (f *fld) fillRefField(rhs, dst ast.Expr, t string) []ast.Stmt {
	switch t := f.x.(type) {

	case *ast.Ident:
		switch t.Name {

		case "bool", "int64", "int", "float64":
			return f.typedFillIn(&ast.UnaryExpr{X: rhs, Op: token.AND}, dst, t.Name)

		default:
			return f.newAndFillIn(rhs, dst, ast.NewIdent(t.Name))

		}

	default:
		return f.newAndFillIn(rhs, dst, f.x)

	}
}

// {dst} = new({t})
// *{dst} = {t}({rhs})
func (f *fld) newAndFillIn(rhs, dst, t ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{Fun: ast.NewIdent("new"), Args: []ast.Expr{t}},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.StarExpr{X: dst}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{Fun: t, Args: []ast.Expr{rhs}},
			},
		},
	}
}

func (f *fld) fillField(rhs, dst ast.Expr, t string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.x.(type) {

	case *ast.Ident:
		return f.typedFillIn(rhs, dst, t.Name)

	case *ast.StructType:
		return result

	case *ast.SelectorExpr:
		// Structures that support the Unmarshaler interface do not require an assignment operation,
		// because they are decoded directly into the target field
		switch t.Sel.Name {

		case "Time":
			fallthrough
		case "UUID":
			result = append(
				result,
				&ast.AssignStmt{
					Lhs: []ast.Expr{dst},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{rhs},
				},
			)
		default:

		}
		return result

	case *ast.ArrayType:
		// s.List = valList
		result = append(
			result,
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{rhs},
			},
		)
		return result

	default:
		return nil
	}
}

func (f *fld) typedFillIn(rhs, dst ast.Expr, t string) []ast.Stmt {
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
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent(t),
					Args: []ast.Expr{rhs},
				}},
			},
		}
	}
}

func (f *fld) typedRefFillIn(rhs, dst ast.Expr, t string) []ast.Stmt {
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
func (f *fld) ifDefault(name string) []ast.Stmt {
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
	if f.isStar {
		return []ast.Stmt{
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names:  []*ast.Ident{ast.NewIdent("x" + name)},
							Type:   f.x,
							Values: []ast.Expr{helpers.BasicLiteralFromType(f.x, f.t.DefaultValue())},
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
			Rhs: []ast.Expr{helpers.BasicLiteralFromType(f.x, f.t.DefaultValue())},
		},
	}
}
