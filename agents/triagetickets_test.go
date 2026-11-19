package agents

import (
	"testing"
)

func TestCreateManualTicketRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "qa_comment": "value",
  "priority": "low"
}`
	assertRoundTrip(t, wire, new(CreateManualTicketRequestModel))
}

func TestCreateAgentConversationTicketRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value",
  "title": "value",
  "qa_comment": "value",
  "turn_comments": [
    {
      "turn_index": 123,
      "comment": "value"
    }
  ],
  "priority": "low"
}`
	assertRoundTrip(t, wire, new(CreateAgentConversationTicketRequestModel))
}

func TestPatchAgentConversationTicketRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "qa_comment": "value",
  "status": "open",
  "assignee_user_id": "value",
  "priority": "low"
}`
	assertRoundTrip(t, wire, new(PatchAgentConversationTicketRequestModel))
}

func TestAddTicketCommentRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "comment": "value"
}`
	assertRoundTrip(t, wire, new(AddTicketCommentRequestModel))
}

func TestMergeAgentConversationTicketsRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_ticket_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(MergeAgentConversationTicketsRequestModel))
}

func TestAddTurnCommentRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "turn_index": 123,
  "comment": "value"
}`
	assertRoundTrip(t, wire, new(AddTurnCommentRequestModel))
}
