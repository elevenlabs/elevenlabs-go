package types

import (
	"testing"
)

func TestProjectVideoThumbnailSheetResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "start_thumbnail_index": 123,
  "thumbnail_count": 123,
  "signed_cloud_url": "value"
}`
	assertRoundTrip(t, wire, new(ProjectVideoThumbnailSheetResponseModel))
}
