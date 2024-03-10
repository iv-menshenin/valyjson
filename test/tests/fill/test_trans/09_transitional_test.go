package test_trans

import (
	"fill/test_any"
	"fill/test_string"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_TestTransitional(t *testing.T) {
	t.Parallel()
	t.Run("UnmarshalJSON", func(t *testing.T) {
		t.Parallel()
		var test TestTransitional
		require.NoError(t, test.UnmarshalJSON([]byte(`{"test-field":12222}`)))
		require.EqualValues(t, 12222, test.TestField)
	})
	t.Run("MarshalJSON", func(t *testing.T) {
		t.Parallel()
		var test = TestTransitional{
			TestField: 55554,
		}
		const expected = `{"test-field":55554}`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
}

func Test_TestExternalNested(t *testing.T) {
	t.Parallel()
	t.Run("Unmarshal", func(t *testing.T) {
		t.Parallel()
		const data = `{"comment":"done","level":13,"command":"test","range":88,"field":"foo-bar","fieldRef":"\"bar\"","defRef":null}`
		var bar = `"bar"`
		var test TestExternalNested
		require.NoError(t, test.UnmarshalJSON([]byte(data)))
		require.Equal(t, TestExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "done",
				Level:   13,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "test",
				Range:   88,
			},
			TestStr01: test_string.TestStr01{
				Field:    "foo-bar",
				FieldRef: &bar,
			},
		}, test)
	})
	t.Run("Marshal", func(t *testing.T) {
		t.Parallel()
		var bar = `"bar"`
		var test = TestExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "Go home, American!",
				Level:   -99,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "go 'away",
			},
			TestStr01: test_string.TestStr01{
				Field:    "foo-bar",
				FieldRef: &bar,
			},
		}
		const expected = `{"comment":"Go home, American!","level":-99,"command":"go 'away","field":"foo-bar","fieldRef":"\"bar\"","defRef":null}`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("IsZero", func(t *testing.T) {
		t.Parallel()
		var test = TestExternalNested{}
		require.True(t, test.IsZero())

		test = TestExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "",
				Level:   1,
			},
		}
		require.False(t, test.IsZero())

		test = TestExternalNested{
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "",
				Range:   -1,
			},
		}
		require.False(t, test.IsZero())

		test = TestExternalNested{
			TestStr01: test_string.TestStr01{
				Field: "zz",
			},
		}
		require.False(t, test.IsZero())
	})
	t.Run("Reset", func(t *testing.T) {
		t.Parallel()
		var bar = `"foo_bar"`
		var test = TestExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "Hello world<?php",
				Level:   -99,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "?>",
				Range:   33342,
			},
			TestStr01: test_string.TestStr01{
				Field:    "foo-bar",
				FieldRef: &bar,
			},
		}
		test.Reset()
		require.True(t, test.IsZero())
		require.Equal(t, TestExternalNested{}, test)
	})
}

func Test_TestExternalNested_Reset_Allocs(t *testing.T) {
	var bar = `"foo_bar"`
	allocs := testing.AllocsPerRun(1000000, func() {
		var test = TestExternalNested{
			TestAllOfSecond: test_any.TestAllOfSecond{
				Comment: "Hello world<?php",
				Level:   -99,
			},
			TestAllOfThird: test_any.TestAllOfThird{
				Command: "?>",
				Range:   33342,
			},
			TestStr01: test_string.TestStr01{
				Field:    "foo-bar",
				FieldRef: &bar,
			},
		}
		test.Reset()
	})
	require.NotEmpty(t, bar)
	require.LessOrEqual(t, allocs, float64(0))
}
