package test_extr

import (
	"testing"

	"github.com/stretchr/testify/require"

	"fill/test_any"
	"fill/test_string"
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
		t.Parallel()
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
		t.Parallel()
		var obj *External
		data, err := obj.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
}

func Test_IsZero(t *testing.T) {
	t.Parallel()
	t.Run("just_created", func(t *testing.T) {
		t.Parallel()
		var v External
		require.True(t, v.IsZero())
	})
	t.Run("not_zero_1", func(t *testing.T) {
		t.Parallel()
		var v External
		v.Test01.Comment = "0"
		require.False(t, v.IsZero())
	})
	t.Run("not_zero_2", func(t *testing.T) {
		t.Parallel()
		var v External
		var refStr = ""
		v.Test02.FieldRef = &refStr
		require.False(t, v.IsZero())
	})
}

func TestExternalNested(t *testing.T) {
	t.Parallel()
	t.Run("MarshalJSON", func(t *testing.T) {
		t.Parallel()
		var refA, refB = "A", "B"
		var v = ExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "someComment",
				Level:   12,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "CMD_test",
				Range:   22,
			},
			TestStr01: test_string.TestStr01{
				Field:    "Fld0001A",
				FieldRef: &refA,
				DefRef:   &refB,
			},
		}
		const expected = "{\"command\":\"CMD_test\", \"comment\":\"someComment\", \"defRef\":\"B\", \"field\":\"Fld0001A\", \"fieldRef\":\"A\", \"level\":12, \"range\":22}"
		data, err := v.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Parallel()
		const data = "{\"command\":\"CMD_test_2\", \"comment\":\"someComment_2\", \"defRef\":\"A\", \"field\":\"Fld0001B\", \"fieldRef\":\"B\", \"level\":21, \"range\":33}"
		var refA, refB = "A", "B"
		var expected = ExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "someComment_2",
				Level:   21,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "CMD_test_2",
				Range:   33,
			},
			TestStr01: test_string.TestStr01{
				Field:    "Fld0001B",
				FieldRef: &refB,
				DefRef:   &refA,
			},
		}
		var actual ExternalNested
		require.NoError(t, actual.UnmarshalJSON([]byte(data)))
		require.EqualValues(t, expected, actual)
	})
	t.Run("IsZero", func(t *testing.T) {
		t.Parallel()
		var empty ExternalNested
		require.True(t, empty.IsZero())

		var s string
		var nonEmpty = ExternalNested{
			TestStr01: test_string.TestStr01{
				FieldRef: &s,
			},
		}
		require.False(t, nonEmpty.IsZero())
	})
}
