package types

import (
	"testing"
)

func TestOpenAiFunctionToolRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "function",
  "name": "value",
  "description": "value",
  "parameters": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(OpenAiFunctionTool))
}
