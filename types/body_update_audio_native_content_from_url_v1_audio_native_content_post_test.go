package types

import (
	"testing"
)

func TestBodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "author": "value",
  "title": "value"
}`
	assertRoundTrip(t, wire, new(BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost))
}
