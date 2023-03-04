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
	return nil
}

func (t *Transitive) AppendJsonFunc() ast.Decl {
	return nil
}

// FillerFunc generates function code that will fill in all fields of the structure with the fastjson.Value attribute
//   func (s *Struct) FillFromJson(v *fastjson.Value, objPath string) (err error) {
//       return (*StructElem)(s).FillFromJson(v, objPath)
//   }
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
