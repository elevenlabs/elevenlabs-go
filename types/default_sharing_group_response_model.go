package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DefaultSharingGroupResponseModel struct {
	Group *WorkspaceGroupResponseModel `json:"group" xml:"group" required:"true"`
	// The permission level to grant to the group
	PermissionLevel DefaultSharingGroupResponseModelPermissionLevel `json:"permission_level" xml:"permission_level" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DefaultSharingGroupResponseModel) GetGroup() (value *WorkspaceGroupResponseModel) {
	if d == nil {
		return
	}
	return d.Group
}

func (d *DefaultSharingGroupResponseModel) GetPermissionLevel() (value DefaultSharingGroupResponseModelPermissionLevel) {
	if d == nil {
		return
	}
	return d.PermissionLevel
}

func (d *DefaultSharingGroupResponseModel) SetGroup(group *WorkspaceGroupResponseModel) {
	d.Group = group
}

func (d *DefaultSharingGroupResponseModel) SetPermissionLevel(permissionLevel DefaultSharingGroupResponseModelPermissionLevel) {
	d.PermissionLevel = permissionLevel
}

func (d *DefaultSharingGroupResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DefaultSharingGroupResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DefaultSharingGroupResponseModel to string"
	}
	return string(jsonData)
}

func (d *DefaultSharingGroupResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DefaultSharingGroupResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DefaultSharingGroupResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DefaultSharingGroupResponseModel) MarshalJSON() ([]byte, error) {
	type embed DefaultSharingGroupResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
