package types

import (
	"testing"
)

func TestKnowledgeBaseBulkDeleteSuccessfulResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "success",
  "data": {
    "id": "value"
  }
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseBulkDeleteSuccessfulResponseModel))
}
