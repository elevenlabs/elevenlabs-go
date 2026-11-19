package types

import (
	"testing"
)

func TestListAssetsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "list_assets"
}`
	assertRoundTrip(t, wire, new(ListAssetsParams))
}
