package test_userdefined

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestUserDefined_Unmarshal(t *testing.T) {
	t.Run("default_strings", func(t *testing.T) {
		var got TestUserDefined
		var expected = TestUserDefined{
			Int32:   32,
			String:  "default_string",
			Float32: 123.01,
		}
		err := got.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("simple", func(t *testing.T) {
		var got TestUserDefined
		var expected = TestUserDefined{
			Int32:   32,
			Int64:   64,
			Float32: 0.32,
			Float64: 0.64,
			String:  "foo-bar",
			Bool:    true,
		}
		err := got.UnmarshalJSON([]byte(`{"f_int32": 32,"f_int64": 64,"f_float32":0.32,"f_float64":0.64,"f_string":"foo-bar","f_bool":true}`))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("zeroed", func(t *testing.T) {
		var got TestUserDefined
		var expected = TestUserDefined{
			Int32:   0,
			Int64:   0,
			Float32: 0.0,
			Float64: 0.0,
			String:  "",
			Bool:    false,
		}
		err := got.UnmarshalJSON([]byte(`{"f_int32": 0,"f_int64": 0,"f_float32":0.00,"f_float64":0.00,"f_string":"","f_bool":false}`))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("error_f_int32_double", func(t *testing.T) {
		var got TestUserDefined
		err := got.UnmarshalJSON([]byte(`{"f_int32": 0,"f_int32": 0}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "f_int32")
		require.ErrorContains(t, err, "field appears in the object twice")
	})
	t.Run("error_f_int32_format", func(t *testing.T) {
		var got TestUserDefined
		err := got.UnmarshalJSON([]byte(`{"f_int32": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "f_int32")
		require.ErrorContains(t, err, "value doesn't contain number")
	})
	t.Run("error_f_float32_double", func(t *testing.T) {
		var got TestUserDefined
		err := got.UnmarshalJSON([]byte(`{"f_float32": 0,"f_float32": 0}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "f_float32")
		require.ErrorContains(t, err, "field appears in the object twice")
	})
	t.Run("error_f_float32_format", func(t *testing.T) {
		var got TestUserDefined
		err := got.UnmarshalJSON([]byte(`{"f_float32": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "f_float32")
		require.ErrorContains(t, err, "value doesn't contain number")
	})
}

func Test_TestUserDefined_Marshal(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test TestUserDefined
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"f_int32":0,"f_float32": 0,"f_string":""}`, string(data))
	})
	t.Run("zero", func(t *testing.T) {
		var test = TestUserDefined{
			Int32:   1,
			Int64:   2,
			Float32: 3,
			Float64: 4,
			String:  "5",
			Bool:    true,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"f_int32":1,"f_int64":2, "f_float32": 3,"f_float64":4, "f_string":"5","f_bool": true}`, string(data))
	})
	t.Run("null", func(t *testing.T) {
		var test *TestUserDefined
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `null`, string(data))
	})
}

func Test_TestUserDefined_Zero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test = TestUserDefined{}
		require.True(t, test.IsZero())
	})
	t.Run("not_zero", func(t *testing.T) {
		var test = TestUserDefined{
			String: "default_string",
		}
		require.False(t, test.IsZero())
	})
}
