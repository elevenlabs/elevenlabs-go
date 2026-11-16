package types

import (
	"testing"
)

func TestPendingExternalAudiosMetadataModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "target_global_offset_ms": 123,
  "external_audio_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(PendingExternalAudiosMetadataModel))
}
