package types

import (
	"testing"
)

func TestAnalysisItemVersionsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "versions": [
    {
      "id": "value",
      "version_number": 123,
      "created_at_unix_secs": 123,
      "definition": {
        "name": "value",
        "source": "user",
        "llm": "gpt-4o-mini",
        "kind": "evaluation",
        "conversation_goal_prompt": "value",
        "use_knowledge_base": true,
        "scoring": {
          "mode": "binary"
        }
      }
    }
  ],
  "current_version_id": "value",
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(AnalysisItemVersionsResponseModel))
}
