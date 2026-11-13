package types

import (
	"testing"
)

func TestTestRunMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "workspace_id": "value",
  "test_name": "value",
  "ran_by_user_email": "value",
  "test_type": "llm"
}`
	assertRoundTrip(t, wire, new(TestRunMetadata))
}
