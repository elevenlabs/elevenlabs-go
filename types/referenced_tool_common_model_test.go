package types

import (
	"testing"
)

func TestReferencedToolCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "type": "system"
}`
	assertRoundTrip(t, wire, new(ReferencedToolCommonModel))
}

func TestReferencedToolCommonModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "type": "system"
}`
	assertRoundTrip(t, wire, new(ReferencedToolCommonModel))
}
