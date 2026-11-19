package types

import (
	"testing"
)

func TestCustomLlmApiKeyRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(CustomLlmApiKey))
}
