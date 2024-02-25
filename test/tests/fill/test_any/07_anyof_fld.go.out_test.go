package test_any

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestAllOfAnyOf01(t *testing.T) {
	t.Parallel()
	t.Run("one_int_", func(t *testing.T) {
		t.Parallel()
		var value = TestAllOf01{
			TestAllOfFirstIsOne: TestAllOfFirstIsOne{
				OneOf: 33,
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"value":33, "comment":""}`, string(data))

		var outValue TestAllOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":332}`), &outValue))
		val, ok := outValue.OneOf.(TestOneOfInteger)
		require.True(t, ok)
		require.EqualValues(t, 332, val)
		require.Empty(t, outValue.Comment)
		require.Empty(t, outValue.Level)

		outValue.Reset()
		require.Empty(t, outValue)
	})
	t.Run("one_string_", func(t *testing.T) {
		t.Parallel()
		var value = TestAllOf01{
			TestAllOfFirstIsOne: TestAllOfFirstIsOne{
				OneOf: "some string",
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"value":"some string", "comment":""}`, string(data))

		var outValue TestAllOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":"some string 1"}`), &outValue))
		val, ok := outValue.OneOf.(TestOneOfString)
		require.True(t, ok)
		require.EqualValues(t, "some string 1", val)
		require.Empty(t, outValue.Comment)
		require.Empty(t, outValue.Level)

		outValue.Reset()
		require.Empty(t, outValue)
	})
	t.Run("one_struct_", func(t *testing.T) {
		t.Parallel()
		var value = TestAllOf01{
			TestAllOfFirstIsOne: TestAllOfFirstIsOne{
				OneOf: TestOneOfStruct{
					Class: "chair",
					Value: 322.5,
				},
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"value":{"class":"chair","width":322.5}, "comment":""}`, string(data))

		var outValue TestAllOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":{"class":"table","width":256.7}}`), &outValue))
		val, ok := outValue.OneOf.(TestOneOfStruct)
		require.True(t, ok)
		require.EqualValues(t, TestOneOfStruct{
			Class: "table",
			Value: 256.7,
		}, val)
		require.Empty(t, outValue.Comment)
		require.Empty(t, outValue.Level)

		outValue.Reset()
		require.Empty(t, outValue)
	})
	t.Run("all_of", func(t *testing.T) {
		t.Parallel()
		var value = TestAllOf01{
			TestAllOfFirstIsOne: TestAllOfFirstIsOne{
				OneOf: TestOneOfStruct{
					Class: "chair",
					Value: 12.5,
				},
			},
			TestAllOfSecond: TestAllOfSecond{
				Comment: "hello",
				Level:   2,
			},
			TestAllOfThird: TestAllOfThird{
				Command: "world",
				Range:   6,
			},
		}
		data, err := value.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"comment":"hello", "command": "world", "level":2, "range":6, "value":{"class":"chair","width":12.5}}`, string(data))

		var outValue TestAllOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":{"class":"table","width":256.7}, "comment":"foo", "command": "bar", "level":1, "range":5}`), &outValue))
		val, ok := outValue.OneOf.(TestOneOfStruct)
		require.True(t, ok)
		require.EqualValues(t, TestOneOfStruct{
			Class: "table",
			Value: 256.7,
		}, val)
		require.Equal(t, "foo", outValue.Comment)
		require.EqualValues(t, 1, outValue.Level)
		require.Equal(t, "bar", outValue.Command)
		require.EqualValues(t, 5, outValue.Range)

		outValue.Reset()
		require.Empty(t, outValue)
	})
	t.Run("all_of_with_integer", func(t *testing.T) {
		t.Parallel()
		var value = TestAllOf01{
			TestAllOfFirstIsOne: TestAllOfFirstIsOne{
				OneOf: TestOneOfInteger(999),
			},
			TestAllOfSecond: TestAllOfSecond{
				Comment: "hello",
				Level:   2,
			},
			TestAllOfThird: TestAllOfThird{
				Command: "world",
				Range:   6,
			},
		}
		data, err := value.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"comment":"hello", "command": "world", "level":2, "range":6, "value":999}`, string(data))

		var outValue TestAllOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":998, "comment":"foo", "command": "bar", "level":1, "range":5}`), &outValue))
		val, ok := outValue.OneOf.(TestOneOfInteger)
		require.True(t, ok)
		require.EqualValues(t, TestOneOfInteger(998), val)
		require.Equal(t, "foo", outValue.Comment)
		require.EqualValues(t, 1, outValue.Level)
		require.Equal(t, "bar", outValue.Command)
		require.EqualValues(t, 5, outValue.Range)

		outValue.Reset()
		require.Empty(t, outValue)
	})
}

func TestReset(t *testing.T) {
	t.Parallel()
	t.Run("TestOneOfInteger", func(t *testing.T) {
		t.Parallel()
		var val TestOneOfInteger = 1
		val.Reset()
		require.Empty(t, val)
	})
	t.Run("TestOneOfString", func(t *testing.T) {
		t.Parallel()
		var val TestOneOfString = "1"
		val.Reset()
		require.Empty(t, val)
	})
	t.Run("TestOneOfStruct", func(t *testing.T) {
		t.Parallel()
		var val = TestOneOfStruct{
			Class: "tst",
			Value: 120,
		}
		val.Reset()
		require.Empty(t, val)
	})
	t.Run("TestAllOfSecond", func(t *testing.T) {
		t.Parallel()
		var val = TestAllOfSecond{
			Comment: "foo",
			Level:   9983,
		}
		val.Reset()
		require.Empty(t, val)
	})
	t.Run("TestAllOfThird", func(t *testing.T) {
		t.Parallel()
		var val = TestAllOfThird{
			Command: "bar",
			Range:   6665,
		}
		val.Reset()
		require.Empty(t, val)
	})
}
