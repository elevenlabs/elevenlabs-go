package types

import (
	"testing"
)

func TestAnalysisItemsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "analysis_items": [
    {
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
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(AnalysisItemsResponseModel))
}
