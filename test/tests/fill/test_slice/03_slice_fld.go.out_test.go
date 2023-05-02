package test_slice

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestSlice01_Unmarshal(t *testing.T) {
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

func Test_TestSlice_Unmarshal_Bad(t *testing.T) {
	t.Run("empty-json", func(t *testing.T) {
		var test1 TestSlice01
		require.Error(t, test1.UnmarshalJSON(nil))
		var test2 TestSlice02
		require.Error(t, test2.UnmarshalJSON(nil))
		var test3 TestSlice03
		require.Error(t, test3.UnmarshalJSON(nil))
	})
	t.Run("bad-json", func(t *testing.T) {
		var test1 TestSlice01
		require.Error(t, test1.UnmarshalJSON([]byte{0, 1, 2, 3}))
		var test2 TestSlice02
		require.Error(t, test2.UnmarshalJSON([]byte{0, 1, 2, 3}))
		var test3 TestSlice03
		require.Error(t, test3.UnmarshalJSON([]byte{0, 1, 2, 3}))
	})
	t.Run("invalid-object", func(t *testing.T) {
		var test1 TestSlice01
		require.Error(t, test1.UnmarshalJSON([]byte(`{"strs":null,"strs":null}`)))
		var test2 TestSlice02
		require.Error(t, test2.UnmarshalJSON([]byte(`{}`)))
		var test3 TestSlice03
		require.Error(t, test3.UnmarshalJSON([]byte(`[]`)))
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

func Test_TestSlice01_Marshal(t *testing.T) {
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
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "ints")
	})
	t.Run("not_object", func(t *testing.T) {
		const jData = `122333`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "value doesn't contain object")
		require.ErrorContains(t, err, "it contains number")
	})
	t.Run("strs_appears_twice", func(t *testing.T) {
		const jData = `{"strs":[],"strs":["1","2"]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "strs")
	})
	t.Run("data_appears_twice", func(t *testing.T) {
		const jData = `{"data":{},"data":{},"extended":{}}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice03
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
		require.ErrorContains(t, err, "data")
	})
	t.Run("both_appears_twice", func(t *testing.T) {
		const jData = `{"strs":[],"ints":[],"strs":["1","2"],"ints":[1,2]}`
		parser := jsonParserTestSlice01.Get()
		v, err := parser.ParseBytes([]byte(jData))
		require.NoError(t, err)
		var test TestSlice01
		err = test.validate(v)
		require.ErrorContains(t, err, "appears in the object twice")
	})
}

func TestTestSlice_IsZero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test1 TestSlice01
		require.True(t, test1.IsZero())
		var test2 TestSlice02
		require.True(t, test2.IsZero())
		var test3 TestSlice03
		require.True(t, test3.IsZero())
	})
	t.Run("nonzero", func(t *testing.T) {
		var test1 = TestSlice01{Field: []string{}}
		require.False(t, test1.IsZero())
		var test1_2 = TestSlice01{FieldRef: []*int{}}
		require.False(t, test1_2.IsZero())
		var test2 = TestSlice02{TestSlice03{Data: 3}}
		require.False(t, test2.IsZero())
		var test3 = TestSlice03{Data: 44}
		require.False(t, test3.IsZero())
	})
}

func Test_TestSlice02_Unmarshal(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`null`))
		require.NoError(t, err)
		require.Nil(t, test1)
		require.Len(t, test1, 0)
	})
	t.Run("empty-array", func(t *testing.T) {
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[]`))
		require.NoError(t, err)
		require.NotNil(t, test1)
		require.Len(t, test1, 0)
	})
	t.Run("filled-array", func(t *testing.T) {
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[{"data": 13},{"data": 54}]`))
		require.NoError(t, err)
		require.NotNil(t, test1)
		require.Len(t, test1, 2)
		require.Equal(t, TestSlice02{{Data: 13}, {Data: 54}}, test1)
	})
	t.Run("null-instead-of-object", func(t *testing.T) {
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[{"data": 13},null]`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("filled-array-2", func(t *testing.T) {
		var test1 TestSlice02
		err := test1.UnmarshalJSON([]byte(`[{"data": 13},{},{"data": 54}]`))
		require.NoError(t, err)
		require.NotNil(t, test1)
		require.Len(t, test1, 3)
		require.Equal(t, TestSlice02{{Data: 13}, {}, {Data: 54}}, test1)
	})
}

func Test_TestSlice02_Marshal(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		var test TestSlice02
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("empty", func(t *testing.T) {
		var test = TestSlice02{}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "[]", string(data))
	})
	t.Run("marshal", func(t *testing.T) {
		var test = TestSlice02{{Data: 190}, {Data: 191}}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `[{"data":190},{"data":191}]`, string(data))
	})
}

func Test_TestSlice03_Marshal(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		var test *TestSlice03
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("empty", func(t *testing.T) {
		var test = TestSlice03{}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"data":0}`, string(data))
	})
}
