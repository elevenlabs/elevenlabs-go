package types

import (
	"testing"
)

func TestWhatsAppTemplateDocumentParamDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "link": "value",
  "filename": "value"
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateDocumentParamDetails))
}
