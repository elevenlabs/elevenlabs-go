package types

import (
	"testing"
)

func TestBodyCreateServiceAccountV1ServiceAccountsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "default_sharing_groups": [
    {
      "group_id": "value",
      "permission_level": "admin"
    }
  ]
}`
	assertRoundTrip(t, wire, new(BodyCreateServiceAccountV1ServiceAccountsPost))
}
