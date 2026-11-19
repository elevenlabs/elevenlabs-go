package types

import (
	"testing"
)

func TestReaderResourceResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resource_type": "read",
  "resource_id": "value"
}`
	assertRoundTrip(t, wire, new(ReaderResourceResponseModel))
}
