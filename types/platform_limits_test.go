package types

import (
	"testing"
)

func TestPlatformLimitsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "credits": {
    "limit": 123,
    "usage": 123
  },
  "pvc": {
    "limit": 123
  },
  "concurrency": {
    "limit": 123
  },
  "dubbing_concurrency": {
    "limit": 123
  },
  "music_concurrency": {
    "limit": 123
  }
}`
	assertRoundTrip(t, wire, new(PlatformLimits))
}
