package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ResourceMetadataResponseModel struct {
	// The ID of the resource.
	ResourceId string `json:"resource_id" xml:"resource_id" required:"true"`
	// The name of the resource, if available.
	ResourceName *string `json:"resource_name" xml:"resource_name" required:"true"`
	// Resource types that can be shared in the workspace. The name always need to match the collection names
	ResourceType WorkspaceResourceType `json:"resource_type" xml:"resource_type" required:"true"`
	// The ID of the user who created the resource.
	CreatorUserId *string `json:"creator_user_id" xml:"creator_user_id" required:"true"`
	// The access level for anonymous users. If None, the resource is not shared publicly.
	AnonymousAccessLevelOverride *ResourceMetadataResponseModelAnonymousAccessLevelOverride `json:"anonymous_access_level_override" xml:"anonymous_access_level_override" required:"true"`
	// A mapping of grant slots to group IDs. Keys are the preset access roles plus workspace custom-role ids ('role_...'). When the resource is shared with a user, the group id is the user's id.
	RoleToGroupIds map[string][]string `json:"role_to_group_ids" required:"true"`
	// List of options for sharing the resource further in the workspace. These are users who don't have access to the resource yet.
	ShareOptions    []*ShareOptionResponseModel `json:"share_options" xml:"share_options" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ResourceMetadataResponseModel) GetResourceId() (value string) {
	if r == nil {
		return
	}
	return r.ResourceId
}

func (r *ResourceMetadataResponseModel) GetResourceName() (value *string) {
	if r == nil {
		return
	}
	return r.ResourceName
}

func (r *ResourceMetadataResponseModel) GetResourceType() (value WorkspaceResourceType) {
	if r == nil {
		return
	}
	return r.ResourceType
}

func (r *ResourceMetadataResponseModel) GetCreatorUserId() (value *string) {
	if r == nil {
		return
	}
	return r.CreatorUserId
}

func (r *ResourceMetadataResponseModel) GetAnonymousAccessLevelOverride() (value *ResourceMetadataResponseModelAnonymousAccessLevelOverride) {
	if r == nil {
		return
	}
	return r.AnonymousAccessLevelOverride
}

func (r *ResourceMetadataResponseModel) GetRoleToGroupIds() (value map[string][]string) {
	if r == nil {
		return
	}
	return r.RoleToGroupIds
}

func (r *ResourceMetadataResponseModel) GetShareOptions() (value []*ShareOptionResponseModel) {
	if r == nil {
		return
	}
	return r.ShareOptions
}

func (r *ResourceMetadataResponseModel) SetResourceId(resourceId string) {
	r.ResourceId = resourceId
}

func (r *ResourceMetadataResponseModel) SetResourceName(resourceName *string) {
	r.ResourceName = resourceName
}

func (r *ResourceMetadataResponseModel) SetResourceType(resourceType WorkspaceResourceType) {
	r.ResourceType = resourceType
}

func (r *ResourceMetadataResponseModel) SetCreatorUserId(creatorUserId *string) {
	r.CreatorUserId = creatorUserId
}

func (r *ResourceMetadataResponseModel) SetAnonymousAccessLevelOverride(anonymousAccessLevelOverride *ResourceMetadataResponseModelAnonymousAccessLevelOverride) {
	r.AnonymousAccessLevelOverride = anonymousAccessLevelOverride
}

func (r *ResourceMetadataResponseModel) SetRoleToGroupIds(roleToGroupIds map[string][]string) {
	r.RoleToGroupIds = roleToGroupIds
}

func (r *ResourceMetadataResponseModel) SetShareOptions(shareOptions []*ShareOptionResponseModel) {
	r.ShareOptions = shareOptions
}

func (r *ResourceMetadataResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResourceMetadataResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResourceMetadataResponseModel to string"
	}
	return string(jsonData)
}

func (r *ResourceMetadataResponseModel) UnmarshalJSON(data []byte) error {
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

func (r *ResourceMetadataResponseModel) MarshalJSON() ([]byte, error) {
	type embed ResourceMetadataResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
