package field

import (
	"fmt"
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"
)

type WriteBlock struct {
	NotZero ast.Expr
	Block   []ast.Stmt
	Default []ast.Stmt
}

func (w WriteBlock) Render(putComma ast.Stmt) []ast.Stmt {
	if w.NotZero == nil {
		return append([]ast.Stmt{putComma}, w.Block...)
	}
	if len(w.Default) > 0 {
		return []ast.Stmt{
			putComma,
			asthlp.IfElse(
				w.NotZero,
				asthlp.Block(w.Block...),
				asthlp.Block(w.Default...),
			),
		}
	}
	return []ast.Stmt{
		asthlp.If(
			w.NotZero,
			append([]ast.Stmt{putComma}, w.Block...)...,
		),
	}
}

// timeMarshal produces block of code that writes date-time format.
//	result.WriteRune(',')
//	if !s.DateBegin.IsZero() {
//		result.WriteString("\"date_begin\":\"")
//		b = s.DateBegin.AppendFormat(buf[:0], time.RFC3339)
//		result.Write(b)
//		result.WriteRune('"')
//	} else {
//		result.WriteString("\"date_begin\":\"0000-00-00T00:00:00Z\"")
//	}
func timeMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var w = WriteBlock{
		NotZero: asthlp.Not(asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "IsZero")))),
		Block: []ast.Stmt{
			// result.WriteString("\"date_begin\":\"")
			asthlp.CallStmt(asthlp.Call(writeStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":"`, jsonName)).Expr())),
			// b = s.DateBegin.AppendFormat(buf[:0], time.RFC3339)
			asthlp.Assign(
				asthlp.VarNames{buffVar},
				asthlp.Assignment,
				asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "AppendFormat")), buffExpr, asthlp.SimpleSelector("time", "RFC3339")),
			),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(writeBytesFn, buffVar)),
			// result.WriteRune('"')
			asthlp.CallStmt(asthlp.Call(writeRuneFn, asthlp.RuneConstant('"').Expr())),
		},
	}
	if !omitempty {
		w.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":"0000-00-00T00:00:00Z"`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

func intMarshal(src ast.Expr, jsonName string, omitempty, needCast bool) WriteBlock {
	var srcInt64 = src
	if needCast {
		srcInt64 = asthlp.ExpressionTypeConvert(src, asthlp.Int64)
	}
	var w = WriteBlock{
		NotZero: asthlp.NotEqual(src, asthlp.IntegerConstant(0).Expr()),
		Block: []ast.Stmt{
			// result.WriteString(`"field":`)
			asthlp.CallStmt(asthlp.Call(writeStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = strconv.AppendInt(buf[:0], int64({src}), 10)
			asthlp.Assign(asthlp.VarNames{buffVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendInt")),
				buffExpr,                          // buf[:0]
				srcInt64,                          // int64({src})
				asthlp.IntegerConstant(10).Expr(), // 10
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(writeBytesFn, buffVar)),
		},
	}
	if !omitempty {
		w.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":0`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

// b = strconv.AppendUint(buf[:0], uint64({src}), 10)
func uintMarshal(src ast.Expr, jsonName string, omitempty, needCast bool) WriteBlock {
	var srcUint64 = src
	if needCast {
		srcUint64 = asthlp.ExpressionTypeConvert(src, asthlp.UInt64)
	}
	var w = WriteBlock{
		NotZero: asthlp.NotEqual(src, asthlp.IntegerConstant(0).Expr()),
		Block: []ast.Stmt{
			// result.WriteString(`"field":`)
			asthlp.CallStmt(asthlp.Call(writeStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = strconv.AppendUint(buf[:0], uint64({src}), 10)
			asthlp.Assign(asthlp.VarNames{buffVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendUint")),
				buffExpr,                          // buf[:0]
				srcUint64,                         // uint64({src})
				asthlp.IntegerConstant(10).Expr(), // 10
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(writeBytesFn, buffVar)),
		},
	}
	if !omitempty {
		w.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":0`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

// b = strconv.AppendFloat(buf[:0], float64({src}), 'f', 10, 64)
func floatMarshal(src ast.Expr, jsonName string, omitempty, needCast bool) WriteBlock {
	var srcFloat64 = src
	if needCast {
		srcFloat64 = asthlp.ExpressionTypeConvert(src, asthlp.Float64)
	}
	var w = WriteBlock{
		NotZero: asthlp.NotEqual(src, asthlp.IntegerConstant(0).Expr()),
		Block: []ast.Stmt{
			// result.WriteString(`"field":`)
			asthlp.CallStmt(asthlp.Call(writeStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = strconv.AppendFloat(buf[:0], float64({src}), 'f', -1, 64)
			asthlp.Assign(asthlp.VarNames{buffVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendFloat")),
				buffExpr,                          // buf[:0]
				srcFloat64,                        // float64({src})
				asthlp.RuneConstant('f').Expr(),   // 'f'
				asthlp.IntegerConstant(-1).Expr(), // -1 // todo @menshenin pass precision through structTags
				asthlp.IntegerConstant(64).Expr(), // 64
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(writeBytesFn, buffVar)),
		},
	}
	if !omitempty {
		w.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":0`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

// b = strconv.AppendBool(buf[:0], {src})
func boolMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var w = WriteBlock{
		NotZero: src,
		Block: []ast.Stmt{
			// buf = buf[:0]
			asthlp.Assign(asthlp.VarNames{buffVar}, asthlp.Assignment, buffExpr),
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":true`, jsonName)).Expr(),
			)),
		},
	}
	if !omitempty {
		w.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":false`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

func refBoolMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var def = []ast.Stmt{
		asthlp.CallStmt(asthlp.Call(
			writeStringFn,
			asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
		)),
	}
	if omitempty {
		def = nil
	}
	return WriteBlock{
		Block: []ast.Stmt{
			// buf = buf[:0]
			asthlp.Assign(asthlp.VarNames{buffVar}, asthlp.Assignment, buffExpr),
			asthlp.IfElse(
				src,
				asthlp.Block(asthlp.CallStmt(asthlp.Call(
					writeStringFn,
					asthlp.StringConstant(fmt.Sprintf(`"%s":true`, jsonName)).Expr(),
				))),
				asthlp.Block(asthlp.CallStmt(asthlp.Call(
					writeStringFn,
					asthlp.StringConstant(fmt.Sprintf(`"%s":false`, jsonName)).Expr(),
				))),
			),
		},
		Default: def,
	}
}

// b := marshalString(s.Field, buf[:0])
func stringMarshal(src ast.Expr, jsonName string, omitempty, needCast bool) WriteBlock {
	var srcString = src
	if needCast {
		srcString = asthlp.ExpressionTypeConvert(src, asthlp.String)
	}
	var w = WriteBlock{
		NotZero: asthlp.NotEqual(src, asthlp.EmptyString),
		Block: []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":"`, jsonName)).Expr(),
			)),
			asthlp.Assign(asthlp.VarNames{buffVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.NewIdent("marshalString")),
				buffExpr, srcString,
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(writeBytesFn, buffVar)),
		},
	}
	if !omitempty {
		w.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":""`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

// 	if b, err = s.Nested1.MarshalAppend(buf[:0]); err != nil {
//		return nil, fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
//	} else {
//		if len(b) > 2 {
//			result.WriteString(`"nested1":`)
//			result.Write(b)
//		}
//	}
func structMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var wrapWriter = func(block ...ast.Stmt) []ast.Stmt { return block }
	if omitempty {
		wrapWriter = func(block ...ast.Stmt) []ast.Stmt {
			return []ast.Stmt{
				// if len(b) > 2 {
				asthlp.If(
					asthlp.Great(
						asthlp.Call(asthlp.LengthFn, buffVar),
						asthlp.IntegerConstant(2).Expr(),
					),
					block...,
				),
			}
		}
	}
	return WriteBlock{
		Block: []ast.Stmt{
			asthlp.IfInitElse(
				asthlp.Assign(asthlp.VarNames{buffVar, asthlp.NewIdent("err")}, asthlp.Assignment, asthlp.Call(
					asthlp.InlineFunc(asthlp.Selector(src, "MarshalAppend")),
					buffExpr,
				)),
				asthlp.NotNil(asthlp.NewIdent("err")),
				asthlp.Block(
					// return nil, fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
					asthlp.Return(
						asthlp.Nil,
						asthlp.Call(
							asthlp.FmtErrorfFn,
							asthlp.StringConstant(`can't marshal "nested1" attribute: %w`).Expr(),
							asthlp.NewIdent("err"),
						),
					),
				),
				asthlp.Block(wrapWriter(
					// result.WriteString(`"nested1":`)
					asthlp.CallStmt(asthlp.Call(
						writeStringFn,
						asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr(),
					)),
					// result.Write(b)
					asthlp.CallStmt(asthlp.Call(writeBytesFn, buffVar)),
				)...),
			),
		},
	}
}

func refStructMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var blockWriter = structMarshal(src, jsonName, omitempty)
	blockWriter.NotZero = asthlp.NotNil(src)
	if !omitempty {
		blockWriter.Default = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				writeStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
			)),
		}
	}
	return blockWriter
}
