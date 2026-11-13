package types

import (
	"testing"
)

func TestUpdateMemberUsageLimitResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(UpdateMemberUsageLimitResponseModel))
}
