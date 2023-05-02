package codegen

import (
	asthlp "github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"go/ast"
)

type (
	Array struct {
		TaggedRenderer
		spec *ast.ArrayType
	}
)

func NewArray(name string, tags []string, spec *ast.ArrayType) *Array {
	return &Array{
		TaggedRenderer: TaggedRenderer{
			name: name,
			tags: tags,
		},
		spec: spec,
	}
}

func (a *Array) UnmarshalFunc() []ast.Decl {
	return NewUnmarshalFunc(a.Name())
}

func (a *Array) FillerFunc() ast.Decl {
	const (
		_v = "v"
		_a = "a"
		_i = "i"
	)
	valFactory := field.New(asthlp.Field("", asthlp.MakeTagsForField(map[string][]string{}), a.spec.Elt))
	valFactory.DontCheckErr()

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " fills the array with the values recognized from fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(a.name))))
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

	makeStmt := asthlp.EmptyStmt()
	if a.spec.Len == nil {
		//	*m = make(map[string]TableOf, len(a)
		makeStmt = asthlp.Assign(asthlp.VarNames{asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))}, asthlp.Assignment, asthlp.Call(
			asthlp.MakeFn,
			a.spec,
			asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(_a)),
		))
	} else {
		//if len(*s) != len(a) {
		//	return fmt.Errorf("error parsing '%s', expected %d elsemens, got %d", objPath, len(*s), len(a))
		//}
		makeStmt = asthlp.If(
			asthlp.NotEqual(
				asthlp.Call(asthlp.LengthFn, asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
				asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(_a)),
			),
			asthlp.Return(asthlp.Call(
				asthlp.FmtErrorfFn,
				asthlp.StringConstant("expected %d elements, got %d").Expr(),
				asthlp.Call(asthlp.LengthFn, asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
				asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(_a)),
			)),
		)
	}

	fn.AppendStmt(
		//	a, err := v.Array()
		//	if err != nil {
		asthlp.Assign(
			asthlp.MakeVarNames(_a, names.VarNameError),
			asthlp.Definition,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(_v, "Array"))),
		),
		asthlp.If(
			asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
			asthlp.Return(asthlp.NewIdent(names.VarNameError)),
		),
		makeStmt,
		asthlp.Range(true, _i, _v, asthlp.NewIdent(_a),
			append(
				field.IsNotEmpty(valFactory.TypedValue(asthlp.NewIdent("value"), _v, asthlp.Call(asthlp.StrconvItoaFn, asthlp.NewIdent(_i)))),
				//if err != nil {
				//	return fmt.Errorf("error parsing '%s[%d]' value: %w", objPath, i, err)
				//}
				asthlp.If(
					asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
					asthlp.Return(
						asthlp.Call(
							asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
							asthlp.Call(
								asthlp.FmtSprintfFn,
								asthlp.StringConstant("%d").Expr(),
								asthlp.NewIdent(_i),
							),
							asthlp.NewIdent(names.VarNameError),
						),
					),
				),
				// (*s)[i] = test_extr.External(value)
				asthlp.Assign(
					asthlp.VarNames{
						asthlp.Index(
							asthlp.ParenExpr(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
							asthlp.FreeExpression(asthlp.NewIdent(_i)),
						),
					},
					asthlp.Assignment,
					asthlp.VariableTypeConvert("value", a.spec.Elt),
				),
			)...,
		),
		asthlp.Return(asthlp.Nil),
	)

	return fn.Decl()
}

// TODO @menshenin

func (a *Array) ValidatorFunc() ast.Decl {
	return nil
}

func (a *Array) MarshalFunc() []ast.Decl {
	return NewMarshalFunc(a.name)
}

//	if s == nil || *s == nil {
//		return []byte("null"), nil
//	}
//
// var (
//
//	err     error
//	wantComma bool
//	buf     = make([]byte, 0, 128)
//	result  = bytes.NewBuffer(dst)
//
// )
// result.WriteRune('[')
//
//	for _, _v := range *s {
//		if wantComma {
//			result.WriteRune(',')
//		}
//		wantComma = true
//		buf = strconv.AppendInt(buf[:0], _v, 10)
//		result.Write(buf)
//	}
//
// result.WriteRune(']')
// return result.Bytes(), err
func (a *Array) AppendJsonFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshalTo)).
		Comments("// " + names.MethodNameMarshalTo + " serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(a.name)))).
		Params(asthlp.Field(names.VarNameWriter, nil, asthlp.NewIdent("Writer"))).
		Results(asthlp.Field("", nil, asthlp.ErrorType))

	var nilCondition = asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver))
	if a.spec.Len == nil {
		nilCondition = asthlp.Or(
			asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver)),
			asthlp.IsNil(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
		)
	}
	fn.AppendStmt(
		// 	if s == nil || *s == nil {
		//		result.WriteString("null")
		//		return nil
		//	}
		asthlp.If(
			nilCondition,
			// result.WriteString("null")
			asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant("null").Expr())),
			asthlp.Return(asthlp.Nil),
		),
	)

	fn.AppendStmt(
		// var (
		// 	err      error
		//  filled bool
		// )
		field.NeedVars(),
		// result.WriteString("{")
		asthlp.CallStmt(asthlp.Call(
			field.WriteStringFn,
			asthlp.StringConstant("[").Expr(),
		)),
	)

	errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+a.name+`" value at position %d: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
	ve := field.GetValueExtractor(denotedType(a.spec.Elt), errExpr)

	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteString(",")
		//	}
		asthlp.If(field.WantCommaVar, asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant(",").Expr()))),
		// filled = true
		field.SetCommaVar,
		// _k = _k
		asthlp.Assign(asthlp.MakeVarNames("_k"), asthlp.Assignment, asthlp.NewIdent("_k")),
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
		makeWriteAndReturn("]")...,
	)
	return fn.Decl()
}

func (a *Array) ZeroFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameZero)).
		Comments("// " + names.MethodNameZero + " shows whether the object is an empty value.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, ast.NewIdent(a.name))).
		Results(asthlp.Field("", nil, asthlp.Bool))

	if a.spec.Len == nil {
		// return len(s) == 0
		fn.AppendStmt(
			asthlp.Return(asthlp.Equal(asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(names.VarNameReceiver)), asthlp.Zero)),
		)
	} else {
		//	for _, _v := range s {
		//		if _v != 0 {
		//			return false
		//		{
		//	}
		//	return true
		var isZero ast.Expr
		if zero := helpers.ZeroValueOfT(a.spec.Elt); zero != nil {
			isZero = asthlp.NotEqual(asthlp.NewIdent("_v"), zero)
		} else {
			isZero = asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector("_v", names.MethodNameZero)))
		}
		fn.AppendStmt(
			asthlp.Range(
				true, "_", "_v", asthlp.NewIdent(names.VarNameReceiver),
				asthlp.If(isZero, asthlp.Return(asthlp.False)),
			),
			asthlp.Return(asthlp.True),
		)
	}
	return fn.Decl()
}
