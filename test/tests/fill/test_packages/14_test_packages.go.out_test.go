package test_packages

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTest01_Smoke(t *testing.T) {
	t.Parallel()
	t.Run("marshal", func(t *testing.T) {
		t.Parallel()
		var x = Test01{Field: 1133}
		data, err := x.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field": 1133}`, string(data))
	})
	t.Run("unmarshal", func(t *testing.T) {
		t.Parallel()
		var x Test01
		require.NoError(t, x.UnmarshalJSON([]byte(`{"field":9988}`)))
		require.Equal(t, Test01{Field: 9988}, x)
	})
}
