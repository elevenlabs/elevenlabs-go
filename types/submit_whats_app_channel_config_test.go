package types

import (
	"testing"
)

func TestSubmitWhatsAppChannelConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "channel": "whatsapp",
  "whatsapp_phone_number_id": "value",
  "mode_config": {
    "mode": "voice",
    "whatsapp_call_permission_request_template_name": "value",
    "whatsapp_call_permission_request_template_language_code": "value"
  }
}`
	assertRoundTrip(t, wire, new(SubmitWhatsAppChannelConfig))
}
