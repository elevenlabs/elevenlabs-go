package types

import (
	"testing"
)

func TestRagIndexOverviewEmbeddingModelResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model": "e5_mistral_7b_instruct",
  "used_bytes": 123
}`
	assertRoundTrip(t, wire, new(RagIndexOverviewEmbeddingModelResponseModel))
}
