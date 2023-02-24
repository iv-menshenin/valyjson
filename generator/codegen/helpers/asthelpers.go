package helpers

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
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

// fmt.Errorf("{format}", {attrs[0]}, {attrs[1]}, ..., {attrs[n]})
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
