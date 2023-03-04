package codegen

import (
	"go/ast"
	"go/token"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

type (
	Map struct {
		TaggedRenderer
		spec *ast.MapType
	}
)

func NewMap(name string, tags []string, spec *ast.MapType) *Map {
	return &Map{
		TaggedRenderer: TaggedRenderer{
			name: name,
			tags: tags,
		},
		spec: spec,
	}
}

func (m *Map) UnmarshalFunc() []ast.Decl {
	return NewUnmarshalFunc(m.Name())
}

//
//func (m *MapTable) FillFromJson(v *fastjson.Value, objPath string) (err error) {
//	o, err := v.Object()
//	if err != nil {
//		return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
//	}
//	*m = make(map[string]TableOf, o.Len())
//	o.Visit(func(key []byte, v *fastjson.Value) {
//		if err != nil {
//			return
//		}
//		var value TableOf
//		err = value.FillFromJson(v, objPath+"tables.")
//		if err == nil {
//			(*m)[string(key)] = TableOf(value)
//		}
//	})
//	if err != nil {
//		return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
//	}
//	return nil
//}
func (m *Map) FillerFunc() ast.Decl {
	const (
		v = "v"
		o = "o"
	)

	valFactory := field.New(asthlp.Field("", asthlp.MakeTagsForField(map[string][]string{}), m.spec.Value))
	valFactory.DontCheckErr()

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " recursively fills the keys with fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(m.name))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)

	fn.AppendStmt(
		//	o, err := v.Object()
		//	if err != nil {
		asthlp.Assign(
			asthlp.MakeVarNames(o, names.VarNameError),
			asthlp.Definition,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, "Object"))),
		),
		asthlp.If(
			asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
			// return fmt.Errorf("error parsing '%stables' value: %w", objPath, err)
			asthlp.Return(asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant("error parsing '%s' value: %w").Expr(), asthlp.NewIdent(names.VarNameObjPath), asthlp.NewIdent(names.VarNameError))),
		),
		//	*m = make(map[string]TableOf, o.Len())
		asthlp.Assign(asthlp.VarNames{asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))}, asthlp.Assignment, asthlp.Call(
			asthlp.MakeFn,
			m.spec,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(o, "Len"))),
		)),
		//	o.Visit(func(key []byte, v *fastjson.Value) {
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector(o, "Visit")),
			asthlp.DeclareFunction(nil).
				Params(
					asthlp.Field("key", nil, asthlp.ArrayType(asthlp.Byte)),
					asthlp.Field(v, nil, asthlp.Star(names.FastJsonValue)),
				).
				AppendStmt(
					// if err != nil {
					//   return
					// }
					asthlp.If(asthlp.NotNil(asthlp.NewIdent(names.VarNameError)), asthlp.ReturnEmpty()),
				).
				AppendStmt(
					// fills one value
					valFactory.TypedValue(asthlp.NewIdent("value"), "v")...,
				).
				AppendStmt(
					// if err == nil {
					//   err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
					//   return
					// }
					asthlp.If(
						asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
						asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
							asthlp.FmtErrorfFn,
							asthlp.StringConstant("error parsing '%s.%s' value: %w").Expr(),
							asthlp.NewIdent(names.VarNameObjPath),
							asthlp.VariableTypeConvert("key", m.spec.Key),
							asthlp.NewIdent(names.VarNameError),
						)),
						asthlp.ReturnEmpty(),
					),
					asthlp.Assign(
						asthlp.VarNames{
							asthlp.Index(
								asthlp.ParenExpr(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
								asthlp.FreeExpression(asthlp.VariableTypeConvert("key", m.spec.Key)),
							),
						},
						asthlp.Assignment,
						asthlp.VariableTypeConvert("value", m.spec.Value),
					),
				).
				Lit(),
		)),
		// return err
		asthlp.Return(asthlp.NewIdent(names.VarNameError)),
	)

	return fn.Decl()
}

// TODO @menshenin

func (m *Map) ValidatorFunc() ast.Decl {
	return nil
}

func (m *Map) MarshalFunc() ast.Decl {
	return asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshal)).
		Comments("// "+names.MethodNameMarshal+" serializes the structure with all its values into JSON format.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(m.name)))).
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

func (m *Map) AppendJsonFunc() ast.Decl {
	const filled = "_filled"
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameAppend)).
		Comments("// "+names.MethodNameAppend+" serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(m.name)))).
		Params(asthlp.Field("dst", nil, asthlp.ArrayType(asthlp.Byte))).
		Results(
			asthlp.Field("", nil, asthlp.ArrayType(asthlp.Byte)),
			asthlp.Field("", nil, asthlp.ErrorType),
		)

	fn.AppendStmt(
		// 	if s == nil {
		//		return []byte("null"), nil
		//	}
		asthlp.If(
			asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver)),
			asthlp.Return(asthlp.ExpressionTypeConvert(asthlp.StringConstant("null").Expr(), asthlp.ArrayType(asthlp.Byte)), asthlp.Nil),
		),
		// var (
		// 	err      error
		//  filled bool
		// 	buf    = make([]byte, 0, 128)
		// 	result = bytes.NewBuffer(dst)
		// )
		asthlp.Var(
			asthlp.VariableType(names.VarNameError, asthlp.ErrorType),
			asthlp.VariableType(filled, asthlp.Bool),
			asthlp.VariableValue("buf", asthlp.FreeExpression(asthlp.Call(
				asthlp.MakeFn,
				asthlp.ArrayType(asthlp.Byte),
				asthlp.IntegerConstant(0).Expr(),
				asthlp.IntegerConstant(128).Expr(),
			))),
			asthlp.VariableValue("result", asthlp.FreeExpression(asthlp.Call(
				asthlp.BytesNewBufferFn,
				ast.NewIdent("dst"),
			))),
		),
		// result.WriteRune('{')
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector("result", "WriteRune")),
			asthlp.RuneConstant('{').Expr(),
		)),
	)

	errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+m.name+`" attribute %q: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
	ve := field.GetValueExtractor(denotedType(m.spec.Value), errExpr)

	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteRune(',')
		//	}
		asthlp.If(asthlp.NewIdent(filled), asthlp.CallStmt(asthlp.Call(field.WriteRuneFn, asthlp.RuneConstant(',').Expr()))),
		// filled = true
		asthlp.Assign(asthlp.MakeVarNames(filled), asthlp.Assignment, asthlp.True),
		// result.WriteRune('"')
		// result.WriteString(string(_k))
		// result.WriteString(`":`)
		asthlp.CallStmt(asthlp.Call(field.WriteRuneFn, asthlp.RuneConstant('"').Expr())),
		asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.VariableTypeConvert("_k", asthlp.String))),
		asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant(`":`).Expr())),
	}
	iterBlock = append(
		iterBlock,
		ve(asthlp.NewIdent("_v"))...,
	)
	iterBlock = append(
		iterBlock,
		// result.Write(buf)
		asthlp.CallStmt(asthlp.Call(field.WriteBytesFn, field.BufVar)),
	)

	fn.AppendStmt(asthlp.Range(
		true,
		"_k", "_v",
		asthlp.Star(asthlp.NewIdent(names.VarNameReceiver)),
		iterBlock...,
	))

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

func denotedType(t ast.Expr) ast.Expr {
	i, ok := t.(*ast.Ident)
	if !ok {
		return t
	}
	if i.Obj != nil {
		ts, ok := i.Obj.Decl.(*ast.TypeSpec)
		if ok {
			return ts.Type
		}
	}
	return i
}
