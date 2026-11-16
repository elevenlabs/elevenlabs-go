package types

import (
	"testing"
)

func TestApiIntegrationWebhookOverridesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "schema_overrides": {
    "key": {
      "source": "constant",
      "constant_value": "value"
    }
  },
  "response_filter_mode": "all",
  "response_filters": [
    "value"
  ],
  "request_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ApiIntegrationWebhookOverrides))
}
