package types

import (
	"testing"
)

func TestGetToolDependentAgentsResponseModelAgentsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "referenced_resource_ids": [
    "value"
  ],
  "id": "value",
  "name": "value",
  "type": "available",
  "created_at_unix_secs": 123,
  "access_level": "admin"
}`
	assertRoundTrip(t, wire, new(GetToolDependentAgentsResponseModelAgentsItem))
}
