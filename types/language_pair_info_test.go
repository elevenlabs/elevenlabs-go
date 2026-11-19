package types

import (
	"testing"
)

func TestLanguagePairInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(LanguagePairInfo))
}
