package types

import (
	"testing"
)

func TestCreateMtlsAuthRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "mtls",
  "provider": "value",
  "client_certificate": "value",
  "client_key": "value",
  "ca_certificate": "value",
  "key_passphrase": "value"
}`
	assertRoundTrip(t, wire, new(CreateMtlsAuthRequest))
}
