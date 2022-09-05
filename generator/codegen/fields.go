package codegen

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type (
	fld struct {
		f *ast.Field
		t Tags
	}
)

// 	if offset := v.Get("offset"); offset != nil {
//		s.Offset, err = offset.Int()
//		if err != nil {
//			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//		}
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
	return result
}

//	s.{name}, err = {v}.(Int|Int64|String|Bool)()
func (f fld) typedValue(name, v string) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.f.Type.(type) {

	case *ast.Ident:
		switch t.Name {

		case "int":
			result = append(result, intExtraction(name, v))
			return result

		case "int64":
			result = append(result, int64Extraction(name, v))
			return result

		case "bool":
			result = append(result, boolExtraction(name, v))
			return result

		case "string":
			result = append(result, stringExtraction(name, v, f.t.jsonName())...)
			return result

		default:
			result = append(result, nestedExtraction(name, v, f.t.jsonName()))
			return result

		}

	case *ast.StructType:
		panic("unsupported field type 'struct'")

	case *ast.SelectorExpr:
		panic(fmt.Errorf("unsupported field type '%s.%s'", t.X.(*ast.Ident).Name, t.Sel))

	}
	panic("unsupported field type")
}

//	if {v}.Type() != fastjson.TypeString {
//		err = fmt.Errorf("value doesn't contain string; it contains %s", {v}.Type())
//		return fmt.Errorf("error parsing '%s{json}' value: %w", objPath, err)
//	}
//	s.{name} = {v}.String()
func stringExtraction(name, v, json string) []ast.Stmt {
	var result []ast.Stmt
	var valueType = &ast.CallExpr{
		Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Type")},
	}
	result = append(result, &ast.IfStmt{
		Cond: &ast.BinaryExpr{
			X:  valueType,
			Op: token.NEQ,
			Y:  &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("TypeString")},
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{ast.NewIdent(errVarName)},
				Tok: token.ASSIGN,
				Rhs: []ast.Expr{fmtError("value doesn't contain string; it contains %s", valueType)},
			},
			&ast.ReturnStmt{Results: []ast.Expr{
				fmtError("error parsing '%s"+json+"' value: %w", ast.NewIdent(objPathVarName), ast.NewIdent(errVarName)),
			}},
		}},
	})
	result = append(result, &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("String")},
		}},
	})
	return result
}

// s.{name}, err = {v}.Int()
func intExtraction(name, v string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}, ast.NewIdent(errVarName)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Int")},
		}},
	}
}

// s.{name}, err = {v}.Int64()
func int64Extraction(name, v string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}, ast.NewIdent(errVarName)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Int64")},
		}},
	}
}

// s.{name}, err = {v}.Bool()
func boolExtraction(name, v string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{&ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(name)}, ast.NewIdent(errVarName)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{&ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("Bool")},
		}},
	}
}

// err = s.{name}.fill({v}, objPath+"{json}.")
func nestedExtraction(name, v, json string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(errVarName)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent(recvVarName),
						Sel: ast.NewIdent(name),
					},
					Sel: ast.NewIdent(fillerFuncName),
				},
				Args: []ast.Expr{
					ast.NewIdent(v),
					&ast.BinaryExpr{
						X:  ast.NewIdent(objPathVarName),
						Op: token.ADD,
						Y: &ast.BasicLit{
							Kind:  token.STRING,
							Value: "\"" + strings.Replace(json, "\"", "\\\"", -1) + ".\"",
						},
					},
				},
			},
		},
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

//	} else {
//		s.{name} = 100
//	}
func (f fld) ifDefault(name string) []ast.Stmt {
	if f.t.defaultValue() == "" {
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
