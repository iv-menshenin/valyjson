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

func Test_TestSlice_Unarshal_Bad(t *testing.T) {
	t.Run("empty-json", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON(nil)
		require.Error(t, err)
	})
	t.Run("bad-json", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte{0, 1, 2, 3})
		require.Error(t, err)
	})
	t.Run("invalid-object", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null,"strs":null}`))
		require.Error(t, err)
	})
	t.Run("invalid-object-field-strs", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":{},"ints":null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain array")
	})
	t.Run("invalid-object-vals-strs", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[{},{}],"ints":null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("invalid-object-field-ints", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[],"ints":2}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain array")
	})
	t.Run("invalid-object-vals-ints", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[],"ints":[{},{}]}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain number")
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
	t.Run("null", func(t *testing.T) {
		var s *TestSlice01
		data, err := s.MarshalJSON()
		require.NoError(t, err)
		const expected = `null`
		require.JSONEq(t, expected, string(data))
	})
}

func TestTestSlice01_validate(t *testing.T) {
	t.Run("ints_appears_twice", func(t *testing.T) {
		const jData = `{"ints":[],"ints":[1,2]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v, "(root)")
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "ints")
	})
	t.Run("not_object", func(t *testing.T) {
		const jData = `122333`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v, "(root)")
		require.ErrorContains(t, err, "value doesn't contain object")
		require.ErrorContains(t, err, "it contains number")
	})
	t.Run("strs_appears_twice", func(t *testing.T) {
		const jData = `{"strs":[],"strs":["1","2"]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v, "(root)")
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "strs")
	})
	t.Run("both_appears_twice", func(t *testing.T) {
		const jData = `{"strs":[],"ints":[],"strs":["1","2"],"ints":[1,2]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v, "(root)")
		require.ErrorContains(t, err, "appears in the object twice")
	})
}

func TestTestSlice01_IsZero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test TestSlice01
		require.True(t, test.IsZero())
	})
	t.Run("nonzero1", func(t *testing.T) {
		var test = TestSlice01{Field: []string{}}
		require.False(t, test.IsZero())
	})
	t.Run("nonzero2", func(t *testing.T) {
		var test = TestSlice01{FieldRef: []*int{}}
		require.False(t, test.IsZero())
	})
}
