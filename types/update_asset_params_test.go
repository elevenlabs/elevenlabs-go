package types

import (
	"testing"
)

func TestUpdateAssetParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_asset",
  "per_resource_timezones": true
}`
	assertRoundTrip(t, wire, new(UpdateAssetParams))
}
