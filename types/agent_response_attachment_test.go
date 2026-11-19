package types

import (
	"testing"
)

func TestAgentResponseAttachmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "name": "value",
  "mime_type": "value"
}`
	assertRoundTrip(t, wire, new(AgentResponseAttachment))
}
