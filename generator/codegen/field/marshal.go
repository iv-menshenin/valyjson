package field

import (
	"fmt"
	"go/ast"
	"go/token"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
)

// putCommaFirst puts comma
//  result.Write([]byte{','})
var putCommaFirst = asthlp.CallStmt(asthlp.Call(
	WriteBytesFn,
	asthlp.SliceByteLiteral{','}.Expr(), // []byte{','}
))

//  if result.Len() > 1 {
//    result.Write([]byte{','})
//  }
var putCommaFirstIf = asthlp.If(
	asthlp.Great(
		asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector("result", "Len"))),
		asthlp.IntegerConstant(1).Expr(),
	),
	putCommaFirst,
)

var (
	WriteStringFn = asthlp.InlineFunc(asthlp.SimpleSelector("result", "WriteString"))
	WriteBytesFn  = asthlp.InlineFunc(asthlp.SimpleSelector("result", "Write"))
)

// result.WriteString("\"{json}\":")
// b = strconv.AppendUint(buf[:0], uint64({src}), 10)
// result.Write(b)
func (f *Field) typeMarshal(src ast.Expr, v, t string) []ast.Stmt {
	var (
		ur = src
		wb WriteBlock
		nc = f.refx != f.expr
	)
	if f.isStar {
		ur = asthlp.Star(src)
	}
	switch t {

	case "int", "int8", "int16", "int32", "int64":
		wb = intMarshal(ur, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", nc || t != "int64")

	case "uint", "uint8", "uint16", "uint32", "uint64":
		wb = uintMarshal(ur, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", nc || t != "uint64")

	case "float32", "float64":
		wb = floatMarshal(ur, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", nc || t != "float64")

	case "bool":
		if f.isStar {
			wb = refBoolMarshal(ur, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty")
		} else {
			wb = boolMarshal(ur, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty")
		}

	case "string":
		wb = stringMarshal(ur, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", nc)

	default:
		panic(fmt.Errorf("unknown base type %q", t))
	}

	if f.isStar {
		wb.NotZero = asthlp.NotNil(src)
		if f.tags.JsonAppendix() != "omitempty" {
			wb.IfZero = []ast.Stmt{
				asthlp.CallStmt(asthlp.Call(
					WriteStringFn,
					asthlp.StringConstant(fmt.Sprintf(`"%s":null`, f.tags.JsonName())).Expr(),
				)),
			}
		}
	}

	return wb.Render(putCommaFirstIf)
}

// result.WriteString("\"field\":\"\"")
func (f *Field) typeMarshalDefault(t string) []ast.Stmt {
	var writeArg ast.Expr
	switch t {

	case "int", "int8", "int16", "int32", "int64":
		writeArg = asthlp.StringConstant(fmt.Sprintf(`"%s":0`, f.tags.JsonName())).Expr()

	case "uint", "uint8", "uint16", "uint32", "uint64":
		writeArg = asthlp.StringConstant(fmt.Sprintf(`"%s":0`, f.tags.JsonName())).Expr()

	case "float32", "float64":
		writeArg = asthlp.StringConstant(fmt.Sprintf(`"%s":0.0`, f.tags.JsonName())).Expr()

	case "bool":
		writeArg = asthlp.StringConstant(fmt.Sprintf(`"%s":false`, f.tags.JsonName())).Expr()

	case "string":
		writeArg = asthlp.StringConstant(fmt.Sprintf(`"%s":""`, f.tags.JsonName())).Expr()

	default:
		panic(fmt.Errorf("default value for %q is not implemented", t))
	}
	return []ast.Stmt{
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector("result", "WriteString")),
			writeArg,
		)),
	}
}

// if s.HeightRef != nil {
//     {v} := *{src}
//     result.WriteString("\"{json}\":")
//     b = strconv.AppendUint(buf[:0], uint64({v}), 10)
//     result.Write(b)
// } else {
//     result.WriteString("\"{json}\":{default}")
// }
func (f *Field) typeRefMarshal(src ast.Expr, v, t string) []ast.Stmt {
	var result = []ast.Stmt{
		define(ast.NewIdent(v), asthlp.Star(src)),
	}
	result = append(result, f.typeMarshal(ast.NewIdent(v), v, t)...)

	if stmt := f.ifNil(); len(stmt) > 0 {
		return []ast.Stmt{
			asthlp.IfElse(
				asthlp.NotNil(src),
				asthlp.Block(result...),
				asthlp.Block(stmt...),
			),
		}
	}
	return []ast.Stmt{
		asthlp.If(
			asthlp.NotNil(src),
			result...,
		),
	}
}

// result.WriteString("\"{name}\":{default}")
func (f *Field) ifNil() []ast.Stmt {
	if f.tags.DefaultValue() == "" {
		if f.tags.JsonTags().Has("omitempty") {
			return nil
		}
		// result.WriteString("\"{name}\":null")
		return []ast.Stmt{
			&ast.ExprStmt{
				X: &ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
					Args: []ast.Expr{
						&ast.BasicLit{Kind: token.STRING, Value: `"\"` + f.tags.JsonName() + `\":null"`},
					},
				},
			},
		}
	}
	return []ast.Stmt{
		&ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
				Args: []ast.Expr{
					&ast.BasicLit{Kind: token.STRING, Value: `"\"` + f.tags.JsonName() + `\":` + helpers.StringFromType(f.expr, f.tags.DefaultValue()) + `"`},
				},
			},
		},
	}
}
