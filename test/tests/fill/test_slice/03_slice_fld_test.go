package test_slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestSlice_Unmarshal(t *testing.T) {
	t.Run("test-all-empty", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[], "ints":null}`))
		require.NoError(t, err)
		require.NotNil(t, test1.Field)
		require.Len(t, test1.Field, 0)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("test-all-nulled", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":null}`))
		require.NoError(t, err)
		require.Nil(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("filled-slice", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":["test1", "test2"], "ints":null}`))
		require.NoError(t, err)
		require.Len(t, test1.Field, 2)
		require.ElementsMatch(t, []string{"test1", "test2"}, test1.Field)
	})
	t.Run("filled-slice-ref", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":[1, 2, 3]}`))
		require.NoError(t, err)
		require.Len(t, test1.FieldRef, 3)
		var expected []*int
		for i := 1; i <= 3; i++ {
			var value = i
			expected = append(expected, &value)
		}
		require.ElementsMatch(t, expected, test1.FieldRef)
	})
	t.Run("filled-slice-ref-with-null", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":[1, 2, null, 3]}`))
		require.NoError(t, err)
		require.Len(t, test1.FieldRef, 4)
		var expected []*int
		for i := 1; i <= 3; i++ {
			var value = i
			expected = append(expected, &value)
			if value == 3 {
				expected = append(expected, nil)
			}
		}
		require.ElementsMatch(t, expected, test1.FieldRef)
	})
}

func Test_TestSlice_Marshal(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var s TestSlice01
		b, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":null,"ints":null}`
		require.JSONEq(t, expected, string(b))
	})
	t.Run("comma_after_filled", func(t *testing.T) {
		var s = TestSlice01{Field: []string{"1"}}
		b, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":["1"],"ints":null}`
		require.JSONEq(t, expected, string(b))
	})
	t.Run("filled_both", func(t *testing.T) {
		var a, b, c = 3, 2, 1
		var s = TestSlice01{
			Field:    []string{"1"},
			FieldRef: []*int{&a, &b, &c},
		}
		data, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":["1"],"ints":[3,2,1]}`
		require.JSONEq(t, expected, string(data))
	})
	t.Run("with_null_elem", func(t *testing.T) {
		var a, b, c = 3, 2, 1
		var s = TestSlice01{
			Field:    []string{"1", ""},
			FieldRef: []*int{&a, nil, &b, &c},
		}
		data, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `{"strs":["1",""],"ints":[3,null,2,1]}`
		require.JSONEq(t, expected, string(data))
	})
}
