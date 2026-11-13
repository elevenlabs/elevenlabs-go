package types

import (
	"testing"
)

func TestBuiltInToolsOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "transfer_to_agent": {
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
  },
  "end_call": {
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
  },
  "language_detection": {
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
  },
  "transfer_to_number": {
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
  },
  "skip_turn": {
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
  },
  "play_keypad_touch_tone": {
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
  },
  "voicemail_detection": {
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
  }
}`
	assertRoundTrip(t, wire, new(BuiltInToolsOutput))
}
