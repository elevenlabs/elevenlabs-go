package types

import (
	"testing"
)

func TestFieldConflictRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "path": "value",
  "section": "conversation_config",
  "base_value": {},
  "source_value": {},
  "target_value": {},
  "id": "value",
  "accepts_custom_value": true
}`
	assertRoundTrip(t, wire, new(FieldConflict))
}
