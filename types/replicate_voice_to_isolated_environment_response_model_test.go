package types

import (
	"testing"
)

func TestReplicateVoiceToIsolatedEnvironmentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value"
}`
	assertRoundTrip(t, wire, new(ReplicateVoiceToIsolatedEnvironmentResponseModel))
}
