package types

import (
	"testing"
)

func TestOpenAiSessionConfigInputToolsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "function",
  "name": "value",
  "description": "value",
  "parameters": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(OpenAiSessionConfigInputToolsItem))
}
