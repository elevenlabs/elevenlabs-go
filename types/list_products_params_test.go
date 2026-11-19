package types

import (
	"testing"
)

func TestListProductsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "list_kwargs": {
    "key": {}
  },
  "smb_tool_type": "list_products"
}`
	assertRoundTrip(t, wire, new(ListProductsParams))
}
