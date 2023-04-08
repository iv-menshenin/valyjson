package vjson

import (
	"bytes"
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_writeInt64(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeInt64(w, 0)
		require.Equal(t, "0", w.String())
	})
	t.Run("positive", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeInt64(w, 19)
		require.Equal(t, "19", w.String())
	})
	t.Run("negative", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeInt64(w, -32456)
		require.Equal(t, "-32456", w.String())
	})
	t.Run("minInt", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeInt64(w, math.MinInt64)
		require.Equal(t, "-9223372036854775808", w.String())
	})
	t.Run("maxInt", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeInt64(w, math.MaxInt64)
		require.Equal(t, "9223372036854775807", w.String())
	})
	t.Run("allocations", func(t *testing.T) {
		var w = nullWriter{}
		n := testing.AllocsPerRun(100, func() {
			writeInt64(w, 19)
		})
		require.EqualValues(t, 0, n)
	})
}

func Test_writeUint64(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeUint64(w, 0)
		require.Equal(t, "0", w.String())
	})
	t.Run("positive", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeUint64(w, 19)
		require.Equal(t, "19", w.String())
	})
	t.Run("minUint", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeUint64(w, math.MaxUint64)
		require.Equal(t, "18446744073709551615", w.String())
	})
	t.Run("maxUint", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeUint64(w, math.MaxInt64)
		require.Equal(t, "9223372036854775807", w.String())
	})
	t.Run("allocations", func(t *testing.T) {
		var w = nullWriter{}
		n := testing.AllocsPerRun(100, func() {
			writeUint64(w, 19)
		})
		require.EqualValues(t, 0, n)
	})
}

func Test_writeFloat64(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeFloat64(w, 0)
		require.Equal(t, "0", w.String())
	})
	t.Run("positive", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeFloat64(w, 19)
		require.Equal(t, "19", w.String())
	})
	t.Run("negative", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeFloat64(w, -32456)
		require.Equal(t, "-32456", w.String())
	})
	t.Run("precision", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeFloat64(w, 0.00000004)
		require.Equal(t, "0.00000004", w.String())
	})
	t.Run("allocations", func(t *testing.T) {
		var w = nullWriter{}
		n := testing.AllocsPerRun(100, func() {
			writeFloat64(w, math.MaxFloat64)
		})
		require.EqualValues(t, 0, n)
	})
}

func Test_writeTime(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeTime(w, time.Time{}, "20060102150405")
		require.Equal(t, `"00010101000000"`, w.String())
	})
	t.Run("RFC3339Nano", func(t *testing.T) {
		var w = bytes.NewBufferString("")
		writeTime(w, time.Date(2023, time.March, 14, 21, 8, 55, 99, time.UTC), time.RFC3339Nano)
		require.Equal(t, `"2023-03-14T21:08:55.000000099Z"`, w.String())
	})
	t.Run("allocations", func(t *testing.T) {
		var w = nullWriter{}
		var d = time.Date(2023, time.March, 14, 21, 12, 31, 734, time.UTC)
		n := testing.AllocsPerRun(100, func() {
			writeTime(w, d, time.RFC3339Nano)
		})
		require.EqualValues(t, 0, n)
	})
}

type (
	nullWriter struct{}
)

func (nullWriter) Write(b []byte) (n int, err error) { return len(b), nil }

func (nullWriter) WriteString(b string) (n int, err error) { return len(b), nil }

func Test_writeString(t *testing.T) {
	t.Run("allocation-without-special", func(t *testing.T) {
		var w = nullWriter{}
		n := testing.AllocsPerRun(100, func() {
			writeString(w, "это не те дроиды,которых вы ищете")
		})
		require.EqualValues(t, 0, n)
	})
	t.Run("allocation-with-special", func(t *testing.T) {
		var w = nullWriter{}
		n := testing.AllocsPerRun(100, func() {
			writeString(w, "эй ты\nиди сюда\n - \"чего?\"")
		})
		require.EqualValues(t, 1, n)
	})
}
