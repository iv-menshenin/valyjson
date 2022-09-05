package codegen

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func getBasicLitFromString(t ast.Expr, val string) ast.Expr {
	switch i := t.(type) {

	case *ast.Ident:
		switch i.Name {

		case "int", "int8", "int16", "int32", "int64":
			fallthrough
		case "uint", "uint8", "uint16", "uint32", "uint64":
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

// fmt.Errorf("{format}", {attrs[0]}, {attrs[1]}, ..., {attrs[n]})
func fmtError(format string, attrs ...ast.Expr) ast.Expr {
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
