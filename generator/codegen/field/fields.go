package field

import (
	"go/ast"
	"go/token"
	"strings"

	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

type (
	// fld render helper for ast.Field
	fld struct {
		// x represents field type expression
		x ast.Expr
		// t contains tag descriptor
		t tags.Tags
		// isStar is type is ref
		isStar bool
	}
)

func New(f *ast.Field) *fld {
	if f.Tag == nil {
		panic("you must fill in all fields with tags")
	}
	var ff = fld{
		x: f.Type,
		t: tags.Parse(f.Tag.Value),
	}
	ff.prepareRef()
	return &ff
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
func (f *fld) FillStatements(name string) []ast.Stmt {
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
func (f *fld) MarshalStatements(name string) []ast.Stmt {
	var mrsh []ast.Stmt
	var v = intermediateVarName(name, f.t)
	var src = &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}
	switch tt := f.x.(type) {

	case *ast.Ident:
		if f.isStar {
			mrsh = append(mrsh, f.typeRefMarshal(src, v, tt.Name)...)
		} else {
			mrsh = append(mrsh, f.typeMarshal(src, v, tt.Name)...)
		}

	default:
		// todo @menshenin panic
	}
	return mrsh
}

func intermediateVarName(name string, t tags.Tags) string {
	return strings.ToLower(name)
}
