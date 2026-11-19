package types

import (
	"testing"
)

func TestWhatsAppTemplateImageParamRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "image",
  "image": {
    "link": "value"
  }
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateImageParam))
}
