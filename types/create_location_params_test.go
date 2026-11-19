package types

import (
	"testing"
)

func TestCreateLocationParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "create_location"
}`
	assertRoundTrip(t, wire, new(CreateLocationParams))
}
