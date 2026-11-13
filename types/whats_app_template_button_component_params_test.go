package types

import (
	"testing"
)

func TestWhatsAppTemplateButtonComponentParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "button",
  "parameters": [
    {
      "parameter_name": "value",
      "type": "text",
      "text": "value"
    }
  ],
  "index": 123,
  "sub_type": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateButtonComponentParams))
}
