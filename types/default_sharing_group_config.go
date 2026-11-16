package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DefaultSharingGroupConfig struct {
	// The ID of the group to share with
	GroupId string `json:"group_id" xml:"group_id" required:"true"`
	// The permission level to grant to the group
	PermissionLevel DefaultSharingGroupResponseModelPermissionLevel `json:"permission_level" xml:"permission_level" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DefaultSharingGroupConfig) GetGroupId() (value string) {
	if d == nil {
		return
	}
	return d.GroupId
}

func (d *DefaultSharingGroupConfig) GetPermissionLevel() (value DefaultSharingGroupResponseModelPermissionLevel) {
	if d == nil {
		return
	}
	return d.PermissionLevel
}

func (d *DefaultSharingGroupConfig) SetGroupId(groupId string) {
	d.GroupId = groupId
}

func (d *DefaultSharingGroupConfig) SetPermissionLevel(permissionLevel DefaultSharingGroupResponseModelPermissionLevel) {
	d.PermissionLevel = permissionLevel
}

func (d *DefaultSharingGroupConfig) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DefaultSharingGroupConfig) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DefaultSharingGroupConfig to string"
	}
	return string(jsonData)
}

func (d *DefaultSharingGroupConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DefaultSharingGroupConfig
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DefaultSharingGroupConfig(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DefaultSharingGroupConfig) MarshalJSON() ([]byte, error) {
	type embed DefaultSharingGroupConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
