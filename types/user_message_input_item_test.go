package types

import (
	"testing"
)

func TestUserMessageInputItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "user_message",
  "content": [
    {
      "type": "text",
      "text": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(UserMessageInputItem))
}
