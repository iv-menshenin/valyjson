package field

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

type (
	// fld render helper for ast.Field
	fld struct {
		// f contains field AST
		f *ast.Field
		// t contains tag descriptor
		t tags.Tags
	}
)

func New(f *ast.Field) *fld {
	if f.Tag == nil {
		panic("you must fill in all fields with tags")
	}
	return &fld{
		f: f,
		t: tags.Parse(f.Tag.Value),
	}
}

// 	if offset := v.Get("offset"); offset != nil {
//      var vOffset int
//		vOffset, err = offset.Int()
//		if err != nil {
//			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//		}
//      s.Offset = vOffset
//	} else {
//		s.Offset = 100
//	}
func (f fld) FillStatements(name string) []ast.Stmt {
	if f.t.JsonName() == "" {
		return nil
	}
	var v = intermediateVarName(name, f.t)
	var body *ast.BlockStmt
	var els ast.Stmt
	if stmt := f.fillFrom(name, v); len(stmt) > 0 {
		body = &ast.BlockStmt{List: stmt}
	}
	if stmt := f.ifDefault(name); len(stmt) > 0 {
		els = &ast.BlockStmt{List: stmt}
	}
	if body == nil {
		// todo @menshenin panic?
		return nil
	}
	return []ast.Stmt{
		&ast.IfStmt{
			Init: f.extract(v),
			Cond: &ast.BinaryExpr{
				X:  ast.NewIdent(v),
				Op: token.NEQ,
				Y:  ast.NewIdent("nil"),
			},
			Body: body,
			Else: els,
		},
	}
}

// result.WriteString("\"field\":")
// b, err = marshalString(s.Field, buf[:0])
// if err != nil {
// 	return nil, err
// }
// result.Write(b)
func (f fld) MarshalStatements(name string) []ast.Stmt {
	var mrsh = []ast.Stmt{
		// result.WriteString("\"field\":")
		&ast.ExprStmt{X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("WriteString")},
			Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"\"%s\":"`, f.t.JsonName())}},
		}},
	}
	var v = intermediateVarName(name, f.t)
	var src = &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}
	switch tt := f.f.Type.(type) {

	case *ast.Ident:
		mrsh = append(mrsh, f.typeMarshal(src, v, tt.Name)...)

	case *ast.StarExpr:
		var tName = "nested"
		if ident, ok := tt.X.(*ast.Ident); ok {
			tName = ident.Name
		}
		mrsh = append(mrsh, f.typeRefMarshal(src, v, tName)...)

	default:
		// todo @menshenin panic
	}
	// result.Write(b)
	mrsh = append(mrsh, &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun:  &ast.SelectorExpr{X: ast.NewIdent("result"), Sel: ast.NewIdent("Write")},
			Args: []ast.Expr{ast.NewIdent("b")},
		},
	})
	return mrsh
}

func intermediateVarName(name string, t tags.Tags) string {
	return strings.ToLower(name)
}
