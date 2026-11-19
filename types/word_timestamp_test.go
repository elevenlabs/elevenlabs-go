package types

import (
	"testing"
)

func TestWordTimestampRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "word": "value",
  "start_ms": 123,
  "end_ms": 123
}`
	assertRoundTrip(t, wire, new(WordTimestamp))
}

func TestWordTimestampRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "word": "value",
  "start_ms": 123,
  "end_ms": 123
}`
	assertRoundTrip(t, wire, new(WordTimestamp))
}
