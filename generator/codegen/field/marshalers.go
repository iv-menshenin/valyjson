package field

import (
	"fmt"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"
)

type WriteBlock struct {
	NotZero ast.Expr
	Block   []ast.Stmt
	IfZero  []ast.Stmt
}

func (w WriteBlock) Render(putComma ast.Stmt) []ast.Stmt {
	if w.NotZero == nil {
		return append([]ast.Stmt{putComma}, w.Block...)
	}
	if len(w.IfZero) > 0 {
		return []ast.Stmt{
			putComma,
			asthlp.IfElse(
				w.NotZero,
				asthlp.Block(w.Block...),
				asthlp.Block(w.IfZero...),
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
//
//	result.WriteRune(',')
//	if !s.DateBegin.IsZero() {
//		result.WriteString("\"date_begin\":")
//		b = marshalTime(buf[:0], s.DateBegin, time.RFC3339Nano)
//		result.Write(b)
//	} else {
//		result.WriteString("\"date_begin\":\"0000-00-00T00:00:00Z\"")
//	}
func timeMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var w = WriteBlock{
		NotZero: asthlp.Not(asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "IsZero")))),
		Block: []ast.Stmt{
			// result.WriteString(`"date_begin":`)
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = marshalTime(buf[:0], s.DateBegin, time.RFC3339Nano)
			asthlp.Assign(
				asthlp.VarNames{BufVar},
				asthlp.Assignment,
				asthlp.Call(names.MarshalTimeFunc, BufExpr, src, names.TimeDefaultLayout),
			),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":"0000-00-00T00:00:00Z"`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

//	if buf, err := s.RayID.MarshalText(); err != nil {
//			return nil, fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
//		} else {
//			result.WriteString(`"ray_id":"`)
//			result.Write(buf)
//			result.WriteRune('"')
//		}
func uuidMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var w = WriteBlock{
		Block: []ast.Stmt{
			asthlp.IfInitElse(
				asthlp.Assign(
					asthlp.VarNames{BufVar, asthlp.NewIdent(names.VarNameError)},
					asthlp.Definition, // must shadow the buf variable
					asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "MarshalText"))),
				),
				asthlp.NotNil(asthlp.NewIdent("err")),
				// return nil, fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
				asthlp.Block(
					asthlp.Return(
						asthlp.Nil,
						asthlp.Call(
							asthlp.InlineFunc(asthlp.SimpleSelector("fmt", "Errorf")),
							asthlp.StringConstant(`can't marshal "nested1" attribute: %w`).Expr(),
							asthlp.NewIdent("err"),
						),
					),
				),
				asthlp.Block(
					// result.WriteString("\"ray_id\":")
					asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":"`, jsonName)).Expr())),
					// result.Write(buf)
					asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
					// result.WriteRune('"')
					asthlp.CallStmt(asthlp.Call(WriteRuneFn, asthlp.RuneConstant('"').Expr())),
				),
			),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
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
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = strconv.AppendInt(buf[:0], int64({src}), 10)
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendInt")),
				BufExpr,                           // buf[:0]
				srcInt64,                          // int64({src})
				asthlp.IntegerConstant(10).Expr(), // 10
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
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
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = strconv.AppendUint(buf[:0], uint64({src}), 10)
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendUint")),
				BufExpr,                           // buf[:0]
				srcUint64,                         // uint64({src})
				asthlp.IntegerConstant(10).Expr(), // 10
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
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
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// b = strconv.AppendFloat(buf[:0], float64({src}), 'f', -1, 64)
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("strconv", "AppendFloat")),
				BufExpr,                           // buf[:0]
				srcFloat64,                        // float64({src})
				asthlp.RuneConstant('f').Expr(),   // 'f'
				asthlp.IntegerConstant(-1).Expr(), // -1 // todo @menshenin pass precision through structTags
				asthlp.IntegerConstant(64).Expr(), // 64
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
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
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, BufExpr),
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":true`, jsonName)).Expr(),
			)),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":false`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

func refBoolMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var def = []ast.Stmt{
		asthlp.CallStmt(asthlp.Call(
			WriteStringFn,
			asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
		)),
	}
	if omitempty {
		def = nil
	}
	return WriteBlock{
		Block: []ast.Stmt{
			// buf = buf[:0]
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, BufExpr),
			asthlp.IfElse(
				src,
				asthlp.Block(asthlp.CallStmt(asthlp.Call(
					WriteStringFn,
					asthlp.StringConstant(fmt.Sprintf(`"%s":true`, jsonName)).Expr(),
				))),
				asthlp.Block(asthlp.CallStmt(asthlp.Call(
					WriteStringFn,
					asthlp.StringConstant(fmt.Sprintf(`"%s":false`, jsonName)).Expr(),
				))),
			),
		},
		IfZero: def,
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
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr(),
			)),
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, asthlp.Call(
				names.MarshalStringFunc,
				BufExpr, srcString,
			)),
			// result.Write(b)
			asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":""`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

//	if b, err = s.Nested1.MarshalAppend(buf[:0]); err != nil {
//		return nil, fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
//	} else {
//
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
						asthlp.Call(asthlp.LengthFn, BufVar),
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
				asthlp.Assign(asthlp.VarNames{BufVar, asthlp.NewIdent("err")}, asthlp.Assignment, asthlp.Call(
					asthlp.InlineFunc(asthlp.Selector(src, "MarshalAppend")),
					BufExpr,
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
						WriteStringFn,
						asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr(),
					)),
					// result.Write(b)
					asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
				)...),
			),
		},
	}
}

func refStructMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var blockWriter = structMarshal(src, jsonName, omitempty)
	blockWriter.NotZero = asthlp.NotNil(src)
	if !omitempty {
		blockWriter.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
			)),
		}
	}
	return blockWriter
}

//	if s.Tags != nil {
//		buf = buf[:0]
//		result.WriteString(`"tags":{`)
//		var filled bool
//		for k, v := range s.Tags {
//			if filled {
//				result.WriteRune(',')
//			}
//			filled = true
//			result.WriteRune('"')
//			result.WriteString(k)
//			result.WriteString(`":"`)
//			result.WriteString(v)
//			result.WriteRune('"')
//		}
//		result.WriteRune('}')
//	} else {
//
//		result.WriteString(`"tags":null`)
//	}
func mapMarshal(src ast.Expr, jsonName string, omitempty, isStringKey bool, ve ValueExtractor) WriteBlock {
	const (
		filled = "_filled"
		key    = "_k"
		val    = "_v"
	)
	var keyAsString ast.Expr = asthlp.NewIdent(key)
	if !isStringKey {
		keyAsString = asthlp.ExpressionTypeConvert(keyAsString, asthlp.String)
	}
	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteRune(',')
		//	}
		asthlp.If(asthlp.NewIdent(filled), asthlp.CallStmt(asthlp.Call(WriteRuneFn, asthlp.RuneConstant(',').Expr()))),
		// filled = true
		asthlp.Assign(asthlp.MakeVarNames(filled), asthlp.Assignment, asthlp.True),
		// result.WriteRune('"')
		asthlp.CallStmt(asthlp.Call(WriteRuneFn, asthlp.RuneConstant('"').Expr())),
		// result.WriteString(key)
		asthlp.CallStmt(asthlp.Call(WriteStringFn, keyAsString)),
		// result.WriteString(`":`)
		asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(`":`).Expr())),
	}
	iterBlock = append(
		iterBlock,
		ve(asthlp.NewIdent(val))...,
	)
	iterBlock = append(
		iterBlock,
		// result.Write(b)
		asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
	)
	var w = WriteBlock{
		NotZero: asthlp.NotNil(src),
		Block: []ast.Stmt{
			// buf = buf[:0]
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, BufExpr),
			// result.WriteString(`"jsonName":{`)
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":{`, jsonName)).Expr())),
			// var filled bool
			asthlp.Var(asthlp.VariableType(filled, asthlp.Bool)),
			// for key, val := range {src} {
			asthlp.Range(true, key, val, src, iterBlock...),
			// result.WriteRune('}')
			asthlp.CallStmt(asthlp.Call(WriteRuneFn, asthlp.RuneConstant('}').Expr())),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
			)),
		}
	}
	return w
}

func arrayMarshal(src ast.Expr, jsonName string, omitempty bool, ve ValueExtractor) WriteBlock {
	const (
		filled = "_filled"
		key    = "_k"
		val    = "_v"
	)
	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteRune(',')
		//	}
		asthlp.If(asthlp.NewIdent(filled), asthlp.CallStmt(asthlp.Call(WriteRuneFn, asthlp.RuneConstant(',').Expr()))),
		// filled = true
		asthlp.Assign(asthlp.MakeVarNames(filled), asthlp.Assignment, asthlp.True),
	}
	iterBlock = append(
		iterBlock,
		ve(asthlp.NewIdent(val))...,
	)
	iterBlock = append(
		iterBlock,
		// result.Write(b)
		asthlp.CallStmt(asthlp.Call(WriteBytesFn, BufVar)),
	)
	var w = WriteBlock{
		NotZero: asthlp.NotNil(src),
		Block: []ast.Stmt{
			// buf = buf[:0]
			asthlp.Assign(asthlp.VarNames{BufVar}, asthlp.Assignment, BufExpr),
			// result.WriteString(`"jsonName":[`)
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":[`, jsonName)).Expr())),
			// var filled bool
			asthlp.Var(asthlp.VariableType(filled, asthlp.Bool)),
			// for key, val := range {src} {
			asthlp.Range(true, key, val, src, iterBlock...),
			// result.WriteRune(']')
			asthlp.CallStmt(asthlp.Call(WriteRuneFn, asthlp.RuneConstant(']').Expr())),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
			)),
		}
	}
	return w
}
