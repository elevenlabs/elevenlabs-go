package types

import (
	"testing"
)

func TestCreateCustomHeaderAuthRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "custom_header_auth",
  "provider": "value",
  "header_name": "value",
  "token": "value"
}`
	assertRoundTrip(t, wire, new(CreateCustomHeaderAuthRequest))
}
