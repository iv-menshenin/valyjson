package test_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestStr01(t *testing.T) {
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

func Test_TestStr02(t *testing.T) {
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