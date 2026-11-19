package types

import (
	"testing"
)

func TestApiIntegrationWebhookToolConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "api_integration_webhook",
  "name": "value",
  "description": "value",
  "response_timeout_secs": 123,
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "tool_error_handling_mode": "auto",
  "dynamic_variables": {
    "dynamic_variable_placeholders": {
      "key": {}
    }
  },
  "execution_mode": "immediate",
  "tool_version": "value",
  "api_integration_id": "value",
  "api_integration_connection_id": "value",
  "api_schema_overrides": {
    "schema_overrides": {
      "key": {
        "source": "constant",
        "constant_value": "value"
      }
    },
    "response_filter_mode": "all",
    "response_filters": [
      "value"
    ],
    "request_headers": {
      "key": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(ApiIntegrationWebhookToolConfigOutput))
}

func TestApiIntegrationWebhookToolConfigOutputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "api_integration_webhook",
  "name": "value",
  "description": "value",
  "response_timeout_secs": 123,
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "tool_error_handling_mode": "auto",
  "dynamic_variables": {
    "dynamic_variable_placeholders": {
      "key": {}
    }
  },
  "execution_mode": "immediate",
  "tool_version": "value",
  "api_integration_id": "value",
  "api_integration_connection_id": "value",
  "api_schema_overrides": {
    "schema_overrides": {
      "key": {
        "source": "constant",
        "constant_value": "value"
      }
    },
    "response_filter_mode": "all",
    "response_filters": [
      "value"
    ],
    "request_headers": {
      "key": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(ApiIntegrationWebhookToolConfigOutput))
}
