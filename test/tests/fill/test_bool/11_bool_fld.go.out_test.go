package test_bool

import (
	"github.com/mailru/easyjson/jwriter"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTestBool01_MarshalJSON(t *testing.T) {
	t.Run("allocations", func(t *testing.T) {
		var (
			True  = true
			False = false
		)
		var test = &TestBool01{
			Bool:     false,
			BlMaybe:  true,
			RefBool:  &True,
			RefMaybe: &False,
		}
		var jw jwriter.Writer
		n := testing.AllocsPerRun(1000, func() {
			err := test.MarshalTo(&jw)
			if err != nil {
				t.Error(err)
			}
		})
		require.LessOrEqual(t, n, float64(0))
	})
	t.Run("filled-all", func(t *testing.T) {
		var (
			True  = true
			False = false
		)
		var test = &TestBool01{
			Bool:     true,
			BlMaybe:  true,
			RefBool:  &True,
			RefMaybe: &False,
			DefBool:  true,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"bl": true, "mb": true, "refBool": true, "refMaybe": false, "defBool": true}`, string(b))
	})
	t.Run("omit-false", func(t *testing.T) {
		var (
			True  = true
			False = false
		)
		var test = &TestBool01{
			Bool:     false,
			BlMaybe:  false,
			RefBool:  &True,
			RefMaybe: &False,
			DefBool:  true,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"bl": false, "refBool": true, "refMaybe": false, "defBool": true}`, string(b))
	})
	t.Run("omit-nil", func(t *testing.T) {
		var True = true
		var test = &TestBool01{
			Bool:     false,
			BlMaybe:  true,
			RefBool:  &True,
			RefMaybe: nil,
			DefBool:  false,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"bl": false, "mb": true, "refBool": true, "defBool": false}`, string(b))
	})
	t.Run("nil-as-nil", func(t *testing.T) {
		var test = &TestBool01{
			Bool:    false,
			BlMaybe: true,
			DefBool: false,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"bl": false, "mb": true, "refBool": null, "defBool": false}`, string(b))
	})
	t.Run("null", func(t *testing.T) {
		var test *TestBool01
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `null`, string(b))
	})
}

func TestTestBool01_UnmarshalJSON(t *testing.T) {
	t.Run("filled-all", func(t *testing.T) {
		var (
			True  = true
			False = false
		)
		var test TestBool01
		var expected = TestBool01{
			Bool:     false,
			BlMaybe:  true,
			RefBool:  &True,
			RefMaybe: &False,
			DefBool:  true,
		}
		err := test.UnmarshalJSON([]byte(`{"bl": false, "mb": true, "refBool": true, "refMaybe": false}`))
		require.NoError(t, err)
		require.Equal(t, expected, test)
		require.NotEmpty(t, test)
		test.Reset()
		require.Empty(t, test)
	})
	t.Run("omit-fields", func(t *testing.T) {
		var True = true
		var test TestBool01
		var expected = TestBool01{
			Bool:     true,
			BlMaybe:  false,
			RefBool:  &True,
			RefMaybe: nil,
			DefBool:  true,
		}
		err := test.UnmarshalJSON([]byte(`{"bl": true, "refBool": true}`))
		require.NoError(t, err)
		require.Equal(t, expected, test)
		require.NotEmpty(t, test)
		test.Reset()
		require.Empty(t, test)
	})
	t.Run("refs-as-null", func(t *testing.T) {
		var test TestBool01
		var expected = TestBool01{
			Bool:    false,
			BlMaybe: true,
			DefBool: true,
		}
		err := test.UnmarshalJSON([]byte(`{"mb": true, "refBool": null, "refMaybe": null}`))
		require.NoError(t, err)
		require.Equal(t, expected, test)
		require.NotEmpty(t, test)
		test.Reset()
		require.Empty(t, test)
	})
	t.Run("def-bool-false", func(t *testing.T) {
		var test TestBool01
		var expected = TestBool01{
			Bool:    false,
			BlMaybe: true,
			DefBool: false,
		}
		err := test.UnmarshalJSON([]byte(`{"mb": true, "defBool": false}`))
		require.NoError(t, err)
		require.Equal(t, expected, test)
		require.NotEmpty(t, test)
		test.Reset()
		require.Empty(t, test)
	})
	t.Run("invalid-type", func(t *testing.T) {
		var test TestBool01

		err := test.UnmarshalJSON([]byte(`{"mb": null, "defBool": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain bool")

		err = test.UnmarshalJSON([]byte(`{"bl": null, "defBool": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain bool")

		err = test.UnmarshalJSON([]byte(`{"bl": true, "refBool": "null"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain bool")

		err = test.UnmarshalJSON([]byte(`{"bl": true, "refMaybe": "null"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain bool")

		err = test.UnmarshalJSON([]byte(`{"bl": true, "defBool": "null"}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "value doesn't contain bool")
	})
	t.Run("invalid-format", func(t *testing.T) {
		var test TestBool01
		err := test.UnmarshalJSON([]byte(`{"mb": f, "defBool": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "unexpected value found")
	})
	t.Run("strict-validation", func(t *testing.T) {
		var test TestBool01
		err := test.UnmarshalJSON([]byte(`{"mb": true, "unknown": 123, "defBool": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "unexpected field")
		require.ErrorContains(t, err, "unknown")
	})
	t.Run("double-fields", func(t *testing.T) {
		var test TestBool01
		err := test.UnmarshalJSON([]byte(`{"mb": true, "defBool": null, "defBool": null}`))
		require.Error(t, err)
		require.ErrorContains(t, err, "field appears in the object twice")
	})
}

func Test_InhBool(t *testing.T) {
	t.Run("unmarhal", func(t *testing.T) {
		var test TestInhBool

		require.NoError(t, test.UnmarshalJSON([]byte("true")))
		require.True(t, bool(test))

		require.NoError(t, test.UnmarshalJSON([]byte("false")))
		require.False(t, bool(test))

		require.Error(t, test.UnmarshalJSON([]byte("-")))
	})
	t.Run("marshal", func(t *testing.T) {
		var test1 TestInhBool = true
		data, err := test1.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "true", string(data))

		var test2 TestInhBool = false
		data, err = test2.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "false", string(data))

		var test3 *TestInhBool
		data, err = test3.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, "null", string(data))
	})
}

func Benchmark_TestBool01_MarshalJSON(b *testing.B) {
	var (
		True  = true
		False = false
	)
	var test = &TestBool01{
		Bool:     false,
		BlMaybe:  true,
		RefBool:  &True,
		RefMaybe: &False,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := test.MarshalJSON()
		if err != nil {
			b.Error(err)
		}
	}
}

func Test_TestBool02(t *testing.T) {
	t.Run("Unmarshal", func(t *testing.T) {
		var got TestBool02
		var expected = TestBool02{
			I: true,
			X: true,
		}
		const data = `{"i": true, "x": true}`
		err := got.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("Marshal", func(t *testing.T) {
		var test = TestBool02{
			I: true,
			X: true,
		}
		const expected = `{"x":true,"i":true}`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
}

func Test_RequiredRequired(t *testing.T) {
	t.Run("NotRequiredUnmarshal", func(t *testing.T) {
		var got TestBool03
		var expected = TestBool03{}
		const data = `{}`
		err := got.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("NotRequiredMarshal", func(t *testing.T) {
		var test = TestBool03{}
		const expected = `{"required":false}`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("RequiredUnmarshal", func(t *testing.T) {
		var got TestBool04
		var expected = TestBool04{}
		const data = `{"required":false}`
		err := got.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("RequiredCheck", func(t *testing.T) {
		var got TestBool04
		const data = `{}`
		err := got.UnmarshalJSON([]byte(data))
		require.Error(t, err)
		require.ErrorContains(t, err, "'required' is missing")
	})
	t.Run("RequiredMarshal", func(t *testing.T) {
		var test = TestBool04{}
		const expected = `{"required":false}`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
	t.Run("OmittedRequiredUnmarshal", func(t *testing.T) {
		var got TestBool05
		var expected = TestBool05{}
		const data = `{}`
		err := got.UnmarshalJSON([]byte(data))
		require.NoError(t, err)
		require.Equal(t, expected, got)
	})
	t.Run("OmittedRequiredMarshal", func(t *testing.T) {
		var test = TestBool05{}
		const expected = `{}`
		data, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, expected, string(data))
	})
}

func TestReset(t *testing.T) {
	t.Parallel()
	t.Run("TestBool01", func(t *testing.T) {
		t.Parallel()
		var v1 = true
		var val = TestBool01{
			Bool:     true,
			BlMaybe:  true,
			RefBool:  &v1,
			RefMaybe: &v1,
			DefBool:  true,
		}
		val.Reset()
		require.Empty(t, val)
	})
	t.Run("TestBool02", func(t *testing.T) {
		t.Parallel()
		var val = TestBool02{
			I: true,
			X: true,
		}
		val.Reset()
		require.Empty(t, val)
	})
	t.Run("TestInhBool", func(t *testing.T) {
		t.Parallel()
		var val TestInhBool = true
		val.Reset()
		require.False(t, bool(val))
	})
	t.Run("TestBool03", func(t *testing.T) {
		t.Parallel()
		var val = TestBool03{Required: true}
		val.Reset()
		require.False(t, val.Required)
	})
	t.Run("TestBool04", func(t *testing.T) {
		t.Parallel()
		var val = TestBool04{Required: true}
		val.Reset()
		require.False(t, val.Required)
	})
	t.Run("TestBool05", func(t *testing.T) {
		t.Parallel()
		var val = TestBool05{Required: true}
		val.Reset()
		require.False(t, val.Required)
	})
}
