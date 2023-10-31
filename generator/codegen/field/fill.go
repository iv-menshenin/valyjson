package field

import (
	"fmt"
	asthlp "github.com/iv-menshenin/go-ast"
	"go/ast"
	"go/token"

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
	f.field = &ast.SelectorExpr{X: ast.NewIdent(names.VarNameReceiver), Sel: ast.NewIdent(name)}
	var bufVariable = makeBufVariable(name)
	var result []ast.Stmt

	result = append(result, IsNotEmpty(f.TypedValue(bufVariable, v, asthlp.StringConstant(f.tags.JsonName()).Expr()))...)
	result = append(result, f.checkErr(bufVariable)...)

	if f.isStar {
		result = append(result, f.fillRefField(bufVariable, f.field)...)
	} else {
		result = append(result, f.fillField(bufVariable, f.field)...)
	}
	return result
}

func makeBufVariable(name string) *ast.Ident {
	return asthlp.NewIdent("val" + name)
}

// var elem int
//
//	if elem, err = listElem.Int(); err != nil {
//		err = newParsingError(strconv.Itoa(listElemNum), err)
//		break
//	}
//
// valList = append(valList, int32(elem))
func (f *Field) fillElem(dst ast.Expr, v string) []ast.Stmt {
	var bufVariable = asthlp.NewIdent("elem")
	var result []ast.Stmt
	if f.isNullable {
		//if !valueIsNotNull(listElem) {
		//	valFieldRef = append(valFieldRef, nil)
		//	continue
		//}
		result = append(result, asthlp.If(
			asthlp.Not(asthlp.Call(valueIsNotNull, ast.NewIdent(v))),
			appendStmt(dst, ast.NewIdent("nil")),
			asthlp.Continue(),
		))
	}
	result = append(result, IsNotEmpty(f.TypedValue(bufVariable, v, asthlp.Call(asthlp.StrconvItoaFn, asthlp.NewIdent("_elemNum"))))...)
	result = append(result, f.breakErr()...)

	elemAsParticularType := asthlp.Call(asthlp.InlineFunc(f.expr), bufVariable)
	// valList = append(valList, int32(elem))
	if f.isStar {
		const newElem = "newElem"
		result = append(
			result,
			asthlp.Assign(asthlp.MakeVarNames(newElem), asthlp.Definition, elemAsParticularType),
			appendStmt(dst, asthlp.Ref(ast.NewIdent(newElem))),
		)
		return result
	}
	return append(result, appendStmt(dst, elemAsParticularType))
}

func appendStmt(dst, el ast.Expr) ast.Stmt {
	return asthlp.Assign(
		asthlp.VarNames{dst},
		asthlp.Assignment,
		asthlp.Call(asthlp.AppendFn, dst, el),
	)
}

//	 var val{name} {type}
//		val{name}, err = {v}.(Int|Int64|String|Bool)()
func (f *Field) TypedValue(dst *ast.Ident, v string, elemPathExpr ast.Expr) []ast.Stmt {
	var result []ast.Stmt
	switch t := f.refx.(type) {

	case *ast.Ident:
		result = append(result, f.typeExtraction(dst, v, t.Name, elemPathExpr)...)

	case *ast.StructType:
		result = append(result, f.typeExtraction(dst, v, "struct", elemPathExpr)...)

	case *ast.SelectorExpr:
		switch t.Sel.Name {

		case "Time":
			result = append(result, timeExtraction(dst, v, f.tags.JsonName(), f.tags.Layout())...)

		case "UUID":
			result = append(result, uuidExtraction(dst, f.refx, v, f.tags.JsonName())...)

		default:
			result = append(result, nestedExtraction(dst, f.expr, v, f.tags.JsonName())...)
		}

	case *ast.ArrayType:
		intF := Field{
			expr: t.Elt,
			tags: tags.Parse(fmt.Sprintf(`json:"%s"`, f.tags.JsonName())),
		}
		intF.prepareRef()
		result = append(result, arrayExtraction(dst, f.field, v, f.tags.JsonName(), t.Elt, intF.fillElem(dst, "listElem"))...)
		return result

	case *ast.MapType:
		result = append(result, f.mapExtraction(dst, t, v, f.tags.JsonName())...)

	case *ast.InterfaceType:
		// TODO

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
		return nestedExtraction(dst, f.expr, v, f.tags.JsonName())

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
	var valueAsValue = asthlp.ExpressionTypeConvert(value, t.Value)
	if _, isStar := t.Value.(*ast.StarExpr); isStar {
		valueAsValue = asthlp.Call(
			asthlp.InlineFunc(asthlp.ParenExpr(t.Value)),
			asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector("unsafe", "Pointer")),
				asthlp.Ref(value),
			),
		)
		// if v.Type() == fastjson.TypeNull {
		//			{dst}[string(key)] = prop
		//			return
		//		}
		ifNullValue = asthlp.If(
			helpers.MakeIfItsNullTypeCondition(),
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
		// var {dst} = make(map[Key]Property, o.Len())
		asthlp.Var(asthlp.VariableValue(dst.Name, asthlp.FreeExpression(asthlp.Call(
			asthlp.MakeFn, t, asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector("o", "Len"))),
		)))),
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
							asthlp.Assign(
								asthlp.VarNames{
									asthlp.Index(dst, asthlp.FreeExpression(asthlp.VariableTypeConvert("key", t.Key))),
								},
								asthlp.Assignment,
								valueAsValue,
							),
						),
					),
				).
				Lit(),
		)),
	).List
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
func (f *Field) breakErr() []ast.Stmt {
	if helpers.IsIdent(f.expr, "string") {
		// no error checking for string
		return nil
	}
	return []ast.Stmt{
		asthlp.If(
			asthlp.NotNil(ast.NewIdent(names.VarNameError)),
			asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
				asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
				asthlp.Call(asthlp.StrconvItoaFn, asthlp.NewIdent("_elemNum")),
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
			return f.typedFillIn(&ast.UnaryExpr{X: rhs, Op: token.AND}, dst, t.Name)

		case "string":
			// s.FieldRef = (*string)(unsafe.Pointer(&valFieldRef))
			return []ast.Stmt{
				asthlp.Assign(
					asthlp.VarNames{dst},
					asthlp.Assignment,
					asthlp.Call(
						asthlp.InlineFunc(asthlp.ParenExpr(asthlp.Star(t))),
						asthlp.Call(
							asthlp.InlineFunc(asthlp.SimpleSelector("unsafe", "Pointer")),
							asthlp.Ref(rhs),
						),
					),
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

// {dst} = new({t})
// *{dst} = {t}({rhs})
func (f *Field) newAndFillIn(rhs, dst, t ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		assign(dst, asthlp.Call(asthlp.NewFn, t)),
		assign(asthlp.Star(dst), asthlp.ExpressionTypeConvert(rhs, t)),
	}
}

func (f *Field) fillField(rhs, dst ast.Expr) []ast.Stmt {
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
		if f.tags.JsonTags().Has("required") {
			// return fmt.Errorf("required element '{json}' is missing", objPath)
			return []ast.Stmt{
				asthlp.Return(
					asthlp.Call(
						asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
						asthlp.StringConstant(f.tags.JsonName()).Expr(),
						helpers.FmtError("\"required element '%s' is missing\"", asthlp.StringConstant(f.tags.JsonName()).Expr()),
					),
				),
			}
		}
		return nil
	}
	if f.isStar {
		var tmpVarName = "__" + name
		return []ast.Stmt{
			// if {tmp} = nil {
			asthlp.If(
				asthlp.IsNil(asthlp.NewIdent(varName)),
				asthlp.Var(
					asthlp.VariableType(tmpVarName, f.expr, asthlp.FreeExpression(helpers.BasicLiteralFromType(f.refx, f.tags.DefaultValue()))),
				),
				assign(asthlp.SimpleSelector(names.VarNameReceiver, name), asthlp.Ref(asthlp.NewIdent(tmpVarName))),
			),
		}
	}
	return []ast.Stmt{
		assign(asthlp.SimpleSelector(names.VarNameReceiver, name), helpers.BasicLiteralFromType(f.refx, f.tags.DefaultValue())),
	}
}
