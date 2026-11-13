package types

import (
	"testing"
)

func TestReplicateVoiceToIsolatedEnvironmentRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "target_workspace_id": "value",
  "preserve_voice_id": true
}`
	assertRoundTrip(t, wire, new(ReplicateVoiceToIsolatedEnvironmentRequestModel))
}
