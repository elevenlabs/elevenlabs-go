package types

import (
	"testing"
)

func TestGetConversationTagsPageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conversation_tags": [
    {
      "tag_id": "value",
      "workspace_id": "value",
      "owner_user_id": "value",
      "title": "value",
      "description": "value",
      "created_at_unix_secs": 123
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(GetConversationTagsPageResponseModel))
}
