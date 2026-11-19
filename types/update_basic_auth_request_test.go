package types

import (
	"testing"
)

func TestUpdateBasicAuthRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_type": "basic_auth",
  "provider": "value",
  "username": "value",
  "password": "value"
}`
	assertRoundTrip(t, wire, new(UpdateBasicAuthRequest))
}
