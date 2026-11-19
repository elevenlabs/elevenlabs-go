package types

import (
	"testing"
)

func TestKnowledgeBaseFolderPathSegmentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseFolderPathSegmentResponseModel))
}
