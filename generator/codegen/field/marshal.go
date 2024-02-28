package field

import (
	"fmt"
	asthlp "github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"go/ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
)

// WantCommaVar wantComma
var WantCommaVar = asthlp.NewIdent("wantComma")

func NeedVars() ast.Stmt {
	return asthlp.Var(
		asthlp.VariableType(names.VarNameError, asthlp.ErrorType),
		asthlp.VariableType(WantCommaVar.Name, asthlp.Bool),
	)
}

var SetCommaVar = asthlp.Assign(asthlp.VarNames{WantCommaVar}, asthlp.Assignment, asthlp.True)

// putCommaStmt puts comma
//
//	if wantComma {
//		result.RawByte(',')
//	}
var putCommaStmt = asthlp.CallStmt(asthlp.Call(
	RawByteFn,
	asthlp.RuneConstant(',').Expr(),
))

//	if wantComma {
//	  result.Write([]byte{','})
//	}
var putCommaFirstIf = asthlp.If(WantCommaVar, putCommaStmt)

var (
	RawStringFn = asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameWriter, "RawString"))
	RawByteFn   = asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameWriter, "RawByte"))
	StringFn    = asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameWriter, "String"))
	// BytesAppendFn
	//	result.Buffer.AppendBytes(buf)
	BytesAppendFn = asthlp.InlineFunc(asthlp.Selector(asthlp.SimpleSelector(names.VarNameWriter, "Buffer"), "AppendBytes"))
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
		if s, ok := src.(*ast.StarExpr); ok {
			wb.NotZero = asthlp.And(asthlp.NotNil(s.X), wb.NotZero)
		}
		if f.tags.JsonAppendix() != "omitempty" {
			wb.IfZero = []ast.Stmt{
				asthlp.CallStmt(asthlp.Call(
					RawStringFn,
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
		asthlp.CallStmt(asthlp.Call(RawStringFn, writeArg)),
	}
}

//	if s.HeightRef != nil {
//	    {v} := *{src}
//	    result.WriteString("\"{json}\":")
//	    b = strconv.AppendUint(buf[:0], uint64({v}), 10)
//	    result.Write(b)
//	} else {
//
//	    result.WriteString("\"{json}\":{default}")
//	}
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
		return asthlp.Block(
			asthlp.CallStmt(asthlp.Call(
				RawStringFn,
				asthlp.StringConstant(`"\"`+f.tags.JsonName()+`\":null"`).Expr(),
			)),
		).List
	}
	return asthlp.Block(
		asthlp.CallStmt(asthlp.Call(
			RawStringFn,
			asthlp.StringConstant(`"\"`+f.tags.JsonName()+`\":`+helpers.StringFromType(f.expr, f.tags.DefaultValue())+`"`).Expr(),
		)),
	).List
}
