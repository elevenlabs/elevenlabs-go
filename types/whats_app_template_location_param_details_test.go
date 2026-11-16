package types

import (
	"testing"
)

func TestWhatsAppTemplateLocationParamDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "latitude": "value",
  "longitude": "value",
  "name": "value",
  "address": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateLocationParamDetails))
}
