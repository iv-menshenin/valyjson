package codegen

import (
	"go/ast"

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
								&ast.ParenExpr{X: asthlp.Star(asthlp.NewIdent(names.VarNameReceiver))},
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

func (m *Map) ValidatorFunc() ast.Decl {
	return nil
}

func (m *Map) MarshalFunc() ast.Decl {
	return nil
}

func (m *Map) AppendJsonFunc() ast.Decl {
	return nil
}
