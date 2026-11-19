package dashboard

import (
	"testing"
)

func TestPatchConvAiDashboardSettingsRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "charts": [
    {
      "name": "value",
      "type": "call_success"
    }
  ]
}`
	assertRoundTrip(t, wire, new(PatchConvAiDashboardSettingsRequest))
}

func TestPatchConvAiDashboardSettingsRequestChartsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "type": "call_success"
}`
	assertRoundTrip(t, wire, new(PatchConvAiDashboardSettingsRequestChartsItem))
}
