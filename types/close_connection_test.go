package types

import (
	"testing"
)

func TestCloseConnectionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": ""
}`
	assertRoundTrip(t, wire, new(CloseConnection))
}
