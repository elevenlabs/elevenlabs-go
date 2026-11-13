package types

import (
	"testing"
)

func TestConversationSignedUrlResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "signed_url": "value"
}`
	assertRoundTrip(t, wire, new(ConversationSignedUrlResponseModel))
}
