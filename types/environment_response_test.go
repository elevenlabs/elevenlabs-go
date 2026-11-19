package types

import (
	"testing"
)

func TestEnvironmentResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "updated_at_unix_secs": 123,
  "created_by_user_id": "value"
}`
	assertRoundTrip(t, wire, new(EnvironmentResponse))
}
