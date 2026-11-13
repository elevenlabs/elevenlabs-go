package types

import (
	"testing"
)

func TestAddKnowledgeBaseResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AddKnowledgeBaseResponseModel))
}
