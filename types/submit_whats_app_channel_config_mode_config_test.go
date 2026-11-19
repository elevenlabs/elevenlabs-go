package types

import (
	"testing"
)

func TestSubmitWhatsAppChannelConfigModeConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "voice",
  "whatsapp_call_permission_request_template_name": "value",
  "whatsapp_call_permission_request_template_language_code": "value"
}`
	assertRoundTrip(t, wire, new(SubmitWhatsAppChannelConfigModeConfig))
}
