package codegen

import (
	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

type Transitive struct {
	TaggedRenderer
	tran ast.Expr
}

func NewTransitive(name string, tags []string, spec ast.Expr) *Transitive {
	return &Transitive{
		TaggedRenderer: TaggedRenderer{
			name: name,
			tags: tags,
		},
		tran: spec,
	}
}

func (t *Transitive) UnmarshalFunc() []ast.Decl {
	return NewUnmarshalFunc(t.name)
}

func (t *Transitive) ValidatorFunc() ast.Decl {
	return nil
}

func (t *Transitive) MarshalFunc() ast.Decl {
	return NewMarshalFunc(t.name)
}

func (t *Transitive) AppendJsonFunc() ast.Decl {
	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshalTo))
	fn.Comments("// " + names.MethodNameMarshalTo + " serializes all fields of the structure using a buffer.")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(t.name))))
	fn.Params(asthlp.Field(names.VarNameWriter, nil, asthlp.NewIdent("Writer")))
	fn.Results(
		asthlp.Field("", nil, asthlp.ErrorType),
	)

	//if s == nil {
	//	result.WriteString("null")
	//	return nil
	//}
	fn.AppendStmt(
		asthlp.If(
			asthlp.Equal(asthlp.NewIdent(names.VarNameReceiver), asthlp.Nil),
			asthlp.CallStmt(asthlp.Call(field.WriteStringFn, asthlp.StringConstant("null").Expr())),
			asthlp.Return(asthlp.Nil),
		),
	)

	sel, ok := t.tran.(*ast.SelectorExpr)
	if ok {
		if sel.Sel.Name == "UUID" {
			bufVar := asthlp.NewIdent("_uuid")
			return fn.AppendStmt(
				asthlp.Assign(
					asthlp.VarNames{bufVar, asthlp.NewIdent(names.VarNameError)},
					asthlp.Definition,
					asthlp.Call(asthlp.InlineFunc(
						asthlp.Selector(asthlp.ExpressionTypeConvert(asthlp.Star(asthlp.NewIdent(names.VarNameReceiver)), t.tran), "MarshalText"),
					)),
				),
				asthlp.If(asthlp.IsNil(asthlp.NewIdent(names.VarNameError)), asthlp.CallStmt(asthlp.Call(field.WriteBytesFn, bufVar))),
				asthlp.Return(asthlp.NewIdent(names.VarNameError)),
			).Decl()
		}
	}

	typed, ok := t.tran.(*ast.Ident)
	if ok {
		src := asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))
		switch typed.Name {
		case "int", "int8", "int16", "int32", "int64":
			fn.AppendStmt(asthlp.CallStmt(asthlp.Call(
				names.WriteInt64Func,
				asthlp.NewIdent(names.VarNameWriter),
				asthlp.ExpressionTypeConvert(src, asthlp.Int64),
			)))
			fn.AppendStmt(asthlp.Return(asthlp.Nil))
			return fn.Decl()

		case "uint", "uint8", "uint16", "uint32", "uint64":
			fn.AppendStmt(asthlp.CallStmt(asthlp.Call(
				names.WriteUint64Func,
				asthlp.NewIdent(names.VarNameWriter),
				asthlp.ExpressionTypeConvert(src, asthlp.UInt64),
			)))
			fn.AppendStmt(asthlp.Return(asthlp.Nil))
			return fn.Decl()

		case "float32", "float64":
			fn.AppendStmt(asthlp.CallStmt(asthlp.Call(
				names.WriteFloat64Func,
				asthlp.NewIdent(names.VarNameWriter),
				asthlp.ExpressionTypeConvert(src, asthlp.Float64),
			)))
			fn.AppendStmt(asthlp.Return(asthlp.Nil))
			return fn.Decl()

		case "string":
			fn.AppendStmt(asthlp.CallStmt(asthlp.Call(
				names.WriteStringFunc,
				asthlp.NewIdent(names.VarNameWriter),
				asthlp.ExpressionTypeConvert(src, asthlp.String),
			)))
			fn.AppendStmt(asthlp.Return(asthlp.Nil))
			return fn.Decl()

		case "rune":
			panic("unimplemented")

		case "bool":
			fn.AppendStmt(asthlp.IfElse(
				src,
				// result.WriteString
				asthlp.Block(asthlp.CallStmt(asthlp.Call(
					asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameWriter, "WriteString")),
					asthlp.StringConstant("true").Expr(),
				))),
				asthlp.Block(asthlp.CallStmt(asthlp.Call(
					asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameWriter, "WriteString")),
					asthlp.StringConstant("false").Expr(),
				))),
			))
			fn.AppendStmt(asthlp.Return(asthlp.Nil))
			return fn.Decl()
		}
	}

	fn.AppendStmt(asthlp.Return(
		asthlp.Call(
			asthlp.InlineFunc(asthlp.Selector(asthlp.VariableTypeConvert("s", asthlp.Star(t.tran)), names.MethodNameMarshalTo)),
			asthlp.NewIdent(names.VarNameWriter),
		),
	))
	return fn.Decl()
}

// FillerFunc generates function code that will fill in all fields of the structure with the fastjson.Value attribute
//
//	func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
//	    return (*StructElem)(s).FillFromJson(v, objPath)
//	}
func (t *Transitive) FillerFunc() ast.Decl {
	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameFill))
	fn.Comments("// " + names.MethodNameFill + " recursively fills the fields with fastjson.Value")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(t.name))))
	fn.Params(
		asthlp.Field(names.VarNameJsonValue, nil, asthlp.Star(names.FastJsonValue)),
		asthlp.Field(names.VarNameObjPath, nil, asthlp.String),
	)
	fn.Results(
		asthlp.Field(names.VarNameError, nil, asthlp.ErrorType),
	)

	typedFiller := field.NewFromType(t.tran, true).TypedValue(asthlp.NewIdent("_val"), names.VarNameJsonValue)
	if len(typedFiller) > 0 {
		fn.AppendStmt(typedFiller...)
		fn.AppendStmt(
			asthlp.If(asthlp.NotNil(asthlp.NewIdent(names.VarNameError)), asthlp.Return(asthlp.NewIdent(names.VarNameError))),
			asthlp.Assign(asthlp.VarNames{asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))}, asthlp.Assignment, asthlp.VariableTypeConvert("_val", asthlp.NewIdent(t.name))),
		)
		asthlp.NewIdent(names.VarNameReceiver)
		return fn.AppendStmt(asthlp.Return(asthlp.Nil)).Decl()
	}

	fn.AppendStmt(asthlp.Return(
		asthlp.Call(
			asthlp.InlineFunc(asthlp.Selector(asthlp.VariableTypeConvert("s", asthlp.Star(t.tran)), names.MethodNameFill)),
			asthlp.NewIdent(names.VarNameJsonValue),
			asthlp.NewIdent(names.VarNameObjPath),
		),
	))
	return fn.Decl()
}

func (t *Transitive) ZeroFunc() ast.Decl {
	var fn = asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameZero)).
		Comments("// " + names.MethodNameZero + " shows whether the object is an empty value.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, ast.NewIdent(t.name))).
		Results(asthlp.Field("", nil, asthlp.Bool))

	zero := helpers.ZeroValueOfT(t.tran)
	if zero != nil {
		fn.AppendStmt(
			asthlp.Return(asthlp.Equal(asthlp.NewIdent(names.VarNameReceiver), asthlp.ExpressionTypeConvert(zero, asthlp.NewIdent(t.name)))),
		)
	} else {
		// return s.IsZero()
		fn.AppendStmt(
			asthlp.Return(asthlp.Call(asthlp.InlineFunc(
				asthlp.Selector(asthlp.ExpressionTypeConvert(asthlp.NewIdent(names.VarNameReceiver), t.tran), names.MethodNameZero),
			))),
		)
	}

	return fn.Decl()
}
