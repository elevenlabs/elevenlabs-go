package types

import (
	"testing"
)

func TestUpdateServiceParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_service"
}`
	assertRoundTrip(t, wire, new(UpdateServiceParams))
}
