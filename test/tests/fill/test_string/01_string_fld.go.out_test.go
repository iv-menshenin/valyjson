package test_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestStr01_Unmarshal(t *testing.T) {
	t.Run("test-all-omitted", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Empty(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("test-default", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Empty(t, test1.Field)
		require.Nil(t, test1.FieldRef)
		require.NotNil(t, test1.DefRef)
		require.EqualValues(t, "default", *test1.DefRef)
	})
	t.Run("nil-ref-omitted", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field_filled", test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("ref-wrong-type", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"fieldRef": 0.123}`))
		require.Error(t, err)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("fill-fields", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field", "fieldRef": "test_fieldRef", "defRef": "000"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field", test1.Field)
		require.NotNil(t, test1.FieldRef)
		require.Equal(t, "test_fieldRef", *test1.FieldRef)
		require.NotNil(t, test1.DefRef)
		require.Equal(t, "000", *test1.DefRef)
	})
	t.Run("extra-fld-omitted", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled", "field-2": "test"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field_filled", test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("null-for-ref-field", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled", "fieldRef": null}`))
		require.NoError(t, err)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("wrong-type-for-field", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": 112}`))
		require.Error(t, err)
	})
	t.Run("test-double", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "foo", "field": "bar"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "the '(root).field' field appears in the object twice")
	})
}

func Test_TestStr01_Marshal(t *testing.T) {
	t.Run("null-fields", func(t *testing.T) {
		const expected = `{"field":"foo-bar","fieldRef":null,"defRef":null}`
		var test = TestStr01{
			Field:    "foo-bar",
			FieldRef: nil,
			DefRef:   nil,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("null-struct", func(t *testing.T) {
		const expected = `null`
		var test1 *TestStr01
		data, err := test1.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("fill-fields", func(t *testing.T) {
		const expected = `{"field":"foo-bar","fieldRef":"nil/null","defRef":"nil/null"}`
		var str = "nil/null"
		var test = TestStr01{
			Field:    "foo-bar",
			FieldRef: &str,
			DefRef:   &str,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("multiline", func(t *testing.T) {
		const expected = `{"field":"test\nmulti\nlined","fieldRef":null,"defRef":null}`
		var test = TestStr01{
			Field: "test\nmulti\nlined",
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("empty-field", func(t *testing.T) {
		const expected = `{"field":"","fieldRef":null,"defRef":null}`
		var test = TestStr01{
			Field: "",
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("special-runes", func(t *testing.T) {
		const expected = `{"field":"\"quoted\\slashed\ttabbed","fieldRef":null,"defRef":null}`
		var test = TestStr01{
			Field: `"quoted\slashed	tabbed`,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
}

func Test_TestStr02_Marshal(t *testing.T) {
	t.Run("null-fields", func(t *testing.T) {
		const expected = `{"field":"foo-bar","fieldRef":null,"string":"nil"}`
		var test = TestStr02{
			Field:    "foo-bar",
			FieldRef: nil,
			String:   "nil",
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("null-struct", func(t *testing.T) {
		const expected = "null"
		var test2 *TestStr02
		data, err := test2.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("fill-fields", func(t *testing.T) {
		const expected = `{"field":"foo-bar","fieldRef":"nil/null","string":""}`
		var str = "nil/null"
		var test = TestStr02{
			Field:    "foo-bar",
			FieldRef: &str,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("ref-spec-symbols", func(t *testing.T) {
		const expected = "{\"field\":\"foo-bar\",\"fieldRef\":\"\\twe\\nwrap\\nall the \\\"world\\\"\",\"string\":\"\"}"
		var str = "\twe\nwrap\nall the \"world\""
		var test = TestStr02{
			Field:    "foo-bar",
			FieldRef: &str,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("empty-field", func(t *testing.T) {
		const expected = `{"field":"","fieldRef":null,"string":""}`
		var test = TestStr02{}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
}

func Test_Strings_Allocations(t *testing.T) {
	t.Parallel()
	t.Run("ref-allocation", func(t *testing.T) {
		t.Parallel()
		n := testing.AllocsPerRun(100, func() {
			var test1 TestStr01
			_ = test1.UnmarshalJSON([]byte(`{"field": "test_field", "defRef": "foo bar"}`))
		})
		require.LessOrEqual(t, n, float64(1)) // one allocation for ref
	})
	t.Run("no-allocations", func(t *testing.T) {
		t.Parallel()
		n := testing.AllocsPerRun(100, func() {
			var test2 TestStr02
			_ = test2.UnmarshalJSON([]byte(`{"field": "test_field", "string": "foo bar ipsum"}`))
		})
		require.LessOrEqual(t, n, float64(0))
	})
}

func Benchmark_TestStr_Alloc(b *testing.B) {
	b.Run("TestStr01", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			var test1 TestStr01
			_ = test1.UnmarshalJSON([]byte(`{"field": "test_field", "defRef": "foo bar"}`))
		}
	})
	b.Run("TestStr02", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			var test1 TestStr02
			_ = test1.UnmarshalJSON([]byte(`{"field": "test_field", "string": "foo bar"}`))
		}
	})
}

func Test_TestStr02_Unmarshal(t *testing.T) {
	t.Run("test-all-omitted", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Empty(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("field-omitted", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field_filled", test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("fill-fields", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field", "fieldRef": "test_fieldRef"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field", test1.Field)
		require.NotNil(t, test1.FieldRef)
		require.Equal(t, "test_fieldRef", *test1.FieldRef)
	})
	t.Run("null-for-ref-field", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled", "fieldRef": null}`))
		require.NoError(t, err)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("wrong-type-for-field", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": 112}`))
		require.Error(t, err)
	})

	t.Run("test-extra-fld", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"badName": "test"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "unexpected field '(root).badName'")
	})
	t.Run("test-double", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "foo", "field": "bar"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "the '(root).field' field appears in the object twice")
	})
}

func Test_TestSubTypeString(t *testing.T) {
	t.Run("check_type", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"string": "filled well"}`))
		require.NoError(t, err)
		require.Equal(t, FieldValueString("filled well"), test1.String)
	})
	t.Run("test-allocs", func(t *testing.T) {
		n := testing.AllocsPerRun(100, func() {
			var test1 TestStr02
			_ = test1.UnmarshalJSON([]byte(`{"string": "filled well foo/bar"}`))
		})
		require.LessOrEqual(t, n, float64(0))
	})
	t.Run("test_default", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "foo", "fieldRef": "bar"}`))
		require.NoError(t, err)
		require.EqualValues(t, "value-foo-bar", test1.String)
	})
}

func Test_Zero(t *testing.T) {
	t.Run("TestStr01", func(t *testing.T) {
		var test TestStr01
		require.True(t, test.IsZero())
	})
	t.Run("TestStr01_not_zero", func(t *testing.T) {
		var test1 TestStr01
		test1.Field = "."
		require.False(t, test1.IsZero())
		var test2 TestStr01
		test2.FieldRef = &test1.Field
		require.False(t, test2.IsZero())
		var test3 TestStr01
		test3.DefRef = &test1.Field
		require.False(t, test3.IsZero())
	})
	t.Run("TestStr02", func(t *testing.T) {
		var test TestStr02
		require.True(t, test.IsZero())
	})
	t.Run("TestStr02_not_zero", func(t *testing.T) {
		var test TestStr02
		test.String = "value-foo-bar"
		require.False(t, test.IsZero())
	})
}
