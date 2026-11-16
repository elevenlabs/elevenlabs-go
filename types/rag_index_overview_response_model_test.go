package types

import (
	"testing"
)

func TestRagIndexOverviewResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total_used_bytes": 123,
  "total_max_bytes": 123,
  "models": [
    {
      "model": "e5_mistral_7b_instruct",
      "used_bytes": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(RagIndexOverviewResponseModel))
}
