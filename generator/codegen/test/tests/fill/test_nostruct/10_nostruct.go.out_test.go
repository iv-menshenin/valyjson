package test_nostruct

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestMap10_MarshalJSON(t *testing.T) {
	t.Run("", func(t *testing.T) {
		const expected = `{"test": 123, "negative": -2, "zero": 0}`
		var test = TestMap10{
			"test":     123,
			"negative": -2,
			"zero":     0,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
}
