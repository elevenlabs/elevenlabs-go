package types

import (
	"testing"
)

func TestOrderMediaResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "media_id": "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa",
  "name": "value",
  "content_type": "value",
  "language": "value",
  "signed_url": "value"
}`
	assertRoundTrip(t, wire, new(OrderMediaResponse))
}
