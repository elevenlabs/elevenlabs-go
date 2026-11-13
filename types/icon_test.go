package types

import (
	"testing"
)

func TestIconRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "src": "value",
  "mimeType": "value",
  "sizes": [
    "value"
  ],
  "theme": "light"
}`
	assertRoundTrip(t, wire, new(Icon))
}
