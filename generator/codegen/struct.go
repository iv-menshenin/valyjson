package codegen

import (
	"go/ast"
	"go/token"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

type Struct struct {
	TaggedRenderer
	spec *ast.StructType
}

func NewStruct(name string, tags []string, spec *ast.StructType) *Struct {
	return &Struct{
		TaggedRenderer: TaggedRenderer{
			name: name,
			tags: tags,
		},
		spec: spec,
	}
}

func (s *Struct) UnmarshalFunc() []ast.Decl {
	return NewUnmarshalFunc(s.name)
}

// FillerFunc generates function code that will fill in all fields of the structure with the fastjson.Value attribute
//   func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) { ... }
func (s *Struct) FillerFunc() ast.Decl {
	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " recursively fills the fields with fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(s.name))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)
	if s.tags.StrictRules() {
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
	for _, f := range s.spec.Fields.List {
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

// ValidatorFunc generates a function declaration for validating a JSON object, taking into account the presence of fields.
//  func validate(v *fastjson.Value, objPath string) error {
func (s *Struct) ValidatorFunc() ast.Decl {
	fastJsonValue := asthlp.Star(names.FastJsonValue)
	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameValidate))
	fn.Comments("// " + names.MethodNameValidate + " checks for correct data structure")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(s.name))))
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
	for i, fieldSpec := range s.spec.Fields.List {
		var fieldTags tags.Tags
		if fieldSpec.Tag != nil {
			fieldTags = tags.Parse(fieldSpec.Tag.Value)
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
	if s.tags.StrictRules() {
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

	if len(s.spec.Fields.List) > 0 {
		// var checkFields [1]int
		fn.AppendStmt(
			asthlp.Var(asthlp.VariableType(checkFieldsVarName, asthlp.ArrayType(asthlp.Int, asthlp.IntegerConstant(len(s.spec.Fields.List)).Expr()))),
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

// MarshalFunc marshal
//   func (s *S) MarshalJSON() ([]byte, error) {
// 	  var buf [128]byte
// 	  return s.marshalAppend(buf[:0])
//   }
func (s *Struct) MarshalFunc() ast.Decl {
	return asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshal)).
		Comments("// "+names.MethodNameMarshal+" serializes the structure with all its values into JSON format.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(s.name)))).
		Results(
			asthlp.Field("", nil, asthlp.ArrayType(asthlp.Byte)),
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

// AppendJsonFunc produces MarshalAppend(dst []byte) ([]byte, error)
func (s *Struct) AppendJsonFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameAppend)).
		Comments("// "+names.MethodNameAppend+" serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(s.name)))).
		Params(asthlp.Field("dst", nil, asthlp.ArrayType(asthlp.Byte))).
		Results(
			asthlp.Field("", nil, asthlp.ArrayType(asthlp.Byte)),
			asthlp.Field("", nil, asthlp.ErrorType),
		)

	fn.AppendStmt(
		// var result = bytes.NewBuffer(dst)
		asthlp.Var(asthlp.VariableValue("result", asthlp.FreeExpression(asthlp.Call(
			asthlp.BytesNewBufferFn,
			ast.NewIdent("dst"),
		)))),
		// var (
		// 	err error
		// 	buf = make([]byte, 0, 128)
		// )
		asthlp.Var(
			asthlp.VariableType(names.VarNameError, asthlp.ErrorType),
			asthlp.VariableValue("buf", asthlp.FreeExpression(asthlp.Call(
				asthlp.MakeFn,
				asthlp.ArrayType(asthlp.Byte),
				asthlp.IntegerConstant(0).Expr(),
				asthlp.IntegerConstant(128).Expr(),
			))),
		),
		// result.WriteRune('{')
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector("result", "WriteRune")),
			asthlp.RuneConstant('{').Expr(),
		)),
	)

	for _, fld := range s.spec.Fields.List {
		fn.AppendStmt(jsonFieldStmts(fld)...)
	}

	fn.AppendStmt(
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
	return fn.Decl()
}

func jsonFieldStmts(fld *ast.Field) []ast.Stmt {
	if len(fld.Names) == 0 {
		factory := field.New(fld)
		name := ""
		switch t := fld.Type.(type) {

		case *ast.Ident:
			name = t.Name

		case *ast.SelectorExpr:
			name = t.Sel.Name
		}
		return factory.MarshalStatements(name)
	}
	var result []ast.Stmt
	factory := field.New(fld)
	for _, name := range fld.Names {
		result = append(result, factory.MarshalStatements(name.Name)...)
	}
	return result
}
