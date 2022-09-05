package codegen

import (
	"go/ast"
)

/*
func (s *Struct) fill(v *fastjson.Value, objPath string) (err error) {
	if filter := v.Get("filter"); filter != nil {
		if filter.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", v.Type())
			return fmt.Errorf("error parsing '%sfilter' value: %w", objPath, err)
		}
		s.Filter = filter.String()
	} else {
		return fmt.Errorf("the '%sfilter' path is required but ommitted", objPath)
	}
	if limit := v.Get("limit"); limit != nil {
		s.Limit, err = limit.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	}
	if offset := v.Get("offset"); offset != nil {
		s.Offset, err = offset.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	} else {
		s.Offset = 100
	}
	if nested := v.Get("nested"); nested != nil {
		err = s.Nested.fill(nested, objPath+"nested.")
		if err != nil {
			return fmt.Errorf("error parsing '%slimit' value: %w", objPath, err)
		}
	}
	return nil
}
*/

const (
	errVarName     = "err"
	recvVarName    = "s"
	valVarName     = "v"
	objPathVarName = "objPath"
	fillerFuncName = "fill"
)

func NewFillerFunc(structName string, body ...ast.Stmt) *ast.FuncDecl {
	fastJsonValue := ast.StarExpr{X: &ast.SelectorExpr{X: ast.NewIdent("fastjson"), Sel: ast.NewIdent("Value")}}
	return &ast.FuncDecl{
		Doc: nil,
		Recv: &ast.FieldList{List: []*ast.Field{
			{Names: []*ast.Ident{ast.NewIdent(recvVarName)}, Type: ast.NewIdent(structName)},
		}},
		Name: ast.NewIdent(fillerFuncName),
		Type: &ast.FuncType{
			Params: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(valVarName)}, Type: &fastJsonValue},
				{Names: []*ast.Ident{ast.NewIdent(objPathVarName)}, Type: ast.NewIdent("string")},
			}},
			Results: &ast.FieldList{List: []*ast.Field{
				{Names: []*ast.Ident{ast.NewIdent(errVarName)}, Type: ast.NewIdent("error")},
			}},
		},
		Body: &ast.BlockStmt{List: body},
	}
}

func NewFieldFillerStmt(fld *ast.Field) []ast.Stmt {
	var result []ast.Stmt
	factory := newField(fld)
	for _, name := range fld.Names {
		result = append(result, factory.Explore(name.Name)...)
	}
	return result
}
