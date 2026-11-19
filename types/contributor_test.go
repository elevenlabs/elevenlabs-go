package types

import (
	"testing"
)

func TestContributorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "role": "value",
  "bio": "value",
  "profile_id": "value"
}`
	assertRoundTrip(t, wire, new(Contributor))
}
