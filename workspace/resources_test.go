package workspace

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "role": "admin",
  "resource_type": "voice",
  "user_email": "value",
  "group_id": "value",
  "workspace_api_key_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost))
}

func TestBodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "resource_type": "voice",
  "user_email": "value",
  "group_id": "value",
  "workspace_api_key_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost))
}

func TestBodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole
		wire    string
	}
	cases := []enumCase{
		{variant: BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleAdmin, wire: "\"admin\""},
		{variant: BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleEditor, wire: "\"editor\""},
		{variant: BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleCommenter, wire: "\"commenter\""},
		{variant: BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleViewer, wire: "\"viewer\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
