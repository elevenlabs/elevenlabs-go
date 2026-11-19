package types

import (
	"testing"
)

func TestCreateAgentConversationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(CreateAgentConversationResponse))
}
