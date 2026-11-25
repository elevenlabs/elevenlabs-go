package types

import (
	"testing"
)

func TestUserMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "type": "user_message",
  "text": "value",
  "user_identifier": "value",
  "source_medium": "audio"
}`
	assertRoundTrip(t, wire, new(UserMessage))
}
