package generator

import (
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
	for _, structDecl := range v.decls {
		switch strc := structDecl.spec.Type.(type) {

		case *ast.StructType:
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewUnmarshalFunc(structDecl.spec.Name.Name, structDecl.tags)...,
			)
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewFillerFunc(structDecl.spec.Name.Name, strc.Fields.List, structDecl.tags),
			)
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewValidatorFunc(structDecl.spec.Name.Name, strc.Fields.List, structDecl.tags),
			)
		}
	}
}

func (g *Gen) BuildJsoners() {
	var v visitor
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.decls {
		switch strc := structDecl.spec.Type.(type) {

		case *ast.StructType:
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewMarshalFunc(structDecl.spec.Name.Name, structDecl.tags),
			)
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewAppendJsonFunc(structDecl.spec.Name.Name, strc.Fields.List, structDecl.tags),
			)
		}
	}
}

func (v *visitor) Visit(node ast.Node) (w ast.Visitor) {
	switch currNode := node.(type) {

	case *ast.GenDecl:
		if tags := extractTags(currNode.Doc); len(tags) > 0 {
			for _, declSpec := range currNode.Specs {
				switch spec := declSpec.(type) {
				case *ast.TypeSpec:
					v.decls = append(v.decls, taggedDecl{
						spec: spec,
						tags: tags,
					})
				}
			}
		}

	case *ast.TypeSpec:
		if tags := extractTags(currNode.Doc); len(tags) > 0 {
			v.decls = append(v.decls, taggedDecl{
				spec: currNode,
				tags: tags,
			})
		}
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
