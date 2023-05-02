package benchmark

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestXLStruct_FillFromJSON(t *testing.T) {
	var x XLStruct
	const data = `{"data":[{"statuses":[{"entities":{"hashtags":[{},{"text":6}]}}]}]}`
	err := x.UnmarshalJSON([]byte(data))
	require.ErrorContains(t, err, "data.0.statuses.0.entities.hashtags.1.text")
}
