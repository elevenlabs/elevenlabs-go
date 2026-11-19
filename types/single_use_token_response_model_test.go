package types

import (
	"testing"
)

func TestSingleUseTokenResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "token": "value"
}`
	assertRoundTrip(t, wire, new(SingleUseTokenResponseModel))
}
