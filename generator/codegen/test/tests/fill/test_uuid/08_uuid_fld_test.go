package test_uuid

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestUUID(t *testing.T) {
	t.Run("marshall", func(t *testing.T) {
		var value = TestUUID{
			UUID: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}
		data, err := json.Marshal(value)
		require.NoError(t, err)
		require.JSONEq(t, `{"uuid":"01020304-0506-0708-090a-0b0c0d0e0f10"}`, string(data))
	})
	t.Run("unmarshall_1", func(t *testing.T) {
		var (
			actual   TestUUID
			expected = TestUUID{
				UUID: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			}
		)
		err := actual.UnmarshalJSON([]byte(`{"uuid":"01020304-0506-0708-090a-0b0c0d0e0f10"}`))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
	t.Run("unmarshall_2", func(t *testing.T) {
		var (
			actual   TestUUID
			expected = TestUUID{
				UUID: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			}
		)
		err := actual.UnmarshalJSON([]byte(`{"uuid":"0102030405060708090a0b0c0d0e0f10"}`))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
