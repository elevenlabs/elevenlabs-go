package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AvatarContextResponseModel struct {
	// The ID of the avatar.
	AvatarId *string `json:"avatar_id,omitempty" xml:"avatar_id,omitempty"`
	// The ID of the avatar style.
	AvatarStyleId *string `json:"avatar_style_id,omitempty" xml:"avatar_style_id,omitempty"`
	// The name of the avatar.
	AvatarName *string `json:"avatar_name,omitempty" xml:"avatar_name,omitempty"`
	// The name of the avatar style.
	AvatarStyleName *string `json:"avatar_style_name,omitempty" xml:"avatar_style_name,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AvatarContextResponseModel) GetAvatarId() (value *string) {
	if a == nil {
		return
	}
	return a.AvatarId
}

func (a *AvatarContextResponseModel) GetAvatarStyleId() (value *string) {
	if a == nil {
		return
	}
	return a.AvatarStyleId
}

func (a *AvatarContextResponseModel) GetAvatarName() (value *string) {
	if a == nil {
		return
	}
	return a.AvatarName
}

func (a *AvatarContextResponseModel) GetAvatarStyleName() (value *string) {
	if a == nil {
		return
	}
	return a.AvatarStyleName
}

func (a *AvatarContextResponseModel) SetAvatarId(avatarId *string) {
	a.AvatarId = avatarId
}

func (a *AvatarContextResponseModel) SetAvatarStyleId(avatarStyleId *string) {
	a.AvatarStyleId = avatarStyleId
}

func (a *AvatarContextResponseModel) SetAvatarName(avatarName *string) {
	a.AvatarName = avatarName
}

func (a *AvatarContextResponseModel) SetAvatarStyleName(avatarStyleName *string) {
	a.AvatarStyleName = avatarStyleName
}

func (a *AvatarContextResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AvatarContextResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AvatarContextResponseModel to string"
	}
	return string(jsonData)
}

func (a *AvatarContextResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AvatarContextResponseModel) MarshalJSON() ([]byte, error) {
	type embed AvatarContextResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
