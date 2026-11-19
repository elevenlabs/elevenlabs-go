package types

import (
	"testing"
)

func TestPlainPlatformMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_initiation_source": "plain_integration",
  "thread_id": "value",
  "workspace_id": "value"
}`
	assertRoundTrip(t, wire, new(PlainPlatformMetadata))
}
