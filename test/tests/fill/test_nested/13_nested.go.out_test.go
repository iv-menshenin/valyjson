package test_nested

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Unmarshal(t *testing.T) {
	t.Run("path_to_error_array", func(t *testing.T) {
		const data = `{"meta":{"count": 2}, "data":[{"name":"Igor", "surname":"Menshenin"},{},{"name":554, "surname":"Petrov"}]}`
		var v Root
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "data.2.name")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_error_inh_array", func(t *testing.T) {
		const data = `[{"name":"Igor", "surname":"Menshenin"},{},{"name":554, "surname":"Petrov"}]`
		var v Middles
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "2.name")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_err_map", func(t *testing.T) {
		const data = `{"meta":{"count": 2}, "data":[{"name":"Igor", "surname":"Menshenin"},{"surname":"Petrov","tags":{"foo":"bar","count": 7}}]}`
		var v Root
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "data.1.tags.count")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
	t.Run("path_to_err_inlined_map", func(t *testing.T) {
		const data = `{"test":"test","foo":"bar","conf":null}`
		var v Tags
		err := v.UnmarshalJSON([]byte(data))
		require.ErrorContains(t, err, "conf")
		require.ErrorContains(t, err, "value doesn't contain string")
	})
}
