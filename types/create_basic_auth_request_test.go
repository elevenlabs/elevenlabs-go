package types

import (
	"testing"
)

func TestCreateBasicAuthRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "basic_auth",
  "provider": "value",
  "username": "value",
  "password": "value"
}`
	assertRoundTrip(t, wire, new(CreateBasicAuthRequest))
}
