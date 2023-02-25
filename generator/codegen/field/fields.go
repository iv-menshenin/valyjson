package field

import (
	"fmt"
	"go/ast"
	"unicode"

	asthlp "github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

type (
	// Field render helper for ast.Field
	Field struct {
		// field
		field ast.Expr
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
		// dontCheckErr do not check error and not to return
		dontCheckErr bool
	}
)

func New(f *ast.Field) *Field {
	if f.Tag == nil && len(f.Names) > 0 {
		panic("you must fill in all fields with tags")
	}
	var ff = Field{
		expr: f.Type,
	}
	if f.Tag != nil {
		ff.tags = tags.Parse(f.Tag.Value)
	}
	ff.prepareRef()
	return &ff
}

func (f *Field) prepareRef() {
	var dstType = f.expr
	star, isStar := dstType.(*ast.StarExpr)
	if isStar {
		f.expr = star.X
		f.isStar = true
	}
	_, isArray := dstType.(*ast.ArrayType)
	_, isMap := dstType.(*ast.MapType)
	_, isStruct := dstType.(*ast.StructType)
	f.isNullable = isStar || isArray || isMap || isStruct
	f.fillDenotedType()
}

func (f *Field) fillDenotedType() {
	if i, ok := f.expr.(*ast.Ident); ok {
		f.refx = denotedType(i)
	} else {
		f.refx = f.expr
	}
}

func denotedType(t *ast.Ident) ast.Expr {
	if t.Obj != nil {
		ts, ok := t.Obj.Decl.(*ast.TypeSpec)
		if ok {
			return ts.Type
		}
	}
	return t
}

func (f *Field) DontCheckErr() {
	f.dontCheckErr = true
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
func (f *Field) FillStatements(name string) []ast.Stmt {
	if f.tags.JsonName() == "-" {
		return nil
	}
	var v = intermediateVarName(name, f.tags)
	var body *ast.BlockStmt
	var els *ast.BlockStmt
	if stmt := f.fillFrom(name, v); len(stmt) > 0 {
		body = &ast.BlockStmt{List: stmt}
	}
	if stmt := f.ifDefault(v, name); len(stmt) > 0 {
		els = &ast.BlockStmt{List: stmt}
	}
	if body == nil {
		panic(fmt.Errorf("can`t prepare AST for '%s'", name))
	}
	var condition = asthlp.NotNil(ast.NewIdent(v))
	if f.isNullable {
		condition = asthlp.Call(asthlp.InlineFunc(asthlp.NewIdent("valueIsNotNull")), ast.NewIdent(v))
	}
	return asthlp.Block(
		asthlp.IfInitElse(
			f.extract(v),
			condition,
			body,
			els,
		),
	).List
}

// result.WriteString("\"field\":")
// b, err = marshalString(s.Field, buf[:0])
// if err != nil {
// 	return nil, err
// }
// result.Write(b)
func (f *Field) MarshalStatements(name string) []ast.Stmt {
	var v = intermediateVarName(name, f.tags)
	var src = asthlp.SimpleSelector(names.VarNameReceiver, name)
	switch tt := f.refx.(type) {

	case *ast.Ident:
		return f.typeMarshal(src, v, tt.Name)

	case *ast.SelectorExpr:
		if tt.Sel.Name == "Time" {
			block := timeMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty")
			return block.Render(putCommaFirstIf)
		}
		panic("d")

	case *ast.StructType:
		var block WriteBlock
		if f.isStar {
			block = refStructMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty")
		} else {
			block = structMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty")
		}
		return block.Render(putCommaFirstIf)

	default:
		// todo @menshenin panic
		panic("not implemented")
	}
}

func (f *Field) notEmptyCondition(src ast.Expr) ast.Expr {
	i, ok := f.refx.(*ast.Ident)
	if !ok {
		return src
		// TODO panic("expected ident")
	}
	switch i.Name {

	case "int", "int8", "int16", "int32", "int64":
		return asthlp.NotEqual(src, asthlp.Zero)

	case "uint", "uint8", "uint16", "uint32", "uint64":
		return asthlp.NotEqual(src, asthlp.Zero)

	case "float32", "float64":
		return asthlp.NotEqual(src, asthlp.Zero)

	case "bool":
		return src

	case "string":
		return asthlp.NotEqual(src, asthlp.EmptyString)

	default:
		panic("not implemented")
	}
}

func intermediateVarName(name string, _ tags.Tags) string {
	return "_" + string(unicode.ToLower([]rune(name)[0])) + string([]rune(name)[1:])
}
