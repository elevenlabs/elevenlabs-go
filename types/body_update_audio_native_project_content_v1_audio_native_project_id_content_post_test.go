package types

import (
	"testing"
)

func TestBodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auto_convert": true,
  "auto_publish": true
}`
	assertRoundTrip(t, wire, new(BodyUpdateAudioNativeProjectContentV1AudioNativeProjectIdContentPost))
}
