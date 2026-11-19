package types

import (
	"testing"
)

func TestCustomLlmRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "model_id": "value",
  "api_key": {
    "secret_id": "value"
  },
  "auth_connection": {
    "auth_connection_id": "value"
  },
  "request_headers": {
    "key": "value"
  },
  "api_version": "value",
  "api_type": "chat_completions"
}`
	assertRoundTrip(t, wire, new(CustomLlm))
}

func TestCustomLlmRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "url": "value",
  "model_id": "value",
  "api_key": {
    "secret_id": "value"
  },
  "auth_connection": {
    "auth_connection_id": "value"
  },
  "request_headers": {
    "key": "value"
  },
  "api_version": "value",
  "api_type": "chat_completions"
}`
	assertRoundTrip(t, wire, new(CustomLlm))
}
