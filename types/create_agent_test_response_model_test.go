package types

import (
	"testing"
)

func TestCreateAgentTestResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value"
}`
	assertRoundTrip(t, wire, new(CreateAgentTestResponseModel))
}
