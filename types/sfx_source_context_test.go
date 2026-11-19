package types

import (
	"testing"
)

func TestSfxSourceContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "sfx",
  "sound_generation_history_item_id": "value",
  "text": "value",
  "generation_config": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(SfxSourceContext))
}
