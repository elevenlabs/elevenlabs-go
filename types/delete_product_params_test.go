package types

import (
	"testing"
)

func TestDeleteProductParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_product"
}`
	assertRoundTrip(t, wire, new(DeleteProductParams))
}
