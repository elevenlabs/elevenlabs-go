package types

import (
	"testing"
)

func TestGetAgentKnowledgebaseSizeResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "number_of_pages": 9.99
}`
	assertRoundTrip(t, wire, new(GetAgentKnowledgebaseSizeResponseModel))
}
