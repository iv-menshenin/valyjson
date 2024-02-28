package vjson

import (
	"testing"
	"time"

	"github.com/mailru/easyjson/jwriter"
	"github.com/stretchr/testify/require"
)

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
		require.LessOrEqual(t, float64(1), n)
	})
}
