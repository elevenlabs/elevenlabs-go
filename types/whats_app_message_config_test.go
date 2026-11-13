package types

import (
	"testing"
)

func TestWhatsAppMessageConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "message",
  "template": {
    "name": "value",
    "language": "value",
    "parameter_format": "POSITIONAL",
    "components": [
      {
        "type": "value",
        "format": "value",
        "text": "value",
        "buttons": [
          {
            "type": "value",
            "text": "value",
            "url": "value"
          }
        ]
      }
    ],
    "category": "value"
  }
}`
	assertRoundTrip(t, wire, new(WhatsAppMessageConfig))
}
