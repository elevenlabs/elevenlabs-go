package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkspaceGroupResponseModel struct {
	Name                         string                                                   `json:"name" xml:"name" required:"true"`
	Id                           string                                                   `json:"id" xml:"id" required:"true"`
	Members                      []string                                                 `json:"members" xml:"members" required:"true"`
	Permissions                  []WorkspaceGroupPermission                               `json:"permissions" xml:"permissions" required:"true"`
	GroupUsageLimit              *WorkspaceGroupResponseModelGroupUsageLimit              `json:"group_usage_limit,omitempty" xml:"group_usage_limit,omitempty"`
	GroupPvcLimit                *WorkspaceGroupResponseModelGroupPvcLimit                `json:"group_pvc_limit,omitempty" xml:"group_pvc_limit,omitempty"`
	GroupConcurrencyLimit        *WorkspaceGroupResponseModelGroupConcurrencyLimit        `json:"group_concurrency_limit,omitempty" xml:"group_concurrency_limit,omitempty"`
	GroupDubbingConcurrencyLimit *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit `json:"group_dubbing_concurrency_limit,omitempty" xml:"group_dubbing_concurrency_limit,omitempty"`
	GroupMusicConcurrencyLimit   *WorkspaceGroupResponseModelGroupMusicConcurrencyLimit   `json:"group_music_concurrency_limit,omitempty" xml:"group_music_concurrency_limit,omitempty"`
	CharacterCount               *int                                                     `json:"character_count,omitempty" xml:"character_count,omitempty"`
	IsScimSynced                 *bool                                                    `json:"is_scim_synced,omitempty" xml:"is_scim_synced,omitempty"`
	ScimGroup                    *ScimGroupResponseModel                                  `json:"scim_group,omitempty" xml:"scim_group,omitempty"`
	ScimFrozen                   *bool                                                    `json:"scim_frozen,omitempty" xml:"scim_frozen,omitempty"`
	extraProperties              map[string]interface{}
}

func (w *WorkspaceGroupResponseModel) GetName() (value string) {
	if w == nil {
		return
	}
	return w.Name
}

func (w *WorkspaceGroupResponseModel) GetId() (value string) {
	if w == nil {
		return
	}
	return w.Id
}

func (w *WorkspaceGroupResponseModel) GetMembers() (value []string) {
	if w == nil {
		return
	}
	return w.Members
}

func (w *WorkspaceGroupResponseModel) GetPermissions() (value []WorkspaceGroupPermission) {
	if w == nil {
		return
	}
	return w.Permissions
}

func (w *WorkspaceGroupResponseModel) GetGroupUsageLimit() (value *WorkspaceGroupResponseModelGroupUsageLimit) {
	if w == nil {
		return
	}
	return w.GroupUsageLimit
}

func (w *WorkspaceGroupResponseModel) GetGroupPvcLimit() (value *WorkspaceGroupResponseModelGroupPvcLimit) {
	if w == nil {
		return
	}
	return w.GroupPvcLimit
}

func (w *WorkspaceGroupResponseModel) GetGroupConcurrencyLimit() (value *WorkspaceGroupResponseModelGroupConcurrencyLimit) {
	if w == nil {
		return
	}
	return w.GroupConcurrencyLimit
}

func (w *WorkspaceGroupResponseModel) GetGroupDubbingConcurrencyLimit() (value *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) {
	if w == nil {
		return
	}
	return w.GroupDubbingConcurrencyLimit
}

func (w *WorkspaceGroupResponseModel) GetGroupMusicConcurrencyLimit() (value *WorkspaceGroupResponseModelGroupMusicConcurrencyLimit) {
	if w == nil {
		return
	}
	return w.GroupMusicConcurrencyLimit
}

func (w *WorkspaceGroupResponseModel) GetCharacterCount() (value *int) {
	if w == nil {
		return
	}
	return w.CharacterCount
}

func (w *WorkspaceGroupResponseModel) GetIsScimSynced() (value *bool) {
	if w == nil {
		return
	}
	return w.IsScimSynced
}

func (w *WorkspaceGroupResponseModel) GetScimGroup() (value *ScimGroupResponseModel) {
	if w == nil {
		return
	}
	return w.ScimGroup
}

func (w *WorkspaceGroupResponseModel) GetScimFrozen() (value *bool) {
	if w == nil {
		return
	}
	return w.ScimFrozen
}

func (w *WorkspaceGroupResponseModel) SetName(name string) {
	w.Name = name
}

func (w *WorkspaceGroupResponseModel) SetId(id string) {
	w.Id = id
}

func (w *WorkspaceGroupResponseModel) SetMembers(members []string) {
	w.Members = members
}

func (w *WorkspaceGroupResponseModel) SetPermissions(permissions []WorkspaceGroupPermission) {
	w.Permissions = permissions
}

func (w *WorkspaceGroupResponseModel) SetGroupUsageLimit(groupUsageLimit *WorkspaceGroupResponseModelGroupUsageLimit) {
	w.GroupUsageLimit = groupUsageLimit
}

func (w *WorkspaceGroupResponseModel) SetGroupPvcLimit(groupPvcLimit *WorkspaceGroupResponseModelGroupPvcLimit) {
	w.GroupPvcLimit = groupPvcLimit
}

func (w *WorkspaceGroupResponseModel) SetGroupConcurrencyLimit(groupConcurrencyLimit *WorkspaceGroupResponseModelGroupConcurrencyLimit) {
	w.GroupConcurrencyLimit = groupConcurrencyLimit
}

func (w *WorkspaceGroupResponseModel) SetGroupDubbingConcurrencyLimit(groupDubbingConcurrencyLimit *WorkspaceGroupResponseModelGroupDubbingConcurrencyLimit) {
	w.GroupDubbingConcurrencyLimit = groupDubbingConcurrencyLimit
}

func (w *WorkspaceGroupResponseModel) SetGroupMusicConcurrencyLimit(groupMusicConcurrencyLimit *WorkspaceGroupResponseModelGroupMusicConcurrencyLimit) {
	w.GroupMusicConcurrencyLimit = groupMusicConcurrencyLimit
}

func (w *WorkspaceGroupResponseModel) SetCharacterCount(characterCount *int) {
	w.CharacterCount = characterCount
}

func (w *WorkspaceGroupResponseModel) SetIsScimSynced(isScimSynced *bool) {
	w.IsScimSynced = isScimSynced
}

func (w *WorkspaceGroupResponseModel) SetScimGroup(scimGroup *ScimGroupResponseModel) {
	w.ScimGroup = scimGroup
}

func (w *WorkspaceGroupResponseModel) SetScimFrozen(scimFrozen *bool) {
	w.ScimFrozen = scimFrozen
}

func (w *WorkspaceGroupResponseModel) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkspaceGroupResponseModel) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkspaceGroupResponseModel to string"
	}
	return string(jsonData)
}

func (w *WorkspaceGroupResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkspaceGroupResponseModel) MarshalJSON() ([]byte, error) {
	type embed WorkspaceGroupResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
