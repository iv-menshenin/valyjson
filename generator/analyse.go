package generator

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"

	"github.com/iv-menshenin/valyjson/generator/codegen"
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
)

type (
	visitor struct {
		g     *Gen
		over  *ast.Ident
		decls []taggedDecl
	}
	taggedDecl struct {
		spec *ast.TypeSpec
		tags []string
	}
	renderer interface {
		Name() string
		Tags() tags.StructTags
		UnmarshalFunc() []ast.Decl
		FillerFunc() ast.Decl
		ValidatorFunc() ast.Decl
		MarshalFunc() ast.Decl
		AppendJsonFunc() ast.Decl
	}
)

func (g *Gen) BuildFillers() {
	var v = visitor{g: g}
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		if structDecl.Tags().Custom() {
			continue
		}
		if unm := structDecl.UnmarshalFunc(); len(unm) > 0 {
			g.result.Decls = append(g.result.Decls, unm...)
		}
		if fil := structDecl.FillerFunc(); fil != nil {
			g.result.Decls = append(g.result.Decls, fil)
		}
		if val := structDecl.ValidatorFunc(); val != nil {
			g.result.Decls = append(g.result.Decls, val)
		}
	}
}

func (g *Gen) BuildJsoners() {
	var v visitor
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		g.result.Decls = append(
			g.result.Decls,
			structDecl.MarshalFunc(),
		)
		g.result.Decls = append(
			g.result.Decls,
			structDecl.AppendJsonFunc(),
		)
	}
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch currNode := node.(type) {

	case *ast.GenDecl:
		for _, declSpec := range currNode.Specs {
			switch spec := declSpec.(type) {
			case *ast.TypeSpec:
				v.decls = append(v.decls, taggedDecl{
					spec: spec,
					tags: extractTags(currNode.Doc),
				})
			}
		}

	case *ast.TypeSpec:
		v.decls = append(v.decls, taggedDecl{
			spec: currNode,
			tags: extractTags(currNode.Doc),
		})
	}
	return v
}

var jsonTag = "json:"

func extractTags(comment *ast.CommentGroup) []string {
	if comment == nil {
		return nil
	}
	var commentTags []string
	for _, text := range comment.List {
		if text == nil {
			continue
		}
		if commentLine := strings.TrimLeft(text.Text, "/ \t"); strings.HasPrefix(commentLine, jsonTag) {
			splitPostfix := strings.Split(commentLine[len(jsonTag):], ",")
			for _, tagRaw := range splitPostfix {
				if tag := strings.ToLower(strings.TrimSpace(tagRaw)); tag != "" {
					commentTags = append(commentTags, tag)
				}
			}
		}
	}
	return commentTags
}

func (v *visitor) getNormalized() []renderer {
	var result []renderer
	for _, decl := range v.decls {
		if len(decl.tags) == 0 {
			// only tagged structures
			continue
		}
		switch typed := decl.spec.Type.(type) {

		case *ast.StructType:
			if tags.StructTags(decl.tags).Has(tags.TransitHandlers) {
				result = append(result, codegen.NewTransitive(decl.spec.Name.Name, decl.tags, typed))
				break
			}
			stct := &ast.StructType{
				Fields: &ast.FieldList{List: v.collectFields(typed.Fields.List)}, // uninline
			}
			result = append(result, codegen.NewStruct(decl.spec.Name.Name, decl.tags, stct))

		case *ast.MapType:
			result = append(result, codegen.NewMap(decl.spec.Name.Name, decl.tags, typed))

		case *ast.ArrayType:
			result = append(result, codegen.NewArray(decl.spec.Name.Name, decl.tags, typed))

		case *ast.Ident, *ast.SelectorExpr:
			result = append(result, codegen.NewTransitive(decl.spec.Name.Name, decl.tags, typed))

		default:
			panic("unsupported")
		}
	}
	return result
}

func (v *visitor) getDeclByName(name string) *taggedDecl {
	for i, decl := range v.decls {
		if decl.spec.Name.Name == name {
			return &v.decls[i]
		}
	}
	return nil
}

func (v *visitor) collectFields(src []*ast.Field) []*ast.Field {
	var flds = make([]*ast.Field, 0, len(src))
	for _, fld := range src {
		var tag tags.Tags
		if fld.Tag != nil {
			tag = tags.Parse(fld.Tag.Value)
		} else {
			if len(fld.Names) > 0 {
				panic("all fields should have tags")
			}
		}
		if tag.JsonAppendix() == "inline" {
			flds = append(flds, v.exploreInlined(fld)...)
			continue
		}
		if v.over != nil {
			if i, ok := fld.Type.(*ast.Ident); ok && unicode.IsUpper([]rune(i.Name)[0]) {
				fld.Type = &ast.SelectorExpr{
					X:   v.over,
					Sel: i,
				}
			}
		}
		flds = append(flds, fld)
	}
	return flds
}

func (v *visitor) exploreInlined(fld *ast.Field) []*ast.Field {
	switch inlined := fld.Type.(type) {

	case *ast.Ident:
		decl := v.getDeclByName(inlined.Name)
		if decl == nil {
			panic("can't resolve inlined field by name")
		}
		stct, ok := decl.spec.Type.(*ast.StructType)
		if !ok {
			panic("can't inline")
		}
		return v.collectFields(stct.Fields.List)

	case *ast.SelectorExpr:
		packIdent, ok := inlined.X.(*ast.Ident)
		if !ok {
			panic(fmt.Errorf("can't inline struct kind %+v; can't recognize %+v expression", fld.Type, inlined.X))
		}
		pkg, err := v.g.discovery.GetPackage(packIdent.Name)
		if err != nil {
			panic(fmt.Errorf("can't inline struct kind %+v; can't parse '%s' package: %+v", fld.Type, packIdent.Name, err))
		}
		var v1 = visitor{g: v.g, over: packIdent}
		ast.Walk(&v1, pkg)
		decl := v1.getDeclByName(inlined.Sel.Name)
		if decl == nil {
			panic("can't resolve inlined field by name")
		}
		stct, ok := decl.spec.Type.(*ast.StructType)
		if !ok {
			panic("can't inline")
		}
		return v.collectFields(stct.Fields.List)

	default:
		panic(fmt.Errorf("can't inline struct kind %+v", fld.Type))
	}
}
