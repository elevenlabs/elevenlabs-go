package types

import (
	"testing"
)

func TestBase64FileSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "base64",
  "filename": "value",
  "media_type": "image/png",
  "data": "value"
}`
	assertRoundTrip(t, wire, new(Base64FileSource))
}
