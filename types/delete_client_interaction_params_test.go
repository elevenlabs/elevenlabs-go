package types

import (
	"testing"
)

func TestDeleteClientInteractionParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "delete_client_interaction"
}`
	assertRoundTrip(t, wire, new(DeleteClientInteractionParams))
}
