package types

import (
	"testing"
)

func TestWhatsAppTemplateDocumentParamRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "document",
  "document": {
    "link": "value",
    "filename": "value"
  }
}`
	assertRoundTrip(t, wire, new(WhatsAppTemplateDocumentParam))
}
