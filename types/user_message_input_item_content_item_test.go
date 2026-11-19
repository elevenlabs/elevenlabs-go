package types

import (
	"testing"
)

func TestUserMessageInputItemContentItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "text",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(UserMessageInputItemContentItem))
}
