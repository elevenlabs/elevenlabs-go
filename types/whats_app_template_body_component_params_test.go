package types

import (
	"testing"
)

func TestWhatsAppTemplateBodyComponentParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "body",
  "parameters": [
    {
      "parameter_name": "value",
      "type": "text",
      "text": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateBodyComponentParams))
}
