package types

import (
	"testing"
)

func TestSubmitBusinessInfoParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "submit_business_info"
}`
	assertRoundTrip(t, wire, new(SubmitBusinessInfoParams))
}
