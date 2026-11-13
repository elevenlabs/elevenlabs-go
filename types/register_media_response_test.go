package types

import (
	"testing"
)

func TestRegisterMediaResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "media_id": "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa"
}`
	assertRoundTrip(t, wire, new(RegisterMediaResponse))
}
