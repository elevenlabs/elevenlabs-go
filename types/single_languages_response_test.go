package types

import (
	"testing"
)

func TestSingleLanguagesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "single",
  "languages": [
    {
      "code": "value",
      "label": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(SingleLanguagesResponse))
}
