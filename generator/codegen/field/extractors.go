package field

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"

	"github.com/iv-menshenin/go-ast"
	"github.com/iv-menshenin/valyjson/generator/codegen/names"
)

// sliceExtraction makes a block of code that extracts values from json array
func sliceExtraction(dst *ast.Ident, fld ast.Expr, v, json string, t ast.Expr, fa fillArrayResult) []ast.Stmt {
	varListA := asthlp.Var(
		asthlp.TypeSpec(names.VarNameListOfArrayValues, asthlp.ArrayType(asthlp.Star(names.FastJsonValue))),
	)
	stmtGetListArray := asthlp.Assign(
		asthlp.MakeVarNames(names.VarNameListOfArrayValues, names.VarNameError),
		asthlp.Assignment, asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, "Array"))),
	)
	valListSliceMake := asthlp.Call(
		asthlp.MakeFn,
		asthlp.ArrayType(t),
		asthlp.IntegerConstant(0).Expr(),
		asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(names.VarNameListOfArrayValues)),
	)
	valListSliceDeclare := asthlp.Assign(
		asthlp.VarNames{dst},
		asthlp.Definition,
		asthlp.SliceExpr(fld, nil, asthlp.IntegerConstant(0)),
	)
	var keyName = asthlp.Blank.Name
	if fa.varNum != nil {
		keyName = fa.varNum.Name
	}
	return asthlp.Block(
		// var listA []*fastjson.Value
		varListA,
		// listA, err = list.Array()
		stmtGetListArray,
		// 	if err != nil {
		// 		return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
		// 	}
		checkErrAndReturnParsingError(asthlp.StringConstant(json).Expr()),
		//	valList := s.Field[:0]
		valListSliceDeclare,
		//	if l := len(listA); cap(valList) < l || (l == 0 && s.Field == nil) {
		//		valList = make([]string, 0, l)
		//	}
		asthlp.IfInit(
			asthlp.Assign(asthlp.MakeVarNames("l"), asthlp.Definition, asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(names.VarNameListOfArrayValues))),
			asthlp.Or(
				asthlp.Binary(asthlp.Call(asthlp.CapFn, dst), asthlp.NewIdent("l"), token.LSS),
				asthlp.ParenExpr(
					asthlp.And(
						asthlp.Equal(asthlp.NewIdent("l"), asthlp.Zero),
						asthlp.IsNil(fld),
					),
				),
			),
			asthlp.Assign(asthlp.VarNames{dst}, asthlp.Assignment, valListSliceMake),
		),
		// for _, listElem := range listA {
		asthlp.Range(true, keyName, fa.varElem.Name, asthlp.NewIdent(names.VarNameListOfArrayValues), fa.body...),
	).List
}

func arrayExtraction(dst *ast.Ident, fld ast.Expr, v, json string, t *ast.ArrayType, fa fillArrayResult) []ast.Stmt {
	varListA := asthlp.Var(
		asthlp.TypeSpec(names.VarNameListOfArrayValues, asthlp.ArrayType(asthlp.Star(names.FastJsonValue))),
	)
	stmtGetListArray := asthlp.Assign(
		asthlp.MakeVarNames(names.VarNameListOfArrayValues, names.VarNameError),
		asthlp.Assignment, asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, "Array"))),
	)
	valListDeclare := asthlp.DeclareVariable().AppendSpec(
		asthlp.TypeSpec(dst.Name, t),
	)
	var keyName = asthlp.Blank.Name
	if fa.varNum != nil {
		keyName = fa.varNum.Name
	}
	return asthlp.Block(
		// var listA []*fastjson.Value
		varListA,
		// listA, err = list.Array()
		stmtGetListArray,
		// 	if err != nil {
		// 		return fmt.Errorf("error parsing '%slist' value: %w", objPath, err)
		// 	}
		checkErrAndReturnParsingError(asthlp.StringConstant(json).Expr()),
		//	var valList []Type
		valListDeclare.Stmt(),
		//	if len(listA) != 3 {
		//		return fmt.Error
		//	}
		asthlp.If(
			asthlp.Binary(asthlp.Call(asthlp.LengthFn, asthlp.NewIdent(names.VarNameListOfArrayValues)), t.Len, token.NEQ),
			asthlp.Return(
				asthlp.Call(
					asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
					asthlp.StringConstant(json).Expr(),
					asthlp.Call(asthlp.FmtErrorfFn, asthlp.StringConstant("array len mismatch").Expr()),
				),
			),
		),
		// for _, listElem := range listA {
		asthlp.Range(true, keyName, fa.varElem.Name, asthlp.NewIdent(names.VarNameListOfArrayValues), fa.body...),
	).List
}

// checkErrAndReturnParsingError generates a decoding error check
//
//	if err != nil {
//	    return fmt.Errorf("error parsing '%s.name' value: %w", objPath, err)
//	}
func checkErrAndReturnParsingError(pathElem ast.Expr) ast.Stmt {
	return asthlp.If(
		asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
		asthlp.Return(
			asthlp.Call(
				asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
				pathElem,
				asthlp.NewIdent(names.VarNameError),
			),
		),
	)
}

// stringExtraction makes a block of code that extracts an string from json element into []byte variable
//
//	var valField []byte
//	if valField, err = field.StringBytes(); err != nil {
//	  return fmt.Errorf("error parsing '%s.{json}' value: %w", objPath, err)
//	}
func stringExtraction(dst *ast.Ident, v string, elemPath ast.Expr) []ast.Stmt {
	return asthlp.Block(
		// var valField []byte
		asthlp.Var(asthlp.VariableType(dst.Name, asthlp.ArrayType(asthlp.Byte))),
		asthlp.IfInit(
			// if valField, err = field.StringBytes(); err != nil
			asthlp.Assign(
				asthlp.MakeVarNames(dst.Name, names.VarNameError),
				asthlp.Assignment,
				asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, "StringBytes"))),
			),
			asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
			// return fmt.Errorf("error parsing ...
			asthlp.Return(
				asthlp.Call(
					asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
					elemPath,
					asthlp.NewIdent(names.VarNameError),
				),
			),
		),
	).List
}

func stringExtractionWithoutErrChecking(dst *ast.Ident, v string) []ast.Stmt {
	return asthlp.Block(
		// var valField []byte
		asthlp.Var(asthlp.VariableType(dst.Name, asthlp.ArrayType(asthlp.Byte))),
		// valField, err = field.StringBytes()
		asthlp.Assign(
			asthlp.MakeVarNames(dst.Name, names.VarNameError),
			asthlp.Assignment,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, "StringBytes"))),
		),
	).List
}

// intExtraction makes a block of code that extracts an integer from json element into int variable
//
//	var {dst} int
//	{dst}, err = {v}.Int()
func intExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return particularTypeExtraction(dst.Name, v, asthlp.Int, "Int")
}

// uintExtraction makes a block of code that extracts an integer from json element into uint variable
//
//	var {dst} uint
//	{dst}, err = {v}.Uint()
func uintExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return particularTypeExtraction(dst.Name, v, asthlp.UInt, "Uint")
}

// int64Extraction makes a block of code that extracts an integer from json element into int64 variable
//
//	var {dst} int64
//	{dst}, err = {v}.Int64()
func int64Extraction(dst *ast.Ident, v string) []ast.Stmt {
	return particularTypeExtraction(dst.Name, v, asthlp.Int64, "Int64")
}

// uint64Extraction makes a block of code that extracts an integer from json element into uint64 variable
//
//	var {dst} uint64
//	{dst}, err = {v}.Uint64()
func uint64Extraction(dst *ast.Ident, v string) []ast.Stmt {
	return particularTypeExtraction(dst.Name, v, asthlp.UInt64, "Uint64")
}

// floatExtraction makes a block of code that extracts numeric value from json element into float64 variable
//
//	var {dst} float64
//	{dst}, err = {v}.Float64()
func floatExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return particularTypeExtraction(dst.Name, v, asthlp.Float64, "Float64")
}

// boolExtraction makes a block of code that extracts boolean value from json element into bool variable
//
//	var {dst} bool
//	{dst}, err = {v}.Bool()
func boolExtraction(dst *ast.Ident, v string) []ast.Stmt {
	return particularTypeExtraction(dst.Name, v, asthlp.Bool, "Bool")
}

// particularTypeExtraction makes a block of code that extracts value from json element into typed variable
//
//	var {dst} {varType}
//	{dst}, err = {v}.{method}()
func particularTypeExtraction(dst, v string, varType ast.Expr, method string) []ast.Stmt {
	return asthlp.Block(
		asthlp.Var(asthlp.VariableType(dst, varType)),
		asthlp.Assign(
			asthlp.MakeVarNames(dst, names.VarNameError),
			asthlp.Assignment,
			asthlp.Call(asthlp.InlineFunc(asthlp.SimpleSelector(v, method))),
		),
	).List
}

func (f *Field) nestedExtraction(dst *ast.Ident, src, t ast.Expr, v ast.Expr) []ast.Stmt {
	var declare ast.Stmt
	if src == nil || f.isStar {
		declare = asthlp.Var(asthlp.VariableType(dst.Name, t))
	} else {
		f.filled = true
		declare = asthlp.Var(asthlp.VariableValue(dst.Name, asthlp.FreeExpression(asthlp.Ref(src))))
	}
	return []ast.Stmt{
		declare,
		asthlp.Assign(asthlp.MakeVarNames(names.VarNameError), asthlp.Assignment, asthlp.Call(
			asthlp.InlineFunc(asthlp.Selector(dst, names.MethodNameFill)), v,
		)),
	}
}

// uuidExtraction generates the code of the standard conversion process from string to UUID format
//
//	var valfield uuid.UUID
//	b, err := field.StringBytes()
//	if err != nil {
//	    return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
//	}
//	valfield, err = uuid.ParseBytes(b)
func uuidExtraction(dst *ast.Ident, t ast.Expr, v, name string) []ast.Stmt {
	var stmt = []ast.Stmt{
		&ast.DeclStmt{
			Decl: &ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{dst},
						Type:  t,
					},
				},
			},
		},
		&ast.AssignStmt{
			Lhs: []ast.Expr{ast.NewIdent("b"), ast.NewIdent("err")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun: &ast.SelectorExpr{X: ast.NewIdent(v), Sel: ast.NewIdent("StringBytes")},
				},
			},
		},
		checkErrAndReturnParsingError(asthlp.StringConstant(name).Expr()),
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent("err")},
			Tok: token.ASSIGN,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun:  asthlp.SimpleSelector("uuid", "ParseBytes"),
					Args: []ast.Expr{ast.NewIdent("b")},
				},
			},
		},
	}
	return stmt
}

// b, err := {v}.StringBytes()
//
//	if err != nil {
//	  return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
//	}
//
// {dst}, err := parseDateTime(string(b))
func timeExtraction(dst *ast.Ident, v, jsonName, layout string) []ast.Stmt {
	var extrStmt []ast.Stmt
	var srcString = asthlp.VariableTypeConvert("b", asthlp.String)
	if layout == "" {
		extrStmt = timeExtractionUnify(dst, srcString)
	} else {
		var layoutExpr ast.Expr
		if l := strings.Split(layout, "."); len(l) == 2 && l[0] == "time" {
			layoutExpr = &ast.SelectorExpr{
				X:   ast.NewIdent(l[0]),
				Sel: ast.NewIdent(l[1]),
			}
		} else {
			layoutExpr = &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf(`"%s"`, strings.Replace(layout, "\"", "\\\"", -1))}
		}
		extrStmt = []ast.Stmt{
			&ast.AssignStmt{
				Lhs: []ast.Expr{dst, ast.NewIdent("err")},
				Tok: token.DEFINE,
				Rhs: []ast.Expr{
					&ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("time"),
							Sel: ast.NewIdent("Parse"),
						},
						Args: []ast.Expr{
							layoutExpr,
							srcString,
						},
					},
				},
			},
		}
	}
	return append(
		//   b, err := {v}.StringBytes()
		//   if err != nil {
		//     return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		//   }
		[]ast.Stmt{
			asthlp.Assign(asthlp.MakeVarNames("b", names.VarNameError), asthlp.Definition, asthlp.Call(
				asthlp.InlineFunc(asthlp.SimpleSelector(v, "StringBytes")),
			)),
			asthlp.If(
				asthlp.NotNil(asthlp.NewIdent(names.VarNameError)),
				asthlp.Return(
					asthlp.Call(
						asthlp.InlineFunc(asthlp.NewIdent(names.ParsingError)),
						asthlp.StringConstant(jsonName).Expr(),
						asthlp.NewIdent(names.VarNameError),
					),
				),
			),
		},
		extrStmt...,
	)
}

// valDateBegin, err := parseDateTime(string(b))
func timeExtractionUnify(dst *ast.Ident, v ast.Expr) []ast.Stmt {
	return []ast.Stmt{
		&ast.AssignStmt{
			Lhs: []ast.Expr{dst, ast.NewIdent("err")},
			Tok: token.DEFINE,
			Rhs: []ast.Expr{
				&ast.CallExpr{
					Fun:  ast.NewIdent("parseDateTime"),
					Args: []ast.Expr{v},
				},
			},
		},
	}
}
