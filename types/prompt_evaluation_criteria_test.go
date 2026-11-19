package types

import (
	"testing"
)

func TestPromptEvaluationCriteriaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "type": "prompt",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scope": "conversation",
  "llm": "gpt-4o-mini",
  "scoring_mode": "binary",
  "max_score": 123,
  "score_instructions": "value"
}`
	assertRoundTrip(t, wire, new(PromptEvaluationCriteria))
}

func TestPromptEvaluationCriteriaRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "type": "prompt",
  "conversation_goal_prompt": "value",
  "use_knowledge_base": true,
  "scope": "conversation",
  "llm": "gpt-4o-mini",
  "scoring_mode": "binary",
  "max_score": 123,
  "score_instructions": "value"
}`
	assertRoundTrip(t, wire, new(PromptEvaluationCriteria))
}
