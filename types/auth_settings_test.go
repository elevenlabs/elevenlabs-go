package types

import (
	"testing"
)

func TestAuthSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enable_auth": true,
  "allowlist": [
    {
      "hostname": "value"
    }
  ],
  "require_origin_header": true,
  "shareable_token": "value"
}`
	assertRoundTrip(t, wire, new(AuthSettings))
}
