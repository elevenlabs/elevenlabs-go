package types

import (
	"testing"
)

func TestApiIntegrationWebhookOverridesSchemaOverridesValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "constant",
  "constant_value": "value"
}`
	assertRoundTrip(t, wire, new(ApiIntegrationWebhookOverridesSchemaOverridesValue))
}
