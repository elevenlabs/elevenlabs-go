package types

import (
	"testing"
)

func TestCloseSocketRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "close_socket": true
}`
	assertRoundTrip(t, wire, new(CloseSocket))
}
