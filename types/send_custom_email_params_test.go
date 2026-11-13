package types

import (
	"testing"
)

func TestSendCustomEmailParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "send_custom_email"
}`
	assertRoundTrip(t, wire, new(SendCustomEmailParams))
}
