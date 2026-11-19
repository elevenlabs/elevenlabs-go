package types

import (
	"testing"
)

func TestFileContentPartRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file",
  "source": {
    "type": "base64",
    "filename": "value",
    "media_type": "image/png",
    "data": "value"
  }
}`
	assertRoundTrip(t, wire, new(FileContentPart))
}
