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

//
//func (m *MapTable) FillFromJson(v *fastjson.Value, objPath string) (err error) {
//	o, err := v.Object()
//	if err != nil {
//		return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
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
//		return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
//	}
//	return nil
//}
func (m *Map) FillerFunc() ast.Decl {
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
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
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
			// return fmt.Errorf("error parsing '%s.tables' value: %w", objPath, err)
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
					ifNullValue,
				).
				AppendStmt(
					// fills one value
					valFactory.TypedValue(value, "v")...,
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

func (m *Map) MarshalFunc() ast.Decl {
	return NewMarshalFunc(m.name)
}

func (m *Map) AppendJsonFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshalTo)).
		Comments("// " + names.MethodNameMarshalTo + " serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(m.name)))).
		Params(asthlp.Field(names.VarNameWriter, nil, asthlp.NewIdent("Writer"))).
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
			// result.WriteString("null")
			asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant("null").Expr())),
			asthlp.Return(asthlp.Nil),
		),
		// var (
		// 	err    error
		//  filled bool
		// )
		field.NeedVars(),
		// result.WriteString("{")
		asthlp.CallStmt(asthlp.Call(
			field.WriteStringFn,
			asthlp.StringConstant("{").Expr(),
		)),
	)

	errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+m.name+`" attribute %q: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
	ve := field.GetValueExtractor(denotedType(m.spec.Value), errExpr)

	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteString(",")
		//	}
		asthlp.If(field.WantCommaVar, asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant(",").Expr()))),
		// filled = true
		field.SetCommaVar,
		// result.WriteString(`"`)
		// result.WriteString(string(_k))
		// result.WriteString(`":`)
		asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant(`"`).Expr())),
		asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.VariableTypeConvert("_k", asthlp.String))),
		asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant(`":`).Expr())),
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
		makeWriteAndReturn("}")...,
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
