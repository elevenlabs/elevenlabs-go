package types

import (
	"testing"
)

func TestConvAiUserSecretDbModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "encrypted_value": "value",
  "nonce": "value",
  "id": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiUserSecretDbModel))
}
