package orders

import (
	"testing"
)

func TestBodyRegisterMediaV1ProductionsOrdersOrderIdMediaPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "declared_language": "value",
  "media_url": "value",
  "media_url_filename": "value",
  "media_url_content_type": "value"
}`
	assertRoundTrip(t, wire, new(BodyRegisterMediaV1ProductionsOrdersOrderIdMediaPost))
}
