package types

import (
	"testing"
)

func TestCreateAgentTestFolderResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(CreateAgentTestFolderResponseModel))
}
