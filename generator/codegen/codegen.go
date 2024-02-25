package codegen

import (
	"fmt"
	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"go/ast"

	"github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/field"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

type TaggedRenderer struct {
	name string
	tags tags.StructTags
}

func (t *TaggedRenderer) Name() string {
	return t.name
}

func (t *TaggedRenderer) Tags() tags.StructTags {
	return t.tags
}

// NewUnmarshalFunc generates code for unmarshalling, function that parses the JSON object and fills all the structure properties.
//
//	func (s *Struct) UnmarshalJSON(data []byte) error {
func NewUnmarshalFunc(structName string) []ast.Decl {
	const parser = "parser"
	poolName := fmt.Sprintf("%s%s", names.VarPrefixPool, structName)

	fn := asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameUnmarshal))
	fn.Comments("// " + names.MethodNameUnmarshal + " implements json.Unmarshaler")
	fn.Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName))))
	fn.Params(asthlp.Field(names.VarNameData, nil, asthlp.ArrayType(asthlp.Byte)))
	fn.Results(asthlp.Field("", nil, asthlp.ErrorType))
	fn.AppendStmt(
		asthlp.Assign(asthlp.MakeVarNames(parser), asthlp.Definition, asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(poolName, "Get")))),
		asthlp.CommentStmt("parses data containing JSON"),
		asthlp.Assign(
			asthlp.MakeVarNames(names.VarNameJsonValue, names.VarNameError),
			asthlp.Definition,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(parser, "ParseBytes")), ast.NewIdent(names.VarNameData)),
		),
		asthlp.If(asthlp.NotNil(asthlp.NewIdent(names.VarNameError)), asthlp.Return(asthlp.NewIdent(names.VarNameError))),
		asthlp.DeferCall(asthlp.InlineFunc(asthlp.SimpleSelector(poolName, "Put")), ast.NewIdent(parser)),
		//	return s.FillFromJson(v, "")
		asthlp.Return(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameReceiver, names.MethodNameFill)),
			ast.NewIdent(names.VarNameJsonValue),
		)),
	)

	return []ast.Decl{
		asthlp.DeclareVariable().
			Comments("// " + poolName + " used for pooling Parsers for " + structName + " JSONs.").
			AppendSpec(asthlp.VariableType(poolName, names.FastJsonParserPool)).
			Decl(),
		fn.Decl(),
	}
}

//		func (s *TestStr01) MarshalJSON() ([]byte, error) {
//		var result jwriter.Writer
//		if err := s.MarshalTo(&result); err != nil {
//			return nil, err
//		}
//		return result.BuildBytes()
//	}
func NewMarshalFunc(structName string) []ast.Decl {
	return []ast.Decl{
		asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshal)).
			Comments("// "+names.MethodNameMarshal+" serializes the structure with all its values into JSON format.").
			Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName)))).
			Results(
				asthlp.Field("", nil, asthlp.ArrayType(asthlp.Byte)),
				asthlp.Field("", nil, asthlp.ErrorType),
			).
			AppendStmt(
				asthlp.Var(
					asthlp.VariableType(names.VarNameWriter, asthlp.SimpleSelector("jwriter", "Writer")),
				),
				asthlp.IfInit(
					asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Definition, asthlp.Call(
						asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameReceiver, names.MethodNameMarshalTo)),
						asthlp.Ref(asthlp.NewIdent(names.VarNameWriter)),
					)),
					asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
					asthlp.Return(asthlp.Nil, asthlp.NewIdent(names.VarNameError)),
				),
				asthlp.Return(
					asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameWriter, "BuildBytes"))),
				),
			).Decl(),
	}
}

func makeWriteAndReturn(r rune) []ast.Stmt {
	return []ast.Stmt{
		// result.RawByte('}')
		asthlp.CallStmt(asthlp.Call(
			field.RawByteFn, asthlp.RuneConstant(r).Expr(),
		)),
		// err = result.Error
		asthlp.Assign(
			asthlp.MakeVarNames(names.VarNameError),
			asthlp.Assignment,
			asthlp.SimpleSelector(names.VarNameWriter, "Error"),
		),
		// return err
		asthlp.Return(
			asthlp.NewIdent(names.VarNameError),
		),
	}
}

func resetStmt(t, name ast.Expr) ast.Stmt {
	switch tt := t.(type) {
	case *ast.ArrayType:
		if tt.Len == nil {
			return asthlp.Assign(asthlp.VarNames{name}, asthlp.Assignment, asthlp.SliceExpr(name, nil, asthlp.IntegerConstant(0)))
		} else {
			return asthlp.Assign(asthlp.VarNames{name}, asthlp.Assignment, asthlp.StructLiteral(tt).Expr())
		}

	case *ast.SelectorExpr:
		// uuid.Nil
		if tt.Sel.Name == "UUID" {
			return asthlp.Assign(asthlp.VarNames{name}, asthlp.Assignment, asthlp.SimpleSelector("uuid", "Nil"))
		}
		if tt.Sel.Name == "Time" {
			return asthlp.Assign(asthlp.VarNames{name}, asthlp.Assignment, asthlp.StructLiteral(tt).Expr())
		}
	}

	d := helpers.DenotedType(t)
	zero := helpers.ZeroValueOfT(d)
	if zero != nil {
		// v.Field = 0
		if t != d {
			return asthlp.Assign(asthlp.VarNames{name}, asthlp.Assignment, asthlp.ExpressionTypeConvert(zero, t))
		}
		return asthlp.Assign(asthlp.VarNames{name}, asthlp.Assignment, zero)
	}

	// v.Field.Reset()
	return asthlp.CallStmt(asthlp.Call(
		asthlp.InlineFunc(asthlp.Selector(name, names.MethodNameReset)),
	))
}
