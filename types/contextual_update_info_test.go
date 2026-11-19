package types

import (
	"testing"
)

func TestContextualUpdateInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "context_id": "value",
  "is_superseded": true
}`
	assertRoundTrip(t, wire, new(ContextualUpdateInfo))
}
