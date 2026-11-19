package types

import (
	"testing"
)

func TestConvAiFileUploadResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "file_id": "value"
}`
	assertRoundTrip(t, wire, new(ConvAiFileUploadResponseModel))
}
