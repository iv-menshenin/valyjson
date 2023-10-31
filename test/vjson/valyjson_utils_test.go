package vjson

import (
	"math"
	"testing"
	"time"

	"github.com/mailru/easyjson/jwriter"
	"github.com/stretchr/testify/require"
)

func Test_writeInt64(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w jwriter.Writer
		writeInt64(&w, 0)
		require.Equal(t, "0", string(w.Buffer.Buf))
	})
	t.Run("positive", func(t *testing.T) {
		var w jwriter.Writer
		writeInt64(&w, 19)
		require.Equal(t, "19", string(w.Buffer.Buf))
	})
	t.Run("negative", func(t *testing.T) {
		var w jwriter.Writer
		writeInt64(&w, -32456)
		require.Equal(t, "-32456", string(w.Buffer.Buf))
	})
	t.Run("minInt", func(t *testing.T) {
		var w jwriter.Writer
		writeInt64(&w, math.MinInt64)
		require.Equal(t, "-9223372036854775808", string(w.Buffer.Buf))
	})
	t.Run("maxInt", func(t *testing.T) {
		var w jwriter.Writer
		writeInt64(&w, math.MaxInt64)
		require.Equal(t, "9223372036854775807", string(w.Buffer.Buf))
	})
	t.Run("allocations", func(t *testing.T) {
		var w jwriter.Writer
		n := testing.AllocsPerRun(100, func() {
			writeInt64(&w, 19)
		})
		require.EqualValues(t, 0, n)
	})
}

func Test_writeUint64(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w jwriter.Writer
		writeUint64(&w, 0)
		require.Equal(t, "0", string(w.Buffer.Buf))
	})
	t.Run("positive", func(t *testing.T) {
		var w jwriter.Writer
		writeUint64(&w, 19)
		require.Equal(t, "19", string(w.Buffer.Buf))
	})
	t.Run("minUint", func(t *testing.T) {
		var w jwriter.Writer
		writeUint64(&w, math.MaxUint64)
		require.Equal(t, "18446744073709551615", string(w.Buffer.Buf))
	})
	t.Run("maxUint", func(t *testing.T) {
		var w jwriter.Writer
		writeUint64(&w, math.MaxInt64)
		require.Equal(t, "9223372036854775807", string(w.Buffer.Buf))
	})
	t.Run("allocations", func(t *testing.T) {
		var w jwriter.Writer
		n := testing.AllocsPerRun(100, func() {
			writeUint64(&w, 19)
		})
		require.EqualValues(t, 0, n)
	})
}

func Test_writeFloat64(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w jwriter.Writer
		writeFloat64(&w, 0)
		require.Equal(t, "0", string(w.Buffer.Buf))
	})
	t.Run("positive", func(t *testing.T) {
		var w jwriter.Writer
		writeFloat64(&w, 19)
		require.Equal(t, "19", string(w.Buffer.Buf))
	})
	t.Run("negative", func(t *testing.T) {
		var w jwriter.Writer
		writeFloat64(&w, -32456)
		require.Equal(t, "-32456", string(w.Buffer.Buf))
	})
	t.Run("precision", func(t *testing.T) {
		var w jwriter.Writer
		writeFloat64(&w, 0.0004)
		require.Equal(t, "0.0004", string(w.Buffer.Buf))
	})
	t.Run("allocations", func(t *testing.T) {
		var w jwriter.Writer
		n := testing.AllocsPerRun(100, func() {
			writeFloat64(&w, math.MaxFloat64)
		})
		require.EqualValues(t, 0, n)
	})
}

func Test_writeTime(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		var w jwriter.Writer
		writeTime(&w, time.Time{}, "20060102150405")
		require.Equal(t, `"00010101000000"`, string(w.Buffer.Buf))
	})
	t.Run("RFC3339Nano", func(t *testing.T) {
		var w jwriter.Writer
		writeTime(&w, time.Date(2023, time.March, 14, 21, 8, 55, 99, time.UTC), time.RFC3339Nano)
		require.Equal(t, `"2023-03-14T21:08:55.000000099Z"`, string(w.Buffer.Buf))
	})
	t.Run("allocations", func(t *testing.T) {
		var w jwriter.Writer
		var d = time.Date(2023, time.March, 14, 21, 12, 31, 734, time.UTC)
		n := testing.AllocsPerRun(100, func() {
			writeTime(&w, d, time.RFC3339Nano)
		})
		require.EqualValues(t, 0, n)
	})
}
