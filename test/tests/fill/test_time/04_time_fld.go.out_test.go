package test_time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_TestTime01_unmarshal(t *testing.T) {
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
	t.Run("wrong-json", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_end": null`))
		require.Error(t, err)
	})
	t.Run("test-src-date_end-twice", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_end": null, "date_end": null, "date_end": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field appears in the object twice")
	})
	t.Run("test-src-date_custom-twice", func(t *testing.T) {
		var test1 TestTime01
		err := test1.UnmarshalJSON([]byte(`{"date_custom": null, "date_custom": null, "date_custom": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field appears in the object twice")
	})
}

func Test_TestTime01_marshal(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var test TestTime01
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"date_begin":"0001-01-01T00:00:00Z", "date_custom":"0001.01.01"}`, string(data))
	})
	t.Run("filled", func(t *testing.T) {
		var dE = time.Date(2023, time.April, 2, 0, 0, 0, 0, time.UTC)
		var test = TestTime01{
			DateBegin:  time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC),
			DateCustom: time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC),
			DateEnd:    &dE,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"date_begin":"2023-04-01T00:00:00Z", "date_custom":"2023.04.01", "date_end":"2023-04-02T00:00:00Z"}`, string(data))
	})
}

func Test_TestTime01_Zero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test TestTime01
		require.True(t, test.IsZero())
	})
	t.Run("not_zero_1", func(t *testing.T) {
		var test TestTime01
		test.DateEnd = &time.Time{}
		require.False(t, test.IsZero())
	})
	t.Run("not_zero_2", func(t *testing.T) {
		var test TestTime01
		test.DateBegin = time.Now()
		require.False(t, test.IsZero())
	})
	t.Run("not_zero_3", func(t *testing.T) {
		var test TestTime01
		test.DateCustom = time.Now()
		require.False(t, test.IsZero())
	})
}
