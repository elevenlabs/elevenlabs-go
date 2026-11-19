package types

import (
	"testing"
)

func TestSpeechEngineConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "ws_url": "value",
  "request_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(SpeechEngineConfig))
}
