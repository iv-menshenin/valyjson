package codegen

import (
	"fmt"
	"go/ast"

	"github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

const (
	marshalObjectBufLen = 512
	marshalFieldBufLen  = 128
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
			asthlp.StringConstant("(root)").Expr(),
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

func NewMarshalFunc(structName string) ast.Decl {
	return asthlp.DeclareFunction(asthlp.NewIdent(names.MethodNameMarshal)).
		Comments("// "+names.MethodNameMarshal+" serializes the structure with all its values into JSON format.").
		Receiver(asthlp.Field(names.VarNameReceiver, nil, asthlp.Star(asthlp.NewIdent(structName)))).
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
