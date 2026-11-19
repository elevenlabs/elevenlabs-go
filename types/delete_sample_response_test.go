package types

import (
	"testing"
)

func TestDeleteSampleResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteSampleResponse))
}
