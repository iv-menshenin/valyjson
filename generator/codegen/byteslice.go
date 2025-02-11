package codegen

import (
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

type (
	ByteSlice struct {
		TaggedRenderer
		spec *ast.ArrayType
	}
)

func NewByteSlice(name string, tags []string, spec *ast.ArrayType) *ByteSlice {
	return &ByteSlice{
		TaggedRenderer: TaggedRenderer{
			name: name,
			tags: tags,
		},
		spec: spec,
	}
}

func (a *ByteSlice) UnmarshalFunc() []ast.Decl {
	return NewUnmarshalFunc(a.Name())
}

//	if v.Type() == fastjson.TypeNull {
//		return nil
//	}
//
//	 // slice of bytes in JSON format is implemented via BASE64 string
//	 b, err := v.StringBytes()
//		if err != nil {
//			return err
//		}
//	 *s = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
//	 n, err := base64.StdEncoding.Decode(*s, b)
//	 if err != nil {
//	    return err
//	 }
//	 *s = (*s)[:n]
//	 return nil
func (a *ByteSlice) FillFromFunc() ast.Decl {
	valFactory := field.New(asthlp.Field("", asthlp.MakeTagsForField(map[string][]string{}), a.spec.Elt))
	valFactory.DontCheckErr()

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " fills the byteslice with the values recognized from fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(a.name))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)
	fn.AppendStmt(field.ByteSliceFillFrom(asthlp.NewIdent(names.VarNameJsonValue), asthlp.Star(asthlp.NewIdent(names.VarNameReceiver)), a.spec)...)
	fn.AppendStmt(asthlp.ReturnEmpty())

	return fn.Decl()
}

// TODO @menshenin

func (a *ByteSlice) ValidatorFunc() ast.Decl {
	return nil
}

func (a *ByteSlice) MarshalFunc() []ast.Decl {
	return NewMarshalFunc(a.name)
}

func (a *ByteSlice) MarshalToFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshalTo)).
		Comments("// " + names.MethodNameMarshalTo + " serializes all fields of the structure using a buffer.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(a.name)))).
		Params(asthlp.Field(names.VarNameWriter, nil, asthlp.Star(asthlp.SimpleSelector("jwriter", "Writer")))).
		Results(asthlp.Field("", nil, asthlp.ErrorType))

	fn.AppendStmt(
		// there is no pointer
		asthlp.If(
			asthlp.IsNil(asthlp.NewIdent(names.VarNameReceiver)),
			asthlp.CallStmt(asthlp.Call(field.RawStringFn, asthlp.StringConstant("null").Expr())),
			asthlp.Return(asthlp.Nil),
		),
	)
	if a.spec.Len == nil {
		fn.AppendStmt(
			// there is no value
			asthlp.If(
				asthlp.IsNil(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))),
				asthlp.CallStmt(asthlp.Call(field.RawStringFn, asthlp.StringConstant("null").Expr())),
				asthlp.Return(asthlp.Nil),
			),
		)
	}
	fn.AppendStmt(field.ByteSliceMarshal(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver)), a.spec)...)
	fn.AppendStmt(asthlp.Return(asthlp.Nil))

	return fn.Decl()
}

func (a *ByteSlice) ZeroFunc() ast.Decl {
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

func (a *ByteSlice) ResetFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameReset)).
		Comments("// " + names.MethodNameReset + " resets the values of all fields of the structure to their initial states, defined by default for the data type of each field.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(ast.NewIdent(a.name))))

	if a.spec.Len != nil {
		fn.AppendStmt(asthlp.Assign(asthlp.VarNames{asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))}, asthlp.Assignment, asthlp.StructLiteral(ast.NewIdent(a.name)).Expr()))
	} else {
		switch helpers.DenotedType(a.spec.Elt).(type) {
		case *ast.StructType:
			sliceVal := asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))
			fn.AppendStmt(
				asthlp.Range(
					true, "i", "", sliceVal,
					asthlp.CallStmt(asthlp.Call(asthlp.InlineFunc(asthlp.Selector(asthlp.Index(sliceVal, asthlp.FreeExpression(asthlp.NewIdent("i"))), names.MethodNameReset)))),
				),
			)
		}
		fn.AppendStmt(
			asthlp.Assign(
				asthlp.VarNames{asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))},
				asthlp.Assignment,
				asthlp.SliceExpr(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver)), nil, asthlp.IntegerConstant(0)),
			),
		)
	}

	return fn.Decl()
}
