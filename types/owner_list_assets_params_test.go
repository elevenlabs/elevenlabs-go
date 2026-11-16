package types

import (
	"testing"
)

func TestOwnerListAssetsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "owner_list_assets"
}`
	assertRoundTrip(t, wire, new(OwnerListAssetsParams))
}
