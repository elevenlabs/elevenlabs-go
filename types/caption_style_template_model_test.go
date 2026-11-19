package types

import (
	"testing"
)

func TestCaptionStyleTemplateModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "key": "value",
  "label": "value",
  "requires_high_fps": true
}`
	assertRoundTrip(t, wire, new(CaptionStyleTemplateModel))
}
