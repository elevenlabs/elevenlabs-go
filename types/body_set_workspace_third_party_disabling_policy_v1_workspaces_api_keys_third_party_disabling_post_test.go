package types

import (
	"testing"
)

func TestBodySetWorkspaceThirdPartyDisablingPolicyV1WorkspacesApiKeysThirdPartyDisablingPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "third_party_disable_allowed": true
}`
	assertRoundTrip(t, wire, new(BodySetWorkspaceThirdPartyDisablingPolicyV1WorkspacesApiKeysThirdPartyDisablingPost))
}
