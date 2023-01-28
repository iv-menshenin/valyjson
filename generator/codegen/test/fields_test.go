package test

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"

	"github.com/iv-menshenin/valyjson/generator"
)

const fillTestsDir = "./tests/fill/"

func Test_fld_FillStatements(t *testing.T) {
	t.Parallel()
	files, err := ioutil.ReadDir(fillTestsDir)
	if err != nil {
		t.Fatal(err)
	}

	var cnt int
	for _, f := range files {
		if path.Ext(f.Name()) != ".go" {
			continue
		}
		if strings.HasSuffix(f.Name(), ".out.go") {
			continue
		}
		if strings.HasSuffix(f.Name(), "_test.go") {
			continue
		}
		if f.Name() == "valyjson_utils.go" {
			continue
		}
		cnt++
		fileName := f.Name()
		t.Run(f.Name(), func(t *testing.T) {
			t.Parallel()
			caseTestFillStatements(t, fillTestsDir+fileName)
		})
	}
	if cnt == 0 {
		t.Error("there is no tests")
	}
}

func caseTestFillStatements(t *testing.T, testFile string) {
	g := generator.New(testFile)
	if err := g.Parse(); err != nil {
		t.Fatal(err)
	}
	g.BuildFillers()
	g.FixImports()
	g.Print(testFile + ".out.go")
}

func Test_GenerateVJson(t *testing.T) {
	caseTestFillStatements(t, "./vjson/types.go")
}
