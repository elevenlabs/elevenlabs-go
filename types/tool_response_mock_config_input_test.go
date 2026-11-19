package types

import (
	"testing"
)

func TestToolResponseMockConfigInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ToolResponseMockConfigInput))
}

func TestToolResponseMockConfigInputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ToolResponseMockConfigInput))
}

func TestToolResponseMockConfigInputRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ToolResponseMockConfigInput))
}

func TestToolResponseMockConfigInputRoundTrip4(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ToolResponseMockConfigInput))
}
