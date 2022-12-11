package field

import (
	"bufio"
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
	"testing"
)

func Test_fld_FillStatements(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name  string
		argm  fld
		fName string
		need  string
	}
	var cases = []testCase{
		{
			name:  "string_as_string",
			argm:  stringFld,
			fName: "field",
			need:  stringFiller,
		},
		{
			name:  "string_as_SubGroup",
			argm:  subStringFld,
			fName: "field",
			need:  subStringFiller,
		},
		{
			name:  "string_as_ref_string",
			argm:  refStringFld,
			fName: "field",
			need:  refStringFiller,
		},
		{
			name:  "int_as_int",
			argm:  intFld,
			fName: "field",
			need:  intFiller,
		},
		{
			name:  "int_as_SubInt",
			argm:  subIntFld,
			fName: "field",
			need:  subIntFiller,
		},
		{
			name:  "uuid_UUID",
			argm:  uuidFld,
			fName: "field",
			need:  uuidFiller,
		},
		{
			name:  "time_Time",
			argm:  timeFld,
			fName: "field",
			need:  timeFiller,
		},
		{
			name:  "ref_time_Time",
			argm:  refTimeFld,
			fName: "field",
			need:  refTimeFiller,
		},
		{
			name:  "array_Struct",
			argm:  arrayFld,
			fName: "items",
			need:  arrayFiller,
		},
	}
	for i := range cases {
		test := cases[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := printAST(t, &ast.BlockStmt{List: test.argm.FillStatements(test.fName)})
			checkCode(t, got, test.need)
		})
	}
}

func printAST(t *testing.T, a ast.Node) string {
	var b strings.Builder
	if err := printer.Fprint(&b, token.NewFileSet(), a); err != nil {
		t.Fatal(err)
	}
	return b.String()
}

func checkCode(t *testing.T, generated, reference string) {
	t.Helper()
	a := bufio.NewScanner(strings.NewReader(generated))
	b := bufio.NewScanner(strings.NewReader(reference))
	var isError bool
	var lines []string
	for {
		if !a.Scan() {
			lines = append(lines, "===END OF GENERATED===")
			for b.Scan() {
				if strings.TrimSpace(b.Text()) != "" {
					isError = true
					lines = append(lines, ">"+b.Text())
				}
			}
			break
		}
		if !b.Scan() {
			lines = append(lines, "===END OF REFERENCED===")
			for a.Scan() {
				if strings.TrimSpace(b.Text()) != "" {
					isError = true
					lines = append(lines, "<"+b.Text())
				}
			}
			break
		}
		if a.Text() != b.Text() {
			lines = append(lines, "-"+b.Text())
			lines = append(lines, "+"+a.Text())
			isError = true
		} else {
			lines = append(lines, "="+a.Text())
		}
	}
	if isError {
		t.Errorf("Matching error:\n%s\n%s", strings.Join(lines, "\n"), generated)
	}
}

var stringFld = fld{
	expr: ast.NewIdent("string"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const stringFiller = `{
	if field := v.Get("field"); field != nil {
		if field.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", field.Type())
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		valfield := field.String()
		s.field = valfield
	}
}`

var refStringFld = fld{
	expr: ast.NewIdent("string"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
	isStar: true,
}

const refStringFiller = `{
	if field := v.Get("field"); field != nil {
		if field.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", field.Type())
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		valfield := field.String()
		s.field = new(string)
		*s.field = string(valfield)
	}
}`

var subStringFld = fld{
	expr: ast.NewIdent("SubGroup"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const subStringFiller = `{
	if field := v.Get("field"); field != nil {
		if field.Type() != fastjson.TypeString {
			err = fmt.Errorf("value doesn't contain string; it contains %s", field.Type())
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		valfield := field.String()
		s.field = SubGroup(valfield)
	}
}`

var intFld = fld{
	expr: ast.NewIdent("int"),
	refx: ast.NewIdent("int"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const intFiller = `{
	if field := v.Get("field"); field != nil {
		var valfield int
		valfield, err = field.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.field = valfield
	}
}`

var subIntFld = fld{
	expr: ast.NewIdent("SubInt"),
	refx: ast.NewIdent("int64"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const subIntFiller = `{
	if field := v.Get("field"); field != nil {
		var valfield int64
		valfield, err = field.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.field = SubInt(valfield)
	}
}`

var uuidFld = fld{
	expr: &ast.SelectorExpr{
		X:   ast.NewIdent("uuid"),
		Sel: ast.NewIdent("UUID"),
	},
	refx: &ast.SelectorExpr{
		X:   ast.NewIdent("uuid"),
		Sel: ast.NewIdent("UUID"),
	},
	tags: map[string][]string{
		"json": {"field"},
	},
}

const uuidFiller = `{
	if field := v.Get("field"); field != nil {
		var valfield uuid.UUID
		b, err := field.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		err = valfield.UnmarshalText(b)
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.field = valfield
	}
}`

var timeFld = fld{
	expr: &ast.SelectorExpr{
		X:   ast.NewIdent("time"),
		Sel: ast.NewIdent("Time"),
	},
	refx: &ast.SelectorExpr{
		X:   ast.NewIdent("time"),
		Sel: ast.NewIdent("Time"),
	},
	tags: map[string][]string{
		"json": {"field"},
	},
}

const timeFiller = `{
	if field := v.Get("field"); field != nil {
		valfield, err := time.Parse(time.RFC3339, field.String())
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.field = valfield
	}
}`

var refTimeFld = fld{
	expr: &ast.SelectorExpr{
		X:   ast.NewIdent("time"),
		Sel: ast.NewIdent("Time"),
	},
	refx: &ast.SelectorExpr{
		X:   ast.NewIdent("time"),
		Sel: ast.NewIdent("Time"),
	},
	tags: map[string][]string{
		"json": {"field"},
	},
	isStar: true,
}

const refTimeFiller = `{
	if field := v.Get("field"); field != nil {
		valfield, err := time.Parse(time.RFC3339, field.String())
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.field = new(time.Time)
		*s.field = time.Time(valfield)
	}
}`

var arrayFld = fld{
	expr: &ast.ArrayType{
		Elt: ast.NewIdent("DatarentPixelItemsValue"),
	},
	refx: &ast.ArrayType{
		Elt: ast.NewIdent("DatarentPixelItemsValue"),
	},
	tags: map[string][]string{
		"json": {"field"},
	},
}

const arrayFiller = `{
	if items := v.Get("field"); items != nil {
		var listA []*fastjson.Value
		listA, err = items.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		valitems := make([]DatarentPixelItemsValue, 0, len(listA))
		for _, listElem := range listA {
			var elem DatarentPixelItemsValue
			err = elem.FillFromJson(listElem, objPath+".")
			if err != nil {
				break
			}
			valitems = append(valitems, DatarentPixelItemsValue(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%sfield' value: %w", objPath, err)
		}
		s.items = valitems
	}
}`
