package types

import (
	"testing"
)

func TestUsageCharactersResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "time": [
    123
  ],
  "usage": {
    "key": [
      9.99
    ]
  }
}`
	assertRoundTrip(t, wire, new(UsageCharactersResponseModel))
}
