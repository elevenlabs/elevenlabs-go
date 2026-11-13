package types

import (
	"testing"
)

func TestZendeskPlatformMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_initiation_source": "zendesk_integration",
  "ticket_id": "value"
}`
	assertRoundTrip(t, wire, new(ZendeskPlatformMetadata))
}
