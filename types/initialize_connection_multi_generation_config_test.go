package types

import (
	"testing"
)

func TestInitializeConnectionMultiGenerationConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunk_length_schedule": [
    123
  ]
}`
	assertRoundTrip(t, wire, new(InitializeConnectionMultiGenerationConfig))
}
