package types

import (
	"testing"
)

func TestAudioNativeCreateProjectResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_id": "value",
  "converting": true,
  "html_snippet": "value"
}`
	assertRoundTrip(t, wire, new(AudioNativeCreateProjectResponseModel))
}
