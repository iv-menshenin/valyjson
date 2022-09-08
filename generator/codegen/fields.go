package codegen

import (
	"go/ast"
	"go/token"
)

type (
	// fld render helper for ast.Field
	fld struct {
		// f contains field AST
		f *ast.Field
		// t contains tag descriptor
		t Tags
	}
)

// 	if offset := v.Get("offset"); offset != nil {
//      var vOffset int
//		vOffset, err = offset.Int()
//		if err != nil {
//			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//		}
//      s.Offset = vOffset
//	} else {
//		s.Offset = 100
//	}
func (f fld) Explore(name string) []ast.Stmt {
	if f.t.jsonName() == "" {
		return nil
	}
	var v = varName(name, f.t)
	var body *ast.BlockStmt
	var els ast.Stmt
	if stmt := f.fillFrom(name, v); len(stmt) > 0 {
		body = &ast.BlockStmt{List: stmt}
	}
	if stmt := f.ifDefault(name); len(stmt) > 0 {
		els = &ast.BlockStmt{List: stmt}
	}
	if body == nil {
		// todo @menshenin panic?
		return nil
	}
	return []ast.Stmt{
		&ast.IfStmt{
			Init: f.extract(v),
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(v),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: body,
			Else: els,
		},
	}
}

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
			X:   ast.NewIdent(valVarName),
			Sel: ast.NewIdent("Get"),
		},
		Args: []ast.Expr{&ast.BasicLit{
			Kind:  token.STRING,
			Value: "\"" + f.t.jsonName() + "\"",
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
		result = append(result, nestedExtraction(name, v, f.t.jsonName())...)
		return result

	case *ast.ArrayType:
		result = append(result, arrayExtraction(name, v, f.t.jsonName())...)
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
		return stringExtraction("x"+name, v, f.t.jsonName())

	default:
		return nestedExtraction(name, v, f.t.jsonName())

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
	format := "error parsing '%s" + f.t.jsonName() + "' value: %w"
	return []ast.Stmt{
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(errVarName),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{fmtError(format, ast.NewIdent(objPathVarName), ast.NewIdent(errVarName))},
				},
			}},
		},
	}
}

func (f fld) fillField(name, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.f.Type.(type) {

	case *ast.Ident:
		return f.typedFillIn(name, t.Name, false)

	case *ast.StructType:
		return result

	case *ast.SelectorExpr:
		return result

	case *ast.ArrayType:
		return result

	case *ast.StarExpr:
		return f.typedFillIn(name, t.X.(*ast.Ident).Name, true)

	}
	return nil
}

func (f fld) typedFillIn(name, t string, amp bool) []ast.Stmt {
	var rhs ast.Expr = ast.NewIdent("x" + name)
	if amp {
		rhs = &ast.UnaryExpr{X: rhs, Op: token.AND}
	}
	switch t {
	case "string", "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{rhs},
			},
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		return []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}},
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

//	} else {
//		s.{name} = 100
//	}
func (f fld) ifDefault(name string) []ast.Stmt {
	if f.t.defaultValue() == "" {
		if f.t.jsonTags().Has("required") {
			// return fmt.Errorf("required element '%s{json}' is missing", objPath)
			return []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent("fmt"), Sel: ast.NewIdent("Errorf")},
						Args: []ast.Expr{
							&ast.BasicLit{Kind: token.STRING, Value: "\"required element '%s" + f.t.jsonName() + "' is missing\""},
							ast.NewIdent(objPathVarName),
						},
					}},
				},
			}
		}
		return nil
	}
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{getBasicLitFromString(f.f.Type, f.t.defaultValue())},
		},
	}
}

func newField(f *ast.Field) *fld {
	if f.Tag == nil {
		panic("you must fill in all fields with tags")
	}
	return &fld{
		f: f,
		t: parseTags(f.Tag.Value),
	}
}
