package codegen

import (
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
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

//	func (m *MapTable) FillFromJson(v *fastjson.Value, objPath string) (err error) {
//		o, err := v.Object()
//		if err != nil {
//			return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
//		}
//		*m = make(map[string]TableOf, o.Len())
//		o.Visit(func(key []byte, v *fastjson.Value) {
//			if err != nil {
//				return
//			}
//			var value TableOf
//			err = value.FillFromJson(v, objPath+"tables.")
//			if err == nil {
//				(*m)[string(key)] = TableOf(value)
//			}
//		})
//		if err != nil {
//			return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
//		}
//		return nil
//	}
func (m *Map) FillFromFunc() ast.Decl {
	const (
		v = "v"
		o = "o"
	)
	var value = asthlp.NewIdent("value")
	var ifNullValue = asthlp.EmptyStmt()
	var valueAsValue = asthlp.ExpressionTypeConvert(value, m.spec.Value)
	if _, isStar := m.spec.Value.(*ast.StarExpr); isStar {
		valueAsValue = asthlp.Call(
			asthlp.InlineFunc(asthlp.ParenExpr(m.spec.Value)),
			asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("unsafe", "Pointer")),
				asthlp.Ref(value),
			),
		)
		ifNullValue = asthlp.If(
			helpers.MakeIfItsNullTypeCondition(),
			asthlp.Assign(
				asthlp.VarNames{
					asthlp.Index(
						asthlp.ParenExpr(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
						asthlp.FreeExpression(asthlp.VariableTypeConvert("key", m.spec.Key)),
					),
				},
				asthlp.Assignment,
				asthlp.Nil,
			),
			asthlp.ReturnEmpty(),
		)
	}

	valFactory := field.New(asthlp.Field("", asthlp.MakeTagsForField(map[string][]string{}), m.spec.Value))
	valFactory.DontCheckErr()

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " recursively fills the keys with fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(m.name))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)

	// 	if v.Type() == fastjson.TypeNull {
	//		return nil
	//	}
	fn.AppendStmt(
		asthlp.If(
			helpers.MakeIfItsNullTypeCondition(),
			asthlp.Return(asthlp.Nil),
		),
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
			asthlp.Return(asthlp.NewIdent(names.VarNameError)),
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
					ifNullValue,
				).
				AppendStmt(
					// fills one value
					field.IsNotEmpty(valFactory.TypedValue(value, "v", asthlp.VariableTypeConvert("key", asthlp.String)))...,
				).
				AppendStmt(
					// if err == nil {
					//   err = fmt.Errorf("error parsing '%s.%s' value: %w", objPath, string(key), err)
					//   return
					// }
					asthlp.If(
						asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
						asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
							asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
							asthlp.VariableTypeConvert("key", asthlp.String),
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
						valueAsValue,
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

func (m *Map) MarshalFunc() []ast.Decl {
	return NewMarshalFunc(m.name)
}

func (m *Map) MarshalToFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshalTo)).
		Comments("// " + names.MethodNameMarshalTo + " serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(m.name)))).
		Params(asthlp.Field(names.VarNameWriter, nil, asthlp.Star(asthlp.SimpleSelector("jwriter", "Writer")))).
		Results(
			asthlp.Field("", nil, asthlp.ErrorType),
		)

	fn.AppendStmt(
		// 	if s == nil || *s == nil {
		//		result.WriteString("null")
		//		return nil
		//	}
		asthlp.If(
			asthlp.Or(
				asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver)),
				asthlp.IsNil(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
			),
			// result.RawString("null")
			asthlp.CallStmt(asthlp.Call(field.RawStringFn, asthlp.StringConstant("null").Expr())),
			asthlp.Return(asthlp.Nil),
		),
		// var (
		// 	err    error
		//  filled bool
		// )
		field.NeedVars(),
		// result.RawByte('{')
		asthlp.CallStmt(asthlp.Call(
			field.RawByteFn,
			asthlp.RuneConstant('{').Expr(),
		)),
	)

	var valType = m.spec.Value
	dec := field.GetDecorExpr(valType)
	if i, ok := valType.(*ast.Ident); ok {
		valType = helpers.DenotedIdent(i)
	}
	errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+m.name+`" attribute %q: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
	ve := field.GetValueExtractor(valType, errExpr, dec)

	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.RawByte(',')
		//	}
		asthlp.If(field.WantCommaVar, asthlp.CallStmt(asthlp.Call(field.RawByteFn, asthlp.RuneConstant(',').Expr()))),
		// filled = true
		field.SetCommaVar,
		// result.String(string(_k))
		asthlp.CallStmt(asthlp.Call(field.StringFn, asthlp.VariableTypeConvert("_k", asthlp.String))),
		asthlp.CallStmt(asthlp.Call(field.RawByteFn, asthlp.RuneConstant(':').Expr())),
	}
	iterBlock = append(
		iterBlock,
		ve(asthlp.NewIdent("_v"))...,
	)

	fn.AppendStmt(asthlp.Range(
		true,
		"_k", "_v",
		asthlp.Star(asthlp.NewIdent(names.VarNameReceiver)),
		iterBlock...,
	))

	fn.AppendStmt(
		makeWriteAndReturn('}')...,
	)
	return fn.Decl()
}

func (m *Map) ZeroFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameZero)).
		Comments("// " + names.MethodNameZero + " shows whether the object is an empty value.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, ast.NewIdent(m.name))).
		Results(asthlp.Field("", nil, asthlp.Bool))

	// return len(s) == 0
	fn.AppendStmt(
		asthlp.Return(asthlp.Equal(asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(names.VarNameReceiver)), asthlp.Zero)),
	)
	return fn.Decl()
}

func (m *Map) ResetFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameReset)).
		Comments("// " + names.MethodNameReset + " resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, ast.NewIdent(m.name)))

	fn.AppendStmt(
		asthlp.Range(
			true,
			"k", "v",
			asthlp.NewIdent(names.VarNameReceiver),
			append(
				resetStmt(m.spec.Value, asthlp.NewIdent("v"), 0),
				asthlp.Assign(
					asthlp.VarNames{asthlp.Index(asthlp.NewIdent(names.VarNameReceiver), asthlp.FreeExpression(asthlp.NewIdent("k")))},
					asthlp.Assignment,
					asthlp.NewIdent("v"),
				),
			)...,
		),
	)

	return fn.Decl()
}
