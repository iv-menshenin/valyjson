package test_inh

import (
	"github.com/stretchr/testify/require"
	"testing"
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
		require.ErrorContains(t, err, "error parsing 'random' value")
	})
	t.Run("test-wrong-nested-type", func(t *testing.T) {
		var test1 TestInh01
		err := test1.UnmarshalJSON([]byte(`{"int_16":543,"random":66,"nested1":{"int_16":888,"random":"999"},"nested2":{"int_16":777,"random":666},"date_begin":"2023-01-28 07:10:05Z"}`))
		require.ErrorContains(t, err, "error parsing 'nested1.random' value")
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
