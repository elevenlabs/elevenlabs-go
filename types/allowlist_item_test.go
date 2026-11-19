package types

import (
	"testing"
)

func TestAllowlistItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "hostname": "value"
}`
	assertRoundTrip(t, wire, new(AllowlistItem))
}

func TestAllowlistItemRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "hostname": "value"
}`
	assertRoundTrip(t, wire, new(AllowlistItem))
}
