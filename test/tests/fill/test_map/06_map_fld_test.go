package test_map

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_TestMap01_Unmarshal(t *testing.T) {
	t.Run("test-all-empty", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
	})
	t.Run("test-all-null", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"tags":null, "properties":null, "key_typed_properties":null}`))
		require.NoError(t, err)
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
	})
	t.Run("test-incorrect-type", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"tags":324, "properties":null, "key_typed_properties":null}`))
		require.ErrorContains(t, err, "error parsing 'tags':")
		require.ErrorContains(t, err, "doesn't contain object")
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
	})
	t.Run("fill-booleans", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"bool":{"foo":true, "bar":false}}`))
		require.NoError(t, err)
		require.NotNil(t, test1.BoolVal)
		require.Equal(t, test1.BoolVal, map[Key]bool{
			"foo": true,
			"bar": false,
		})
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
	})
	t.Run("fill-simple-strings", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"tags":{"test":"maps", "foo":"bar"}}`))
		require.NoError(t, err)
		require.NotNil(t, test1.Tags)
		require.Equal(t, test1.Tags, map[string]string{
			"test": "maps",
			"foo":  "bar",
		})
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
	})
	t.Run("fill-key-strings", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"properties":{"test":{"value":"maps","name":"test"}, "foo":{"name": "foo","value":"bar"}}}`))
		require.NoError(t, err)
		require.NotNil(t, test1.Properties)
		require.Equal(t, test1.Properties, map[string]Property{
			"test": {Name: "test", Value: "maps"},
			"foo":  {Name: "foo", Value: "bar"},
		})
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.KeyTypedProperties)
	})
	t.Run("fill-key-typed", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"key_typed_properties":{"test":{"value":"maps","name":"test"}, "foo":{"name": "foo","value":"bar"}}}`))
		require.NoError(t, err)
		require.NotNil(t, test1.KeyTypedProperties)
		require.Equal(t, test1.KeyTypedProperties, map[Key]Property{
			"test": {Name: "test", Value: "maps"},
			"foo":  {Name: "foo", Value: "bar"},
		})
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
	})
	t.Run("test-int-values", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"integerVal":{"123": 123, "0": 0, "-1": -1}}`))
		require.NoError(t, err)
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
		require.NotEmpty(t, test1.IntegerVal)
		require.EqualValues(t, test1.IntegerVal["123"], 123)
		require.EqualValues(t, test1.IntegerVal["0"], 0)
		require.EqualValues(t, test1.IntegerVal["-1"], -1)
	})
	t.Run("test-uint-error", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"uintVal":{"123": 123, "0": 0, "-1": -1}}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "uintVal")
		require.ErrorContains(t, err, "-1")
	})
	t.Run("test-uint-values", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"uintVal":{"123": 123, "0": 0,"nil": null}}`))
		require.NoError(t, err)
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
		require.NotEmpty(t, test1.UintVal)
		require.NotNil(t, *test1.UintVal["123"])
		require.EqualValues(t, *test1.UintVal["123"], 123)
		require.NotNil(t, *test1.UintVal["0"])
		require.EqualValues(t, *test1.UintVal["0"], 0)
		val, ok := test1.UintVal["nil"]
		require.True(t, ok)
		require.Nil(t, val)
	})
	t.Run("test-float-values", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"floatVal":{"123": 123, "0": 0, "-1": -1}}`))
		require.NoError(t, err)
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
		require.NotEmpty(t, test1.FloatVal)
		require.EqualValues(t, test1.FloatVal["123"], 123)
		require.EqualValues(t, test1.FloatVal["0"], 0)
		require.EqualValues(t, test1.FloatVal["-1"], -1)
	})
	t.Run("test-typed-error", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"typed-val":{"123": 123, "0": 0, "minus": -2}}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "typed-val")
		require.ErrorContains(t, err, "-2")
	})
	t.Run("test-typed-values", func(t *testing.T) {
		var test1 TestMap01
		err := test1.UnmarshalJSON([]byte(`{"typed-val":{"123": 123, "0": 0, "-1": 1}}`))
		require.NoError(t, err)
		require.Nil(t, test1.Tags)
		require.Nil(t, test1.Properties)
		require.Nil(t, test1.KeyTypedProperties)
		require.NotEmpty(t, test1.TypedVal)
		require.EqualValues(t, test1.TypedVal["123"], 123)
		require.EqualValues(t, test1.TypedVal["0"], 0)
		require.EqualValues(t, test1.TypedVal["-1"], 1)
	})
}

func Test_TestMap01_Marshal(t *testing.T) {
	t.Run("fulfill", func(t *testing.T) {
		const expected = `{"tags":{"test":"maps", "lorem":"ipsum"},"properties":{"test":{"value":"foo","name":"bar"}, "foo":{"name": "foo","value":"lorem ipsum"}},"key_typed_properties":null}`
		var test = TestMap01{
			Tags: map[string]string{
				"test":  "maps",
				"lorem": "ipsum",
			},
			Properties: map[string]Property{
				"test": {
					Name:  "bar",
					Value: "foo",
				},
				"foo": {
					Name:  "foo",
					Value: "lorem ipsum",
				},
			},
			KeyTypedProperties: nil,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("omit-properties", func(t *testing.T) {
		const expected = `{"tags":{"test":"maps", "lorem":"ipsum"},"key_typed_properties":{"test":{"value":"foo","name":"bar"}, "foo":{"name": "foo","value":"lorem ipsum"}}}`
		var test = TestMap01{
			Tags: map[string]string{
				"test":  "maps",
				"lorem": "ipsum",
			},
			KeyTypedProperties: map[Key]Property{
				"test": {
					Name:  "bar",
					Value: "foo",
				},
				"foo": {
					Name:  "foo",
					Value: "lorem ipsum",
				},
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("null", func(t *testing.T) {
		const expected = `{"tags":null,"key_typed_properties":null}`
		var test = TestMap01{}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("empty", func(t *testing.T) {
		const expected = `{"tags":{},"properties":{},"key_typed_properties":{}}`
		var test = TestMap01{
			Tags:               map[string]string{},
			Properties:         map[string]Property{},
			KeyTypedProperties: map[Key]Property{},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("int-values", func(t *testing.T) {
		const expected = `{"tags":null,"key_typed_properties":null,"integerVal":{"344": 344, "345": 345, "0": 0}}`
		var test = TestMap01{
			IntegerVal: map[Key]int32{
				"344": 344, "345": 345, "0": 0,
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("float-values", func(t *testing.T) {
		const expected = `{"tags":null,"key_typed_properties":null,"floatVal":{"344": 344, "345": 345, "0": 0}}`
		var test = TestMap01{
			FloatVal: map[Key]float64{
				"344": 344, "345": 345, "0": 0,
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("uint-values", func(t *testing.T) {
		const expected = `{"tags":null,"key_typed_properties":null,"uintVal":{"344": 344, "345": 345, "0": 0}}`
		var (
			v344 uint16 = 344
			v345 uint16 = 345
			v000 uint16 = 0
			test        = TestMap01{
				UintVal: map[Key]*uint16{
					"344": &v344, "345": &v345, "0": &v000,
				},
			}
		)
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("typed-values", func(t *testing.T) {
		const expected = `{"tags":null,"key_typed_properties":null,"typed-val":{"344": 344, "345": 345, "0": 0}}`
		var test = TestMap01{
			TypedVal: map[Key]Val{
				"344": 344, "345": 345, "0": 0,
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("boolean-values", func(t *testing.T) {
		const expected = `{"tags":null,"key_typed_properties":null,"bool":{"foo": true, "bar": false}}`
		var test = TestMap01{
			BoolVal: map[Key]bool{
				"foo": true, "bar": false,
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
}
