package types

import (
	"testing"
)

func TestSubmitWhatsAppMessageConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "message",
  "template": {
    "name": "value",
    "language": "value"
  }
}`
	assertRoundTrip(t, wire, new(SubmitWhatsAppMessageConfig))
}
