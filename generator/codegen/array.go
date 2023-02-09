package codegen

import (
	asthlp "github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/field"
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
					asthlp.Return(asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant("error parsing '%s[%d]' value: %w").Expr(), asthlp.NewIdent(names.VarNameObjPath), asthlp.NewIdent(i), asthlp.NewIdent(names.VarNameError))),
				),
				// (*s)[i] = test_extr.External(value)
				asthlp.Assign(
					asthlp.VarNames{
						asthlp.Index(
							&ast.ParenExpr{X: asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))},
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

func (m *Array) ValidatorFunc() ast.Decl {
	return nil
}

func (m *Array) MarshalFunc() ast.Decl {
	return nil
}

func (m *Array) AppendJsonFunc() ast.Decl {
	return nil
}
