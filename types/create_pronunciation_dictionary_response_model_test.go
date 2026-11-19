package types

import (
	"testing"
)

func TestCreatePronunciationDictionaryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(CreatePronunciationDictionaryResponseModel))
}
