package types

import (
	"testing"
)

func TestResponseFilterRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "all",
  "filters": [
    "value"
  ],
  "content_type": "application/json"
}`
	assertRoundTrip(t, wire, new(ResponseFilter))
}
