package test_nostruct

import (
	"fill/test_any"
	"fill/test_extr"
	"fill/test_string"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Map_UnmarshalJSON(t *testing.T) {
	t.Run("map[string]int64", func(t *testing.T) {
		const input = `{"test": 123, "negative": -2, "zero": 0}`
		var test TestMap10
		var expected = TestMap10{
			"test":     123,
			"negative": -2,
			"zero":     0,
		}
		err := test.UnmarshalJSON([]byte(input))
		require.NoError(t, err)
		require.Equal(t, expected, test)
	})
	t.Run("map[string]int64-struct-error", func(t *testing.T) {
		const input = `["test", "negative", "zero"]`
		var test TestMap10
		err := test.UnmarshalJSON([]byte(input))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("map[string]int64-struct-error", func(t *testing.T) {
		const input = `{"test": {}, "negative": -2, "zero": 0}`
		var test TestMap10
		err := test.UnmarshalJSON([]byte(input))
		require.Error(t, err)
		require.ErrorContains(t, err, "(root).test")
		require.ErrorContains(t, err, "value doesn't contain number")
	})

	t.Run("map[string]external", func(t *testing.T) {
		t.SkipNow() // FIXME default values in nested structures
		const input = `{"test": {"test1":{"comment": "foo bar"}}}`
		var def = "default"
		var test TestMap11
		var expected = TestMap11{
			"test": {
				Test01: test_any.TestAllOfSecond{
					Comment: "foo bar",
				},
				Test02: test_string.TestStr01{
					DefRef: &def,
				},
			},
		}
		err := test.UnmarshalJSON([]byte(input))
		require.NoError(t, err)
		require.Equal(t, expected, test)
	})
	t.Run("map[string]external-wrong-null", func(t *testing.T) {
		const input = `{"test": null}`
		var test TestMap11
		err := test.UnmarshalJSON([]byte(input))
		require.Error(t, err)
		require.ErrorContains(t, err, "(root).test")
		require.ErrorContains(t, err, "value doesn't contain object")
	})

	t.Run("map[string]*external", func(t *testing.T) {
		const input = `{"test": {"test1":{"comment": "this isn't a comment", "level": 99}, "test2":{"field": "bar", "fieldRef": "foo", "defRef": "test"}}}`
		var (
			fieldRef = "foo"
			defRef   = "test"
		)
		var test TestMap11Ref
		var expected = TestMap11Ref{
			"test": &test_extr.External{
				Test01: test_any.TestAllOfSecond{
					Comment: "this isn't a comment",
					Level:   99,
				},
				Test02: test_string.TestStr01{
					Field:    "bar",
					FieldRef: &fieldRef,
					DefRef:   &defRef,
				},
			},
		}
		err := test.UnmarshalJSON([]byte(input))
		require.NoError(t, err)
		require.Equal(t, expected, test)
	})

	t.Run("map[string]*external-null-value", func(t *testing.T) {
		const input = `{"empty":null, "test": {"test1":{"comment": "this isn't a comment", "level": 99}, "test2":{"field": "bar", "fieldRef": "foo", "defRef": "test"}}}`
		var (
			fieldRef = "foo"
			defRef   = "test"
		)
		var test TestMap11Ref
		var expected = TestMap11Ref{
			"test": &test_extr.External{
				Test01: test_any.TestAllOfSecond{
					Comment: "this isn't a comment",
					Level:   99,
				},
				Test02: test_string.TestStr01{
					Field:    "bar",
					FieldRef: &fieldRef,
					DefRef:   &defRef,
				},
			},
			"empty": nil,
		}
		err := test.UnmarshalJSON([]byte(input))
		require.NoError(t, err)
		require.Equal(t, expected, test)
	})
	t.Run("map[string]*external-wrong-value", func(t *testing.T) {
		const input = `{"wrong": [1, 2, 3], "test": null}`
		var test TestMap11Ref
		err := test.UnmarshalJSON([]byte(input))
		require.Error(t, err)
		require.ErrorContains(t, err, "(root).wrong")
		require.ErrorContains(t, err, "doesn't contain object")
	})
}

func Test_Map_MarshalJSON(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		const expected = `null`
		var test TestMap10
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("map[string]int64", func(t *testing.T) {
		const expected = `{"test": 123, "negative": -2, "zero": 0}`
		var test = TestMap10{
			"test":     123,
			"negative": -2,
			"zero":     0,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("map[string]external", func(t *testing.T) {
		const expected = `{"test": {"test1":{"comment": "this is a comment", "level": 100}, "test2":{"field": "foo", "fieldRef": "bar", "defRef": ""}}}`
		var (
			fieldRef = "bar"
			defRef   = ""
		)
		var test = TestMap11{
			"test": test_extr.External{
				Test01: test_any.TestAllOfSecond{
					Comment: "this is a comment",
					Level:   100,
				},
				Test02: test_string.TestStr01{
					Field:    "foo",
					FieldRef: &fieldRef,
					DefRef:   &defRef,
				},
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
	t.Run("map[string]*external", func(t *testing.T) {
		const expected = `{"empty":null, "test": {"test1":{"comment": "this isn't a comment", "level": 99}, "test2":{"field": "bar", "fieldRef": "foo", "defRef": "test"}}}`
		var (
			fieldRef = "foo"
			defRef   = "test"
		)
		var test = TestMap11Ref{
			"test": &test_extr.External{
				Test01: test_any.TestAllOfSecond{
					Comment: "this isn't a comment",
					Level:   99,
				},
				Test02: test_string.TestStr01{
					Field:    "bar",
					FieldRef: &fieldRef,
					DefRef:   &defRef,
				},
			},
			"empty": nil,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(b))
	})
}