package types

import (
	"testing"
)

func TestScimGroupResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "scim_external_id": "value",
  "display_name": "value",
  "created_at_unix": 123,
  "updated_at_unix": 123,
  "seat_type": "workspace_admin"
}`
	assertRoundTrip(t, wire, new(ScimGroupResponseModel))
}
