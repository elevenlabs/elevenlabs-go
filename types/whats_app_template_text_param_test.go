package types

import (
	"testing"
)

func TestWhatsAppTemplateTextParamRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "parameter_name": "value",
  "type": "text",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateTextParam))
}
