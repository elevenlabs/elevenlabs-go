package types

import (
	"testing"
)

func TestCreateServiceQuoteRequestParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_service_quote_request"
}`
	assertRoundTrip(t, wire, new(CreateServiceQuoteRequestParams))
}
