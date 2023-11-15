package test_nested

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	t.Parallel()
	t.Run("path_to_error_array", func(t *testing.T) {
		t.Parallel()
		const data = `{"meta":{"count": 2}, "data":[{"name":"Igor", "surname":"Menshenin"},{},{"name":554, "surname":"Petrov"}]}`
		var v Root
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "data.2.name")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_error_inh_array", func(t *testing.T) {
		t.Parallel()
		const data = `[{"name":"Igor", "surname":"Menshenin"},{},{"name":554, "surname":"Petrov"}]`
		var v Middles
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "2.name")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_err_map", func(t *testing.T) {
		t.Parallel()
		const data = `{"meta":{"count": 2}, "data":[{"name":"Igor", "surname":"Menshenin"},{"surname":"Petrov","tags":{"foo":"bar","count": 7}}]}`
		var v Root
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "data.1.tags.count")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_err_inlined_map", func(t *testing.T) {
		t.Parallel()
		const data = `{"test":"test","foo":"bar","conf":null}`
		var v Tags
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "conf")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
}

func Test_InlinedNestedStructures(t *testing.T) {
	t.Parallel()
	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		var ce = CustomEvent{
			WRRetry{
				WRRetry: 3,
			},
		}
		data, err := ce.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"WR-Retry": 3}`, string(data))
	})
	t.Run("marshal-the-same-as-std", func(t *testing.T) {
		t.Parallel()
		var ce = CustomEvent{WRRetry{WRRetry: 312}}
		data1, err := ce.MarshalJSON()
		require.NoError(t, err)

		data2, err := json.Marshal(ce)
		require.NoError(t, err)

		require.JSONEq(t, string(data1), string(data2))
	})
	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		const jsonData = `{"WR-Retry": 304}`
		var expected = CustomEvent{
			WRRetry{
				WRRetry: 304,
			},
		}
		var got CustomEvent
		require.NoError(t, got.UnmarshalJSON([]byte(jsonData)))
		require.Equal(t, expected, got)
	})
	t.Run("unmarshal-the-same-as-std", func(t *testing.T) {
		t.Parallel()
		const jsonData = `{"WR-Retry": 5666}`
		var expected = CustomEvent{WRRetry{WRRetry: 5666}}
		var got1 CustomEvent

		require.NoError(t, got1.UnmarshalJSON([]byte(jsonData)))
		require.Equal(t, expected, got1)

		var got2 CustomEvent
		require.NoError(t, json.Unmarshal([]byte(jsonData), &got2))
		require.Equal(t, expected, got2)
	})
}
