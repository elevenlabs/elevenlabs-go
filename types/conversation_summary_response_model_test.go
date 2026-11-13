package types

import (
	"testing"
)

func TestConversationSummaryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "branch_id": "value",
  "version_id": "value",
  "agent_name": "value",
  "conversation_id": "value",
  "start_time_unix_secs": 123,
  "call_duration_secs": 123,
  "cost": 123,
  "cost_fiat": 9.99,
  "message_count": 123,
  "user_message_count": 123,
  "agent_message_count": 123,
  "status": "initiated",
  "termination_reason": "value",
  "call_successful": "success",
  "call_success_score": 9.99,
  "transcript_summary": "value",
  "call_summary_title": "value",
  "main_language": "value",
  "conversation_initiation_source": "unknown",
  "tool_names": [
    "value"
  ],
  "direction": "inbound",
  "rating": 9.99,
  "sentiment_analysis": {
    "overall_label": "positive",
    "overall_sentiment_score": 0.5,
    "overall_frustration_score": 0.5,
    "min_user_sentiment_score": 0.5,
    "max_user_frustration_score": 0.5,
    "num_scored_user_turns": 123
  },
  "data_collection_results": {
    "key": {}
  },
  "evaluation_criteria_results": {
    "key": {
      "result": "success",
      "score": 123,
      "max_score": 123
    }
  },
  "tag_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ConversationSummaryResponseModel))
}
