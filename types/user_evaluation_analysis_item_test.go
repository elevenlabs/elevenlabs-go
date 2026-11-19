package types

import (
	"testing"
)

func TestUserEvaluationAnalysisItemRoundTrip(t *testing.T) {
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
	assertRoundTrip(t, wire, new(UserEvaluationAnalysisItem))
}
