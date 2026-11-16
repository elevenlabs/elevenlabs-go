package types

import (
	"testing"
)

func TestGetAgentConversationTicketsPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_conversation_tickets": [
    {
      "agentqa_ticket_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "agent_id": "value",
      "needs_clustering": true,
      "title": "value",
      "issue_type": "knowledge_gap",
      "labels": [
        "value"
      ],
      "conversation_ids": [
        "value"
      ],
      "first_seen_unix_secs": 123,
      "last_seen_unix_secs": 123,
      "qa_comment": "value",
      "ticket_comments": [
        {
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "turn_comments": [
        {
          "turn_index": 123,
          "comment": "value",
          "created_at_unix_secs": 123,
          "owner_user_id": "value"
        }
      ],
      "status": "open",
      "priority": "low",
      "priority_changes": [
        {
          "priority": "low",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "status_changes": [
        {
          "status": "open",
          "changed_by_user_id": "value",
          "changed_at_unix_secs": 123
        }
      ],
      "merged_into_ticket_id": "value",
      "source": "qa",
      "assignee_user_id": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "search_match": {
        "field": "title",
        "snippet": "value",
        "highlight_start": 123,
        "highlight_end": 123,
        "turn_index": 123
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetAgentConversationTicketsPageResponseModel))
}
