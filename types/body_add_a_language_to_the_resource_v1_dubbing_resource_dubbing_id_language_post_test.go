package types

import (
	"testing"
)

func TestBodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language": "value"
}`
	assertRoundTrip(t, wire, new(BodyAddALanguageToTheResourceV1DubbingResourceDubbingIdLanguagePost))
}
