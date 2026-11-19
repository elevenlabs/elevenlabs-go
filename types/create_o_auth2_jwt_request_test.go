package types

import (
	"testing"
)

func TestCreateOAuth2JwtRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "oauth2_jwt",
  "provider": "value",
  "algorithm": "HS256",
  "key_id": "value",
  "issuer": "value",
  "audience": "value",
  "subject": "value",
  "expiration_seconds": 123,
  "extra_params": {
    "key": "value"
  },
  "extra_headers": {
    "key": "value"
  },
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "token_response_field": "access_token",
  "jwt_usage": "authorization_grant",
  "secret_key": "value"
}`
	assertRoundTrip(t, wire, new(CreateOAuth2JwtRequest))
}
