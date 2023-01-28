package generator

import (
	"github.com/iv-menshenin/valyjson/generator/codegen/tags"
	"go/ast"
	"strings"

	"github.com/iv-menshenin/valyjson/generator/codegen"
)

type (
	visitor struct {
		decls []taggedDecl
	}
	taggedDecl struct {
		spec *ast.TypeSpec
		tags []string
	}
)

func (g *Gen) BuildFillers() {
	var v visitor
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		g.result.Decls = append(
			g.result.Decls,
			codegen.NewUnmarshalFunc(structDecl.name, structDecl.tags)...,
		)
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
	var tags []string
	for _, text := range comment.List {
		if text == nil {
			continue
		}
		if commentLine := strings.TrimLeft(text.Text, "/ \t"); strings.HasPrefix(commentLine, jsonTag) {
			splitPostfix := strings.Split(commentLine[len(jsonTag):], ",")
			for _, tagRaw := range splitPostfix {
				if tag := strings.ToLower(strings.TrimSpace(tagRaw)); tag != "" {
					tags = append(tags, tag)
				}
			}
		}
	}
	return tags
}

func (v *visitor) getNormalized() []taggedStruct {
	var result []taggedStruct
	for _, decl := range v.decls {
		if len(decl.tags) == 0 {
			// only tagged structures
			continue
		}
		s := v.structFromDecl(decl)
		if s == nil {
			continue
		}
		result = append(result, *s)
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

func (v *visitor) collectFields(src []*ast.Field) []*ast.Field {
	var flds = make([]*ast.Field, 0, len(src))
	for _, fld := range src {
		tag := tags.Parse(fld.Tag.Value)
		if tag.JsonAppendix() == "inline" {
			inlined := v.getDeclByName(fld.Type.(*ast.Ident).Name)
			if inlined == nil {
				panic("can't resolve inlined field by name")
			}
			inlStruct := v.structFromDecl(*inlined)
			if inlStruct == nil {
				panic("can't inline")
			}
			flds = append(flds, v.collectFields(inlStruct.spec.Fields.List)...)
			continue
		}
		flds = append(flds, fld)
	}
	return flds
}

type taggedStruct struct {
	tags []string
	name string
	spec *ast.StructType
}
