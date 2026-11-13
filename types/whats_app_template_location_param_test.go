package types

import (
	"testing"
)

func TestWhatsAppTemplateLocationParamRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "location",
  "location": {
    "latitude": "value",
    "longitude": "value",
    "name": "value",
    "address": "value"
  }
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateLocationParam))
}
