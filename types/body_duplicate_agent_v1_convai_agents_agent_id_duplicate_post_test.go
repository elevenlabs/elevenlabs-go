package types

import (
	"testing"
)

func TestBodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(BodyDuplicateAgentV1ConvaiAgentsAgentIdDuplicatePost))
}
