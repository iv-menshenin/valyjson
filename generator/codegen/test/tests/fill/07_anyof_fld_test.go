package testo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestAllOfAnyOf01(t *testing.T) {
	t.Run("any_int_", func(t *testing.T) {
		var value = TestAnyOf01{
			TestAnyOfUplevel: TestAnyOfUplevel{
				Any: 33,
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"value":33, "comment":""}`, string(data))

		var outValue TestAnyOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":332}`), &outValue))
		val, ok := outValue.Any.(TestAnyOfInteger)
		require.True(t, ok)
		require.EqualValues(t, 332, val)
		require.Empty(t, outValue.Comment)
		require.Empty(t, outValue.Level)
	})
	t.Run("any_string_", func(t *testing.T) {
		var value = TestAnyOf01{
			TestAnyOfUplevel: TestAnyOfUplevel{
				Any: "some string",
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"value":"some string", "comment":""}`, string(data))

		var outValue TestAnyOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":"some string 1"}`), &outValue))
		val, ok := outValue.Any.(TestAnyOfString)
		require.True(t, ok)
		require.EqualValues(t, "some string 1", val)
		require.Empty(t, outValue.Comment)
		require.Empty(t, outValue.Level)
	})
	t.Run("any_struct_", func(t *testing.T) {
		var value = TestAnyOf01{
			TestAnyOfUplevel: TestAnyOfUplevel{
				Any: TestAnyOfStruct{
					Class: "chair",
					Value: 322.5,
				},
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"value":{"class":"chair","width":322.5}, "comment":""}`, string(data))

		var outValue TestAnyOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":{"class":"table","width":256.7}}`), &outValue))
		val, ok := outValue.Any.(TestAnyOfStruct)
		require.True(t, ok)
		require.EqualValues(t, TestAnyOfStruct{
			Class: "table",
			Value: 256.7,
		}, val)
		require.Empty(t, outValue.Comment)
		require.Empty(t, outValue.Level)
	})
	t.Run("all_of", func(t *testing.T) {
		var value = TestAnyOf01{
			TestAnyOfUplevel: TestAnyOfUplevel{
				Any: TestAnyOfStruct{
					Class: "chair",
					Value: 12.5,
				},
			},
			TestAllOfUplevel1: TestAllOfUplevel1{
				Comment: "hello",
				Level:   2,
			},
			TestAllOfUplevel2: TestAllOfUplevel2{
				Command: "world",
				Range:   6,
			},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"comment":"hello", "command": "world", "level":2, "range":6, "value":{"class":"chair","width":12.5}}`, string(data))

		var outValue TestAnyOf01
		require.NoError(t, json.Unmarshal([]byte(`{"value":{"class":"table","width":256.7}, "comment":"foo", "command": "bar", "level":1, "range":5}`), &outValue))
		val, ok := outValue.Any.(TestAnyOfStruct)
		require.True(t, ok)
		require.EqualValues(t, TestAnyOfStruct{
			Class: "table",
			Value: 256.7,
		}, val)
		require.Equal(t, "foo", outValue.Comment)
		require.EqualValues(t, 1, outValue.Level)
		require.Equal(t, "bar", outValue.Command)
		require.EqualValues(t, 5, outValue.Range)
	})
}
