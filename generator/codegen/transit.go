package codegen

import (
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

	// return s.Zero()
	fn.AppendStmt(
		asthlp.Return(asthlp.Call(asthlp.InlineFunc(
			asthlp.Selector(asthlp.ExpressionTypeConvert(asthlp.NewIdent(names.VarNameReceiver), t.tran), names.MethodNameZero),
		))),
	)
	return fn.Decl()
}
