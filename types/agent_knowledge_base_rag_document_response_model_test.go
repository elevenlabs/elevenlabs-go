package types

import (
	"testing"
)

func TestAgentKnowledgeBaseRagDocumentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_id": "value",
  "document_name": "value",
  "source_url": "value",
  "content": "value",
  "truncated": true,
  "content_format": "html",
  "document_type": "file"
}`
	assertRoundTrip(t, wire, new(AgentKnowledgeBaseRagDocumentResponseModel))
}
