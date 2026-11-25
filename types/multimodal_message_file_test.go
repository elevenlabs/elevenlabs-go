package types

import (
	"testing"
)

func TestMultimodalMessageFileRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "file_input",
  "file_id": "value"
}`
	assertRoundTrip(t, wire, new(MultimodalMessageFile))
}
