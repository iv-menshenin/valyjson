package test_trans

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestTransitional(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var test TestTransitional
		require.NoError(t, test.UnmarshalJSON([]byte(`{"test-field":12222}`)))
		require.EqualValues(t, 12222, test.TestField)
	})
}
