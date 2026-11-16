package types

import (
	"testing"
)

func TestAnalysisItemResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
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
  },
  "access_info": {
    "is_creator": true,
    "creator_name": "value",
    "creator_email": "value",
    "role": "admin",
    "anonymous_access_level_override": "admin",
    "access_source": "creator"
  },
  "created_at_unix_secs": 123,
  "updated_at_unix_secs": 123,
  "current_version_id": "value",
  "legacy_source_id": "value"
}`
	assertRoundTrip(t, wire, new(AnalysisItemResponseModel))
}
