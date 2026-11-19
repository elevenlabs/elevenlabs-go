package types

import (
	"testing"
)

func TestGetOrCreateRagIndexRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_id": "value",
  "create_if_missing": true,
  "model": "e5_mistral_7b_instruct"
}`
	assertRoundTrip(t, wire, new(GetOrCreateRagIndexRequestModel))
}
