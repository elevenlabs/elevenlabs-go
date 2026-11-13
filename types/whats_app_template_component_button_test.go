package types

import (
	"testing"
)

func TestWhatsAppTemplateComponentButtonRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "value",
  "text": "value",
  "url": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateComponentButton))
}
