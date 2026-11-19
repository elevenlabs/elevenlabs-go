package types

import (
	"testing"
)

func TestDownloadHistoryRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "history_item_ids": [
    "value"
  ],
  "output_format": "value"
}`
	assertRoundTrip(t, wire, new(DownloadHistoryRequest))
}
