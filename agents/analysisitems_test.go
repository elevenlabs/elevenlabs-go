package agents

import (
	"testing"
)

func TestCreateAnalysisItemsRequestBodyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "evaluation",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scoring": {
    "mode": "binary"
  }
}`
	assertRoundTrip(t, wire, new(CreateAnalysisItemsRequestBody))
}

func TestUpdateAnalysisItemAttachmentsRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "user",
  "attach": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ],
  "detach": [
    {
      "agent_id": "value",
      "branch_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(UpdateAnalysisItemAttachmentsRequestModel))
}

func TestUpdateAnalysisItemsRequestBodyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "source": "user",
  "llm": "gpt-4o-mini",
  "kind": "evaluation",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scoring": {
    "mode": "binary"
  }
}`
	assertRoundTrip(t, wire, new(UpdateAnalysisItemsRequestBody))
}
