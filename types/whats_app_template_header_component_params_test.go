package types

import (
	"testing"
)

func TestWhatsAppTemplateHeaderComponentParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "header",
  "parameters": [
    {
      "parameter_name": "value",
      "type": "text",
      "text": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateHeaderComponentParams))
}
