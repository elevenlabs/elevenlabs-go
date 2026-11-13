package types

import (
	"testing"
)

func TestDeleteChapterResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteChapterResponseModel))
}
