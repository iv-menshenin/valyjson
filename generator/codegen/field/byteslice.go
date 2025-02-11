package field

import (
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

func ByteSliceFillFrom(src, dest ast.Expr, spec *ast.ArrayType) []ast.Stmt {
	const (
		_b = "b"
		_n = "n"
	)
	// array cases
	decodeDest := asthlp.SliceExpr(dest, nil, nil) // (*s)[:]
	//if n != len(s) {
	//	return errors.New("incomplete data")
	//}
	finalAlign := asthlp.If(
		asthlp.NotEqual(asthlp.NewIdent(_n), asthlp.Call(asthlp.LengthFn, dest)),
		asthlp.Return(asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector("errors", "New")), asthlp.StringConstant("incomplete data").Expr())),
	)
	makeStmt := asthlp.EmptyStmt()
	if spec.Len == nil {
		//	dest = make([]byte, base64.StdEncoding.DecodedLen(len(b)))
		makeStmt = asthlp.Assign(asthlp.VarNames{dest}, asthlp.Assignment, asthlp.Call(
			asthlp.MakeFn,
			spec,
			asthlp.Call(
				asthlp.InlineFunc(asthlp.Selector(asthlp.SimpleSelector("base64", "StdEncoding"), "DecodedLen")),
				asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(_b)),
			),
		))
		decodeDest = dest
		//	 dest = dest[:n]
		finalAlign = asthlp.Assign(
			asthlp.VarNames{dest},
			asthlp.Assignment,
			asthlp.SliceExpr(dest, nil, asthlp.FreeExpression(asthlp.NewIdent(_n))),
		)
	}
	return []ast.Stmt{
		// 	if v.Type() != fastjson.TypeNull {
		asthlp.If(
			helpers.MakeIfItsNotNullTypeCondition(src),
			//  // slice of bytes in JSON format is implemented via BASE64 string
			asthlp.CommentStmt("slice of bytes in JSON format is implemented via BASE64 string"),
			//  b, err := v.StringBytes()
			//	if err != nil {
			asthlp.Assign(
				asthlp.MakeVarNames(_b, names.VarNameError),
				asthlp.Definition,
				asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "StringBytes"))),
			),
			asthlp.If(
				asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
				asthlp.Return(asthlp.NewIdent(names.VarNameError)),
			),
			makeStmt,
			//	 n, err := base64.StdEncoding.Decode(*s, b)
			//	 if err != nil {
			//	    return err
			//	 }
			asthlp.Assign(
				asthlp.MakeVarNames(_n, names.VarNameError),
				asthlp.Definition,
				asthlp.Call(
					asthlp.InlineFunc(asthlp.Selector(asthlp.SimpleSelector("base64", "StdEncoding"), "Decode")),
					decodeDest,
					asthlp.NewIdent(_b),
				),
			),
			asthlp.If(
				asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
				asthlp.Return(asthlp.NewIdent(names.VarNameError)),
			),
			finalAlign,
		),
	}
}

func ByteSliceMarshal(src ast.Expr, spec *ast.ArrayType) []ast.Stmt {
	const _buf = "buf"
	encodeSrc := src
	if spec.Len != nil {
		encodeSrc = asthlp.SliceExpr(src, nil, nil)
	}
	return []ast.Stmt{
		// // slice of bytes in JSON format is implemented via BASE64 string
		asthlp.CommentStmt("slice of bytes in JSON format is implemented via BASE64 string"),
		// buf := make([]byte, base64.StdEncoding.EncodedLen(len(*s)))
		asthlp.Assign(asthlp.VarNames{asthlp.NewIdent(_buf)}, asthlp.Definition, asthlp.Call(
			asthlp.MakeFn,
			asthlp.ArrayType(spec.Elt),
			asthlp.Call(
				asthlp.InlineFunc(asthlp.Selector(asthlp.SimpleSelector("base64", "StdEncoding"), "EncodedLen")),
				asthlp.Call(asthlp.LengthFn, src),
			),
		)),
		// base64.StdEncoding.Encode(buf, *s)
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.Selector(asthlp.SimpleSelector("base64", "StdEncoding"), "Encode")),
			asthlp.NewIdent(_buf),
			encodeSrc,
		)),
		// result.String(string(buf))
		asthlp.CallStmt(asthlp.Call(StringFn, asthlp.VariableTypeConvert(_buf, asthlp.String))),
	}
}
