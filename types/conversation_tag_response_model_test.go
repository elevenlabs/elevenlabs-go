package types

import (
	"testing"
)

func TestConversationTagResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tag_id": "value",
  "workspace_id": "value",
  "owner_user_id": "value",
  "title": "value",
  "description": "value",
  "created_at_unix_secs": 123
}`
	assertRoundTrip(t, wire, new(ConversationTagResponseModel))
}
