package types

import (
	"testing"
)

func TestWhatsAppCallPermissionRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "state": "blocked",
  "requested_at_unix": 123,
  "responded_at_unix": 123,
  "permanent": true,
  "granted_by_user_call": true
}`
	assertRoundTrip(t, wire, new(WhatsAppCallPermissionRequest))
}
