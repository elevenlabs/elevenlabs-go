package types

import (
	"testing"
)

func TestUserActivityRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "user_activity"
}`
	assertRoundTrip(t, wire, new(UserActivity))
}
