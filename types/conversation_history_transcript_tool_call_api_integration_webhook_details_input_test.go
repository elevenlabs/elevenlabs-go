package types

import (
	"testing"
)

func TestConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "api_integration_webhook",
  "integration_id": "value",
  "credential_id": "value",
  "integration_connection_id": "value",
  "webhook_details": {
    "type": "webhook",
    "method": "value",
    "url": "value",
    "headers": {
      "key": "value"
    },
    "path_params": {
      "key": "value"
    },
    "query_params": {
      "key": "value"
    },
    "body": "value"
  },
  "function_details": {
    "resolved_parameters": {
      "key": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallApiIntegrationWebhookDetailsInput))
}
