package types

import (
	"testing"
)

func TestMcpToolConfigOverrideInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_name": "value",
  "force_pre_tool_speech": true,
  "pre_tool_speech": "auto",
  "disable_interruptions": true,
  "interruption_mode": "allow",
  "tool_call_sound": "typing",
  "tool_call_sound_behavior": "auto",
  "execution_mode": "immediate",
  "response_timeout_secs": 123,
  "assignments": [
    {
      "source": "response",
      "dynamic_variable": "value",
      "value_path": "value",
      "sanitize": true,
      "preserve_native_type": true
    }
  ],
  "input_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_mocks": [
    {
      "parameter_conditions": [
        {
          "eval": {
            "type": "llm",
            "description": "value"
          },
          "path": "value"
        }
      ],
      "mock_result": "value",
      "is_error": true
    }
  ]
}`
	assertRoundTrip(t, wire, new(McpToolConfigOverrideInput))
}
