package types

import (
	"testing"
)

func TestSystemToolConfigOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "system",
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
  "params": {
    "system_tool_type": "end_call"
  }
}`
	assertRoundTrip(t, wire, new(SystemToolConfigOutput))
}

func TestSystemToolConfigOutputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "system",
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
  "params": {
    "system_tool_type": "end_call"
  }
}`
	assertRoundTrip(t, wire, new(SystemToolConfigOutput))
}

func TestSystemToolConfigOutputRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "system",
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
  "params": {
    "system_tool_type": "end_call"
  }
}`
	assertRoundTrip(t, wire, new(SystemToolConfigOutput))
}
