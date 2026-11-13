package types

import (
	"testing"
)

func TestAgentAnalysisItemsInputDataCollectionItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "user",
  "analysis_item_id": "value",
  "version_id": "value",
  "scope": "conversation"
}`
	assertRoundTrip(t, wire, new(AgentAnalysisItemsInputDataCollectionItem))
}
