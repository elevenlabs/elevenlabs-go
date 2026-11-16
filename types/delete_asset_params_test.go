package types

import (
	"testing"
)

func TestDeleteAssetParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_asset"
}`
	assertRoundTrip(t, wire, new(DeleteAssetParams))
}
