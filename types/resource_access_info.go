package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResourceAccessInfo struct {
	// Whether the user making the request is the creator of the agent
	IsCreator bool `json:"is_creator" xml:"is_creator" required:"true"`
	// Name of the agent's creator
	CreatorName string `json:"creator_name" xml:"creator_name" required:"true"`
	// Email of the agent's creator
	CreatorEmail string `json:"creator_email" xml:"creator_email" required:"true"`
	// The role of the user making the request
	Role ResourceAccessInfoRole `json:"role" xml:"role" required:"true"`
	// The access level for anonymous users. If None, the resource is not shared publicly.
	AnonymousAccessLevelOverride *ResourceAccessInfoAnonymousAccessLevelOverride `json:"anonymous_access_level_override,omitempty" xml:"anonymous_access_level_override,omitempty"`
	// Why the requesting user has access to this resource. 'creator' = caller is the owner. 'explicit' = caller (or one of their workspace groups) is listed in role_to_group_ids beyond the workspace-wide everyone group. 'workspace_default' = the workspace-wide everyone group is listed in role_to_group_ids (every non-anon workspace member, including admins, sees this resource). 'workspace_admin' = caller is a workspace admin and the admin seat is the *only* path to access; reserved for docs nobody else can see. Lets the UI disclose why an admin-bypass viewer sees a doc that wasn't explicitly shared with them.
	AccessSource    *ResourceAccessInfoAccessSource `json:"access_source,omitempty" xml:"access_source,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ResourceAccessInfo) GetIsCreator() (value bool) {
	if r == nil {
		return
	}
	return r.IsCreator
}

func (r *ResourceAccessInfo) GetCreatorName() (value string) {
	if r == nil {
		return
	}
	return r.CreatorName
}

func (r *ResourceAccessInfo) GetCreatorEmail() (value string) {
	if r == nil {
		return
	}
	return r.CreatorEmail
}

func (r *ResourceAccessInfo) GetRole() (value ResourceAccessInfoRole) {
	if r == nil {
		return
	}
	return r.Role
}

func (r *ResourceAccessInfo) GetAnonymousAccessLevelOverride() (value *ResourceAccessInfoAnonymousAccessLevelOverride) {
	if r == nil {
		return
	}
	return r.AnonymousAccessLevelOverride
}

func (r *ResourceAccessInfo) GetAccessSource() (value *ResourceAccessInfoAccessSource) {
	if r == nil {
		return
	}
	return r.AccessSource
}

func (r *ResourceAccessInfo) SetIsCreator(isCreator bool) {
	r.IsCreator = isCreator
}

func (r *ResourceAccessInfo) SetCreatorName(creatorName string) {
	r.CreatorName = creatorName
}

func (r *ResourceAccessInfo) SetCreatorEmail(creatorEmail string) {
	r.CreatorEmail = creatorEmail
}

func (r *ResourceAccessInfo) SetRole(role ResourceAccessInfoRole) {
	r.Role = role
}

func (r *ResourceAccessInfo) SetAnonymousAccessLevelOverride(anonymousAccessLevelOverride *ResourceAccessInfoAnonymousAccessLevelOverride) {
	r.AnonymousAccessLevelOverride = anonymousAccessLevelOverride
}

func (r *ResourceAccessInfo) SetAccessSource(accessSource *ResourceAccessInfoAccessSource) {
	r.AccessSource = accessSource
}

func (r *ResourceAccessInfo) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResourceAccessInfo) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResourceAccessInfo to string"
	}
	return string(jsonData)
}

func (r *ResourceAccessInfo) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResourceAccessInfo) MarshalJSON() ([]byte, error) {
	type embed ResourceAccessInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
