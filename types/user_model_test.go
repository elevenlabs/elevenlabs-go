package types

import (
	"testing"
)

func TestUserModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "uid": "value",
  "type_id": 123,
  "type": "value",
  "email_addr": "value",
  "full_name": "value",
  "domain": "value"
}`
	assertRoundTrip(t, wire, new(UserModel))
}
