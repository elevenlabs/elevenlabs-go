package types

import (
	"testing"
)

func TestDubbingRenderResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123,
  "render_id": "value"
}`
	assertRoundTrip(t, wire, new(DubbingRenderResponseModel))
}
