package types

import (
	"testing"
)

func TestAdditionalFormatResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "requested_format": "value",
  "file_extension": "value",
  "content_type": "value",
  "is_base64_encoded": true,
  "content": "value"
}`
	assertRoundTrip(t, wire, new(AdditionalFormatResponseModel))
}
