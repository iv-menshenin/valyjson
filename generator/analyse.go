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
)

func (g *Gen) BuildFillers() {
	var v = visitor{g: g}
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		if tags.StructTags(structDecl.tags).Custom() {
			continue
		}
		g.result.Decls = append(
			g.result.Decls,
			codegen.NewUnmarshalFunc(structDecl.name, structDecl.tags)...,
		)
		if structDecl.tran != nil {
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewFillerTranFunc(structDecl.name, structDecl.tran, structDecl.tags),
			)
		}
		if structDecl.spec == nil {
			continue
		}
		g.result.Decls = append(
			g.result.Decls,
			codegen.NewFillerFunc(structDecl.name, structDecl.spec.Fields.List, structDecl.tags),
		)
		g.result.Decls = append(
			g.result.Decls,
			codegen.NewValidatorFunc(structDecl.name, structDecl.spec.Fields.List, structDecl.tags),
		)
	}
}

func (g *Gen) BuildJsoners() {
	var v visitor
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		g.result.Decls = append(
			g.result.Decls,
			codegen.NewMarshalFunc(structDecl.name, structDecl.tags),
		)
		g.result.Decls = append(
			g.result.Decls,
			codegen.NewAppendJsonFunc(structDecl.name, structDecl.spec.Fields.List, structDecl.tags),
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

func (v *visitor) getNormalized() []taggedStruct {
	var result []taggedStruct
	for _, decl := range v.decls {
		if len(decl.tags) == 0 {
			// only tagged structures
			continue
		}
		if structObj := v.structFromDecl(decl); structObj != nil {
			result = append(result, *structObj)
		}
		if transitObj := v.transitFromDecl(decl); transitObj != nil {
			result = append(result, *transitObj)
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

func (v *visitor) structFromDecl(decl taggedDecl) *taggedStruct {
	stct, ok := decl.spec.Type.(*ast.StructType)
	if !ok {
		return nil
	}
	return &taggedStruct{
		tags: decl.tags,
		name: decl.spec.Name.Name,
		spec: &ast.StructType{
			Fields: &ast.FieldList{List: v.collectFields(stct.Fields.List)},
		},
	}
}

func (v *visitor) transitFromDecl(decl taggedDecl) *taggedStruct {
	_, ok := decl.spec.Type.(*ast.StructType)
	if ok {
		return nil
	}
	if tags.StructTags(decl.tags).Has(tags.TransitHandlers) {
		return &taggedStruct{
			tags: decl.tags,
			name: decl.spec.Name.Name,
			tran: decl.spec.Type,
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
		inlStruct := v.structFromDecl(*decl)
		if inlStruct == nil {
			panic("can't inline")
		}
		return v.collectFields(inlStruct.spec.Fields.List)

	case *ast.SelectorExpr:
		packIdent, ok := inlined.X.(*ast.Ident)
		if !ok {
			panic(fmt.Errorf("can't inline struct kind %+v; can't recognize %+v expression", fld.Type, inlined.X))
		}
		pkg, err := v.g.discovery.GetPackage(packIdent.Name)
		if err != nil {
			panic(fmt.Errorf("can't inline struct kind %+v; can't parse '%s' package", fld.Type, packIdent.Name))
		}
		var v1 = visitor{g: v.g, over: packIdent}
		ast.Walk(&v1, pkg)
		decl := v1.getDeclByName(inlined.Sel.Name)
		if decl == nil {
			panic("can't resolve inlined field by name")
		}
		inlStruct := v1.structFromDecl(*decl)
		if inlStruct == nil {
			panic("can't inline")
		}
		return v.collectFields(inlStruct.spec.Fields.List)

	default:
		panic(fmt.Errorf("can't inline struct kind %+v", fld.Type))
	}
}

type taggedStruct struct {
	tags []string
	name string
	spec *ast.StructType
	tran ast.Expr
}
