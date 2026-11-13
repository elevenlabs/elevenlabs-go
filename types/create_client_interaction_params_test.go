package types

import (
	"testing"
)

func TestCreateClientInteractionParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_client_interaction"
}`
	assertRoundTrip(t, wire, new(CreateClientInteractionParams))
}
