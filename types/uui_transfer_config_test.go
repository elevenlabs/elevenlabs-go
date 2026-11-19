package types

import (
	"testing"
)

func TestUuiTransferConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "data": "value",
  "protocol_discriminator": "00",
  "protocol_discriminator_mode": "prefix"
}`
	assertRoundTrip(t, wire, new(UuiTransferConfig))
}
