package codegen

import (
	asthlp "github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"go/ast"
	"go/token"
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

func (m *Array) UnmarshalFunc() []ast.Decl {
	return NewUnmarshalFunc(m.Name())
}

func (m *Array) FillerFunc() ast.Decl {
	const (
		v = "v"
		a = "a"
		i = "i"
	)
	valFactory := field.New(asthlp.Field("", asthlp.MakeTagsForField(map[string][]string{}), m.spec.Elt))
	valFactory.DontCheckErr()

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " fills the array with the values recognized from fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(m.name))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)

	makeStmt := asthlp.EmptyStmt()
	if m.spec.Len == nil {
		//	*m = make(map[string]TableOf, len(a)
		makeStmt = asthlp.Assign(asthlp.VarNames{asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))}, asthlp.Assignment, asthlp.Call(
			asthlp.MakeFn,
			m.spec,
			asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(a)),
		))
	} else {
		//if len(*s) != len(a) {
		//	return fmt.Errorf("error parsing '%s', expected %d elsemens, got %d", objPath, len(*s), len(a))
		//}
		makeStmt = asthlp.If(
			asthlp.NotEqual(
				asthlp.Call(asthlp.LengthFn, asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
				asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(a)),
			),
			asthlp.Return(asthlp.Call(
				asthlp.FmtErrorfFn,
				asthlp.StringConstant("error parsing '%s', expected %d elements, got %d").Expr(),
				asthlp.NewIdent(names.VarNameObjPath),
				asthlp.Call(asthlp.LengthFn, asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
				asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(a)),
			)),
		)
	}

	fn.AppendStmt(
		//	a, err := v.Array()
		//	if err != nil {
		asthlp.Assign(
			asthlp.MakeVarNames(a, names.VarNameError),
			asthlp.Definition,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, "Array"))),
		),
		asthlp.If(
			asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
			// return fmt.Errorf("error parsing '%s' value: %w", objPath, err)
			asthlp.Return(asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant("error parsing '%s' value: %w").Expr(), asthlp.NewIdent(names.VarNameObjPath), asthlp.NewIdent(names.VarNameError))),
		),
		makeStmt,
		asthlp.Range(true, i, v, asthlp.NewIdent(a),
			append(
				valFactory.TypedValue(asthlp.NewIdent("value"), "v"),
				//if err != nil {
				//	return fmt.Errorf("error parsing '%s[%d]' value: %w", objPath, i, err)
				//}
				asthlp.If(
					asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
					asthlp.Return(asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant("error parsing '%s.[%d]' value: %w").Expr(), asthlp.NewIdent(names.VarNameObjPath), asthlp.NewIdent(i), asthlp.NewIdent(names.VarNameError))),
				),
				// (*s)[i] = test_extr.External(value)
				asthlp.Assign(
					asthlp.VarNames{
						asthlp.Index(
							asthlp.ParenExpr(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
							asthlp.FreeExpression(asthlp.NewIdent(i)),
						),
					},
					asthlp.Assignment,
					asthlp.VariableTypeConvert("value", m.spec.Elt),
				),
			)...,
		),
		asthlp.Return(asthlp.Nil),
	)

	return fn.Decl()
}

// TODO @menshenin

func (m *Array) ValidatorFunc() ast.Decl {
	return nil
}

func (m *Array) MarshalFunc() ast.Decl {
	return asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshal)).
		Comments("// "+names.MethodNameMarshal+" serializes the structure with all its values into JSON format.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(m.name)))).
		Results(
			asthlp.Field("", nil, asthlp.ArrayType(asthlp.Byte)),
			asthlp.Field("", nil, asthlp.ErrorType),
		).
		AppendStmt(
			// todo @menshenin calculate buffer lengthv
			asthlp.Var(asthlp.VariableType(names.VarNameBuf, asthlp.ArrayType(asthlp.Byte, asthlp.IntegerConstant(marshalObjectBufLen).Expr()))),
			asthlp.Return(
				asthlp.Call(
					asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameReceiver, names.MethodNameAppend)),
					asthlp.Slice(names.VarNameBuf, nil, asthlp.IntegerConstant(0)),
				),
			),
		).Decl()
}

// 	if s == nil || *s == nil {
//		return []byte("null"), nil
//	}
//	var (
//		err     error
//		_filled bool
//		buf     = make([]byte, 0, 128)
//		result  = bytes.NewBuffer(dst)
//	)
//	result.WriteRune('[')
//	for _, _v := range *s {
//		if _filled {
//			result.WriteRune(',')
//		}
//		_filled = true
//		buf = strconv.AppendInt(buf[:0], _v, 10)
//		result.Write(buf)
//	}
//	result.WriteRune(']')
//	return result.Bytes(), err
func (m *Array) AppendJsonFunc() ast.Decl {
	const filled = "_filled"
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameAppend)).
		Comments("// "+names.MethodNameAppend+" serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(m.name)))).
		Params(asthlp.Field("dst", nil, asthlp.ArrayType(asthlp.Byte))).
		Results(
			asthlp.Field("", nil, asthlp.ArrayType(asthlp.Byte)),
			asthlp.Field("", nil, asthlp.ErrorType),
		)

	if m.spec.Len == nil {
		fn.AppendStmt(
			// 	if s == nil || *s == nil {
			//		return []byte("null"), nil
			//	}
			asthlp.If(
				asthlp.Or(
					asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver)),
					asthlp.IsNil(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
				),
				asthlp.Return(asthlp.ExpressionTypeConvert(asthlp.StringConstant("null").Expr(), asthlp.ArrayType(asthlp.Byte)), asthlp.Nil),
			),
		)
	} else {
		fn.AppendStmt(
			// 	if s == nil {
			//		return []byte("null"), nil
			//	}
			asthlp.If(
				asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver)),
				asthlp.Return(asthlp.ExpressionTypeConvert(asthlp.StringConstant("null").Expr(), asthlp.ArrayType(asthlp.Byte)), asthlp.Nil),
			),
		)
	}

	fn.AppendStmt(
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
				asthlp.IntegerConstant(marshalFieldBufLen).Expr(),
			))),
			asthlp.VariableValue("result", asthlp.FreeExpression(asthlp.Call(
				asthlp.BytesNewBufferFn,
				ast.NewIdent("dst"),
			))),
		),
		// result.WriteRune('{')
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector("result", "WriteRune")),
			asthlp.RuneConstant('[').Expr(),
		)),
	)

	errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+m.name+`" value at position %d: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
	ve := field.GetValueExtractor(denotedType(m.spec.Elt), errExpr)

	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteRune(',')
		//	}
		asthlp.If(asthlp.NewIdent(filled), asthlp.CallStmt(asthlp.Call(field.WriteRuneFn, asthlp.RuneConstant(',').Expr()))),
		// filled = true
		asthlp.Assign(asthlp.MakeVarNames(filled), asthlp.Assignment, asthlp.True),
		// _k = _k
		asthlp.Assign(asthlp.MakeVarNames("_k"), asthlp.Assignment, asthlp.NewIdent("_k")),
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
		// result.WriteRune(']')
		// return result.Bytes(), err
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteRune")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.CHAR, Value: "']'"}},
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
