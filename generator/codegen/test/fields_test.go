package test

import (
	"io/ioutil"
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
		cnt++
		t.Run(f.Name(), func(t *testing.T) {
			t.Parallel()
			caseTestFillStatements(t, fillTestsDir+f.Name())
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
	g.Print(testFile + ".out")
}
