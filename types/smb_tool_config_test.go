package types

import (
	"testing"
)

func TestSmbToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "smb",
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
  "human_description": "value",
  "enabled": true,
  "params": {
    "smb_tool_type": "search_clients"
  }
}`
	assertRoundTrip(t, wire, new(SmbToolConfig))
}

func TestSmbToolConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "smb",
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
  "human_description": "value",
  "enabled": true,
  "params": {
    "smb_tool_type": "search_clients"
  }
}`
	assertRoundTrip(t, wire, new(SmbToolConfig))
}
