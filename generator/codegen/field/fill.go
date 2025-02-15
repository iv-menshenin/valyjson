package field

import (
	"fmt"
	"go/ast"

	asthlp "github.com/iv-menshenin/go-ast"

	"github.com/iv-menshenin/valyjson/generator/codegen/helpers"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

var (
	// jsonGet gets json attribute value from json object
	//  v.Get(name)
	jsonGet        = asthlp.InlineFunc(asthlp.SimpleSelector(names.VarNameJsonValue, "Get"))
	valueIsNotNull = asthlp.InlineFunc(ast.NewIdent("valueIsNotNull"))
)

// offset := v.Get("offset")
func (f *Field) extract(v string) ast.Stmt {
	if f.tags.JsonName() == "" {
		// inlined structure; so there is no json name
		return asthlp.Assign(asthlp.MakeVarNames(v), asthlp.Definition, asthlp.NewIdent(names.VarNameJsonValue))
	}
	return asthlp.Assign(asthlp.MakeVarNames(v), asthlp.Definition, f.getValue())
}

// v.Get("offset")
func (f *Field) getValue() ast.Expr {
	return asthlp.Call(jsonGet, asthlp.StringConstant(f.tags.JsonName()).Expr())
}

// fillFrom makes statements to fill some field according to its type
//
//	s.Offset, err = offset.Int()
//	if err != nil {
//	    return fmt.Errorf("error parsing '%s.limit' value: %w", objPath, err)
//	}
func (f *Field) fillFrom(name, v string) []ast.Stmt {
	f.field = asthlp.SimpleSelector(names.VarNameReceiver, name)
	var bufVariable = makeBufVariable(name)

	var result = make([]ast.Stmt, 0, 10)
	result = append(result, IsNotEmpty(f.TypedValue(bufVariable, v, asthlp.StringConstant(f.tags.JsonName()).Expr()))...)
	result = append(result, f.checkErr(bufVariable)...)

	if f.isStarType {
		return append(result, f.fillFieldRef(bufVariable, f.field)...)
	}
	if f.isStar {
		return append(result, f.fillRefField(bufVariable, f.field)...)
	}
	return append(result, f.fillField(bufVariable, f.field)...)
}

func makeBufVariable(name string) *ast.Ident {
	return asthlp.NewIdent("val" + name)
}

type fillArrayResult struct {
	varElem *ast.Ident
	varNum  *ast.Ident
	body    []ast.Stmt
}

func (f *fillArrayResult) append(stmt ...ast.Stmt) {
	f.body = append(f.body, stmt...)
}

//	 valData = valData[:len(valData)+1]
//	 err = valData[i].FillFromJSON(listElem)
//		if valData, err = listElem.Int(); err != nil {
//			err = newParsingError(strconv.Itoa(listElemNum), err)
//			break
//		}
func (f *Field) appendElem(dst ast.Expr, valVariableName string) fillArrayResult {
	r := fillArrayResult{
		varElem: asthlp.NewIdent(valVariableName),
		varNum:  asthlp.NewIdent(fmt.Sprintf("_key%d", f.level)),
	}
	elemVar := asthlp.NewIdent(fmt.Sprintf("_elem%d", f.level))
	bufVariable := asthlp.NewIdent(fmt.Sprintf("_tmp%d", f.level))
	if ident, ok := dst.(*ast.Ident); ok {
		bufVariable = ident
	}
	var (
		bufLenValue = asthlp.Call(asthlp.LengthFn, bufVariable)
		reference   = asthlp.Index(bufVariable, asthlp.FreeExpression(asthlp.Sub(bufLenValue, asthlp.IntegerConstant(1).Expr())))
	)
	r.append(
		// every iteration must extend this slice
		asthlp.Assign(
			asthlp.VarNames{bufVariable},
			asthlp.Assignment,
			asthlp.SliceExpr(bufVariable, nil, asthlp.FreeExpression(asthlp.Add(bufLenValue, asthlp.IntegerConstant(1).Expr()))),
		),
	)
	if f.isNullable {
		// if !valueIsNotNull(listElem) {
		//  valData[len(valData)-1] = nil
		//  continue
		// }
		r.append(asthlp.If(
			asthlp.Not(asthlp.Call(valueIsNotNull, r.varElem)),
			asthlp.Assign(asthlp.VarNames{reference}, asthlp.Assignment, asthlp.Nil),
			asthlp.Continue(),
		))
	}

	r.append(IsNotEmpty(
		f.TypedValue(elemVar, r.varElem.Name, asthlp.Call(asthlp.StrconvItoaFn, r.varNum)),
	)...)
	r.append(f.breakErr(r.varNum)...)

	elemAsParticularType := asthlp.Call(asthlp.InlineFunc(f.expr), elemVar)
	// valList = append(valList, int32(elem))
	if f.isStar {
		const newElem = "newElem"
		r.append(
			asthlp.Assign(asthlp.MakeVarNames(newElem), asthlp.Definition, elemAsParticularType),
			asthlp.Assign(
				asthlp.VarNames{asthlp.Index(bufVariable, asthlp.FreeExpression(r.varNum))},
				asthlp.Assignment,
				asthlp.Ref(ast.NewIdent(newElem)),
			),
		)
		return r
	}
	if !f.filled {
		// valExcluded[_elemNum] = FieldValueString(elem)
		r.append(asthlp.Assign(
			asthlp.VarNames{asthlp.Index(bufVariable, asthlp.FreeExpression(r.varNum))},
			asthlp.Assignment,
			elemAsParticularType,
		))
	}
	return r
}

func appendStmt(dst, el ast.Expr) ast.Stmt {
	return asthlp.Assign(
		asthlp.VarNames{dst},
		asthlp.Assignment,
		asthlp.Call(asthlp.AppendFn, dst, el),
	)
}

func (f *Field) fillElem(dst ast.Expr, valVariableName string) fillArrayResult {
	r := fillArrayResult{
		varElem: asthlp.NewIdent(valVariableName),
		varNum:  asthlp.NewIdent(fmt.Sprintf("_key%d", f.level)),
	}
	bufVariable := asthlp.NewIdent(fmt.Sprintf("_tmp%d", f.level))
	if f.isNullable {
		// if !valueIsNotNull(listElem) {
		//  valFieldRef = append(valFieldRef, nil)
		//   continue
		// }
		r.append(asthlp.If(
			asthlp.Not(asthlp.Call(valueIsNotNull, r.varElem)),
			appendStmt(dst, ast.NewIdent("nil")),
			asthlp.Continue(),
		))
	}
	r.append(IsNotEmpty(f.TypedValue(bufVariable, r.varElem.Name, asthlp.Call(asthlp.StrconvItoaFn, r.varNum)))...)
	r.append(f.breakErr(r.varNum)...)

	elemAsParticularType := asthlp.Call(asthlp.InlineFunc(f.expr), bufVariable)
	// valList[_elemNum] = int32(elem)
	if f.isStar {
		const newElem = "newElem"
		r.append(
			asthlp.Assign(asthlp.MakeVarNames(newElem), asthlp.Definition, elemAsParticularType),
			asthlp.Assign(asthlp.VarNames{asthlp.Index(dst, asthlp.FreeExpression(r.varNum))}, asthlp.Assignment, asthlp.Ref(ast.NewIdent(newElem))),
		)
		return r
	}
	// valExcluded[_elemNum] = FieldValueString(elem)
	r.append(asthlp.Assign(
		asthlp.VarNames{asthlp.Index(dst, asthlp.FreeExpression(r.varNum))},
		asthlp.Assignment,
		elemAsParticularType,
	))
	return r
}

// var val{name} {type}
// val{name}, err = {v}.(Int|Int64|String|Bool)()
func (f *Field) TypedValue(dst *ast.Ident, v string, elemPathExpr ast.Expr) []ast.Stmt {
	var result []ast.Stmt
	var refx = f.refx
	if dt := helpers.DenotedType(f.refx); refx != dt && helpers.IsOrdinal(dt) {
		refx = dt
	}
	switch t := refx.(type) {

	case *ast.Ident:
		result = append(result, f.typeExtraction(dst, v, t.Name, elemPathExpr)...)

	case *ast.StarExpr:
		f.isStarType = true
		intF := *f
		intF.refx = t.X
		intF.level = f.level + 1
		return intF.TypedValue(dst, v, elemPathExpr)

	case *ast.StructType:
		result = append(result, f.typeExtraction(dst, v, "struct", elemPathExpr)...)

	case *ast.SelectorExpr:
		switch t.Sel.Name {

		case "Time":
			result = append(result, timeExtraction(dst, v, f.tags.JsonName(), f.tags.Layout())...)

		case "UUID":
			result = append(result, uuidExtraction(dst, f.refx, v, f.tags.JsonName())...)

		default:
			result = append(result, f.nestedExtraction(dst, f.field, f.expr, asthlp.NewIdent(v))...)
		}

	case *ast.ArrayType:
		if helpers.IsIdent(t.Elt, "byte") {
			result = append(result, asthlp.Var(
				asthlp.TypeSpec(dst.Name, asthlp.ArrayType(t.Elt, t.Len)),
			))
			return append(result, ByteSliceFillFrom(asthlp.NewIdent(v), dst, t)...)
		}
		intF := Field{
			// &valData[len(valData)-1]
			field: asthlp.Index(
				dst,
				asthlp.FreeExpression(asthlp.Sub(asthlp.Call(asthlp.LengthFn, dst), asthlp.IntegerConstant(1).Expr())),
			),
			expr:  t.Elt,
			tags:  tags.Parse(fmt.Sprintf(`json:"%s"`, f.tags.JsonName())),
			level: f.level + 1,
		}
		intF.prepareRef()
		if t.Len == nil {
			result = append(result, sliceExtraction(dst, f.field, v, f.tags.JsonName(), t.Elt, intF.appendElem(dst, fmt.Sprintf("_val%d", intF.level)))...)
		} else {
			result = append(result, arrayExtraction(dst, f.field, v, f.tags.JsonName(), t, intF.fillElem(dst, fmt.Sprintf("_val%d", intF.level)))...)
		}
		return result

	case *ast.MapType:
		result = append(result, f.mapExtraction(dst, t, v, f.tags.JsonName())...)

	case *ast.InterfaceType:
		// supported empty interface only
		if len(t.Methods.List) > 0 {
			panic("interface is not supported")
		}
		result = append(result, f.interfaceExtraction(dst, t, v, f.tags.JsonName())...)

	default:
		return nil
	}
	return result
}

func IsNotEmpty(in []ast.Stmt) []ast.Stmt {
	if len(in) == 0 {
		panic("unsupported data type")
	}
	return in
}

func (f *Field) typeExtraction(dst *ast.Ident, v, t string, elemPathExpr ast.Expr) []ast.Stmt {
	switch t {

	case "int", "int8", "int16", "int32":
		return intExtraction(dst, v)

	case "int64":
		return int64Extraction(dst, v)

	case "uint", "uint8", "uint16", "uint32":
		return uintExtraction(dst, v)

	case "uint64":
		return uint64Extraction(dst, v)

	case "float32", "float64":
		return floatExtraction(dst, v)

	case "bool":
		return boolExtraction(dst, v)

	case "string":
		if f.dontCheckErr {
			return stringExtractionWithoutErrChecking(dst, v)
		}
		return stringExtraction(dst, v, elemPathExpr)

	default:
		return f.nestedExtraction(dst, f.field, f.expr, asthlp.NewIdent(v))

	}
}

// o, err := keytypedproperties.Object()
//
//	if err != nil {
//		return fmt.Errorf("error parsing '%s.key_typed_properties' value: %w", objPath, err)
//	}
//
// var valKeyTypedProperties = make(map[Key]Property, o.Len())
//
//	o.Visit(func(key []byte, v *fastjson.Value) {
//		if err != nil {
//			return
//		}
//		var prop Property
//		err = prop.FillFromJson(v, objPath+"properties.")
//		if err != nil {
//			err = fmt.Errorf("error parsing '%s.key_typed_properties.%s' value: %w", objPath, string(key), err)
//			return
//		}
//		valKeyTypedProperties[Key(key)] = prop
//	})
//
//	if err != nil {
//		return err
//	}
//
// s.KeyTypedProperties = valKeyTypedProperties
func (f *Field) mapExtraction(dst *ast.Ident, t *ast.MapType, v, json string) []ast.Stmt {
	var value = asthlp.NewIdent("value")
	var ifNullValue = asthlp.EmptyStmt()
	var valueAsValue = []ast.Stmt{
		asthlp.Assign(
			asthlp.VarNames{
				asthlp.Index(dst, asthlp.FreeExpression(asthlp.VariableTypeConvert("key", t.Key))),
			},
			asthlp.Assignment,
			asthlp.ExpressionTypeConvert(value, t.Value),
		),
	}
	if itemType, isStar := t.Value.(*ast.StarExpr); isStar {
		// 				var valRef = uint16(value)
		//				valUintVal[Key(key)] = &valRef
		valueAsValue = []ast.Stmt{
			asthlp.Var(asthlp.VariableValue("valRef", asthlp.FreeExpression(asthlp.ExpressionTypeConvert(value, itemType.X)))),
			asthlp.Assign(
				asthlp.VarNames{
					asthlp.Index(dst, asthlp.FreeExpression(asthlp.VariableTypeConvert("key", t.Key))),
				},
				asthlp.Assignment,
				asthlp.Ref(asthlp.NewIdent("valRef")),
			),
		}
		// if v.Type() == fastjson.TypeNull {
		//			{dst}[string(key)] = prop
		//			return
		//		}
		ifNullValue = asthlp.If(
			helpers.MakeIfItsNullTypeCondition(asthlp.NewIdent(names.VarNameJsonValue)),
			asthlp.Assign(
				[]ast.Expr{
					asthlp.Index(dst, asthlp.FreeExpression(asthlp.VariableTypeConvert("key", t.Key))),
				},
				asthlp.Assignment,
				asthlp.Nil,
			),
			asthlp.ReturnEmpty(),
		)
	}
	valFactory := New(asthlp.Field("", asthlp.MakeTagsForField(map[string][]string{
		"json": {f.tags.JsonName()},
	}), t.Value))
	valFactory.DontCheckErr()
	return asthlp.Block(
		//	o, err := {v}.Object()
		asthlp.Assign(asthlp.MakeVarNames("o", names.VarNameError), asthlp.Definition, asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector(v, "Object")),
		)),
		checkErrAndReturnParsingError(asthlp.StringConstant(json).Expr()),
		//	{dst} := s.Tags
		//	if {dst} == nil {
		//		{dst} = make(map[Key]Property, o.Len())
		//	}
		asthlp.Assign(asthlp.MakeVarNames(dst.Name), asthlp.Definition, f.field),
		asthlp.If(
			asthlp.IsNil(asthlp.NewIdent(dst.Name)),
			asthlp.Assign(asthlp.MakeVarNames(dst.Name), asthlp.Assignment, asthlp.Call(
				asthlp.MakeFn, t, asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector("o", "Len"))),
			)),
		),
		// o.Visit(func(key []byte, v *fastjson.Value) {
		asthlp.CallStmt(asthlp.Call(
			asthlp.InlineFunc(asthlp.SimpleSelector("o", "Visit")),
			asthlp.DeclareFunction(nil).
				Params(
					asthlp.Field("key", nil, asthlp.ArrayType(asthlp.Byte)),
					asthlp.Field("v", nil, asthlp.Star(names.FastJsonValue)),
				).
				AppendStmt(
					asthlp.If(asthlp.NotNil(asthlp.NewIdent(names.VarNameError)), asthlp.ReturnEmpty()),
					ifNullValue,
				).
				AppendStmt(
					// fills one value
					IsNotEmpty(valFactory.TypedValue(value, "v", asthlp.StringConstant(f.tags.JsonName()).Expr()))...,
				).
				AppendStmt(
					asthlp.IfElse(
						asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
						// err = newParsingError(string(key), err)
						asthlp.Block(
							asthlp.Assign(
								asthlp.MakeVarNames(names.VarNameError),
								asthlp.Assignment,
								asthlp.Call(
									asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
									asthlp.VariableTypeConvert("key", asthlp.String),
									asthlp.NewIdent(names.VarNameError),
								),
							),
						),
						// {dst}[string(key)] = prop
						asthlp.Block(
							valueAsValue...,
						),
					),
				).
				Lit(),
		)),
	).List
}

func (f *Field) interfaceExtraction(*ast.Ident, *ast.InterfaceType, string, string) []ast.Stmt {
	// TODO
	panic("not implemented")
}

//	if err != nil {
//		return fmt.Errorf("error parsing '%s.limit' value: %w", objPath, err)
//	}
//
//	if valIntFld8 > math.MaxInt8 {
//		return fmt.Errorf("error parsing '%s.int_fld8' value %d exceeds maximum for data type uint8", objPath, valIntFld8)
//	}
func (f *Field) checkErr(val *ast.Ident) []ast.Stmt {
	var checkOverflow = asthlp.EmptyStmt()
	ident, isIdent := f.refx.(*ast.Ident)
	if isIdent && ident.Name == "string" {
		return nil
	}
	if maxExp := getMaxByType(ident); maxExp != nil {
		phldr := "%d"
		if ident.Name == "float32" {
			phldr = "%f"
		}
		maxExceeded := phldr + " exceeds maximum for data type " + ident.Name
		checkOverflow = asthlp.If(
			asthlp.Great(val, maxExp),
			asthlp.Return(asthlp.Call(
				asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
				asthlp.StringConstant(f.tags.JsonName()).Expr(),
				helpers.FmtError(maxExceeded, val),
			)),
		)
	}

	return []ast.Stmt{
		asthlp.If(
			asthlp.NotNil(ast.NewIdent(names.VarNameError)),
			asthlp.Return(
				asthlp.Call(
					asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
					asthlp.StringConstant(f.tags.JsonName()).Expr(),
					asthlp.NewIdent(names.VarNameError),
				),
			),
		),
		checkOverflow,
	}
}
func getMaxByType(t *ast.Ident) ast.Expr {
	if t == nil {
		return nil
	}
	switch t.Name {

	case "float32":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxFloat32")}

	case "int8":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxInt8")}

	case "int16":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxInt16")}

	case "int32":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxInt32")}

	case "uint8":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxUint8")}

	case "uint16":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxUint16")}

	case "uint32":
		return &ast.SelectorExpr{X: ast.NewIdent("math"), Sel: ast.NewIdent("MaxUint32")}

	default:
		return nil
	}
}

//	if err != nil {
//		err = newParsingError(strconv.Itoa(_elemNum), err)
//		break
//	}
func (f *Field) breakErr(elemPathChain ast.Expr) []ast.Stmt {
	if helpers.IsIdent(f.expr, "string") {
		// no error checking for string
		return nil
	}
	return []ast.Stmt{
		asthlp.If(
			asthlp.NotNil(ast.NewIdent(names.VarNameError)),
			asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
				asthlp.Call(asthlp.StrconvItoaFn, elemPathChain),
				ast.NewIdent(names.VarNameError),
			)),
			asthlp.Break(),
		),
	}
}

func (f *Field) fillRefField(rhs, dst ast.Expr) []ast.Stmt {
	switch t := f.expr.(type) {

	case *ast.Ident:
		switch t.Name {

		case "bool", "int64", "int", "float64":
			return f.typedFillIn(asthlp.Ref(rhs), dst, t.Name)

		case "string":
			// 		var valFieldStr = string(valFieldRef)
			//		s.FieldRef = &valFieldStr
			return []ast.Stmt{
				asthlp.Var(asthlp.VariableValue("valFieldStr", asthlp.FreeExpression(asthlp.ExpressionTypeConvert(rhs, asthlp.String)))),
				asthlp.Assign(
					asthlp.VarNames{dst},
					asthlp.Assignment,
					asthlp.Ref(asthlp.NewIdent("valFieldStr")),
				),
			}

		default:
			return f.newAndFillIn(rhs, dst, ast.NewIdent(t.Name))

		}

	case *ast.MapType:
		return []ast.Stmt{
			asthlp.Assign(asthlp.VarNames{dst}, asthlp.Assignment, asthlp.Ref(rhs)),
		}

	default:
		return f.newAndFillIn(rhs, dst, f.expr)
	}
}

// var _tmp = int32(valInt32Ref)
// s.Int32Ref = DefinedRefInt32(&_tmp)
func (f *Field) fillFieldRef(rhs, dst ast.Expr) []ast.Stmt {
	var tmpName = asthlp.NewIdent("_ref")
	var result = []ast.Stmt{
		asthlp.Var(asthlp.VariableType(tmpName.Name, f.refx.(*ast.StarExpr).X)),
	}
	switch t := f.refx.(*ast.StarExpr).X.(type) {

	case *ast.Ident:
		result = append(result, f.typedFillIn(rhs, tmpName, t.Name)...)

	case *ast.StructType:
		panic("not implemented")

	case *ast.SelectorExpr:
		if helpers.IsOrdinal(f.refx) {
			return []ast.Stmt{assign(tmpName, asthlp.ExpressionTypeConvert(rhs, f.expr))}
		}
		result = append(result, assign(tmpName, rhs))

	case *ast.MapType, *ast.ArrayType:
		// {dst} = {rhs}
		result = append(result, assign(tmpName, rhs))

	default:
		return nil
	}
	if f.isStar {
		return append(
			result,
			asthlp.Assign(
				asthlp.VarNames{dst},
				asthlp.Assignment,
				asthlp.Call(asthlp.NewFn, f.expr),
			),
			asthlp.Assign(
				asthlp.VarNames{asthlp.Star(dst)},
				asthlp.Assignment,
				asthlp.ExpressionTypeConvert(asthlp.Ref(tmpName), f.expr),
			),
		)
	}
	return append(
		result,
		asthlp.Assign(
			asthlp.VarNames{dst},
			asthlp.Assignment,
			asthlp.ExpressionTypeConvert(asthlp.Ref(tmpName), f.expr),
		),
	)
}

// {dst} = new({t})
// *{dst} = {t}({rhs})
func (f *Field) newAndFillIn(rhs, dst, t ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		assign(dst, asthlp.Call(asthlp.NewFn, t)),
		assign(asthlp.Star(dst), asthlp.ExpressionTypeConvert(rhs, t)),
	}
}

func (f *Field) fillField(rhs, dst ast.Expr) []ast.Stmt {
	if f.filled {
		return nil
	}
	var result []ast.Stmt
	switch t := f.expr.(type) {

	case *ast.Ident:
		return f.typedFillIn(rhs, dst, t.Name)

	case *ast.StructType:
		return result

	case *ast.SelectorExpr:
		if helpers.IsOrdinal(f.refx) {
			return []ast.Stmt{assign(dst, asthlp.ExpressionTypeConvert(rhs, f.expr))}
		}
		result = append(result, assign(dst, rhs))
		return result

	case *ast.MapType, *ast.ArrayType:
		// {dst} = {rhs}
		result = append(result, assign(dst, rhs))
		return result

	default:
		return nil
	}
}

func (f *Field) typedFillIn(rhs, dst ast.Expr, t string) []ast.Stmt {
	switch t {
	case "string":
		return []ast.Stmt{
			assign(dst, asthlp.ExpressionTypeConvert(rhs, asthlp.NewIdent("string"))),
		}

	case "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			assign(dst, rhs),
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		return []ast.Stmt{
			assign(dst, asthlp.ExpressionTypeConvert(rhs, ast.NewIdent(t))),
		}

	default:
		return []ast.Stmt{
			assign(dst, asthlp.ExpressionTypeConvert(rhs, ast.NewIdent(t))),
		}
	}
}

func assign(dst, rhs ast.Expr) ast.Stmt {
	return asthlp.Assign(asthlp.VarNames{dst}, asthlp.Assignment, rhs)
}

func define(dst, rhs ast.Expr) ast.Stmt {
	return asthlp.Assign(asthlp.VarNames{dst}, asthlp.Definition, rhs)
}

func (f *Field) typedRefFillIn(rhs, dst ast.Expr, t string) []ast.Stmt {
	switch t {
	case "string", "int", "uint", "int64", "uint64", "float64", "bool", "byte", "rune":
		return []ast.Stmt{
			assign(dst, asthlp.Star(rhs)),
		}

	case "int8", "uint8", "int16", "uint16", "int32", "uint32", "float32":
		return []ast.Stmt{
			// s.HeightRef = new(uint32)
			assign(dst, asthlp.Call(asthlp.NewFn, ast.NewIdent(t))),
			// *s.HeightRef = uint32(xHeightRef)
			assign(asthlp.Star(dst), asthlp.ExpressionTypeConvert(rhs, ast.NewIdent(t))),
		}

	default:
		return nil
	}
}

//	} else {
//		s.{name} = 100
//	}
func (f *Field) ifDefault(varName, name string) []ast.Stmt {
	if f.tags.DefaultValue() == "" {
		if f.tags.JsonAppendix() == "required" {
			// return fmt.Errorf("required element '{json}' is missing", objPath)
			return []ast.Stmt{
				asthlp.Return(
					asthlp.Call(
						asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
						asthlp.StringConstant(f.tags.JsonName()).Expr(),
						helpers.FmtError("required element '%s' is missing", asthlp.StringConstant(f.tags.JsonName()).Expr()),
					),
				),
			}
		}
		return nil
	}
	var refVarName = "__" + name
	if f.isStarType {
		if f.isStar {
			// var __RefInt32Ref int32 = 32
			// s.RefInt32Ref = new(DefinedRefInt32)
			// *s.RefInt32Ref = DefinedRefInt32(&__RefInt32Ref)
			return []ast.Stmt{
				asthlp.Var(asthlp.VariableType(refVarName, f.refx.(*ast.StarExpr).X, asthlp.FreeExpression(helpers.BasicLiteralFromType(f.refx, f.tags.DefaultValue())))),
				assign(asthlp.SimpleSelector(names.VarNameReceiver, name), asthlp.Call(asthlp.NewFn, f.expr)),
				assign(asthlp.Star(asthlp.SimpleSelector(names.VarNameReceiver, name)), asthlp.ExpressionTypeConvert(asthlp.Ref(asthlp.NewIdent(refVarName)), f.expr)),
			}
		}
		// var _ref int32 = 32
		// s.Int32Ref = DefinedRefInt32(&_ref)
		return []ast.Stmt{
			asthlp.Var(asthlp.VariableType(refVarName, f.refx.(*ast.StarExpr).X, asthlp.FreeExpression(helpers.BasicLiteralFromType(f.refx, f.tags.DefaultValue())))),
			assign(asthlp.SimpleSelector(names.VarNameReceiver, name), asthlp.ExpressionTypeConvert(asthlp.Ref(asthlp.NewIdent(refVarName)), f.expr)),
		}
	}
	if f.isStar {
		return []ast.Stmt{
			// if {tmp} = nil {
			asthlp.If(
				asthlp.IsNil(asthlp.NewIdent(varName)),
				asthlp.Var(
					asthlp.VariableType(refVarName, f.expr, asthlp.FreeExpression(helpers.BasicLiteralFromType(f.refx, f.tags.DefaultValue()))),
				),
				assign(asthlp.SimpleSelector(names.VarNameReceiver, name), asthlp.Ref(asthlp.NewIdent(refVarName))),
			),
		}
	}
	return []ast.Stmt{
		assign(asthlp.SimpleSelector(names.VarNameReceiver, name), helpers.BasicLiteralFromType(f.refx, f.tags.DefaultValue())),
	}
}
