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
	t.Run("marshal", func(t *testing.T) {
		const expected = `{"breakFirst":-1,"injected":{"int_32":123},"int_16":16,"random":-9,"date_begin":"2023-04-06T00:00:00Z","nested1":{"int_16":17,"random":-8},"nested2":{"int_16":22,"random":88}}`
		var n = TestInh03{
			Int16:  22,
			Random: 88,
		}
		var test1 = TestInh01{
			BreakFirst: -1,
			TestInh02:  TestInh02{Int32: 123},
			TestInh03: TestInh03{
				Int16:  16,
				Random: -9,
			},
			DateBegin: time.Date(2023, time.April, 6, 0, 0, 0, 0, time.UTC),
			Nested1: TestInh03{
				Int16:  17,
				Random: -8,
			},
			Nested2: &n,
		}
		data, err := test1.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("marshal_omitted", func(t *testing.T) {
		const expected = `{"int_16":16,"random":-9,"date_begin":"2023-04-06T00:00:00Z","nested1":{"int_16":17,"random":-8},"nested2":null}`
		var test1 = TestInh01{
			TestInh03: TestInh03{
				Int16:  16,
				Random: -9,
			},
			DateBegin: time.Date(2023, time.April, 6, 0, 0, 0, 0, time.UTC),
			Nested1: TestInh03{
				Int16:  17,
				Random: -8,
			},
		}
		data, err := test1.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
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
	t.Run("unmarshal_err", func(t *testing.T) {
		var got TestNested01
		err := got.UnmarshalJSON([]byte(`{"field_31": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_31")
	})
	t.Run("unmarshal_json_err", func(t *testing.T) {
		var got TestNested01
		err := got.UnmarshalJSON([]byte(`{"field_32: 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "cannot parse object")
	})
	t.Run("unmarshal_not_object", func(t *testing.T) {
		var got TestNested01
		err := got.UnmarshalJSON([]byte(`5`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("unmarshal_json_err_double", func(t *testing.T) {
		var got TestNested01
		err := got.UnmarshalJSON([]byte(`{"field_32": 4444, "field_32": 4444, "field_32": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_32")
		require.ErrorContains(t, err, "appears in the object twice")
	})
	t.Run("marshal", func(t *testing.T) {
		var test = TestNested01{
			TestNested02{
				TestNested03{
					Field32: 1324,
				},
			},
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 1324}`, string(data))
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
	t.Run("omit-injected-check-comma", func(t *testing.T) {
		var nested = TestInh03{
			Int16:  2222,
			Random: 44443,
		}
		var test = TestInh01{
			BreakFirst: 1230000,
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
		require.JSONEq(t, `{"breakFirst":1230000,"int_16":16003,"random":45222,"date_begin":"2001-12-31T12:11:10Z","nested1":{"int_16":4120,"random":9889},"nested2":{"int_16":2222,"random":44443}}`, string(b))
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
		require.JSONEq(t, `{"int_16":0,"random":0,"date_begin":"0001-01-01T00:00:00Z","nested1":{"int_16":0,"random":0},"nested2": null}`, string(b))
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

func Test_Null(t *testing.T) {
	t.Run("TestInh01", func(t *testing.T) {
		var test *TestInh01
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("TestInh02", func(t *testing.T) {
		var test *TestInh02
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("TestInh03", func(t *testing.T) {
		var test *TestInh03
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("TestNested01", func(t *testing.T) {
		var test *TestNested01
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("TestNested02", func(t *testing.T) {
		var test *TestNested02
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("TestNested03", func(t *testing.T) {
		var test *TestNested03
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
	t.Run("TestNested04", func(t *testing.T) {
		var test *TestNested04
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
}

func Test_IsZero(t *testing.T) {
	t.Run("TestNested", func(t *testing.T) {
		var test1 TestNested01
		require.True(t, test1.IsZero())
		var test2 TestNested02
		require.True(t, test2.IsZero())
		var test3 TestNested03
		require.True(t, test3.IsZero())
		var test4 TestNested04
		require.True(t, test4.IsZero())
	})
	t.Run("TestInh", func(t *testing.T) {
		var test1 TestInh01
		require.True(t, test1.IsZero())
		var test2 TestInh02
		require.True(t, test2.IsZero())
		var test3 TestInh03
		require.True(t, test3.IsZero())
	})
	t.Run("TestInh01_notZero_1", func(t *testing.T) {
		var test1 TestInh01
		test1.TestInh03.Int16 = 1
		require.False(t, test1.IsZero())

		var test2 TestInh01
		test2.Nested1.Int16 = 1
		require.False(t, test2.IsZero())
	})
	t.Run("TestInh01_notZero_2", func(t *testing.T) {
		var test TestInh01
		test.Random = 1
		require.False(t, test.IsZero())
	})
	t.Run("TestInh01_notZero_3", func(t *testing.T) {
		var test TestInh01
		test.Nested2 = &TestInh03{}
		require.False(t, test.IsZero())
	})
	t.Run("TestInh01_notZero_4", func(t *testing.T) {
		var test TestInh01
		test.DateBegin = time.Now()
		require.False(t, test.IsZero())
	})
	t.Run("TestInh01_notZero_5", func(t *testing.T) {
		var test TestInh01
		test.Int32 = 2
		require.False(t, test.IsZero())
	})
}

func Test_TestNested04(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		var test = TestNested04{
			Field32: 2222,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 2222}`, string(data))
	})
	t.Run("unmarshal", func(t *testing.T) {
		var expected = TestNested04{
			Field32: 4444,
		}
		var got TestNested04
		require.NoError(t, got.UnmarshalJSON([]byte(`{"field_32": 4444}`)))
		require.Equal(t, expected, got)
	})
	t.Run("unmarshal_err", func(t *testing.T) {
		var got TestNested04
		err := got.UnmarshalJSON([]byte(`{"field_31": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_31")
	})
	t.Run("unmarshal_json_err", func(t *testing.T) {
		var got TestNested04
		err := got.UnmarshalJSON([]byte(`{"field_32: 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "cannot parse object")
	})
	t.Run("unmarshal_not_object", func(t *testing.T) {
		var got TestNested04
		err := got.UnmarshalJSON([]byte(`5`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("unmarshal_json_err_double", func(t *testing.T) {
		var got TestNested04
		err := got.UnmarshalJSON([]byte(`{"field_32": 4444, "field_32": 4444, "field_32": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_32")
		require.ErrorContains(t, err, "appears in the object twice")
	})
}

func Test_TestNested03(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		var test = TestNested03{
			Field32: 2222,
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 2222}`, string(data))
	})
	t.Run("unmarshal", func(t *testing.T) {
		var expected = TestNested03{
			Field32: 4444,
		}
		var got TestNested03
		require.NoError(t, got.UnmarshalJSON([]byte(`{"field_32": 4444}`)))
		require.Equal(t, expected, got)
	})
	t.Run("unmarshal_err", func(t *testing.T) {
		var got TestNested03
		err := got.UnmarshalJSON([]byte(`{"field_31": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_31")
	})
	t.Run("unmarshal_json_err", func(t *testing.T) {
		var got TestNested03
		err := got.UnmarshalJSON([]byte(`{"field_32: 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "cannot parse object")
	})
	t.Run("unmarshal_not_object", func(t *testing.T) {
		var got TestNested03
		err := got.UnmarshalJSON([]byte(`5`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("unmarshal_json_err_double", func(t *testing.T) {
		var got TestNested03
		err := got.UnmarshalJSON([]byte(`{"field_32": 4444, "field_32": 4444, "field_32": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_32")
		require.ErrorContains(t, err, "appears in the object twice")
	})
}

func Test_TestNested02(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		var test = TestNested02{
			TestNested03{
				Field32: 2222,
			},
		}
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"field_32": 2222}`, string(data))
	})
	t.Run("unmarshal", func(t *testing.T) {
		var expected = TestNested02{
			TestNested03{
				Field32: 4444,
			},
		}
		var got TestNested02
		require.NoError(t, got.UnmarshalJSON([]byte(`{"field_32": 4444}`)))
		require.Equal(t, expected, got)
	})
	t.Run("unmarshal_err", func(t *testing.T) {
		var got TestNested02
		err := got.UnmarshalJSON([]byte(`{"field_31": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_31")
	})
	t.Run("unmarshal_json_err", func(t *testing.T) {
		var got TestNested02
		err := got.UnmarshalJSON([]byte(`{"field_32: 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "cannot parse object")
	})
	t.Run("unmarshal_not_object", func(t *testing.T) {
		var got TestNested02
		err := got.UnmarshalJSON([]byte(`5`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain object")
	})
	t.Run("unmarshal_json_err_double", func(t *testing.T) {
		var got TestNested02
		err := got.UnmarshalJSON([]byte(`{"field_32": 4444, "field_32": 4444, "field_32": 4444}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field_32")
		require.ErrorContains(t, err, "appears in the object twice")
	})
}
