package conversations

import (
	"testing"
)

func TestAssignConversationTagsRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tag_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(AssignConversationTagsRequestModel))
}

func TestCreateConversationTagRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(CreateConversationTagRequestModel))
}

func TestPatchConversationTagRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(PatchConversationTagRequestModel))
}
