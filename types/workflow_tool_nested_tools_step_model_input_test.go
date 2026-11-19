package types

import (
	"testing"
)

func TestWorkflowToolNestedToolsStepModelInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "step_latency_secs": 9.99,
  "type": "nested_tools",
  "node_id": "value",
  "requests": [
    {
      "type": "system",
      "request_id": "value",
      "tool_name": "value",
      "tool_id": "value",
      "params_as_json": "value",
      "tool_has_been_called": true,
      "tool_details": {
        "type": "webhook",
        "method": "value",
        "url": "value",
        "headers": {
          "key": "value"
        },
        "path_params": {
          "key": "value"
        },
        "query_params": {
          "key": "value"
        },
        "body": "value"
      }
    }
  ],
  "results": [
    {
      "request_id": "value",
      "tool_name": "value",
      "tool_id": "value",
      "result_value": "value",
      "is_error": true,
      "is_blocked": true,
      "tool_has_been_called": true,
      "tool_latency_secs": 9.99,
      "error_type": "value",
      "raw_error_message": "value",
      "dynamic_variable_updates": [
        {
          "variable_name": "value",
          "old_value": "value",
          "new_value": "value",
          "updated_at": 9.99,
          "tool_name": "value",
          "tool_request_id": "value"
        }
      ],
      "type": "client"
    }
  ],
  "is_successful": true
}`
	assertRoundTrip(t, wire, new(WorkflowToolNestedToolsStepModelInput))
}
