package types

import (
	"testing"
)

func TestOpenAiToolChoiceFunctionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "function",
  "function": {
    "name": "value"
  }
}`
	assertRoundTrip(t, wire, new(OpenAiToolChoiceFunction))
}
