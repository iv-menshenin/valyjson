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
		return marshalTransit(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", f.isStar).Render(putCommaFirstIf)

	case *ast.StructType:
		return marshalTransit(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", f.isStar).Render(putCommaFirstIf)

	case *ast.MapType:
		var isString bool
		if i, ok := tt.Key.(*ast.Ident); ok {
			isString = i.Name == "string"
		}
		var valType = tt.Value
		if i, ok := valType.(*ast.Ident); ok {
			valType = denotedType(i)
		}
		block := mapMarshal(src, f.tags.JsonName(), f.tags.JsonAppendix() == "omitempty", isString, getValueExtractor(valType, f.tags.JsonName()))
		return block.Render(putCommaFirstIf)

	default:
		// todo @menshenin panic
		panic("not implemented")
	}
}

type valueExtractor func(src ast.Expr) []ast.Stmt

func getValueExtractor(t ast.Expr, name string) valueExtractor {
	transitMarshaller := func(src ast.Expr) []ast.Stmt {
		return []ast.Stmt{
			// buf, err = v.MarshalAppend(buf[:0])
			asthlp.Assign(
				asthlp.VarNames{bufVar, asthlp.NewIdent("err")},
				asthlp.Assignment,
				asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "MarshalAppend")), bufExpr),
			),
			//	if err != nil {
			//		return nil, fmt.Errorf(`can't marshal "name" attrbute %q: %w`, k, err)
			//	}
			asthlp.If(asthlp.NotNil(asthlp.NewIdent("err")), asthlp.Return(
				asthlp.Nil,
				asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant(`can't marshal "`+name+`" attrbute %q: %w`).Expr(), asthlp.NewIdent("_k"), asthlp.NewIdent("err")),
			)),
		}
	}
	switch tt := t.(type) {

	case *ast.Ident:
		switch tt.Name {
		case "int", "int8", "int16", "int32", "int64":
			return func(src ast.Expr) []ast.Stmt {
				int64Expression := asthlp.ExpressionTypeConvert(src, asthlp.Int64)
				return []ast.Stmt{
					// b = strconv.AppendInt(buf[:0], int64({src}), 10)
					asthlp.Assign(asthlp.VarNames{bufVar}, asthlp.Assignment, asthlp.Call(
						asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendInt")),
						bufExpr,                           // buf[:0]
						int64Expression,                   // int64({src})
						asthlp.IntegerConstant(10).Expr(), // 10
					)),
				}
			}

		case "uint", "uint8", "uint16", "uint32", "uint64":
			return func(src ast.Expr) []ast.Stmt {
				uint64Expression := asthlp.ExpressionTypeConvert(src, asthlp.UInt64)
				return []ast.Stmt{
					// buf = strconv.AppendUint(buf[:0], uint64({src}), 10)
					asthlp.Assign(asthlp.VarNames{bufVar}, asthlp.Assignment, asthlp.Call(
						asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendUint")),
						bufExpr,                           // buf[:0]
						uint64Expression,                  // uint64({src})
						asthlp.IntegerConstant(10).Expr(), // 10
					)),
				}
			}

		case "float32", "float64":
			return func(src ast.Expr) []ast.Stmt {
				float64Expression := asthlp.ExpressionTypeConvert(src, asthlp.Float64)
				return []ast.Stmt{
					// buf = strconv.AppendFloat(buf[:0], float64({src}), 'f', -1, 64)
					asthlp.Assign(asthlp.VarNames{bufVar}, asthlp.Assignment, asthlp.Call(
						asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendFloat")),
						bufExpr,                           // buf[:0]
						float64Expression,                 // float64({src})
						asthlp.RuneConstant('f').Expr(),   // 'f'
						asthlp.IntegerConstant(-1).Expr(), // -1
						asthlp.IntegerConstant(64).Expr(), // 64
					)),
				}
			}

		case "bool":
			return func(src ast.Expr) []ast.Stmt {
				return []ast.Stmt{
					asthlp.IfElse(
						src,
						// buf = append(buf[:0], []byte(`"name":true`)...)
						asthlp.Block(asthlp.CallStmt(asthlp.CallEllipsis(
							asthlp.AppendFn,
							bufExpr,
							asthlp.ExpressionTypeConvert(
								asthlp.StringConstant(fmt.Sprintf(`"%s":true`, name)).Expr(),
								asthlp.ArrayType(asthlp.Byte),
							),
						))),
						// buf = append(buf[:0], []byte(`"name":false`)...)
						asthlp.Block(asthlp.CallStmt(asthlp.CallEllipsis(
							asthlp.AppendFn,
							bufExpr,
							asthlp.ExpressionTypeConvert(
								asthlp.StringConstant(fmt.Sprintf(`"%s":false`, name)).Expr(),
								asthlp.ArrayType(asthlp.Byte),
							),
						))),
					),
				}
			}

		case "string":
			return func(src ast.Expr) []ast.Stmt {
				// buf = marshalString(buf[:0], _v)
				return []ast.Stmt{
					asthlp.Assign(asthlp.VarNames{bufVar}, asthlp.Assignment, asthlp.Call(
						asthlp.InlineFunc(asthlp.NewIdent("marshalString")),
						bufExpr, src,
					)),
				}
			}

		default:
			return transitMarshaller
		}

	case *ast.SelectorExpr:
		if tt.Sel.Name == "Time" {
			// fixme @menshenin need quotation
			panic("fixme")
			return func(src ast.Expr) []ast.Stmt {
				return []ast.Stmt{
					// b = s.DateBegin.AppendFormat(buf[:0], time.RFC3339)
					asthlp.Assign(
						asthlp.VarNames{bufVar},
						asthlp.Assignment,
						asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "AppendFormat")), bufExpr, asthlp.SimpleSelector("time", "RFC3339")),
					),
				}
			}
		}
		return transitMarshaller

	case *ast.StructType:
		return transitMarshaller

	default:
		panic("not implemented")
	}
}

func marshalTransit(src ast.Expr, name string, omitempty, isStar bool) WriteBlock {
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
