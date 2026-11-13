package types

import (
	"testing"
)

func TestOpenAimcpToolRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp",
  "server_label": "value",
  "server_url": "value",
  "authorization": "value",
  "allowed_tools": [
    "value"
  ],
  "require_approval": "never"
}`
	assertRoundTrip(t, wire, new(OpenAimcpTool))
}
