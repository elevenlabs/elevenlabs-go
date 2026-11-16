package types

import (
	"testing"
)

func TestSmbToolConfigParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "search_clients"
}`
	assertRoundTrip(t, wire, new(SmbToolConfigParams))
}
