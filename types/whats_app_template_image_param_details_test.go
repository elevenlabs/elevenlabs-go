package types

import (
	"testing"
)

func TestWhatsAppTemplateImageParamDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "link": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateImageParamDetails))
}
