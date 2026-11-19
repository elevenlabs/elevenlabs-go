package types

import (
	"testing"
)

func TestOwnerListProductsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "owner_list_products"
}`
	assertRoundTrip(t, wire, new(OwnerListProductsParams))
}
