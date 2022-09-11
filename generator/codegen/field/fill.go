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
	var result []ast.Stmt
	result = append(result, f.typedValue(name, v)...)
	result = append(result, f.checkErr()...)
	result = append(result, f.fillField(name, v)...)
	return result
}

//	s.{name}, err = {v}.(Int|Int64|String|Bool)()
func (f fld) typedValue(name, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.f.Type.(type) {

	case *ast.Ident:
		return f.typeExtraction(name, v, t.Name)

	case *ast.StructType:
		panic("unsupported field type 'struct'")

	case *ast.SelectorExpr:
		result = append(result, nestedExtraction(name, v, f.t.JsonName())...)
		return result

	case *ast.ArrayType:
		result = append(result, arrayExtraction(name, v, f.t.JsonName())...)
		return result

	case *ast.StarExpr:
		return f.typeExtraction(name, v, t.X.(*ast.Ident).Name)

	}
	panic("unsupported field type")
}

func (f fld) typeExtraction(name, v, t string) []ast.Stmt {
	switch t {

	case "int", "int8", "int16", "int32":
		return intExtraction("x"+name, v)

	case "int64":
		return int64Extraction("x"+name, v)

	case "uint", "uint8", "uint16", "uint32":
		return uintExtraction("x"+name, v)

	case "uint64":
		return uint64Extraction("x"+name, v)

	case "float32", "float64":
		return floatExtraction("x"+name, v)

	case "bool":
		return boolExtraction("x"+name, v)

	case "string":
		return stringExtraction("x"+name, v, f.t.JsonName())

	default:
		return nestedExtraction(name, v, f.t.JsonName())

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

func (f fld) fillField(name, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.f.Type.(type) {

	case *ast.Ident:
		return f.typedFillIn(name, t.Name)

	case *ast.StructType:
		return result

	case *ast.SelectorExpr:
		return result

	case *ast.ArrayType:
		return result

	case *ast.StarExpr:
		return f.typedRefFillIn(name, t.X.(*ast.Ident).Name)

	}
	return nil
}

func (f fld) typedFillIn(name, t string) []ast.Stmt {
	var rhs ast.Expr = ast.NewIdent("x" + name)
	switch t {
	case "string", "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{rhs},
			},
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent(t),
					Args: []ast.Expr{rhs},
				}},
			},
		}

	default:
		return nil
	}
}

func (f fld) typedRefFillIn(name, t string) []ast.Stmt {
	switch t {
	case "string", "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.UnaryExpr{X: ast.NewIdent("x" + name), Op: token.AND}},
			},
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		var result []ast.Stmt
		result = append(
			result,
			// s.HeightRef = new(uint32)
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent("new"),
					Args: []ast.Expr{ast.NewIdent(t)},
				}},
			},
			// *s.HeightRef = uint32(xHeightRef)
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.StarExpr{
					X: &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)},
				}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{&ast.CallExpr{
					Fun:  ast.NewIdent(t),
					Args: []ast.Expr{ast.NewIdent("x" + name)},
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
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{helpers.BasicLiteralFromType(f.f.Type, f.t.DefaultValue())},
		},
	}
}
