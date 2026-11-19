package types

import (
	"testing"
)

func TestUpdateProductParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "update_product"
}`
	assertRoundTrip(t, wire, new(UpdateProductParams))
}
