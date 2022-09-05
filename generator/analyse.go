package generator

import (
	"go/ast"
	"strings"
	"valyjson/generator/codegen"
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

func (g *Gen) Analyse() {
	var v visitor
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.decls {
		var fillers []ast.Stmt
		switch strc := structDecl.spec.Type.(type) {

		case *ast.StructType:
			for _, fld := range strc.Fields.List {
				fillers = append(fillers, codegen.NewFieldFillerStmt(fld)...)
			}
			g.result.Decls = append(
				g.result.Decls,
				codegen.NewFillerFunc(structDecl.spec.Name.Name, fillers...),
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

func extractTags(comment *ast.CommentGroup) []string {
	if comment == nil {
		return nil
	}
	var tags []string
	for _, text := range comment.List {
		if text == nil {
			continue
		}
		if commentLine := strings.TrimLeft(text.Text, "/ \t"); strings.HasPrefix(commentLine, "valyjson:") {
			splitPostfix := strings.Split(commentLine[9:], ",")
			for _, tagRaw := range splitPostfix {
				if tag := strings.ToLower(strings.TrimSpace(tagRaw)); tag != "" {
					tags = append(tags, tag)
				}
			}
		}
	}
	return tags
}
