package types

import (
	"testing"
)

func TestCreateAssetParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_asset",
  "per_resource_timezones": true
}`
	assertRoundTrip(t, wire, new(CreateAssetParams))
}
