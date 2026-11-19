package conversations

import (
	"testing"
)

func TestRunConversationEvaluationsRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "evaluation_id": "value",
  "scope": "conversation",
  "evaluated_agent_id": "value",
  "evaluated_branch_id": "value"
}`
	assertRoundTrip(t, wire, new(RunConversationEvaluationsRequest))
}
