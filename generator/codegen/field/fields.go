package field

import (
	"fmt"
	"go/ast"
	"unicode"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
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
		// filled if already filled this field (by reference for example)
		filled bool
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

func NewFromType(expr ast.Expr, isStar bool) *Field {
	var ff = Field{
		expr:   expr,
		isStar: isStar,
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
	dt := helpers.DenotedType(f.expr)
	if i, ok := dt.(*ast.Ident); ok && helpers.Ordinal(i.Name) {
		f.refx = dt
		return
	}
	if i, ok := f.expr.(*ast.Ident); ok {
		f.refx = denotedType(i)
		return
	}
	f.refx = f.expr
}

// denotedType allows to explore real type
//
// deprecated
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
//
//		if offset := v.Get("offset"); offset != nil {
//	     var vOffset int
//	     vOffset, err = offset.Int()
//	     if err != nil {
//	         return fmt.Errorf("error parsing '%s.limit' value: %w", objPath, err)
//	     }
//	     s.Offset = vOffset
//		} else {
//	     s.Offset = 100
//		}
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
//
//	if err != nil {
//		return nil, err
//	}
//
// result.Write(b)
func (f *Field) MarshalStatements(name string) []ast.Stmt {
	var v = intermediateVarName(name, f.tags)
	var src = asthlp.SimpleSelector(names.VarNameReceiver, name)
	switch tt := f.refx.(type) {

	case *ast.Ident:
		return f.typeMarshal(src, v, tt.Name)

	case *ast.SelectorExpr:
		if tt.Sel.Name == "Time" {
			block := timeMarshal(src, f.tags.JsonName(), f.tags.Layout(), f.tags.JsonAppendix() == "omitempty", f.isStar)
			return block.Render(putCommaFirstIf)
		}
		if tt.Sel.Name == "UUID" {
			block := uuidMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty")
			return block.Render(putCommaFirstIf)
		}
		return marshalTransit(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", f.isStar).Render(putCommaFirstIf)

	case *ast.StructType:
		return marshalTransit(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", f.isStar).Render(putCommaFirstIf)

	case *ast.MapType:
		var isString bool
		if i, ok := tt.Key.(*ast.Ident); ok {
			isString = i.Name == "string"
		}
		var valType = tt.Value
		dec := GetDecorExpr(valType)
		if i, ok := valType.(*ast.Ident); ok {
			valType = denotedType(i)
		}
		errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+f.tags.JsonName()+`" attribute %q: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
		block := mapMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", isString, GetValueExtractor(valType, errExpr, dec))
		return block.Render(putCommaFirstIf)

	case *ast.ArrayType:
		var valType = tt.Elt
		dec := GetDecorExpr(valType)
		if i, ok := valType.(*ast.Ident); ok {
			valType = denotedType(i)
		}
		errExpr := asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+f.tags.JsonName()+`" item at position %d: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err"))
		block := arrayMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", GetValueExtractor(valType, errExpr, dec), tt.Len == nil)
		return block.Render(putCommaFirstIf)

	default:
		// todo @menshenin panic
		panic("not implemented")
	}
}

type ValueExtractor func(src ast.Expr) []ast.Stmt
type DecorSrc func(e ast.Expr) ast.Expr

func EmptyDecorSrc(e ast.Expr) ast.Expr {
	return e
}

func GetDecorExpr(valType ast.Expr) DecorSrc {
	if i, ok := valType.(*ast.Ident); ok {
		valType = denotedType(i)
		switch t := valType.(type) {
		case *ast.Ident:
			if i.Name != "string" && t.Name == "string" {
				return func(e ast.Expr) ast.Expr {
					return asthlp.ExpressionTypeConvert(e, asthlp.String)
				}
			}
		}
	}
	return EmptyDecorSrc
}

func GetValueExtractor(t, errExpr ast.Expr, initDecorSrc DecorSrc) ValueExtractor {
	if errExpr == nil {
		errExpr = asthlp.NewIdent(names.VarNameError)
	}
	transitMarshaller := func(src ast.Expr) []ast.Stmt {
		return []ast.Stmt{
			// err = v.MarshalTo(result)
			asthlp.Assign(
				asthlp.VarNames{asthlp.NewIdent(names.VarNameError)},
				asthlp.Assignment,
				asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, names.MethodNameMarshalTo)), asthlp.NewIdent(names.VarNameWriter)),
			),
			//	if err != nil {
			//		return fmt.Errorf(`can't marshal "name" attribute %q: %w`, k, err)
			//	}
			asthlp.If(asthlp.NotNil(asthlp.NewIdent(names.VarNameError)), asthlp.Return(errExpr)),
		}
	}

	if initDecorSrc == nil {
		initDecorSrc = EmptyDecorSrc
	}
	decorSrc := initDecorSrc
	decorStmt := func(_ ast.Expr, stmt []ast.Stmt) []ast.Stmt {
		return stmt
	}
	star, isStar := t.(*ast.StarExpr)
	if isStar {
		t = star.X
	}
	if isStar {
		decorSrc = func(e ast.Expr) ast.Expr {
			return initDecorSrc(asthlp.Star(e))
		}
		decorStmt = func(src ast.Expr, stmt []ast.Stmt) []ast.Stmt {
			return []ast.Stmt{
				asthlp.IfElse(
					asthlp.IsNil(src),
					// result.RawString("null")
					asthlp.Block(asthlp.CallStmt(asthlp.Call(
						RawStringFn,
						asthlp.StringConstant("null").Expr(),
					))),
					asthlp.Block(stmt...),
				),
			}
		}
	}

	switch tt := t.(type) {

	case *ast.Ident:
		switch tt.Name {
		case "int", "int8", "int16", "int32", "int64":
			return func(src ast.Expr) []ast.Stmt {
				int64Expression := asthlp.ExpressionTypeConvert(decorSrc(src), asthlp.Int64)
				return decorStmt(src, []ast.Stmt{
					// result.Int64(int64({src}))
					asthlp.CallStmt(asthlp.Call(names.WriteInt64Func, int64Expression)),
				})
			}

		case "uint", "uint8", "uint16", "uint32", "uint64":
			return func(src ast.Expr) []ast.Stmt {
				uint64Expression := asthlp.ExpressionTypeConvert(decorSrc(src), asthlp.UInt64)
				return decorStmt(src, []ast.Stmt{
					// result.Uint64(uint64({src}))
					asthlp.CallStmt(asthlp.Call(names.WriteUint64Func, uint64Expression)),
				})
			}

		case "float32", "float64":
			return func(src ast.Expr) []ast.Stmt {
				float64Expression := asthlp.ExpressionTypeConvert(decorSrc(src), asthlp.Float64)
				return decorStmt(src, []ast.Stmt{
					// result.Float64(float64({src}))
					asthlp.CallStmt(asthlp.Call(names.WriteFloat64Func, float64Expression)),
				})
			}

		case "bool":
			return func(src ast.Expr) []ast.Stmt {
				return decorStmt(src, []ast.Stmt{
					asthlp.IfElse(
						decorSrc(src),
						// result.RawString(`true`)
						asthlp.Block(asthlp.CallStmt(asthlp.Call(
							RawStringFn,
							asthlp.StringConstant(`true`).Expr(),
						))),
						// result.WriteString(`false`)
						asthlp.Block(asthlp.CallStmt(asthlp.Call(
							RawStringFn,
							asthlp.StringConstant(`false`).Expr(),
						))),
					),
				})
			}

		case "string":
			return func(src ast.Expr) []ast.Stmt {
				// result.String(_v)
				return decorStmt(src, []ast.Stmt{
					asthlp.CallStmt(asthlp.Call(
						names.WriteStringFunc,
						decorSrc(src),
					)),
				})
			}

		default:
			return transitMarshaller
		}

	case *ast.SelectorExpr:
		if tt.Sel.Name == "Time" {
			return func(src ast.Expr) []ast.Stmt {
				return decorStmt(src, []ast.Stmt{
					// writeTime(result, s.DateBegin, time.RFC3339Nano)
					asthlp.CallStmt(
						asthlp.Call(names.WriteTimeFunc, asthlp.NewIdent(names.VarNameWriter), src, names.TimeDefaultLayout),
					),
				})
			}
		}
		return transitMarshaller

	case *ast.StructType:
		return transitMarshaller

	case *ast.InterfaceType:
		// TODO
		panic("not implemented")

	default:
		panic("not implemented")
	}
}

func marshalTransit(src ast.Expr, name string, omitempty, isStar bool) WriteBlock {
	if name == "" {
		// inlined
		// result.Raw(unpackObject(s.Field.MarshalJSON()))
		return inlineStructMarshal(src)
	}
	var block WriteBlock
	if isStar {
		block = refStructMarshal(src, name, omitempty)
	} else {
		block = structMarshal(src, name, omitempty)
	}
	return block
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
