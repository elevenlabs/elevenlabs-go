package types

import (
	"testing"
)

func TestReceiveUpstreamMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "init",
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(ReceiveUpstreamMessage))
}
