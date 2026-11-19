package types

import (
	"testing"
)

func TestErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message": "value",
  "error": "value",
  "code": 123,
  "param": "value"
}`
	assertRoundTrip(t, wire, new(Error))
}

func TestErrorRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message": "value",
  "error": "value",
  "code": 123,
  "param": "value"
}`
	assertRoundTrip(t, wire, new(Error))
}
