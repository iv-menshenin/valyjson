package test_map

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestMap01(t *testing.T) {
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
		require.ErrorContains(t, err, "error parsing 'tags' value")
		require.ErrorContains(t, err, "doesn't contain object")
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
}
