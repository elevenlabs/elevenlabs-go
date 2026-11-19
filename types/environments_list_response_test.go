package types

import (
	"testing"
)

func TestEnvironmentsListResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "environments": [
    {
      "name": "value",
      "updated_at_unix_secs": 123,
      "created_by_user_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(EnvironmentsListResponse))
}
