package types

import (
	"testing"
)

func TestCreateProductParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_product"
}`
	assertRoundTrip(t, wire, new(CreateProductParams))
}
