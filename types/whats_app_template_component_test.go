package types

import (
	"testing"
)

func TestWhatsAppTemplateComponentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateComponent))
}
