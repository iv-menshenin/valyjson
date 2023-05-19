package test

import (
	"crypto/md5" //nolint:gosec // it`s not for security
	"fmt"
	"io"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/iv-menshenin/valyjson/generator"
)

const fillTestsDir = "./tests/fill/"

func Test_fld_FillStatements(t *testing.T) {
	hashes1 := hashDir(fillTestsDir)
	dirs, err := os.ReadDir(fillTestsDir)
	if err != nil {
		t.Fatal(err)
	}

	var cnt int
	for _, d := range dirs {
		if !d.IsDir() {
			sort.Strings(hashes1)
			continue
		}
		dirPath := path.Join(fillTestsDir, d.Name())
		files, err := os.ReadDir(dirPath)
		if err != nil {
			t.Fatal(err)
		}
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
				caseTestFillStatements(t, path.Join(dirPath, fileName))
			})
		}
	}
	if cnt == 0 {
		t.Error("there is no tests")
	}
	hashes2 := hashDir(fillTestsDir)
	sort.Strings(hashes1)
	sort.Strings(hashes2)
	if !reflect.DeepEqual(hashes1, hashes2) {
		t.Error("tests are not regenerated")
	}
}

func caseTestFillStatements(t *testing.T, testFile string) {
	g := generator.New(testFile)
	if err := g.Parse(); err != nil {
		t.Fatal(err)
	}
	g.BuildDecoders()
	g.BuildEncoders()
	g.FixImports()
	g.Print(testFile + ".out.go")
}

func Test_GenerateVJson(t *testing.T) {
	caseTestFillStatements(t, "./vjson/types.go")
}

func hashDir(dirPath string) []string {
	dirs, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	var hashes = make([]string, 0)
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		subDirPath := path.Join(dirPath, d.Name())
		files, err := os.ReadDir(subDirPath)
		if err != nil {
			panic(err)
		}
		for _, f := range files {
			if f.IsDir() {
				hashes = append(hashes, hashDir(path.Join(subDirPath, f.Name()))...)
				continue
			}
			hashes = append(hashes, hashFile(path.Join(subDirPath, f.Name())))
		}
	}
	return hashes
}

func hashFile(fileName string) string {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	h := md5.New()
	if _, err = io.Copy(h, f); err != nil {
		panic(err)
	}
	var buf [md5.Size]byte
	return fmt.Sprintf("%X", h.Sum(buf[:0]))
}
