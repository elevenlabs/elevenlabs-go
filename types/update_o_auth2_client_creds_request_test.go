package types

import (
	"testing"
)

func TestUpdateOAuth2ClientCredsRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "client_secret": "value",
  "custom_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(UpdateOAuth2ClientCredsRequest))
}
