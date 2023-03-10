package test_inh

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Inheritance(t *testing.T) {
	t.Run("test-all-empty", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Zero(t, test1.TestInh02.Int32)
		require.Zero(t, test1.TestInh03.Int16)
		require.Zero(t, test1.Nested1.Int16)
		require.Nil(t, test1.Nested2)
	})
	t.Run("test-filled-hierarchy", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{"injected":{"int_32": 123}}`))
		require.NoError(t, err)
		require.EqualValues(t, test1.TestInh02.Int32, 123)
		require.Zero(t, test1.TestInh03.Int16)
		require.Zero(t, test1.Nested1.Int16)
		require.Nil(t, test1.Nested2)
	})
	t.Run("test-filled-inline", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{"int_16":543,"random":66}`))
		require.NoError(t, err)
		require.Zero(t, test1.TestInh02.Int32)
		require.EqualValues(t, test1.Int16, 543)
		require.EqualValues(t, test1.Random, 66)
		require.Zero(t, test1.Nested1.Int16)
		require.Nil(t, test1.Nested2)
	})
	t.Run("test-fulfilled", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{"int_16":543,"random":66,"nested1":{"int_16":888,"random":999},"nested2":{"int_16":777,"random":666},"date_begin":"2023-01-28 07:10:05Z"}`))
		require.NoError(t, err)
		require.Zero(t, test1.TestInh02.Int32)
		require.EqualValues(t, test1.Int16, 543)
		require.EqualValues(t, test1.Random, 66)
		require.EqualValues(t, test1.Nested1.Int16, 888)
		require.EqualValues(t, test1.Nested1.Random, 999)
		require.NotNil(t, test1.Nested2)
		require.EqualValues(t, test1.Nested2.Int16, 777)
		require.EqualValues(t, test1.Nested2.Random, 666)
		require.False(t, test1.DateBegin.IsZero())
	})
	t.Run("test-wrong-inline-type", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{"int_16":543,"random":"66","nested1":{"int_16":888,"random":999},"nested2":{"int_16":777,"random":666},"date_begin":"2023-01-28 07:10:05Z"}`))
		require.ErrorContains(t, err, "error parsing '(root).random' value")
	})
	t.Run("test-wrong-nested-type", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{"int_16":543,"random":66,"nested1":{"int_16":888,"random":"999"},"nested2":{"int_16":777,"random":666},"date_begin":"2023-01-28 07:10:05Z"}`))
		require.ErrorContains(t, err, "error parsing '(root).nested1.random' value")
	})
}

func Test_TestNested01(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var test1 TestNested01
		err := test1.UnmarshalJSON([]byte(`{}`))
		require.NoError(t, err)
		require.Zero(t, test1.TestNested02.TestNested03)
	})
	t.Run("check-type", func(t *testing.T) {
		var test1 TestNested01
		err := test1.UnmarshalJSON([]byte(`{"field_32": 2147483648}`))
		require.Error(t, err)
		require.Zero(t, test1.TestNested02.TestNested03.Field32)
	})
	t.Run("filled-1", func(t *testing.T) {
		var test1 TestNested01
		err := test1.UnmarshalJSON([]byte(`{"field_32": 490}`))
		require.NoError(t, err)
		require.EqualValues(t, 490, test1.TestNested02.TestNested03.Field32)
	})
	t.Run("filled-2", func(t *testing.T) {
		var test1 TestNested02
		err := test1.UnmarshalJSON([]byte(`{"field_32": 491}`))
		require.NoError(t, err)
		require.EqualValues(t, 491, test1.TestNested03.Field32)
	})
	t.Run("filled-3", func(t *testing.T) {
		var test1 TestNested03
		err := test1.UnmarshalJSON([]byte(`{"field_32": 492}`))
		require.NoError(t, err)
		require.EqualValues(t, 492, test1.Field32)
	})
}

func Test_JsonTestNested01(t *testing.T) {
	t.Run("nested-3", func(t *testing.T) {
		var test3 = TestNested03{Field32: 22}
		b, err := test3.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 22}`, string(b))
	})
	t.Run("nested-2", func(t *testing.T) {
		var test2 = TestNested02{TestNested03{Field32: 33}}
		b, err := test2.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 33}`, string(b))
	})
	t.Run("nested-1", func(t *testing.T) {
		var test1 = TestNested01{TestNested02{TestNested03{Field32: 44}}}
		b, err := test1.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 44}`, string(b))
	})
}

func Test_JsonTestInh01(t *testing.T) {
	t.Run("TestInh02", func(t *testing.T) {
		var test = TestInh02{
			Int32: 1112,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"int_32":1112}`, string(b))
	})
	t.Run("omit-injected", func(t *testing.T) {
		var nested = TestInh03{
			Int16:  2222,
			Random: 44443,
		}
		var test = TestInh01{
			TestInh03: TestInh03{
				Int16:  16003,
				Random: 45222,
			},
			DateBegin: time.Date(2001, time.December, 31, 12, 11, 10, 0, time.UTC),
			Nested1: TestInh03{
				Int16:  4120,
				Random: 9889,
			},
			Nested2: &nested,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"int_16":16003,"random":45222,"date_begin":"2001-12-31T12:11:10Z","nested1":{"int_16":4120,"random":9889},"nested2":{"int_16":2222,"random":44443}}`, string(b))
	})
	t.Run("omit-nested2", func(t *testing.T) {
		var test = TestInh01{
			TestInh02: TestInh02{
				Int32: 1112,
			},
			TestInh03: TestInh03{
				Int16:  16003,
				Random: 45222,
			},
			DateBegin: time.Date(2001, time.December, 31, 12, 11, 10, 0, time.UTC),
			Nested1: TestInh03{
				Int16:  4120,
				Random: 9889,
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"injected":{"int_32":1112},"int_16":16003,"random":45222,"date_begin":"2001-12-31T12:11:10Z","nested1":{"int_16":4120,"random":9889},"nested2": null}`, string(b))
	})
	t.Run("empty-struct", func(t *testing.T) {
		var test = TestInh01{}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"int_16":0,"random":0,"date_begin":"0000-00-00T00:00:00Z","nested1":{"int_16":0,"random":0},"nested2": null}`, string(b))
	})
	t.Run("TestNested01", func(t *testing.T) {
		var test = TestNested01{
			TestNested02{
				TestNested03{
					Field32: 32,
				},
			},
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 32}`, string(b))
	})
	t.Run("whole-struct", func(t *testing.T) {
		var nested = TestInh03{
			Int16:  2222,
			Random: 44443,
		}
		var test = TestInh01{
			TestInh02: TestInh02{
				Int32: 1112,
			},
			TestInh03: TestInh03{
				Int16:  16003,
				Random: 45222,
			},
			DateBegin: time.Date(2001, time.December, 31, 12, 11, 10, 0, time.UTC),
			Nested1: TestInh03{
				Int16:  4120,
				Random: 9889,
			},
			Nested2: &nested,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"injected":{"int_32":1112},"int_16":16003,"random":45222,"date_begin":"2001-12-31T12:11:10Z","nested1":{"int_16":4120,"random":9889},"nested2":{"int_16":2222,"random":44443}}`, string(b))
	})
}
