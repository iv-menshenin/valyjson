package test_extr

import (
	"fill/test_any"
	"fill/test_string"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExternalFill(t *testing.T) {
	t.Parallel()
	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var test External
		var def = "default"
		err := test.UnmarshalJSON([]byte(`{"test1":{"comment":"test","level":444},"test2":{"field":"foo","fieldRef":"bar"}}`))
		require.NoError(t, err)
		var expected = External{
			Test01: test_any.TestAllOfSecond{
				Comment: "test",
				Level:   444,
			},
			Test02: test_string.TestStr01{
				Field:    "foo",
				FieldRef: nil,
				DefRef:   &def,
			},
		}
		expected.Test02.FieldRef = new(string)
		*expected.Test02.FieldRef = "bar"
		require.Equal(t, expected, test)
	})
	t.Run("marshal", func(t *testing.T) {
		const expected = `{"test1":{"comment":"test_2","level":456},"test2":{"field":"bar","fieldRef":"foo","defRef":null}}`
		var rF = "foo"
		var obj = External{
			Test01: test_any.TestAllOfSecond{
				Comment: "test_2",
				Level:   456,
			},
			Test02: test_string.TestStr01{
				Field:    "bar",
				FieldRef: &rF,
			},
		}
		data, err := obj.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("marshal_null", func(t *testing.T) {
		var obj *External
		data, err := obj.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
}

func Test_IsZero(t *testing.T) {
	t.Run("just_created", func(t *testing.T) {
		var v External
		require.True(t, v.IsZero())
	})
	t.Run("not_zero_1", func(t *testing.T) {
		var v External
		v.Test01.Comment = "0"
		require.False(t, v.IsZero())
	})
	t.Run("not_zero_2", func(t *testing.T) {
		var v External
		var refStr = ""
		v.Test02.FieldRef = &refStr
		require.False(t, v.IsZero())
	})
}
