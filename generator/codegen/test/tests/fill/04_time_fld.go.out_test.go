package testo

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTestTime01(t *testing.T) {
	t.Run("test-all-empty", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.True(t, test1.DateBegin.IsZero())
		require.Nil(t, test1.DateEnd)
	})
	t.Run("test-RFC3339", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_begin": "2023-01-26T16:00:01Z", "date_end": "2023-01-26T16:00:01Z"}`))
		require.NoError(t, err)
		expected := time.Date(2023, time.January, 26, 16, 0, 1, 0, time.UTC)
		require.Equal(t, expected, test1.DateBegin)
		require.NotNil(t, test1.DateEnd)
		require.Equal(t, expected, *test1.DateEnd)
	})
	t.Run("test-custom", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_custom": "2023.01.26"}`))
		require.NoError(t, err)
		expected := time.Date(2023, time.January, 26, 0, 0, 0, 0, time.UTC)
		require.Equal(t, expected, test1.DateCustom)
	})
	t.Run("test-error-fmt-1", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_begin": "2023-0x-26T16:00:01Z"}`))
		require.Error(t, err)
	})
	t.Run("test-error-fmt-2", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_end": "2023-0x-26T16:00:01Z"}`))
		require.Error(t, err)
	})
	t.Run("test-error-custom", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_custom": "2023-01-26T16:00:01Z"}`))
		require.Error(t, err)
	})
	t.Run("test-src-data-1", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_begin": 2342}`))
		require.Error(t, err)
	})
	t.Run("test-src-data-2", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_end": 2342}`))
		require.Error(t, err)
	})
	t.Run("test-ref-nil", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_end": null}`))
		require.NoError(t, err)
		require.Nil(t, test1.DateEnd)
	})
}
