package types

import (
	"testing"
)

func TestListSimilarVoicesRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "similarity_threshold": 9.99,
  "top_k": 123
}`
	assertRoundTrip(t, wire, new(ListSimilarVoicesRequest))
}
