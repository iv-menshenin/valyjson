package helpers

import "go/ast"

func DenotedType(expr ast.Expr) ast.Expr {
	switch typed := expr.(type) {
	case *ast.Ident:
		return DenotedIdent(typed)

	case *ast.SelectorExpr:
		if typed.Sel.Name != "Time" && typed.Sel.Name != "UUID" {
			return DenotedIdent(typed.Sel)
		}
	}
	return expr
}

func DenotedIdent(t *ast.Ident) ast.Expr {
	if t.Obj != nil {
		ts, ok := t.Obj.Decl.(*ast.TypeSpec)
		if ok {
			return ts.Type
		}
	}
	return t
}

func Ordinal(s string) bool {
	switch s {
	case "int", "int8", "int16", "int32", "int64":
		return true
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return true
	case "float32", "float64":
		return true
	case "bool":
		return true
	case "string":
		return true
	default:
		return false
	}
}
