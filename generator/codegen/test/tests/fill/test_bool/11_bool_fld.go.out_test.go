package test_bool

import (
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
		n := testing.AllocsPerRun(100, func() {
			_, err := test.MarshalJSON()
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
			Bool:     false,
			BlMaybe:  true,
			RefBool:  &True,
			RefMaybe: &False,
			DefBool:  true,
		}
		b, err := test.MarshalJSON()
		require.NoError(t, err)
		require.JSONEq(t, `{"bl": false, "mb": true, "refBool": true, "refMaybe": false, "defBool": true}`, string(b))
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
