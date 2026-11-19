package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserVoiceCollectionResponseModel struct {
	// The ID of the voice collection.
	CollectionId string `json:"collection_id" xml:"collection_id" required:"true"`
	// The title of the voice collection.
	Title string `json:"title" xml:"title" required:"true"`
	// The icon of the voice collection.
	Icon string `json:"icon" xml:"icon" required:"true"`
	// The caller's access role on the voice collection.
	PermissionOnResource UserVoiceCollectionResponseModelPermissionOnResource `json:"permission_on_resource" xml:"permission_on_resource" required:"true"`
	extraProperties      map[string]interface{}
}

func (u *UserVoiceCollectionResponseModel) GetCollectionId() (value string) {
	if u == nil {
		return
	}
	return u.CollectionId
}

func (u *UserVoiceCollectionResponseModel) GetTitle() (value string) {
	if u == nil {
		return
	}
	return u.Title
}

func (u *UserVoiceCollectionResponseModel) GetIcon() (value string) {
	if u == nil {
		return
	}
	return u.Icon
}

func (u *UserVoiceCollectionResponseModel) GetPermissionOnResource() (value UserVoiceCollectionResponseModelPermissionOnResource) {
	if u == nil {
		return
	}
	return u.PermissionOnResource
}

func (u *UserVoiceCollectionResponseModel) SetCollectionId(collectionId string) {
	u.CollectionId = collectionId
}

func (u *UserVoiceCollectionResponseModel) SetTitle(title string) {
	u.Title = title
}

func (u *UserVoiceCollectionResponseModel) SetIcon(icon string) {
	u.Icon = icon
}

func (u *UserVoiceCollectionResponseModel) SetPermissionOnResource(permissionOnResource UserVoiceCollectionResponseModelPermissionOnResource) {
	u.PermissionOnResource = permissionOnResource
}

func (u *UserVoiceCollectionResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserVoiceCollectionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserVoiceCollectionResponseModel to string"
	}
	return string(jsonData)
}

func (u *UserVoiceCollectionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UserVoiceCollectionResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UserVoiceCollectionResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserVoiceCollectionResponseModel) MarshalJSON() ([]byte, error) {
	type embed UserVoiceCollectionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
