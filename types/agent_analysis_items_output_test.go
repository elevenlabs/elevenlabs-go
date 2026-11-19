package types

import (
	"testing"
)

func TestAgentAnalysisItemsOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "evaluation_criteria": [
    {
      "source": "user",
      "analysis_item_id": "value",
      "version_id": "value",
      "additional_version_ids": [
        "value"
      ],
      "scope": "conversation",
      "weight": 9.99
    }
  ],
  "data_collection": [
    {
      "source": "user",
      "analysis_item_id": "value",
      "version_id": "value",
      "scope": "conversation"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AgentAnalysisItemsOutput))
}
