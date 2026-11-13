package types

import (
	"testing"
)

func TestKnowledgeBaseDeletedResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseDeletedResponseModel))
}
