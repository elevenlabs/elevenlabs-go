package types

import (
	"testing"
)

func TestLanguageAddedResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(LanguageAddedResponse))
}
