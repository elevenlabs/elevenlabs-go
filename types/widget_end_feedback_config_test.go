package types

import (
	"testing"
)

func TestWidgetEndFeedbackConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "rating"
}`
	assertRoundTrip(t, wire, new(WidgetEndFeedbackConfig))
}

func TestWidgetEndFeedbackConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "rating"
}`
	assertRoundTrip(t, wire, new(WidgetEndFeedbackConfig))
}
