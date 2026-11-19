package types

import (
	"testing"
)

func TestPairedLanguagesResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "pair",
  "language_pairs": [
    {
      "source_language": {
        "code": "value",
        "label": "value"
      },
      "destination_languages": [
        {
          "code": "value",
          "label": "value"
        }
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(PairedLanguagesResponse))
}
