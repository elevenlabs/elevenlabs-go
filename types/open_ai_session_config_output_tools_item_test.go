package types

import (
	"testing"
)

func TestOpenAiSessionConfigOutputToolsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "function",
  "name": "value",
  "description": "value",
  "parameters": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(OpenAiSessionConfigOutputToolsItem))
}
