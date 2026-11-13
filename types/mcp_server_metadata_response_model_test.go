package types

import (
	"testing"
)

func TestMcpServerMetadataResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "created_at": 123,
  "owner_user_id": "value"
}`
	assertRoundTrip(t, wire, new(McpServerMetadataResponseModel))
}
