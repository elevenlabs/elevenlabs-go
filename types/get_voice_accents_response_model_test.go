package types

import (
	"testing"
)

func TestGetVoiceAccentsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "accents": [
    {
      "accent": "value",
      "language": "value",
      "code": "value",
      "name": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(GetVoiceAccentsResponseModel))
}
