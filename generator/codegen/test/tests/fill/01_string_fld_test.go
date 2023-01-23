package testo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTestStr01(t *testing.T) {
	t.Run("test-all-omitted", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Empty(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("nil-ref-omitted", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field_filled", test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("fill-fields", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field", "fieldRef": "test_fieldRef"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field", test1.Field)
		require.NotNil(t, test1.FieldRef)
		require.Equal(t, "test_fieldRef", *test1.FieldRef)
	})
	t.Run("extra-fld-omitted", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled", "field-2": "test"}`))
		require.NoError(t, err)
		require.Equal(t, "test_field_filled", test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("wrong-type-for-ref-field", func(t *testing.T) {
		var test1 TestStr01
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled", "fieldRef": nil}`))
		require.Error(t, err)
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
		require.ErrorContains(t, err, "the 'field' field appears in the object twice")
	})
}

func TestTestStr02(t *testing.T) {
	t.Run("test-all-omitted", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Empty(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("nil-ref-omitted", func(t *testing.T) {
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
	t.Run("wrong-type-for-ref-field", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "test_field_filled", "fieldRef": nil}`))
		require.Error(t, err)
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
		require.ErrorContains(t, err, "unexpected field 'badName'")
	})
	t.Run("test-double", func(t *testing.T) {
		var test1 TestStr02
		err := test1.UnmarshalJSON([]byte(`{"field": "foo", "field": "bar"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "the 'field' field appears in the object twice")
	})
}
