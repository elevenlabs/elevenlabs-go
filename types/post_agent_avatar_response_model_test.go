package types

import (
	"testing"
)

func TestPostAgentAvatarResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "avatar_url": "value"
}`
	assertRoundTrip(t, wire, new(PostAgentAvatarResponseModel))
}
