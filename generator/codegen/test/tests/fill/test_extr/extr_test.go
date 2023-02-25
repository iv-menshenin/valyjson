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
}
