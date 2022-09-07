package codegen

import (
	"fmt"
	"go/ast"
	"go/token"
)

/*
func (s *Struct) fill(v *fastjson.Value, objPath string) (err error) {
	if filter := v.Get("filter"); filter != nil {
		if filter.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", v.Type())
			return fmt.Errorf("error parsing '%sfilter' value: %w", objPath, err)
		}
		s.Filter = filter.String()
	} else {
		return fmt.Errorf("the '%sfilter' path is required but ommitted", objPath)
	}
	if limit := v.Get("limit"); limit != nil {
		s.Limit, err = limit.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	}
	if offset := v.Get("offset"); offset != nil {
		s.Offset, err = offset.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		err = s.Nested.fill(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	}
	return nil
}
*/

const (
	errVarName        = "err"
	recvVarName       = "s"
	valVarName        = "v"
	structPoolName    = "jsonParser"
	objPathVarName    = "objPath"
	byteDataVarName   = "data"
	fillerFuncName    = "FillFromJson"
	unmarshalFuncName = "UnmarshalJSON"
	validateFuncName  = "validate"
)

// func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
func NewFillerFunc(structName string, fields []*ast.Field, tags StructTags) *ast.FuncDecl {
	fastJsonValue := ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")}}
	var body []ast.Stmt
	if tags.Has(strictRules) {
		body = append(
			body,
			&ast.ExprStmt{X: &ast.BasicLit{Kind: token.COMMENT, Value: "// only if there is a strict rules"}},
			// if err = validateStructKeys(v, ""); err != nil {
			//		return err
			//	}
			&ast.IfStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent(errVarName)},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(validateFuncName)},
						Args: []ast.Expr{
							ast.NewIdent(valVarName),
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"\"",
							},
						},
					}},
				},
				Cond: &ast.BinaryExpr{
					X:  ast.NewIdent(errVarName),
					Op: token.NEQ,
					Y:  ast.NewIdent("nil"),
				},
				Body: &ast.BlockStmt{List: []ast.Stmt{
					&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent(errVarName)}},
				}},
			},
		)
	}
	for _, field := range fields {
		body = append(body, NewFieldFillerStmt(field)...)
	}
	body = append(
		body,
		// return nil
		&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent("nil")}},
	)
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// " + fillerFuncName + " recursively fills the fields with fastjson.Value"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(recvVarName)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(fillerFuncName),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(valVarName)}, Type: &fastJsonValue},
				{Names: []*ast.Ident{ast.NewIdent(objPathVarName)}, Type: ast.NewIdent("string")},
			}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(errVarName)}, Type: ast.NewIdent("error")},
			}},
		},
		Body: &ast.BlockStmt{List: body},
	}
}

func NewFieldFillerStmt(fld *ast.Field) []ast.Stmt {
	var result []ast.Stmt
	factory := newField(fld)
	for _, name := range fld.Names {
		result = append(result, factory.Explore(name.Name)...)
	}
	return result
}

// func (s *Struct) UnmarshalJSON(data []byte) error {
func NewUnmarshalFunc(structName string, tags StructTags) []ast.Decl {
	const (
		parser = "parser"
	)
	poolName := fmt.Sprintf("%s%s", structPoolName, structName)
	var body = []ast.Stmt{
		//	parser := structPool.Get()
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(parser)},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(poolName), Sel: ast.NewIdent("Get")},
			}},
		},
		&ast.ExprStmt{X: &ast.BasicLit{Kind: token.COMMENT, Value: "// parses data containing JSON"}},
		//	v, err := parser.ParseBytes(data)
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(valVarName), ast.NewIdent(errVarName)},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{&ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: ast.NewIdent(parser), Sel: ast.NewIdent("ParseBytes")},
				Args: []ast.Expr{ast.NewIdent(byteDataVarName)},
			}},
		},
		//	if err != nil {
		//		return err
		//	}
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(errVarName),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent(errVarName)}},
			}},
		},
		//	defer structPool.Put(parser)
		&ast.DeferStmt{
			Call: &ast.CallExpr{
				Fun:  &ast.SelectorExpr{X: ast.NewIdent(poolName), Sel: ast.NewIdent("Put")},
				Args: []ast.Expr{ast.NewIdent(parser)},
			},
		},
		//	return s.FillFromJson(v, "")
		&ast.ReturnStmt{
			Results: []ast.Expr{&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(fillerFuncName)},
				Args: []ast.Expr{ast.NewIdent(valVarName), &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"\"",
				}},
			}},
		},
	}
	return []ast.Decl{
		// var structPool fastjson.ParserPool
		&ast.GenDecl{
			Tok: token.VAR,
			Doc: &ast.CommentGroup{List: []*ast.Comment{
				{Text: "// " + poolName + "used for pooling Parsers for " + structName + " JSONs."},
			}},
			Specs: []ast.Spec{
				&ast.ValueSpec{
					Names: []*ast.Ident{ast.NewIdent(poolName)},
					Type:  &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("ParserPool")},
				},
			},
		},
		// func (s *Struct) UnmarshalJSON(data []byte) error {
		&ast.FuncDecl{
			Doc: &ast.CommentGroup{List: []*ast.Comment{
				{Text: "\n// " + unmarshalFuncName + " implements json.Unmarshaler"},
			}},
			Recv: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(recvVarName)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
			}},
			Name: ast.NewIdent(unmarshalFuncName),
			Type: &ast.FuncType{
				Params: &ast.FieldList{List: []*ast.Field{
					{Names: []*ast.Ident{ast.NewIdent(byteDataVarName)}, Type: &ast.ArrayType{Elt: ast.NewIdent("byte")}},
				}},
				Results: &ast.FieldList{List: []*ast.Field{
					{Type: ast.NewIdent("error")},
				}},
			},
			Body: &ast.BlockStmt{List: body},
		},
	}
}

// func validateStructKeys(v *fastjson.Value, objPath string) error {
func NewValidatorFunc(structName string, fields []*ast.Field, tags StructTags) *ast.FuncDecl {
	const (
		o   = "o"
		v   = "_"
		key = "key"
	)
	fastJsonValue := ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")}}
	var visitBody = []ast.Stmt{
		//		if err != nil {
		//			return
		//		}
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(errVarName),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{}}},
		},
	}
	for _, field := range fields {
		fTags := parseTags(field.Tag.Value)
		var runeArgs []ast.Expr
		for name, i := []rune(fTags.jsonName()), 0; i < len(name); i++ {
			runeArgs = append(runeArgs, &ast.BasicLit{
				Kind:  token.CHAR,
				Value: "'" + string(name[i]) + "'",
			})
		}
		visitBody = append(
			visitBody,
			//		if bytes.Equal(key, []byte{'f', 'i', 'l', 't', 'e', 'r'}) {
			//			return
			//		}
			&ast.IfStmt{
				Cond: &ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("bytes"), Sel: ast.NewIdent("Equal")},
					Args: []ast.Expr{
						ast.NewIdent(key),
						&ast.CompositeLit{
							Type: &ast.ArrayType{Elt: ast.NewIdent("byte")},
							Elts: runeArgs,
						},
					},
				},
				Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{}}},
			},
		)
	}
	//		if objPath == "" {
	//			err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
	//		} else {
	//			err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
	//		}
	visitBody = append(
		visitBody,
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(objPathVarName),
				Op: token.EQL,
				Y: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"\"",
				},
			},
			// err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent(errVarName)},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fmt"),
								Sel: ast.NewIdent("Errorf"),
							},
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: "\"unexpected field '%s' in the root of the object\"",
								},
								&ast.CallExpr{
									Fun:  ast.NewIdent("string"),
									Args: []ast.Expr{ast.NewIdent(key)},
								},
							},
						},
					},
				},
			}},
			// err = fmt.Errorf("unexpected field '%s' in the '%s' path", string(key), objPath)
			Else: &ast.BlockStmt{List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent(errVarName)},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("fmt"),
								Sel: ast.NewIdent("Errorf"),
							},
							Args: []ast.Expr{
								&ast.BasicLit{
									Kind:  token.STRING,
									Value: "\"unexpected field '%s' in the '%s' path\"",
								},
								&ast.CallExpr{
									Fun:  ast.NewIdent("string"),
									Args: []ast.Expr{ast.NewIdent(key)},
								},
								ast.NewIdent(objPathVarName),
							},
						},
					},
				},
			}},
		},
	)
	//	o, err := v.Object()
	//	if err != nil {
	//		return err
	//	}
	var body = []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent(o), ast.NewIdent(errVarName)},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{Fun: &ast.SelectorExpr{X: ast.NewIdent(valVarName), Sel: ast.NewIdent("Object")}},
			},
		},
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(errVarName),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent(errVarName)}},
			}},
		},
		//	o.Visit(func(key []byte, _ *fastjson.Value) {
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun: &ast.SelectorExpr{X: ast.NewIdent(o), Sel: ast.NewIdent("Visit")},
			Args: []ast.Expr{
				&ast.FuncLit{
					Type: &ast.FuncType{
						Params: &ast.FieldList{List: []*ast.Field{
							{
								Names: []*ast.Ident{ast.NewIdent(key)},
								Type:  &ast.ArrayType{Elt: ast.NewIdent("byte")},
							},
							{
								Names: []*ast.Ident{ast.NewIdent(v)},
								Type: &ast.StarExpr{
									X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")},
								},
							},
						}},
					},
					Body: &ast.BlockStmt{List: visitBody},
				},
			},
		}},
		// return nil
		&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent("nil")}},
	}
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// " + validateFuncName + " checks for correct data structure"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(recvVarName)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(validateFuncName),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(valVarName)}, Type: &fastJsonValue},
				{Names: []*ast.Ident{ast.NewIdent(objPathVarName)}, Type: ast.NewIdent("string")},
			}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Type: ast.NewIdent("error")},
			}},
		},
		Body: &ast.BlockStmt{List: body},
	}
}
