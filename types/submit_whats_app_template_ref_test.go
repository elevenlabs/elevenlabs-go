package types

import (
	"testing"
)

func TestSubmitWhatsAppTemplateRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "language": "value"
}`
	assertRoundTrip(t, wire, new(SubmitWhatsAppTemplateRef))
}
