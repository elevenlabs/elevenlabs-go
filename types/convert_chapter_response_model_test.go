package types

import (
	"testing"
)

func TestConvertChapterResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(ConvertChapterResponseModel))
}
