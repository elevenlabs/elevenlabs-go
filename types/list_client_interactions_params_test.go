package types

import (
	"testing"
)

func TestListClientInteractionsParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "list_client_interactions"
}`
	assertRoundTrip(t, wire, new(ListClientInteractionsParams))
}
