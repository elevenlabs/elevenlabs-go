package types

import (
	"testing"
)

func TestSipTrunkCredentialsRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "username": "value",
  "password": "value"
}`
	assertRoundTrip(t, wire, new(SipTrunkCredentialsRequestModel))
}
