package types

import (
	"testing"
)

func TestGetConversationUsersPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "users": [
    {
      "user_id": "value",
      "last_contact_unix_secs": 123,
      "first_contact_unix_secs": 123,
      "conversation_count": 123,
      "last_contact_agent_id": "value",
      "last_contact_conversation_id": "value",
      "last_contact_agent_name": "value",
      "sentiment": {
        "scored_conversation_count": 123,
        "positive_count": 123,
        "neutral_count": 123,
        "negative_count": 123,
        "average_sentiment_score": 9.99,
        "average_frustration_score": 9.99,
        "recent_scored_conversation_count": 123,
        "recent_positive_count": 123,
        "recent_neutral_count": 123,
        "recent_negative_count": 123,
        "recent_average_sentiment_score": 9.99,
        "recent_average_frustration_score": 9.99
      },
      "most_frustrated_conversations": [
        {
          "conversation_id": "value",
          "agent_id": "value",
          "start_time_unix_secs": 123,
          "overall_label": "positive",
          "overall_sentiment_score": 9.99,
          "overall_frustration_score": 9.99
        }
      ]
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetConversationUsersPageResponseModel))
}
