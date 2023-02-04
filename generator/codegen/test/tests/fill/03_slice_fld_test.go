package testo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestSlice01(t *testing.T) {
	t.Run("test-all-empty", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":[], "ints":null}`))
		require.NoError(t, err)
		require.NotNil(t, test1.Field)
		require.Len(t, test1.Field, 0)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("test-all-nulled", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":null}`))
		require.NoError(t, err)
		require.Nil(t, test1.Field)
		require.Nil(t, test1.FieldRef)
	})
	t.Run("filled-slice", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":["test1", "test2"], "ints":null}`))
		require.NoError(t, err)
		require.Len(t, test1.Field, 2)
		require.ElementsMatch(t, []string{"test1", "test2"}, test1.Field)
	})
	t.Run("filled-slice-ref", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":[1, 2, 3]}`))
		require.NoError(t, err)
		require.Len(t, test1.FieldRef, 3)
		var expected []*int
		for i := 1; i <= 3; i++ {
			var value = i
			expected = append(expected, &value)
		}
		require.ElementsMatch(t, expected, test1.FieldRef)
	})
	t.Run("filled-slice-ref-with-null", func(t *testing.T) {
		var test1 TestSlice01
		err := test1.UnmarshalJSON([]byte(`{"strs":null, "ints":[1, 2, null, 3]}`))
		require.NoError(t, err)
		require.Len(t, test1.FieldRef, 4)
		var expected []*int
		for i := 1; i <= 3; i++ {
			var value = i
			expected = append(expected, &value)
			if value == 3 {
				expected = append(expected, nil)
			}
		}
		require.ElementsMatch(t, expected, test1.FieldRef)
	})
}
