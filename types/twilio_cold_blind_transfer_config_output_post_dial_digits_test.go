package types

import (
	"testing"
)

func TestTwilioColdBlindTransferConfigOutputPostDialDigitsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "static",
  "value": "0"
}`
	assertRoundTrip(t, wire, new(TwilioColdBlindTransferConfigOutputPostDialDigits))
}
