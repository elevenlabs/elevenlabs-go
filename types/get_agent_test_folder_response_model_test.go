package types

import (
	"testing"
)

func TestGetAgentTestFolderResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "folder_path": [
    {
      "id": "value",
      "name": "value"
    }
  ],
  "children_count": 123
}`
	assertRoundTrip(t, wire, new(GetAgentTestFolderResponseModel))
}
