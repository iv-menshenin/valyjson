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
		// expr represents field type expression
		expr ast.Expr
		// refx represents field type denoted value
		refx ast.Expr
		// tags contains tag descriptor
		tags tags.Tags
		// isStar is type is ref
		isStar bool
		// isNullable is type nullable
		isNullable bool
	}
)

func New(f *ast.Field) *fld {
	if f.Tag == nil {
		panic("you must fill in all fields with tags")
	}
	var ff = fld{
		expr: f.Type,
		tags: tags.Parse(f.Tag.Value),
	}
	ff.prepareRef()
	return &ff
}

// FillStatements makes statements processed data-filling for struct field
// 	if offset := v.Get("offset"); offset != nil {
//      var vOffset int
//      vOffset, err = offset.Int()
//      if err != nil {
//          return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
//      }
//      s.Offset = vOffset
//	} else {
//      s.Offset = 100
//	}
func (f *fld) FillStatements(name string) []ast.Stmt {
	if f.tags.JsonName() == "" {
		return nil
	}
	var v = intermediateVarName(name, f.tags)
	var body *ast.BlockStmt
	var els ast.Stmt
	if stmt := f.fillFrom(name, v); len(stmt) > 0 {
		body = &ast.BlockStmt{List: stmt}
	}
	if stmt := f.ifDefault(name); len(stmt) > 0 {
		els = &ast.BlockStmt{List: stmt}
	}
	if body == nil {
		panic(fmt.Errorf("can`t prepare AST for '%s'", name))
	}
	var condition ast.Expr = &ast.BinaryExpr{
		X:  ast.NewIdent(v),
		Op: token.NEQ,
		Y:  ast.NewIdent("nil"),
	}
	if f.isNullable {
		condition = &ast.CallExpr{
			Fun:  ast.NewIdent("valueIsNotNull"),
			Args: []ast.Expr{ast.NewIdent(v)},
		}
	}
	return []ast.Stmt{
		&ast.IfStmt{
			Init: f.extract(v),
			Cond: condition,
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
	var elseStmt ast.Stmt
	var v = intermediateVarName(name, f.tags)
	var src = &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}
	switch tt := f.expr.(type) {

	case *ast.Ident:
		if f.isStar {
			mrsh = append(mrsh, f.typeRefMarshal(src, v, tt.Name)...)
		} else {
			mrsh = append(mrsh, f.typeMarshal(src, v, tt.Name)...)
		}
		if !f.tags.JsonTags().Has("omitempty") {
			elseStmt = &ast.BlockStmt{
				List: f.typeMarshalDefault(src, v, tt.Name),
			}
		}

	default:
		// todo @menshenin panic
	}
	return []ast.Stmt{
		&ast.IfStmt{
			Cond: &ast.BinaryExpr{
				X:  src,
				Op: token.NEQ,
				Y:  &ast.BasicLit{Kind: token.STRING, Value: "\"\""},
			},
			Body: &ast.BlockStmt{List: mrsh},
			Else: elseStmt,
		},
	}
}

func intermediateVarName(name string, t tags.Tags) string {
	varName := strings.ToLower(name)
	switch varName {
	// reserved words
	case "break", "case", "chan", "const", "continue", "default", "defer", "else", "fallthrough", "for", "func",
		"go", "goto", "if", "import", "interface", "map", "package", "range", "return", "select", "struct", "switch",
		"type", "var":
		varName = "_" + varName
	}
	return varName
}
