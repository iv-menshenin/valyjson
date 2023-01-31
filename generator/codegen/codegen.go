package codegen

import (
	"fmt"
	"github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
	"go/ast"
	"go/token"
)

// NewFillerFunc generates function code that will fill in all fields of the structure with the fastjson.Value attribute
//   func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) { ... }
func NewFillerFunc(structName string, fields []*ast.Field, structTags tags.StructTags) ast.Decl {
	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " recursively fills the fields with fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)
	if structTags.StrictRules() {
		fn.AppendStmt(asthlp.CommentStmt("strict rules"))
	}
	fn.AppendStmt(
		// 	if err = s.validate(v, ""); err != nil {
		//		return err
		//	}
		asthlp.MakeCallReturnIfError(nil, asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameReceiver, names.MethodNameValidate)),
			asthlp.NewIdent(names.VarNameJsonValue),
			asthlp.NewIdent(names.VarNameObjPath),
		)),
	)
	for _, f := range fields {
		fn.AppendStmt(fillFieldStmts(f)...)
	}
	fn.AppendStmt(asthlp.Return(asthlp.Nil))
	return fn.Decl()
}

func fillFieldStmts(fld *ast.Field) []ast.Stmt {
	var result []ast.Stmt
	factory := field.New(fld)
	for _, name := range fld.Names {
		result = append(result, factory.FillStatements(name.Name)...)
	}
	if len(fld.Names) == 0 {
		// composited struct
		var tag tags.Tags
		if fld.Tag != nil {
			tag = tags.Parse(fld.Tag.Value)
		}
		if tag.JsonAppendix() == "inline" {
			// panic("dfs")
		}
		if i, ok := fld.Type.(*ast.Ident); ok {
			result = append(result, factory.FillStatements(i.Name)...)
		}
	}
	return result
}

// NewUnmarshalFunc generates code for unmarshalling, function that parses the JSON object and fills all the structure properties.
//  func (s *Struct) UnmarshalJSON(data []byte) error {
func NewUnmarshalFunc(structName string, _ tags.StructTags) []ast.Decl {
	const parser = "parser"
	poolName := fmt.Sprintf("%s%s", names.VarPrefixPool, structName)

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameUnmarshal))
	fn.Comments("// " + names.MethodNameUnmarshal + " implements json.Unmarshaler")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName))))
	fn.Params(asthlp.Field(names.VarNameData, nil, asthlp.ArrayType(asthlp.Byte)))
	fn.Results(asthlp.Field("", nil, asthlp.ErrorType))
	fn.AppendStmt(
		asthlp.Assign(asthlp.MakeVarNames(parser), asthlp.Definition, asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(poolName, "Get")))),
		asthlp.CommentStmt("parses data containing JSON"),
		asthlp.Assign(
			asthlp.MakeVarNames(names.VarNameJsonValue, names.VarNameError),
			asthlp.Definition,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(parser, "ParseBytes")), ast.NewIdent(names.VarNameData)),
		),
		asthlp.If(asthlp.NotNil(asthlp.NewIdent(names.VarNameError)), asthlp.Return(asthlp.NewIdent(names.VarNameError))),
		asthlp.DeferCall(asthlp.InlineFunc(asthlp.SimpleSelector(poolName, "Put")), ast.NewIdent(parser)),
		//	return s.FillFromJson(v, "")
		asthlp.Return(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameReceiver, names.MethodNameFill)),
			ast.NewIdent(names.VarNameJsonValue),
			asthlp.EmptyString,
		)),
	)

	return []ast.Decl{
		asthlp.DeclareVariable().
			Comments("// " + poolName + "used for pooling Parsers for " + structName + " JSONs.").
			AppendSpec(asthlp.VariableType(poolName, names.FastJsonParserPool)).
			Decl(),
		fn.Decl(),
	}
}

// NewValidatorFunc generates a function declaration for validating a JSON object, taking into account the presence of fields.
//  func validate(v *fastjson.Value, objPath string) error {
func NewValidatorFunc(structName string, fields []*ast.Field, structTags tags.StructTags) ast.Decl {
	fastJsonValue := asthlp.Star(names.FastJsonValue)
	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameValidate))
	fn.Comments("// " + names.MethodNameValidate + " checks for correct data structure")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, fastJsonValue),
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
	)
	fn.Results(
		asthlp.Field("", nil, asthlp.ErrorType),
	)

	const (
		keyVarName         = "key"
		jsonObjectVarName  = "o"
		checkFieldsVarName = "checkFields"
	)
	visitFunc := asthlp.DeclareFunction(nil).Params(
		asthlp.Field(keyVarName, nil, asthlp.ArrayType(asthlp.Byte)),
		asthlp.Field(asthlp.Blank.Name, nil, asthlp.Star(names.FastJsonValue)),
	)
	visitFunc.AppendStmt(
		//		if err != nil {
		//			return
		//		}
		asthlp.If(
			asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
			asthlp.ReturnEmpty(),
		),
	)
	for i, field := range fields {
		var fieldTags tags.Tags
		if field.Tag != nil {
			fieldTags = tags.Parse(field.Tag.Value)
		}
		visitFunc.AppendStmt(
			//		if bytes.Equal(key, []byte{'f', 'i', 'l', 't', 'e', 'r'}) {
			//			. . .
			//		}
			asthlp.If(
				asthlp.Call(asthlp.BytesEqualFn, asthlp.NewIdent(keyVarName), asthlp.SliceByteLiteral(fieldTags.JsonName()).Expr()),
				asthlp.Increment(asthlp.Index(ast.NewIdent(checkFieldsVarName), asthlp.IntegerConstant(i))),
				asthlp.If(
					asthlp.Great(
						asthlp.Index(ast.NewIdent(checkFieldsVarName), asthlp.IntegerConstant(i)),
						asthlp.IntegerConstant(1).Expr(),
					),
					// err = fmt.Errorf("the '%s' field appears in the object twice [%s]", string(key), objPath)
					asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
						asthlp.FmtErrorfFn,
						asthlp.StringConstant("the '%s%s' field appears in the object twice").Expr(),
						ast.NewIdent(names.VarNameObjPath),
						asthlp.ExpressionTypeConvert(asthlp.NewIdent(keyVarName), asthlp.String),
					)),
				),
				asthlp.ReturnEmpty(),
			),
		)
	}
	//	err = fmt.Errorf("unexpected field '%s%s'", objPath, string(key))
	if structTags.StrictRules() {
		// If there were unregistered data fields in the JSON object, execution will surely get to that point.
		// With strict rules it is necessary to register an error
		visitFunc.AppendStmt(
			asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
				asthlp.FmtErrorfFn,
				asthlp.StringConstant("unexpected field '%s%s'").Expr(),
				asthlp.NewIdent(names.VarNameObjPath),
				asthlp.ExpressionTypeConvert(asthlp.NewIdent(keyVarName), asthlp.String),
			)),
		)
	}
	fn.AppendStmt(
		//	o, err := v.Object()
		//	if err != nil {
		//		return err
		//	}
		asthlp.Assign(
			asthlp.MakeVarNames(jsonObjectVarName, names.VarNameError),
			asthlp.Definition,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameJsonValue, "Object"))),
		),
		asthlp.If(asthlp.NotEqual(asthlp.NewIdent(names.VarNameError), asthlp.Nil), asthlp.Return(asthlp.NewIdent(names.VarNameError))),
	)

	if len(fields) > 0 {
		// var checkFields [1]int
		fn.AppendStmt(
			asthlp.Var(asthlp.VariableType(checkFieldsVarName, asthlp.ArrayType(asthlp.Int, asthlp.IntegerConstant(len(fields)).Expr()))),
		)
	}

	fn.AppendStmt(
		//	o.Visit(func(key []byte, _ *fastjson.Value) {
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector(jsonObjectVarName, "Visit")),
			visitFunc.Lit(),
		)),
		// return err
		asthlp.Return(asthlp.NewIdent(names.VarNameError)),
	)
	return fn.Decl()
}

// func (s *S) MarshalJSON() ([]byte, error) {
// 	var buf [128]byte
// 	return s.marshalAppend(buf[:0])
// }
func NewMarshalFunc(structName string, _ tags.StructTags) ast.Decl {
	return asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshal)).
		Comments(names.MethodNameMarshal+" serializes the structure with all its values into JSON format").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName)))).
		Results(
			asthlp.Field("", nil, asthlp.Byte),
			asthlp.Field("", nil, asthlp.ErrorType),
		).
		AppendStmt(
			// todo @menshenin calculate buffer lengthv
			asthlp.Var(asthlp.VariableType(names.VarNameBuf, asthlp.ArrayType(asthlp.Byte, asthlp.IntegerConstant(128).Expr()))),
			asthlp.Return(
				asthlp.Call(
					asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameReceiver, names.MethodNameAppend)),
					asthlp.Slice(names.VarNameBuf, nil, asthlp.IntegerConstant(0)),
				),
			),
		).Decl()
}

// func (s *S) MarshalAppend(dst []byte) ([]byte, error) {
func NewAppendJsonFunc(structName string, fields []*ast.Field, _ tags.StructTags) ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameAppend)).
		Comments(names.MethodNameAppend + " serializes all fields of the structure using a buffer")

	fn.AppendStmt(
		// var result = bytes.NewBuffer(dst)
		asthlp.Var(asthlp.VariableValue("result", asthlp.FreeExpression(asthlp.Call(
			asthlp.BytesNewBufferFn,
			ast.NewIdent("dst"),
		)))),
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
	)

	var body = []ast.Stmt{}

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
			{Text: "\n// " + names.MethodNameAppend + " serializes all fields of the structure using a buffer"},
		}},
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(names.VarNameReceiver)}, Type: &ast.StarExpr{X: ast.NewIdent(structName)}},
		}},
		Name: ast.NewIdent(names.MethodNameAppend),
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
