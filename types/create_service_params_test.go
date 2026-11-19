package types

import (
	"testing"
)

func TestCreateServiceParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_service"
}`
	assertRoundTrip(t, wire, new(CreateServiceParams))
}
