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

	putCommaCustom bool
}

func (w WriteBlock) Render(putComma ast.Stmt) []ast.Stmt {
	if w.putCommaCustom {
		putComma = asthlp.EmptyStmt()
	}
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
//		writeTime(result, s.DateBegin, time.RFC3339Nano)
//	} else {
//		result.WriteString("\"date_begin\":\"0000-00-00T00:00:00Z\"")
//	}
func timeMarshal(src ast.Expr, jsonName string, omitempty, isStar bool) WriteBlock {
	var notZero = asthlp.Not(asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "IsZero"))))
	if isStar {
		notZero = asthlp.NotNil(src)
		src = asthlp.Star(src)
	}
	var w = WriteBlock{
		NotZero: notZero,
		Block: []ast.Stmt{
			// result.WriteString(`"date_begin":`)
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr())),
			// writeTime(result, s.DateBegin, time.RFC3339Nano)
			asthlp.CallStmt(
				asthlp.Call(names.WriteTimeFunc, asthlp.NewIdent(names.VarNameWriter), src, names.TimeDefaultLayout),
			),
			SetCommaVar,
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":"0000-00-00T00:00:00Z"`, jsonName)).Expr(),
			)),
			SetCommaVar,
		}
	}
	return w
}

//	if buf, err := s.RayID.MarshalText(); err != nil {
//			return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
//		} else {
//			result.WriteString(`"ray_id":"`)
//			result.Write(buf)
//			result.WriteRune('"')
//		}
func uuidMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var bufVar = asthlp.NewIdent("buf")
	var w = WriteBlock{
		Block: []ast.Stmt{
			asthlp.IfInitElse(
				asthlp.Assign(
					asthlp.VarNames{bufVar, asthlp.NewIdent(names.VarNameError)},
					asthlp.Definition, // must shadow the buf variable
					asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, "MarshalText"))),
				),
				asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
				// return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
				asthlp.Block(
					asthlp.Return(
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
					asthlp.CallStmt(asthlp.Call(WriteBytesFn, bufVar)),
					// result.WriteString(`"`)
					putQuoteStmt,
					SetCommaVar,
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
			SetCommaVar,
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
			// writeInt64(result, int64(s.Field))
			asthlp.CallStmt(asthlp.Call(
				names.WriteInt64Func,
				asthlp.NewIdent(names.VarNameWriter), // result
				srcInt64,                             // int64({src})
			)),
			SetCommaVar,
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":0`, jsonName)).Expr(),
			)),
			SetCommaVar,
		}
	}
	return w
}

// writeUint64(result, uint64({src}))
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
			// writeUint64(result, uint64({src}))
			asthlp.CallStmt(asthlp.Call(
				names.WriteUint64Func,
				asthlp.NewIdent(names.VarNameWriter),
				srcUint64,
			)),
			SetCommaVar,
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":0`, jsonName)).Expr(),
			)),
			SetCommaVar,
		}
	}
	return w
}

// writeFloat64(result, float64({src}))
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
			// writeFloat64(result, float64({src}))
			asthlp.CallStmt(asthlp.Call(
				names.WriteFloat64Func, asthlp.NewIdent(names.VarNameWriter), srcFloat64,
			)),
			SetCommaVar,
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":0`, jsonName)).Expr(),
			)),
			SetCommaVar,
		}
	}
	return w
}

func boolMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var w = WriteBlock{
		NotZero: src,
		Block: []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":true`, jsonName)).Expr(),
			)),
			SetCommaVar,
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":false`, jsonName)).Expr(),
			)),
			SetCommaVar,
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
		SetCommaVar,
	}
	if omitempty {
		def = nil
	}
	return WriteBlock{
		Block: []ast.Stmt{
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
			SetCommaVar,
		},
		IfZero: def,
	}
}

// writeString(result, s.Field)
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
			asthlp.CallStmt(asthlp.Call(
				names.WriteStringFunc, asthlp.NewIdent(names.VarNameWriter), srcString,
			)),
			SetCommaVar,
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":""`, jsonName)).Expr(),
			)),
			SetCommaVar,
		}
	}
	return w
}

// structMarshal makes WriteBlock for structure marshalling
//
//	result.WriteString(`"injected":`)
//	if err = s.URL.MarshalTo(result); err != nil {
//		return fmt.Errorf(`can't marshal "injected" attribute: %w`, err)
//	}
func structMarshal(src ast.Expr, jsonName string, omitempty bool) WriteBlock {
	var notZero ast.Expr
	if omitempty {
		notZero = asthlp.Not(asthlp.Call(asthlp.InlineFunc(asthlp.Selector(src, names.MethodNameZero))))
	}
	return WriteBlock{
		NotZero: notZero,
		Block: []ast.Stmt{
			// result.WriteString(`"injected":`)
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":`, jsonName)).Expr(),
			)),
			// if err = s.TestInh02.MarshalTo(_injected); err != nil {
			asthlp.IfInit(
				asthlp.Assign(asthlp.VarNames{asthlp.NewIdent(names.VarNameError)}, asthlp.Assignment, asthlp.Call(
					asthlp.InlineFunc(asthlp.Selector(src, names.MethodNameMarshalTo)), asthlp.NewIdent(names.VarNameWriter),
				)),
				asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
				// return fmt.Errorf(`can't marshal "nested1" attribute: %w`, err)
				asthlp.Return(
					asthlp.Call(
						asthlp.FmtErrorfFn,
						asthlp.StringConstant(`can't marshal "nested1" attribute: %w`).Expr(),
						asthlp.NewIdent(names.VarNameError),
					),
				),
			),
			SetCommaVar,
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
		key = "_k"
		val = "_v"
	)
	var keyAsString ast.Expr = asthlp.NewIdent(key)
	if !isStringKey {
		keyAsString = asthlp.ExpressionTypeConvert(keyAsString, asthlp.String)
	}
	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteString(",")
		//	}
		asthlp.If(WantCommaVar, putCommaStmt),
		// filled = true
		SetCommaVar,
		// result.WriteString(`"`)
		putQuoteStmt,
		// result.WriteString(key)
		asthlp.CallStmt(asthlp.Call(WriteStringFn, keyAsString)),
		// result.WriteString(`":`)
		asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(`":`).Expr())),
	}
	iterBlock = append(
		iterBlock,
		ve(asthlp.NewIdent(val))...,
	)
	var w = WriteBlock{
		NotZero: asthlp.NotNil(src),
		Block: []ast.Stmt{
			SetCommaVar,
			// result.WriteString(`"jsonName":{`)
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":{`, jsonName)).Expr())),
			// var filled bool
			asthlp.Var(asthlp.VariableType(WantCommaVar.Name, asthlp.Bool)),
			// for key, val := range {src} {
			asthlp.Range(true, key, val, src, iterBlock...),
			// result.WriteString("}")
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant("}").Expr())),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			SetCommaVar,
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
		key = "_k"
		val = "_v"
	)
	var iterBlock = []ast.Stmt{
		//	if filled {
		//		result.WriteString(",")
		//	}
		asthlp.If(WantCommaVar, asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(",").Expr()))),
		// filled = true
		SetCommaVar,
		// key = key
		asthlp.Assign(asthlp.MakeVarNames(key), asthlp.Assignment, asthlp.NewIdent(key)),
	}
	iterBlock = append(
		iterBlock,
		ve(asthlp.NewIdent(val))...,
	)
	var w = WriteBlock{
		NotZero: asthlp.NotNil(src),
		Block: []ast.Stmt{
			SetCommaVar,
			// result.WriteString(`"jsonName":[`)
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant(fmt.Sprintf(`"%s":[`, jsonName)).Expr())),
			// var filled bool
			asthlp.Var(asthlp.VariableType(WantCommaVar.Name, asthlp.Bool)),
			// for key, val := range {src} {
			asthlp.Range(true, key, val, src, iterBlock...),
			// result.WriteString("]")
			asthlp.CallStmt(asthlp.Call(WriteStringFn, asthlp.StringConstant("]").Expr())),
		},
	}
	if !omitempty {
		w.IfZero = []ast.Stmt{
			asthlp.CallStmt(asthlp.Call(
				WriteStringFn,
				asthlp.StringConstant(fmt.Sprintf(`"%s":null`, jsonName)).Expr(),
			)),
			SetCommaVar,
		}
	}
	return w
}
