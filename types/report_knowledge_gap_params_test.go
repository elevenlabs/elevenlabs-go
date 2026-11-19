package types

import (
	"testing"
)

func TestReportKnowledgeGapParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "smb_tool_type": "report_knowledge_gap"
}`
	assertRoundTrip(t, wire, new(ReportKnowledgeGapParams))
}
