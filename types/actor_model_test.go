package types

import (
	"testing"
)

func TestActorModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "user": {
    "name": "value",
    "uid": "value",
    "type_id": 123,
    "type": "value",
    "email_addr": "value",
    "full_name": "value",
    "domain": "value"
  },
  "app_name": "value",
  "app_uid": "value",
  "session": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(ActorModel))
}
