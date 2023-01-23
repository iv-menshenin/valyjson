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
			Value: "\"" + f.tags.JsonName() + "\"",
		}},
	}
}

func (f *fld) prepareRef() {
	var dstType = f.expr
	star, isStar := dstType.(*ast.StarExpr)
	if isStar {
		f.expr = star.X
		f.isStar = true
	}
	_, isArray := dstType.(*ast.ArrayType)
	_, isMap := dstType.(*ast.MapType)
	_, isStruct := dstType.(*ast.StructType)
	f.isNullable = isStar || isArray || isMap || isStruct
	f.fillDenotedType()
}

func (f *fld) fillDenotedType() {
	if i, ok := f.expr.(*ast.Ident); ok {
		f.refx = denotedType(i)
	} else {
		f.refx = f.expr
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
	result = append(result, f.checkErr(bufVariable)...)

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
	if f.isNullable {
		//if !valueIsNotNull(listElem) {
		//	valFieldRef = append(valFieldRef, nil)
		//	continue
		//}
		result = append(result, &ast.IfStmt{
			Cond: &ast.UnaryExpr{
				Op: token.NOT,
				X:  &ast.CallExpr{Fun: ast.NewIdent("valueIsNotNull"), Args: []ast.Expr{ast.NewIdent("listElem")}},
			},
			Body: &ast.BlockStmt{
				List: []ast.Stmt{
					appendStmt(dst, ast.NewIdent("nil")),
					&ast.BranchStmt{Tok: token.CONTINUE},
				},
			},
		})
	}
	result = append(result, f.typedValue(bufVariable, v)...)
	result = append(result, f.breakErr()...)

	// valList = append(valList, int32(elem))
	if f.isStar {
		result = append(
			result,
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent("newElem")},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun:  f.expr,
						Args: []ast.Expr{ast.NewIdent("elem")},
					},
				},
			},
			appendStmt(dst, &ast.UnaryExpr{X: ast.NewIdent("newElem"), Op: token.AND}),
		)
		return result
	}
	result = append(result, appendStmt(dst, &ast.CallExpr{
		Fun:  f.expr,
		Args: []ast.Expr{ast.NewIdent("elem")},
	}))
	return result
}

func appendStmt(dst, el ast.Expr) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{dst},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun:  ast.NewIdent("append"),
			Args: []ast.Expr{dst, el},
		}},
	}
}

//  var val{name} {type}
//	val{name}, err = {v}.(Int|Int64|String|Bool)()
func (f *fld) typedValue(dst *ast.Ident, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.refx.(type) {

	case *ast.Ident:
		result = append(result, f.typeExtraction(dst, v, t.Name)...)

	case *ast.StructType:
		result = append(result, f.typeExtraction(dst, v, "?")...)

	case *ast.SelectorExpr:
		switch t.Sel.Name {

		case "Time":
			result = append(result, timeExtraction(dst, v, f.tags.Layout())...)

		case "UUID":
			result = append(result, uuidExtraction(dst, f.refx, v, f.tags.JsonName())...)

		default:
			result = append(result, nestedExtraction(dst, f.expr, v, f.tags.JsonName())...)
		}

	case *ast.ArrayType:
		intF := fld{
			expr: t.Elt,
			tags: tags.Parse(fmt.Sprintf(`json:"%s"`, f.tags.JsonName())),
		}
		intF.prepareRef()
		result = append(result, arrayExtraction(dst, v, f.tags.JsonName(), t.Elt, intF.fillElem(dst, "listElem"))...)
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
		return stringExtraction(dst, v, f.tags.JsonName())

	default:
		return nestedExtraction(dst, f.expr, v, f.tags.JsonName())

	}
}

//	if err != nil {
//		return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//	}
//	if valIntFld8 > math.MaxInt8 {
//		return fmt.Errorf("error parsing '%sint_fld8' value %d exceeds maximum for data type uint8", objPath, valIntFld8)
//	}
func (f *fld) checkErr(val *ast.Ident) []ast.Stmt {
	var checkOverflow ast.Stmt = &ast.EmptyStmt{}
	ident, isIdent := f.refx.(*ast.Ident)
	if isIdent && ident.Name == "string" {
		return nil
	}
	if maxExp := getMaxByType(ident); maxExp != nil {
		phldr := "%d"
		if ident.Name == "float32" {
			phldr = "%f"
		}
		maxExceeded := "error parsing '%s" + f.tags.JsonName() + "' value " + phldr + " exceeds maximum for data type " + ident.Name
		checkOverflow = &ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  val,
				Op: token.GTR,
				Y:  maxExp,
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{helpers.FmtError(maxExceeded, ast.NewIdent(names.VarNameObjPath), val)},
				},
			}},
		}
	}

	format := "error parsing '%s" + f.tags.JsonName() + "' value: %w"
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
		checkOverflow,
	}
}
func getMaxByType(t *ast.Ident) ast.Expr {
	if t == nil {
		return nil
	}
	switch t.Name {
	case "float32":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxFloat32")}
	case "int8":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxInt8")}
	case "int16":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxInt16")}
	case "int32":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxInt32")}
	case "uint8":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxUint8")}
	case "uint16":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxUint16")}
	case "uint32":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxUint32")}
	}
	return nil
}

//	if err != nil {
//		break
//	}
func (f *fld) breakErr() []ast.Stmt {
	if t, ok := f.expr.(*ast.Ident); ok && t.Name == "string" {
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
	switch t := f.expr.(type) {

	case *ast.Ident:
		switch t.Name {

		case "bool", "int64", "int", "float64":
			return f.typedFillIn(&ast.UnaryExpr{X: rhs, Op: token.AND}, dst, t.Name)

		default:
			return f.newAndFillIn(rhs, dst, ast.NewIdent(t.Name))

		}

	default:
		return f.newAndFillIn(rhs, dst, f.expr)

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
	switch t := f.expr.(type) {

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
	case "string":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent("string"),
					Args: []ast.Expr{rhs},
				}},
			},
		}

	case "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
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
	if f.tags.DefaultValue() == "" {
		if f.tags.JsonTags().Has("required") {
			// return fmt.Errorf("required element '%s{json}' is missing", objPath)
			return []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent("fmt"), Sel: ast.NewIdent("Errorf")},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"required element '%s" + f.tags.JsonName() + "' is missing\""},
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
							Type:   f.expr,
							Values: []ast.Expr{helpers.BasicLiteralFromType(f.expr, f.tags.DefaultValue())},
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
			Rhs: []ast.Expr{helpers.BasicLiteralFromType(f.expr, f.tags.DefaultValue())},
		},
	}
}
