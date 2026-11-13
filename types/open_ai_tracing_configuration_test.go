package types

import (
	"testing"
)

func TestOpenAiTracingConfigurationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "group_id": "value",
  "workflow_name": "value",
  "metadata": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(OpenAiTracingConfiguration))
}
