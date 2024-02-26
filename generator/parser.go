package generator

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path"
	"strings"

	"github.com/iv-menshenin/go-ast/explorer"
	"github.com/iv-menshenin/valyjson/generator/discoverer"
	"github.com/iv-menshenin/valyjson/generator/static"
)

type (
	Gen struct {
		fileName string
		parsed   ast.Node
		result   ast.File
		packageN string

		discovery *discoverer.Discoverer
		packages  map[string][]explorer.Package
	}
)

func (g *Gen) Parse() (err error) {
	if err = g.discovery.Discover(); err != nil {
		panic(err)
	}
	g.parsed, err = parseGo(g.fileName)
	if err != nil {
		return err
	}
	g.packageN = g.parsed.(*ast.File).Name.Name
	if err == nil {
		g.result.Name = g.parsed.(*ast.File).Name
	}
	for _, imp := range g.parsed.(*ast.File).Imports {
		var name string
		if imp.Name != nil {
			name = imp.Name.Name
		}
		pkgPath := strings.Trim(imp.Path.Value, `"`)
		if !strings.HasPrefix(pkgPath, g.discovery.GetModuleName()) {
			continue
		}
		if name == "" {
			// FIXME @menshenin the name of package must be read from the package itself
			split := strings.Split(pkgPath, "/")
			name = split[len(split)-1]
		}
		g.packages[name] = append(
			g.packages[name],
			explorer.Package{
				Path: pkgPath,
				Kind: explorer.PkgKindInternal,
			},
		)
	}
	return
}

func (g *Gen) FixImports(internals ...string) {
	explorer.RegisterPackage("jwriter", explorer.Package{
		Path: "github.com/mailru/easyjson/jwriter",
		Kind: explorer.PkgKindExternal,
	})
	// discovery used imports and build their declaration
	for name, pkg := range g.packages {
		for i := range pkg {
			explorer.RegisterPackage(name, pkg[i])
		}
	}
	for i := 0; i < len(internals); i += 2 {
		explorer.RegisterPackage(internals[i], explorer.Package{
			Path: internals[i+1],
			Kind: explorer.PkgKindInternal,
		})
	}
	g.imports()
}

func (g *Gen) imports() {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("PANIC at %q: %+v", g.fileName, r))
		}
	}()
	discovery := explorer.New()
	discovery.Explore(&g.result)
	var decls = make([]ast.Decl, 0, len(g.result.Decls)+1)

	imports := discovery.ImportSpec()
	if len(imports) > 0 {
		decls = append(decls, &ast.GenDecl{
			Tok:   token.IMPORT,
			Specs: discovery.ImportSpec(),
		})
	}
	decls = append(decls, g.result.Decls...)
	g.result.Decls = decls
}

func (g *Gen) Print(name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fmtGOFile(name)
	defer f.Close()
	_, err = fmt.Fprint(f, "// Code generated [github.com/iv-menshenin/valyjson]; DO NOT EDIT.\n")
	if err != nil {
		panic(err)
	}

	if err := printer.Fprint(f, token.NewFileSet(), &g.result); err != nil {
		panic(err)
	}
	err = static.Print(static.Context{
		Package: g.packageN,
	}, path.Dir(name))
	if err != nil {
		panic(err)
	}
}

func (g *Gen) ExploreType(p explorer.Package, objName string) *ast.TypeSpec {
	parsed, err := parsePackage(g.discovery.PackagePathFromPath(p.Path))
	if err != nil {
		return nil
	}
	for name, pack := range parsed {
		if strings.HasSuffix(name, "_test") {
			continue
		}
		for _, file := range pack.Files {
			for _, decl := range file.Decls {
				genDecl, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				for _, spec := range genDecl.Specs {
					t, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if t.Name.Name == objName {
						return t
					}
				}
			}
		}
	}
	return nil
}

func fmtGOFile(fileName string) (err error) {
	var fileSet = token.NewFileSet()
	f, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if e := out.Close(); e != nil && err == nil {
			err = e
		}
	}()
	return format.Node(out, fileSet, f)
}

func parseGo(file string) (ast.Node, error) {
	f := token.NewFileSet()
	return parser.ParseFile(f, file, nil, parser.ParseComments|parser.AllErrors)
}

func parsePackage(path string) (map[string]*ast.Package, error) {
	f := token.NewFileSet()
	return parser.ParseDir(f, path, nil, parser.ParseComments|parser.AllErrors)
}

func New(file string) *Gen {
	return &Gen{
		fileName:  file,
		discovery: discoverer.New(path.Dir(file)),
		packages:  make(map[string][]explorer.Package),
	}
}
