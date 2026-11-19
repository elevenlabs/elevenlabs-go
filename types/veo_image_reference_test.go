package types

import (
	"testing"
)

func TestVeoImageReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "image": {
    "type": "generation",
    "generation_id": "value"
  },
  "role": "subject"
}`
	assertRoundTrip(t, wire, new(VeoImageReference))
}
