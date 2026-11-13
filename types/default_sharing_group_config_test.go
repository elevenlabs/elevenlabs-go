package types

import (
	"testing"
)

func TestDefaultSharingGroupConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "group_id": "value",
  "permission_level": "admin"
}`
	assertRoundTrip(t, wire, new(DefaultSharingGroupConfig))
}
