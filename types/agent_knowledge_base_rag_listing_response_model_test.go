package types

import (
	"testing"
)

func TestAgentKnowledgeBaseRagListingResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "entries": [
    {
      "id": "value",
      "name": "value",
      "is_folder": true
    }
  ],
  "total_count": 123
}`
	assertRoundTrip(t, wire, new(AgentKnowledgeBaseRagListingResponseModel))
}
