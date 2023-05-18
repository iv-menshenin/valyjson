package helpers

import (
	"fmt"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"go/ast"
	"go/token"
	"strings"

	asthlp "github.com/iv-menshenin/go-ast"
)

func IsIdent(expr ast.Expr, ident string) bool {
	if i, ok := expr.(*ast.Ident); ok {
		return i.Name == ident
	}
	return false
}

func BasicLiteralFromType(t ast.Expr, val string) ast.Expr {
	switch i := t.(type) {

	case *ast.Ident:
		switch i.Name {

		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return &ast.BasicLit{
				Kind:  token.INT,
				Value: val,
			}

		case "float32", "float64":
			return &ast.BasicLit{
				Kind:  token.FLOAT,
				Value: val,
			}

		case "string":
			return &ast.BasicLit{
				Kind:  token.STRING,
				Value: "\"" + strings.Replace(val, "\"", "\\\"", -1) + "\"",
			}

		case "rune":
			return &ast.BasicLit{
				Kind:  token.CHAR,
				Value: "'" + strings.Replace(val, "'", "\\'", -1) + "'",
			}

		case "bool":
			return &ast.BasicLit{
				Kind:  token.IDENT,
				Value: val,
			}

		default:
			panic(fmt.Errorf("can't process default value for datatype: %s", i.Name))
		}

	default:
		panic("can't process default value")
	}
}

func StringFromType(t ast.Expr, val string) string {
	switch i := t.(type) {

	case *ast.Ident:
		switch i.Name {

		case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return val

		case "float32", "float64":
			return val

		case "string":
			val = strings.ReplaceAll(val, `\`, `\\`)
			val = strings.ReplaceAll(val, `"`, `\"`)
			return `"` + val + `"`

		case "rune":
			if val == "'" {
				return `'\''`
			}
			return "'" + val + "'"

		case "bool":
			return val

		default:
			panic(fmt.Errorf("can't process default value for datatype: %s", i.Name))
		}

	case *ast.StarExpr:
		return StringFromType(i.X, val)

	default:
		panic("can't process default value")
	}
}

// FmtError produces an error constructor
//
//	fmt.Errorf("{format}", {attrs[0]}, {attrs[1]}, ..., {attrs[n]})
func FmtError(format string, attrs ...ast.Expr) ast.Expr {
	var fmtAttrs []ast.Expr
	fmtAttrs = append(fmtAttrs, &ast.BasicLit{
		Kind:  token.STRING,
		Value: "\"" + format + "\"",
	})
	fmtAttrs = append(fmtAttrs, attrs...)
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   ast.NewIdent("fmt"),
			Sel: ast.NewIdent("Errorf"),
		},
		Args: fmtAttrs,
	}
}

func ZeroValueOfT(x ast.Expr) ast.Expr {
	switch t := x.(type) {

	case *ast.Ident:
		switch t.Name {

		case "float32", "float64", "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
			return asthlp.Zero

		case "string":
			return asthlp.EmptyString

		case "rune":
			return asthlp.RuneConstant(0).Expr()

		case "bool":
			return asthlp.False

		default:
			if t.Obj == nil {
				break
			}
			if e, ok := t.Obj.Decl.(*ast.TypeSpec); ok {
				return ZeroValueOfT(e.Type)
			}
		}

	case *ast.SelectorExpr:
		if t.Sel.Name == "UUID" {
			return asthlp.SimpleSelector("uuid", "Nil")
		}

	case *ast.StarExpr, *ast.MapType:
		return asthlp.Nil

	case *ast.ArrayType:
		if t.Len == nil {
			return asthlp.Nil
		}
		return nil // FIXME can't check, always nonzero
	}
	return nil
}

func MakeIfItsNullTypeCondition() ast.Expr {
	return asthlp.Equal(
		asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameJsonValue, "Type"))),
		asthlp.SimpleSelector("fastjson", "TypeNull"),
	)
}

func DenotedType(expr ast.Expr) ast.Expr {
	switch typed := expr.(type) {
	case *ast.Ident:
		return denotedIdent(typed)

	case *ast.SelectorExpr:
		if typed.Sel.Name != "Time" && typed.Sel.Name != "UUID" {
			return denotedIdent(typed.Sel)
		}
	}
	return expr
}

func denotedIdent(t *ast.Ident) ast.Expr {
	if t.Obj != nil {
		ts, ok := t.Obj.Decl.(*ast.TypeSpec)
		if ok {
			return ts.Type
		}
	}
	return t
}
