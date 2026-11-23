package types

import (
	"testing"
)

func TestConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "label": "value",
  "voice_id": "value",
  "description": "value",
  "language": "value",
  "model_family": "turbo",
  "optimize_streaming_latency": 123,
  "stability": 0.5,
  "speed": 1.2,
  "similarity_boost": 0.5,
  "clone_strength": 0.5
}`
	assertRoundTrip(t, wire, new(ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem))
}
