package types

import (
	"testing"
)

func TestTokenResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "token": "value",
  "conversation_id": "value"
}`
	assertRoundTrip(t, wire, new(TokenResponseModel))
}
