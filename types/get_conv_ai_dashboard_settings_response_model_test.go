package types

import (
	"testing"
)

func TestGetConvAiDashboardSettingsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "charts": [
    {
      "name": "value",
      "type": "call_success"
    }
  ]
}`
	assertRoundTrip(t, wire, new(GetConvAiDashboardSettingsResponseModel))
}
