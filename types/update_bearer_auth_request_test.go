package types

import (
	"testing"
)

func TestUpdateBearerAuthRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_type": "bearer_auth",
  "provider": "value",
  "token": "value"
}`
	assertRoundTrip(t, wire, new(UpdateBearerAuthRequest))
}
