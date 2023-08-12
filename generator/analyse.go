package generator

import (
	"fmt"
	"go/ast"
	"strings"
	"unicode"

	"github.com/iv-menshenin/go-ast/explorer"
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
		MarshalFunc() []ast.Decl
		AppendJsonFunc() ast.Decl
		ZeroFunc() ast.Decl
	}
)

func (g *Gen) BuildDecoders() {
	var v = visitor{g: g}
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		if structDecl.Tags().Custom() {
			continue
		}
		if structDecl.Tags().EncodersOnly() {
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

func (g *Gen) BuildEncoders() {
	var v = visitor{g: g}
	ast.Walk(&v, g.parsed)
	for _, structDecl := range v.getNormalized() {
		if structDecl.Tags().Custom() {
			continue
		}
		if structDecl.Tags().DecodersOnly() {
			continue
		}
		if marshalFn := structDecl.MarshalFunc(); marshalFn != nil {
			g.result.Decls = append(g.result.Decls, marshalFn...)
		}
		if appendFn := structDecl.AppendJsonFunc(); appendFn != nil {
			g.result.Decls = append(g.result.Decls, appendFn)
		}
		if appendFn := structDecl.ZeroFunc(); appendFn != nil {
			g.result.Decls = append(g.result.Decls, appendFn)
		}
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
	var declProcessed taggedDecl
	defer func() {
		r := recover()
		if r != nil {
			panic(fmt.Errorf("processing error in file %q, processed structure %q raised error: %+v", v.g.fileName, declProcessed.spec.Name.Name, r))
		}
	}()
	for _, decl := range v.decls {
		declProcessed = decl
		if len(decl.tags) == 0 {
			// only tagged structures
			continue
		}
		result = v.processDecl(decl, result)
	}
	return result
}

func (v *visitor) processDecl(decl taggedDecl, result []renderer) []renderer {
	put := func(r renderer) {
		for _, ex := range result {
			if ex.Name() == r.Name() {
				return
			}
		}
		result = append(result, r)
	}
	switch typed := decl.spec.Type.(type) {

	case *ast.StructType:
		if tags.StructTags(decl.tags).Has(tags.TransitHandlers) {
			put(codegen.NewTransitive(decl.spec.Name.Name, decl.tags, typed))
			break
		}
		stct := &ast.StructType{
			Fields: &ast.FieldList{List: v.collectFields(typed.Fields.List)}, // uninline
		}
		put(codegen.NewStruct(decl.spec.Name.Name, decl.tags, stct))

	case *ast.MapType:
		put(codegen.NewMap(decl.spec.Name.Name, decl.tags, typed))

	case *ast.ArrayType:
		if el, ok := typed.Elt.(*ast.Ident); ok {
			// elements of the array have no tags, so we will consider the tags of the child structure
			elDecl := v.getDeclByName(el.Name)
			if elDecl != nil && len(elDecl.tags) == 0 {
				var elDeclT = *elDecl
				elDeclT.tags = decl.tags // inherits tags from an heir
				result = v.processDecl(elDeclT, result)
			}
		}
		put(codegen.NewArray(decl.spec.Name.Name, decl.tags, typed))

	case *ast.Ident:
		put(codegen.NewTransitive(decl.spec.Name.Name, decl.tags, typed))

	case *ast.SelectorExpr:
		put(codegen.NewTransitive(decl.spec.Name.Name, decl.tags, typed))

	default:
		panic("unsupported")
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
				panic(fmt.Errorf("all fields should have tags, but %q haven't", fld.Names[0].Name))
			}
		}
		if sel, ok := fld.Type.(*ast.SelectorExpr); ok && sel.Sel.Obj == nil {
			// try to resolve external type
			sel.Sel.Obj = v.resolveExternal(sel)
		}
		if v.over != nil {
			if i, ok := fld.Type.(*ast.Ident); ok && unicode.IsUpper([]rune(i.Name)[0]) {
				fld.Type = &ast.SelectorExpr{
					X:   v.over,
					Sel: i,
				}
			}
		}
		if tag.JsonAppendix() == "inline" {
			flds = append(flds, v.exploreInlined(fld)...)
			continue
		}
		flds = append(flds, fld)
	}
	return flds
}

func (v *visitor) resolveExternal(sel *ast.SelectorExpr) *ast.Object {
	packages, ok := v.g.packages[sel.X.(*ast.Ident).String()]
	if ok {
		for _, p := range packages {
			if p.Kind == explorer.PkgKindSystem {
				continue
			}
			if decl := v.g.ExploreType(p, sel.Sel.Name); decl != nil {
				return &ast.Object{
					Kind: ast.Typ,
					Name: sel.Sel.Name,
					Decl: decl,
				}
			}
		}
	}
	return nil
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
		pkg, err := v.g.discovery.GetPackage(packIdent.Name, inlined.Sel.Name)
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
		if len(fld.Names) == 0 {
			return v1.collectFields(stct.Fields.List)
		}
		if len(fld.Names) > 1 {
			panic("can't inline named fields")
		}
		return []*ast.Field{fld}

	default:
		panic(fmt.Errorf("can't inline struct kind %+v", fld.Type))
	}
}
