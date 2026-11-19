package agents

import (
	"testing"
)

func TestListPhoneNumbersResponseItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`
	assertRoundTrip(t, wire, new(ListPhoneNumbersResponseItem))
}

func TestCreatePhoneNumbersRequestBodyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "provider": "twilio",
  "agent_id": "value",
  "sid": "value",
  "token": "value",
  "account_auth_token": "value",
  "region_config": {
    "region_id": "us1",
    "token": "value",
    "edge_location": "ashburn"
  },
  "enable_sms": true
}`
	assertRoundTrip(t, wire, new(CreatePhoneNumbersRequestBody))
}

func TestGetPhoneNumbersResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`
	assertRoundTrip(t, wire, new(GetPhoneNumbersResponse))
}

func TestUpdatePhoneNumbersResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "phone_number": "value",
  "label": "value",
  "supports_inbound": true,
  "supports_outbound": true,
  "phone_number_id": "value",
  "assigned_agent": {
    "agent_id": "value",
    "agent_name": "value",
    "environment": "value",
    "branch_id": "value"
  },
  "provider": "twilio"
}`
	assertRoundTrip(t, wire, new(UpdatePhoneNumbersResponse))
}

func TestUpdatePhoneNumberRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "label": "value",
  "inbound_trunk_config": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "outbound_trunk_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "credentials": {
      "username": "value",
      "password": "value"
    },
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "livekit_stack": "standard",
  "store_sip_messages": true,
  "environment": "value",
  "branch_id": "value"
}`
	assertRoundTrip(t, wire, new(UpdatePhoneNumberRequest))
}
