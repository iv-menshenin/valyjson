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
	t.Run("error_f_int32_overflow", func(t *testing.T) {
		var got TestUserDefined
		err := got.UnmarshalJSON([]byte(`{"f_int32": 2147483648}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "f_int32")
		require.ErrorContains(t, err, "exceeds maximum for data type")
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
	t.Run("error_f_int64_type_error", func(t *testing.T) {
		var got TestUserDefined
		err := got.UnmarshalJSON([]byte(`{"f_int64": "34"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "f_int64")
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

func Test_TestUserDefined_Marshal_RInt(t *testing.T) {
	t.Parallel()
	t.Run("ref_int", func(t *testing.T) {
		t.Parallel()
		var (
			f  DefinedInt32 = 3
			ff DefinedInt64 = 44
		)
		var test = TestUserDefined{
			RefInt32: &f,
			RefInt64: &ff,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"f_float32":0, "f_int32":0, "f_string":"","r_int32":3,"r_int64": 44}`, string(data))
	})
	t.Run("ref_int", func(t *testing.T) {
		t.Parallel()
		var (
			f  DefinedInt32 = 3
			ff DefinedInt64 = 44
		)
		var test = TestUserDefined{
			RefInt32: &f,
			RefInt64: &ff,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"f_float32":0, "f_int32":0, "f_string":"","r_int32":3,"r_int64": 44}`, string(data))
	})
}

func Test_TestUserDefined_Unmarshal_RDef(t *testing.T) {
	t.Parallel()
	t.Run("ref_int", func(t *testing.T) {
		t.Parallel()
		var (
			f  DefinedInt32 = 3
			ff DefinedInt64 = 2147483647
		)
		var actual TestUserDefined
		var expected = TestUserDefined{
			RefInt32: &f,
			RefInt64: &ff,
		}
		const data = `{"f_float32":0, "f_int32":0, "f_string":"","r_int32":3,"r_int64": 2147483647}`
		err := actual.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
	t.Run("int32_overload", func(t *testing.T) {
		t.Parallel()
		var actual TestUserDefined
		const data = `{"f_float32":0, "f_int32":0, "f_string":"","r_int32":2147483648,"r_int64": 44}`
		err := actual.UnmarshalJSON([]byte(data))
		require.Error(t, err)
		require.ErrorContains(t, err, "exceeds maximum for data type")
	})
	t.Run("ref_float", func(t *testing.T) {
		t.Parallel()
		var (
			f  DefinedFloat32 = 123.4
			ff DefinedFloat64 = 34.5
		)
		var actual TestUserDefined
		var expected = TestUserDefined{
			RefFloat32: &f,
			RefFloat64: &ff,
		}
		const data = `{"f_float32":0, "f_int32":0, "f_string":"","r_float32":123.4, "r_float64":34.5}`
		err := actual.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
	t.Run("float32_overload", func(t *testing.T) {
		t.Parallel()
		var actual TestUserDefined
		const data = `{"f_float32":3.50282346638528859811704183484516925440e+38, "f_int32":0, "f_string":"","r_int32":4433,"r_int64": 44}`
		err := actual.UnmarshalJSON([]byte(data))
		require.Error(t, err)
		require.ErrorContains(t, err, "exceeds maximum for data type")
	})
	t.Run("ref_string", func(t *testing.T) {
		t.Parallel()
		var (
			f DefinedString = "2147483647"
		)
		var actual TestUserDefined
		var expected = TestUserDefined{
			RefString: &f,
		}
		const data = `{"f_float32":0, "f_int32":0, "f_string":"","r_string": "2147483647"}`
		err := actual.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
	t.Run("r_bool", func(t *testing.T) {
		t.Parallel()
		var (
			f DefinedBool = true
		)
		var actual TestUserDefined
		var expected = TestUserDefined{
			RefBool: &f,
		}
		const data = `{"f_float32":0, "f_int32":0, "f_string":"","r_bool": true}`
		err := actual.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, actual)
	})
}
