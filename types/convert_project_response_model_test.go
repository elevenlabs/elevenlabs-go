package types

import (
	"testing"
)

func TestConvertProjectResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(ConvertProjectResponseModel))
}
