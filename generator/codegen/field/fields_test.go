package field

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fld_FillStatements(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name  string
		argm  Field
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
			require.Equal(t, test.need, got)
		})
	}
}

func Test_MarshalStatements(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name  string
		argm  Field
		fName string
		need  string
	}
	var cases = []testCase{
		{
			name:  "string_as_string",
			argm:  stringFldMrsh,
			fName: "Field",
			need:  stringFillerMrsh,
		},
		{
			name:  "string_as_string_omit",
			argm:  omitStringFldMrsh,
			fName: "Field",
			need:  omitStringFillerMrsh,
		},
	}
	for i := range cases {
		test := cases[i]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := printAST(t, &ast.BlockStmt{List: test.argm.MarshalStatements(test.fName)})
			require.Equal(t, test.need, got)
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

var stringFld = Field{
	expr: ast.NewIdent("string"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const stringFiller = `{
	if _field := v.Get("field"); _field != nil {
		var valfield []byte
		if valfield, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = *(*string)(unsafe.Pointer(&valfield))
	}
}`

var refStringFld = Field{
	expr: ast.NewIdent("string"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
	isStar: true,
}

const refStringFiller = `{
	if _field := v.Get("field"); _field != nil {
		var valfield []byte
		if valfield, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = (*string)(unsafe.Pointer(&valfield))
	}
}`

var subStringFld = Field{
	expr: ast.NewIdent("SubGroup"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const subStringFiller = `{
	if _field := v.Get("field"); _field != nil {
		var valfield []byte
		if valfield, err = _field.StringBytes(); err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = *(*SubGroup)(unsafe.Pointer(&valfield))
	}
}`

var intFld = Field{
	expr: ast.NewIdent("int"),
	refx: ast.NewIdent("int"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const intFiller = `{
	if _field := v.Get("field"); _field != nil {
		var valfield int
		valfield, err = _field.Int()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = valfield
	}
}`

var subIntFld = Field{
	expr: ast.NewIdent("SubInt"),
	refx: ast.NewIdent("int64"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const subIntFiller = `{
	if _field := v.Get("field"); _field != nil {
		var valfield int64
		valfield, err = _field.Int64()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = SubInt(valfield)
	}
}`

var uuidFld = Field{
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
	if _field := v.Get("field"); _field != nil {
		var valfield uuid.UUID
		b, err := _field.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		valfield, err = uuid.ParseBytes(b)
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = valfield
	}
}`

var timeFld = Field{
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
	if _field := v.Get("field"); _field != nil {
		b, err := _field.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		valfield, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = valfield
	}
}`

var refTimeFld = Field{
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
	if _field := v.Get("field"); _field != nil {
		b, err := _field.StringBytes()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		valfield, err := parseDateTime(string(b))
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.field = new(time.Time)
		*s.field = time.Time(valfield)
	}
}`

var arrayFld = Field{
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
	if _items := v.Get("field"); _items != nil {
		var listA []*fastjson.Value
		listA, err = _items.Array()
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		valitems := s.items[:0]
		if l := len(listA); cap(valitems) < l || (l == 0 && s.items == nil) {
			valitems = make([]DatarentPixelItemsValue, 0, len(listA))
		}
		for _, listElem := range listA {
			var elem DatarentPixelItemsValue
			err = elem.FillFromJSON(listElem, objPath+".")
			if err != nil {
				break
			}
			valitems = append(valitems, DatarentPixelItemsValue(elem))
		}
		if err != nil {
			return fmt.Errorf("error parsing '%s.field' value: %w", objPath, err)
		}
		s.items = valitems
	}
}`

var stringFldMrsh = Field{
	expr: ast.NewIdent("string"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field"},
	},
}

const stringFillerMrsh = "{\n\tif result.Len() > 1 {\n\t\tresult.WriteRune(',')\n\t}\n\tif s.Field != \"\" {\n\t\tresult.WriteString(`\"field\":`)\n\t\tbuf = marshalString(buf[:0], string(s.Field))\n\t\tresult.Write(buf)\n\t} else {\n\t\tresult.WriteString(`\"field\":\"\"`)\n\t}\n}"

var omitStringFldMrsh = Field{
	expr: ast.NewIdent("string"),
	refx: ast.NewIdent("string"),
	tags: map[string][]string{
		"json": {"field", "omitempty"},
	},
}

const omitStringFillerMrsh = "{\n\tif s.Field != \"\" {\n\t\tif result.Len() > 1 {\n\t\t\tresult.WriteRune(',')\n\t\t}\n\t\tresult.WriteString(`\"field\":`)\n\t\tbuf = marshalString(buf[:0], string(s.Field))\n\t\tresult.Write(buf)\n\t}\n}"
