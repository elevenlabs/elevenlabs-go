package types

import (
	"testing"
)

func TestResellerWorkspaceOnboardingResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "onboarding_completed_at_unix": 123,
  "customer_facing_agent_id": "value"
}`
	assertRoundTrip(t, wire, new(ResellerWorkspaceOnboardingResponse))
}
