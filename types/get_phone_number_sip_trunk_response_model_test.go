package types

import (
	"testing"
)

func TestGetPhoneNumberSipTrunkResponseModelRoundTrip(t *testing.T) {
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
  "provider": "sip_trunk",
  "provider_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "has_auth_credentials": true,
    "username": "value",
    "has_outbound_trunk": true,
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "outbound_trunk": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "has_auth_credentials": true,
    "username": "value",
    "has_outbound_trunk": true,
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "inbound_trunk": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "has_auth_credentials": true,
    "username": "value",
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "livekit_stack": "standard",
  "store_sip_messages": true
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberSipTrunkResponseModel))
}

func TestGetPhoneNumberSipTrunkResponseModelRoundTrip2(t *testing.T) {
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
  "provider": "sip_trunk",
  "provider_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "has_auth_credentials": true,
    "username": "value",
    "has_outbound_trunk": true,
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "outbound_trunk": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "has_auth_credentials": true,
    "username": "value",
    "has_outbound_trunk": true,
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "inbound_trunk": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "has_auth_credentials": true,
    "username": "value",
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "livekit_stack": "standard",
  "store_sip_messages": true
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberSipTrunkResponseModel))
}

func TestGetPhoneNumberSipTrunkResponseModelRoundTrip3(t *testing.T) {
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
  "provider": "sip_trunk",
  "provider_config": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "has_auth_credentials": true,
    "username": "value",
    "has_outbound_trunk": true,
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "outbound_trunk": {
    "address": "value",
    "transport": "auto",
    "media_encryption": "disabled",
    "headers": {
      "key": "value"
    },
    "attributes_to_headers": {
      "key": "value"
    },
    "has_auth_credentials": true,
    "username": "value",
    "has_outbound_trunk": true,
    "enabled_codecs": [
      "G722/8000"
    ]
  },
  "inbound_trunk": {
    "allowed_addresses": [
      "value"
    ],
    "allowed_numbers": [
      "value"
    ],
    "media_encryption": "disabled",
    "has_auth_credentials": true,
    "username": "value",
    "remote_domains": [
      "value"
    ],
    "attributes_to_headers": {
      "key": "value"
    }
  },
  "livekit_stack": "standard",
  "store_sip_messages": true
}`
	assertRoundTrip(t, wire, new(GetPhoneNumberSipTrunkResponseModel))
}
