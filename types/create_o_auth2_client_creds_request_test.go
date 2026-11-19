package types

import (
	"testing"
)

func TestCreateOAuth2ClientCredsRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
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
	assertRoundTrip(t, wire, new(CreateOAuth2ClientCredsRequest))
}
