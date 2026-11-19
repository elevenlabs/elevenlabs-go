package types

import (
	"testing"
)

func TestAgentKnowledgeBaseRagListingEntryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "is_folder": true
}`
	assertRoundTrip(t, wire, new(AgentKnowledgeBaseRagListingEntryResponseModel))
}
