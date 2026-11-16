package types

import (
	"testing"
)

func TestDummyToolResultModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "dummy"
}`
	assertRoundTrip(t, wire, new(DummyToolResultModel))
}

func TestDummyToolResultModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "dummy"
}`
	assertRoundTrip(t, wire, new(DummyToolResultModel))
}
