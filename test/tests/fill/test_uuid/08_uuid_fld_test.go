package test_uuid

import (
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/require"
)

func Test_TestUUID_UnmarshalJSON(t *testing.T) {
	t.Parallel()
	t.Run("unmarshall_1", func(t *testing.T) {
		t.Parallel()
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
		t.Parallel()
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

func Test_TestUUID_Marshal(t *testing.T) {
	t.Parallel()
	t.Run("marshall", func(t *testing.T) {
		t.Parallel()
		var value = TestUUID{
			UUID: [16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		}
		data, err := value.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"uuid":"01020304-0506-0708-090a-0b0c0d0e0f10"}`, string(data))
	})
	t.Run("null", func(t *testing.T) {
		t.Parallel()
		var value *TestUUID
		data, err := value.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `null`, string(data))
	})
}

func Test_Inherits_UUID(t *testing.T) {
	t.Parallel()
	t.Run("uuid_inline_marshal", func(t *testing.T) {
		t.Parallel()
		var value = InheritUUID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
		data, err := value.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `"01020304-0506-0708-090a-0b0c0d0e0f10"`, string(data))
	})
	t.Run("uuid_inline_unmarshal", func(t *testing.T) {
		t.Parallel()
		var (
			actual   InheritUUID
			expected = InheritUUID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
		)
		err := actual.UnmarshalJSON([]byte(`"01020304-0506-0708-090a-0b0c0d0e0f10"`))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
	t.Run("error_format", func(t *testing.T) {
		t.Parallel()
		var UUID InheritUUID
		err := UUID.UnmarshalJSON([]byte(`"__020304-0506-0708-090a-0b0c0d0e0f10"`))
		require.Error(t, err)
		require.ErrorContains(t, err, "invalid UUID format")
	})
	t.Run("error_type", func(t *testing.T) {
		t.Parallel()
		var UUID InheritUUID
		err := UUID.UnmarshalJSON([]byte(`20304`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain string")
	})
}

func Test_Inherits_UUID2(t *testing.T) {
	t.Parallel()
	t.Run("uuid_inline_marshal", func(t *testing.T) {
		t.Parallel()
		var value = InheritUUID2([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
		data, err := value.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `"01020304-0506-0708-090a-0b0c0d0e0f10"`, string(data))
	})
	t.Run("uuid_inline_unmarshal", func(t *testing.T) {
		t.Parallel()
		var (
			actual   InheritUUID2
			expected = InheritUUID2([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16})
		)
		err := actual.UnmarshalJSON([]byte(`"01020304-0506-0708-090a-0b0c0d0e0f10"`))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
	t.Run("error_format", func(t *testing.T) {
		t.Parallel()
		var UUID InheritUUID2
		err := UUID.UnmarshalJSON([]byte(`"__020304-0506-0708-090a-0b0c0d0e0f10"`))
		require.Error(t, err)
		require.ErrorContains(t, err, "invalid UUID format")
	})
	t.Run("error_type", func(t *testing.T) {
		t.Parallel()
		var UUID InheritUUID2
		err := UUID.UnmarshalJSON([]byte(`20304`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain string")
	})
}

func Test_IsZero_Reset(t *testing.T) {
	t.Parallel()
	t.Run("is_zero", func(t *testing.T) {
		require.True(t, TestUUID{}.IsZero())
		require.True(t, InheritUUID2{}.IsZero())
		require.True(t, InheritUUID{}.IsZero())
	})
	t.Run("reset", func(t *testing.T) {
		var (
			var1 = TestUUID{UUID: uuid.Must(uuid.NewUUID())}
			var2 = InheritUUID2(uuid.Must(uuid.NewUUID()))
			var3 = InheritUUID(uuid.Must(uuid.NewUUID()))
		)
		require.False(t, var1.IsZero())
		require.False(t, var2.IsZero())
		require.False(t, var3.IsZero())
		var1.Reset()
		var2.Reset()
		var3.Reset()
		require.True(t, var1.IsZero())
		require.True(t, var2.IsZero())
		require.True(t, var3.IsZero())
	})
}
