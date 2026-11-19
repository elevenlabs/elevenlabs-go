package types

import (
	"testing"
)

func TestRequiredConstraintsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "any_of": [
    {
      "required": [
        "value"
      ]
    }
  ],
  "all_of": [
    {
      "required": [
        "value"
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(RequiredConstraints))
}
