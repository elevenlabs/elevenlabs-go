package types

import (
	"testing"
)

func TestSegmentMigrationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SegmentMigrationResponse))
}
