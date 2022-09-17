package codegen

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

const (
	bufVarName        = "buf"
	structPoolName    = "jsonParser"
	byteDataVarName   = "data"
	unmarshalFuncName = "UnmarshalJSON"
	marshalFuncName   = "MarshalJSON"
	validateFuncName  = "validate"
	jsonerFuncName    = "MarshalAppend"
)

// func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
func NewFillerFunc(structName string, fields []*ast.Field, structTags tags.StructTags) *ast.FuncDecl {
	fastJsonValue := ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")}}
	var body []ast.Stmt
	if structTags.StrictRules() {
		body = append(
			body,
			&ast.ExprStmt{X: &ast.BasicLit{Kind: token.COMMENT, Value: "// only if there is a strict rules"}},
			// if err = validateStructKeys(v, ""); err != nil {
			//		return err
			//	}
			&ast.IfStmt{
				Init: &ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent(names.VarNameError)},
					Tok: token.ASSIGN,
					Rhs: []ast.Expr{&ast.CallExpr{
						Fun: &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(validateFuncName)},
						Args: []ast.Expr{
							ast.NewIdent(names.VarNameJsonValue),
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: "\"\"",
							},
						},
					}},
				},
				Cond: &ast.BinaryExpr{
					X:  ast.NewIdent(names.VarNameError),
					Op: token.NEQ,
					Y:  ast.NewIdent("nil"),
				},
				Body: &ast.BlockStmt{List: []ast.Stmt{
					&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent(names.VarNameError)}},
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
			{Text: "\n// " + names.FuncNameFill + " recursively fills the fields with fastjson.Value"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(names.VarNameReceiver)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(names.FuncNameFill),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(names.VarNameJsonValue)}, Type: &fastJsonValue},
				{Names: []*ast.Ident{ast.NewIdent(names.VarNameObjPath)}, Type: ast.NewIdent("string")},
			}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(names.VarNameError)}, Type: ast.NewIdent("error")},
			}},
		},
		Body: &ast.BlockStmt{List: body},
	}
}

func fillFieldStmts(fld *ast.Field) []ast.Stmt {
	var result []ast.Stmt
	factory := field.New(fld)
	for _, name := range fld.Names {
		result = append(result, factory.FillStatements(name.Name)...)
	}
	return result
}

// func (s *Struct) UnmarshalJSON(data []byte) error {
func NewUnmarshalFunc(structName string, structTags tags.StructTags) []ast.Decl {
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
			Lhs: []ast.Expr{ast.NewIdent(names.VarNameJsonValue), ast.NewIdent(names.VarNameError)},
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
				X:  ast.NewIdent(names.VarNameError),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent(names.VarNameError)}},
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
				Fun: &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(names.FuncNameFill)},
				Args: []ast.Expr{ast.NewIdent(names.VarNameJsonValue), &ast.BasicLit{
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
				{Names: []*ast.Ident{ast.NewIdent(names.VarNameReceiver)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
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
func NewValidatorFunc(structName string, fields []*ast.Field, structTags tags.StructTags) *ast.FuncDecl {
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
				X:  ast.NewIdent(names.VarNameError),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{}}},
		},
	}
	for _, field := range fields {
		fieldTags := tags.Parse(field.Tag.Value)
		var runeArgs []ast.Expr
		for name, i := []rune(fieldTags.JsonName()), 0; i < len(name); i++ {
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
				X:  ast.NewIdent(names.VarNameObjPath),
				Op: token.EQL,
				Y: &ast.BasicLit{
					Kind:  token.STRING,
					Value: "\"\"",
				},
			},
			// err = fmt.Errorf("unexpected field '%s' in the root of the object", string(key))
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{ast.NewIdent(names.VarNameError)},
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
					Lhs: []ast.Expr{ast.NewIdent(names.VarNameError)},
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
								ast.NewIdent(names.VarNameObjPath),
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
			Lhs: []ast.Expr{ast.NewIdent(o), ast.NewIdent(names.VarNameError)},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{Fun: &ast.SelectorExpr{X: ast.NewIdent(names.VarNameJsonValue), Sel: ast.NewIdent("Object")}},
			},
		},
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(names.VarNameError),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: &ast.BlockStmt{List: []ast.Stmt{
				&ast.ReturnStmt{Results: []ast.Expr{ast.NewIdent(names.VarNameError)}},
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
			{Names: []*ast.Ident{ast.NewIdent(names.VarNameReceiver)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(validateFuncName),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(names.VarNameJsonValue)}, Type: &fastJsonValue},
				{Names: []*ast.Ident{ast.NewIdent(names.VarNameObjPath)}, Type: ast.NewIdent("string")},
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
func NewMarshalFunc(structName string, structTags tags.StructTags) *ast.FuncDecl {
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// " + marshalFuncName + " serializes the structure with all its values into JSON format"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(names.VarNameReceiver)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
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
						Fun: &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(jsonerFuncName)},
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
func NewAppendJsonFunc(structName string, fields []*ast.Field, structTags tags.StructTags) *ast.FuncDecl {
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
						Names: []*ast.Ident{ast.NewIdent(names.VarNameError)},
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
		// return result.Bytes(), err
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteRune")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.CHAR, Value: "'}'"}},
		}},
		&ast.ReturnStmt{Results: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("Bytes")},
			},
			ast.NewIdent("err"),
		}},
	)
	return &ast.FuncDecl{
		Doc: &ast.CommentGroup{List: []*ast.Comment{
			{Text: "\n// " + jsonerFuncName + " serializes all fields of the structure using a buffer"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(names.VarNameReceiver)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
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
	factory := field.New(fld)
	for _, name := range fld.Names {
		result = append(result, factory.MarshalStatements(name.Name)...)
	}
	return result
}
