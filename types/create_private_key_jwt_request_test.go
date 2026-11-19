package types

import (
	"testing"
)

func TestCreatePrivateKeyJwtRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "private_key_jwt",
  "provider": "value",
  "algorithm": "HS256",
  "key_id": "value",
  "issuer": "value",
  "audience": "value",
  "subject": "value",
  "expiration_seconds": 123,
  "extra_params": {
    "key": "value"
  },
  "extra_headers": {
    "key": "value"
  },
  "secret_key": "value"
}`
	assertRoundTrip(t, wire, new(CreatePrivateKeyJwtRequest))
}
