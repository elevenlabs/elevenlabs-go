package types

import (
	"testing"
)

func TestUpdateVoicesRequestLabelsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "key": "value"
}`
	assertRoundTrip(t, wire, new(UpdateVoicesRequestLabels))
}

func TestUpdateVoicesRequestLabelsRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "key": "value"
}`
	assertRoundTrip(t, wire, new(UpdateVoicesRequestLabels))
}
