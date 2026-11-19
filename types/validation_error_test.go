package types

import (
	"testing"
)

func TestValidationErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "loc": [
    "value"
  ],
  "msg": "value",
  "type": "value"
}`
	assertRoundTrip(t, wire, new(ValidationError))
}
