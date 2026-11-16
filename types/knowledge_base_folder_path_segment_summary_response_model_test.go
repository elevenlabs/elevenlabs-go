package types

import (
	"testing"
)

func TestKnowledgeBaseFolderPathSegmentSummaryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseFolderPathSegmentSummaryResponseModel))
}

func TestKnowledgeBaseFolderPathSegmentSummaryResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseFolderPathSegmentSummaryResponseModel))
}

func TestKnowledgeBaseFolderPathSegmentSummaryResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseFolderPathSegmentSummaryResponseModel))
}
