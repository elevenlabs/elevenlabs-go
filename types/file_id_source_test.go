package types

import (
	"testing"
)

func TestFileIdSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file_id",
  "file_id": "value"
}`
	assertRoundTrip(t, wire, new(FileIdSource))
}
