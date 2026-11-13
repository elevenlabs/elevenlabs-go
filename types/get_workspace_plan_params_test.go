package types

import (
	"testing"
)

func TestGetWorkspacePlanParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_workspace_plan"
}`
	assertRoundTrip(t, wire, new(GetWorkspacePlanParams))
}
