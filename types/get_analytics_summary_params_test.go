package types

import (
	"testing"
)

func TestGetAnalyticsSummaryParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "get_analytics_summary"
}`
	assertRoundTrip(t, wire, new(GetAnalyticsSummaryParams))
}
