package types

import (
	"testing"
)

func TestAsyncConversationMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "delivery_status": "pending",
  "delivery_timestamp": 123,
  "delivery_error": "value",
  "external_system": "value",
  "external_id": "value",
  "external_link": "value",
  "retry_count": 123,
  "last_retry_timestamp": 123,
  "last_processed_external_message_id": "value",
  "integration_trigger_id": "value",
  "platform_metadata": {
    "conversation_initiation_source": "plain_integration",
    "thread_id": "value",
    "workspace_id": "value"
  }
}`
	assertRoundTrip(t, wire, new(AsyncConversationMetadata))
}
