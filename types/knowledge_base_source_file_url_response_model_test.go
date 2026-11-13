package types

import (
	"testing"
)

func TestKnowledgeBaseSourceFileUrlResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "signed_url": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseSourceFileUrlResponseModel))
}
