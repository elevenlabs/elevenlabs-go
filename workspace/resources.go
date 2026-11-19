package workspace

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost struct {
	// The ID of the target resource.
	ResourceId string `json:"-" url:"-"`
	// Role to grant to the target: one of 'admin', 'editor', 'commenter', or 'viewer'.
	Role BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole `json:"role" xml:"role" required:"true"`
	// Resource types that can be shared in the workspace. The name always need to match the collection names
	ResourceType types.WorkspaceResourceType `json:"resource_type" xml:"resource_type" required:"true"`
	// The email of the user or service account.
	UserEmail *string `json:"user_email,omitempty" xml:"user_email,omitempty"`
	// The ID of the target group. Use 'default' to set the resource's baseline role — every workspace member receives this role unless they hold a higher one through a direct user grant, group membership, or workspace (service account) API key.
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The ID of the target workspace (service account) API key. This is not the API key string itself that you pass in the header for authentication — it is the key's ID, which workspace admins can find under Developers → Service Accounts.
	WorkspaceApiKeyId *string `json:"workspace_api_key_id,omitempty" xml:"workspace_api_key_id,omitempty"`
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetResourceId() (value string) {
	if b == nil {
		return
	}
	return b.ResourceId
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetRole() (value BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole) {
	if b == nil {
		return
	}
	return b.Role
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetResourceType() (value types.WorkspaceResourceType) {
	if b == nil {
		return
	}
	return b.ResourceType
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetUserEmail() (value *string) {
	if b == nil {
		return
	}
	return b.UserEmail
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetGroupId() (value *string) {
	if b == nil {
		return
	}
	return b.GroupId
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetWorkspaceApiKeyId() (value *string) {
	if b == nil {
		return
	}
	return b.WorkspaceApiKeyId
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) SetResourceId(resourceId string) {
	b.ResourceId = resourceId
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) SetRole(role BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole) {
	b.Role = role
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) SetResourceType(resourceType types.WorkspaceResourceType) {
	b.ResourceType = resourceType
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) SetUserEmail(userEmail *string) {
	b.UserEmail = userEmail
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) SetGroupId(groupId *string) {
	b.GroupId = groupId
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) SetWorkspaceApiKeyId(workspaceApiKeyId *string) {
	b.WorkspaceApiKeyId = workspaceApiKeyId
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost to string"
	}
	return string(jsonData)
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost) MarshalJSON() ([]byte, error) {
	type embed BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

// Role to grant to the target: one of 'admin', 'editor', 'commenter', or 'viewer'.
type BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole string

const (
	BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleAdmin     BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole = "admin"
	BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleEditor    BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole = "editor"
	BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleCommenter BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole = "commenter"
	BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleViewer    BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole = "viewer"
)

func NewBodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleFromString(s string) (BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole, error) {
	switch s {
	case "admin":
		return BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleAdmin, nil
	case "editor":
		return BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleEditor, nil
	case "commenter":
		return BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleCommenter, nil
	case "viewer":
		return BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRoleViewer, nil
	}
	var t BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole) Ptr() *BodyShareWorkspaceResourceV1WorkspaceResourcesResourceIdSharePostRole {
	return &e
}

type BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost struct {
	// The ID of the target resource.
	ResourceId string `json:"-" url:"-"`
	// Resource types that can be shared in the workspace. The name always need to match the collection names
	ResourceType types.WorkspaceResourceType `json:"resource_type" xml:"resource_type" required:"true"`
	// The email of the user or service account.
	UserEmail *string `json:"user_email,omitempty" xml:"user_email,omitempty"`
	// The ID of the target group. Use 'default' to set the resource's baseline role — every workspace member receives this role unless they hold a higher one through a direct user grant, group membership, or workspace (service account) API key.
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The ID of the target workspace (service account) API key. This is not the API key string itself that you pass in the header for authentication — it is the key's ID, which workspace admins can find under Developers → Service Accounts.
	WorkspaceApiKeyId *string `json:"workspace_api_key_id,omitempty" xml:"workspace_api_key_id,omitempty"`
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) GetResourceId() (value string) {
	if b == nil {
		return
	}
	return b.ResourceId
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) GetResourceType() (value types.WorkspaceResourceType) {
	if b == nil {
		return
	}
	return b.ResourceType
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) GetUserEmail() (value *string) {
	if b == nil {
		return
	}
	return b.UserEmail
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) GetGroupId() (value *string) {
	if b == nil {
		return
	}
	return b.GroupId
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) GetWorkspaceApiKeyId() (value *string) {
	if b == nil {
		return
	}
	return b.WorkspaceApiKeyId
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) SetResourceId(resourceId string) {
	b.ResourceId = resourceId
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) SetResourceType(resourceType types.WorkspaceResourceType) {
	b.ResourceType = resourceType
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) SetUserEmail(userEmail *string) {
	b.UserEmail = userEmail
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) SetGroupId(groupId *string) {
	b.GroupId = groupId
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) SetWorkspaceApiKeyId(workspaceApiKeyId *string) {
	b.WorkspaceApiKeyId = workspaceApiKeyId
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost to string"
	}
	return string(jsonData)
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost) MarshalJSON() ([]byte, error) {
	type embed BodyUnshareWorkspaceResourceV1WorkspaceResourcesResourceIdUnsharePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	getResourcesRequestFieldResourceType = big.NewInt(1 << 0)
)

type GetResourcesRequest struct {
	ResourceType types.WorkspaceResourceType `json:"-" url:"resource_type,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetResourcesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetResourceType sets the ResourceType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetResourcesRequest) SetResourceType(resourceType types.WorkspaceResourceType) {
	g.ResourceType = resourceType
	g.require(getResourcesRequestFieldResourceType)
}
