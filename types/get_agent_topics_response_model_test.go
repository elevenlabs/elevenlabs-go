package types

import (
	"testing"
)

func TestGetAgentTopicsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "topics": [
    {
      "topic_id": "value",
      "label": "value",
      "description": "value",
      "conversation_count": 123,
      "parent_topic_id": "value",
      "x_2d": 9.99,
      "y_2d": 9.99,
      "metrics": {
        "conversation_count": 123,
        "sentiment": {
          "sentiment": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "frustration": {
            "count": 123,
            "sum": 9.99,
            "min": 9.99,
            "max": 9.99
          },
          "positive_count": 123,
          "neutral_count": 123,
          "negative_count": 123
        },
        "evaluation_criteria": [
          {
            "criteria_id": "value",
            "success_count": 123,
            "failure_count": 123,
            "unknown_count": 123
          }
        ]
      },
      "success_rate": 9.99
    }
  ],
  "window_start_unix_secs": 123,
  "window_end_unix_secs": 123,
  "aggregated_run_count": 123,
  "has_more": true,
  "next_cursor": "value"
}`
	assertRoundTrip(t, wire, new(GetAgentTopicsResponseModel))
}
