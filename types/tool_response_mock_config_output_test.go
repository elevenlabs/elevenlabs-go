package types

import (
	"testing"
)

func TestToolResponseMockConfigOutputRoundTrip(t *testing.T) {
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
	assertRoundTrip(t, wire, new(ToolResponseMockConfigOutput))
}

func TestToolResponseMockConfigOutputRoundTrip2(t *testing.T) {
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
	assertRoundTrip(t, wire, new(ToolResponseMockConfigOutput))
}

func TestToolResponseMockConfigOutputRoundTrip3(t *testing.T) {
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
	assertRoundTrip(t, wire, new(ToolResponseMockConfigOutput))
}

func TestToolResponseMockConfigOutputRoundTrip4(t *testing.T) {
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
	assertRoundTrip(t, wire, new(ToolResponseMockConfigOutput))
}

func TestToolResponseMockConfigOutputRoundTrip5(t *testing.T) {
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
	assertRoundTrip(t, wire, new(ToolResponseMockConfigOutput))
}
