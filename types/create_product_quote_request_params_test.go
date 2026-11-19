package types

import (
	"testing"
)

func TestCreateProductQuoteRequestParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_product_quote_request"
}`
	assertRoundTrip(t, wire, new(CreateProductQuoteRequestParams))
}
