package types

import (
	"testing"
)

func TestLanguageInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "code": "value",
  "label": "value"
}`
	assertRoundTrip(t, wire, new(LanguageInfo))
}
