package types

import (
	"testing"
)

func TestAudioNativeEditContentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_id": "value",
  "converting": true,
  "publishing": true,
  "html_snippet": "value"
}`
	assertRoundTrip(t, wire, new(AudioNativeEditContentResponseModel))
}
