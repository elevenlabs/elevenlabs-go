package templates

import (
	"testing"
)

func TestTemplateRunCreateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "inputs": {
    "key": "value"
  },
  "version_id": "value",
  "webhook": {
    "type": "all"
  }
}`
	assertRoundTrip(t, wire, new(TemplateRunCreateRequest))
}
