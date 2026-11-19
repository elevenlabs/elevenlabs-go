package types

import (
	"testing"
)

func TestWhatsAppTemplateRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplate))
}
