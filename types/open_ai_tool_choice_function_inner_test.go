package types

import (
	"testing"
)

func TestOpenAiToolChoiceFunctionInnerRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value"
}`
	assertRoundTrip(t, wire, new(OpenAiToolChoiceFunctionInner))
}
