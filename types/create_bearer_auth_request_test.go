package types

import (
	"testing"
)

func TestCreateBearerAuthRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "bearer_auth",
  "provider": "value",
  "token": "value"
}`
	assertRoundTrip(t, wire, new(CreateBearerAuthRequest))
}
