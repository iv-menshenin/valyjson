package main

import (
	"testing"

	"github.com/iv-menshenin/valyjson/generator"
)

func Test_Nested_Generation(t *testing.T) {
	g := generator.New("./test/tests/fill/test_extr/nested.go")
	if err := g.Parse(); err != nil {
		t.Fatal(err)
	}

	g.BuildDecoders()
	g.BuildEncoders()
	g.FixImports()
	g.Print("./test/tests/fill/test_extr/nested.go.out.go")
}
