package types

import (
	"testing"
)

func TestGetLiveCountResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "count": 123
}`
	assertRoundTrip(t, wire, new(GetLiveCountResponse))
}
