package types

import (
	"testing"
)

func TestGetWhatsAppAccountResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "business_account_id": "value",
  "phone_number_id": "value",
  "business_account_name": "value",
  "phone_number_name": "value",
  "phone_number": "value",
  "account_type": "cloud_api",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true,
  "assigned_agent_name": "value",
  "is_token_expired": true
}`
	assertRoundTrip(t, wire, new(GetWhatsAppAccountResponse))
}

func TestGetWhatsAppAccountResponseRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "business_account_id": "value",
  "phone_number_id": "value",
  "business_account_name": "value",
  "phone_number_name": "value",
  "phone_number": "value",
  "account_type": "cloud_api",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true,
  "assigned_agent_name": "value",
  "is_token_expired": true
}`
	assertRoundTrip(t, wire, new(GetWhatsAppAccountResponse))
}

func TestGetWhatsAppAccountResponseRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "business_account_id": "value",
  "phone_number_id": "value",
  "business_account_name": "value",
  "phone_number_name": "value",
  "phone_number": "value",
  "account_type": "cloud_api",
  "assigned_agent_id": "value",
  "enable_messaging": true,
  "enable_audio_message_response": true,
  "enable_typing_indicator": true,
  "assigned_agent_name": "value",
  "is_token_expired": true
}`
	assertRoundTrip(t, wire, new(GetWhatsAppAccountResponse))
}
