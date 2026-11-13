package types

import (
	"testing"
)

func TestWorkspaceCreateServiceAccountResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "service-account-user-id": "value"
}`
	assertRoundTrip(t, wire, new(WorkspaceCreateServiceAccountResponseModel))
}
