package types

import (
	"testing"
)

func TestClipAnimationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enter_effect": "none",
  "enter_duration_ms": 123,
  "exit_effect": "none",
  "exit_duration_ms": 123
}`
	assertRoundTrip(t, wire, new(ClipAnimation))
}
