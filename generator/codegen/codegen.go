package codegen

import (
	"fmt"
	"go/ast"
	"go/token"
)

const (
	errVarName        = "err"
	recvVarName       = "s"
	valVarName        = "v"
	bufVarName        = "buf"
	structPoolName    = "jsonParser"
	objPathVarName    = "objPath"
	byteDataVarName   = "data"
	fillerFuncName    = "FillFromJson"
	unmarshalFuncName = "UnmarshalJSON"
	marshalFuncName   = "MarshalJSON"
	validateFuncName  = "validate"
	jsonerFuncName    = "MarshalAppend"
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
		body = append(body, fillFieldStmts(field)...)
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

func fillFieldStmts(fld *ast.Field) []ast.Stmt {
	var result []ast.Stmt
	factory := newField(fld)
	for _, name := range fld.Names {
		result = append(result, factory.FillField(name.Name)...)
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

// func (s *S) MarshalJSON() ([]byte, error) {
// 	var buf [128]byte
// 	return s.marshalAppend(buf[:0])
// }
func NewMarshalFunc(structName string, tags StructTags) *ast.FuncDecl {
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// " + marshalFuncName + " serializes the structure with all its values into JSON format"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(recvVarName)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(marshalFuncName),
		Type: &ast.FuncType{
			Results: &ast.FieldList{List: []*ast.Field{
				{Type: &ast.ArrayType{Elt: ast.NewIdent("byte")}},
				{Type: ast.NewIdent("error")},
			}},
		},
		Body: &ast.BlockStmt{List: []ast.Stmt{
			&ast.DeclStmt{
				Decl: &ast.GenDecl{
					Tok: token.VAR,
					Specs: []ast.Spec{
						&ast.ValueSpec{
							Names: []*ast.Ident{ast.NewIdent(bufVarName)},
							Type: &ast.ArrayType{
								// todo @menshenin calculate buffer length
								Len: &ast.BasicLit{Kind: token.INT, Value: "128"},
								Elt: ast.NewIdent("byte"),
							},
						},
					},
				},
			},
			&ast.ReturnStmt{
				Results: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent(recvVarName), Sel: ast.NewIdent(jsonerFuncName)},
						Args: []ast.Expr{
							&ast.SliceExpr{X: ast.NewIdent(bufVarName), High: &ast.BasicLit{Kind: token.INT, Value: "0"}},
						},
					},
				},
			},
		}},
	}
}

// func (s *S) MarshalAppend(dst []byte) ([]byte, error) {
func NewAppendJsonFunc(structName string, fields []*ast.Field, tags StructTags) *ast.FuncDecl {
	var body = []ast.Stmt{
		// var result = bytes.NewBuffer(dst)
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent("result")},
						Values: []ast.Expr{
							&ast.CallExpr{
								Fun:  &ast.SelectorExpr{X: ast.NewIdent("bytes"), Sel: ast.NewIdent("NewBuffer")},
								Args: []ast.Expr{ast.NewIdent("dst")},
							},
						},
					},
				},
			},
		},
		// var (
		// 	b   []byte
		// 	buf [128]byte
		// 	err error
		// )
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent("b")},
						Type:  &ast.ArrayType{Elt: ast.NewIdent("byte")},
					},
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent("buf")},
						Type:  &ast.ArrayType{Elt: ast.NewIdent("byte"), Len: &ast.BasicLit{Kind: token.INT, Value: "128"}},
					},
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent(errVarName)},
						Type:  ast.NewIdent("error"),
					},
				},
			},
		},
		// result.WriteRune('{')
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteRune")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.CHAR, Value: "'{'"}},
		}},
	}
	for _, field := range fields {
		body = append(body, jsonFieldStmts(field)...)
	}
	body = append(
		body,
		// result.WriteRune('}')
		// return result.Bytes(), nil
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteRune")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.CHAR, Value: "'}'"}},
		}},
		&ast.ReturnStmt{Results: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("Bytes")},
			},
			ast.NewIdent("nil"),
		}},
	)
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// " + jsonerFuncName + " serializes all fields of the structure using a buffer"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(recvVarName)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(jsonerFuncName),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent("dst")}, Type: &ast.ArrayType{Elt: ast.NewIdent("byte")}},
			}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Type: &ast.ArrayType{Elt: ast.NewIdent("byte")}},
				{Type: ast.NewIdent("error")},
			}},
		},
		Body: &ast.BlockStmt{List: body},
	}
}

func jsonFieldStmts(fld *ast.Field) []ast.Stmt {
	var result []ast.Stmt
	factory := newField(fld)
	for _, name := range fld.Names {
		result = append(result, factory.MarshalField(name.Name)...)
	}
	return result
}
