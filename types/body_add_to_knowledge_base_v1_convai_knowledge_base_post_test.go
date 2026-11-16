package types

import (
	"testing"
)

func TestBodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "url": "value"
}`
	assertRoundTrip(t, wire, new(BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost))
}
