package types

import (
	"testing"
)

func TestConfirmChangeParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "smb_confirm_change"
}`
	assertRoundTrip(t, wire, new(ConfirmChangeParams))
}
