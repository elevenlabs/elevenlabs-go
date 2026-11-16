package types

import (
	"testing"
)

func TestContentConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sexual": {
    "is_enabled": true,
    "threshold": 9.99
  },
  "violence": {
    "is_enabled": true,
    "threshold": 9.99
  },
  "harassment": {
    "is_enabled": true,
    "threshold": 9.99
  },
  "self_harm": {
    "is_enabled": true,
    "threshold": 9.99
  },
  "profanity": {
    "is_enabled": true,
    "threshold": 9.99
  },
  "religion_or_politics": {
    "is_enabled": true,
    "threshold": 9.99
  },
  "medical_and_legal_information": {
    "is_enabled": true,
    "threshold": 9.99
  }
}`
	assertRoundTrip(t, wire, new(ContentConfig))
}
