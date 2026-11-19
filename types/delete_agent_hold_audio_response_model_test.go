package types

import (
	"testing"
)

func TestDeleteAgentHoldAudioResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value"
}`
	assertRoundTrip(t, wire, new(DeleteAgentHoldAudioResponseModel))
}
