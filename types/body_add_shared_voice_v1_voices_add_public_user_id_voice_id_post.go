package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost struct {
	// ID of the voice to be used. You can use the [Get voices](/docs/api-reference/voices/search) endpoint list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// Public user ID used to publicly identify ElevenLabs users.
	PublicUserId string `json:"-" url:"-"`
	// The name that identifies this voice. This will be displayed in the dropdown of the website.
	NewName    string `json:"new_name" xml:"new_name" required:"true"`
	Bookmarked *bool  `json:"bookmarked,omitempty" xml:"bookmarked,omitempty"`
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) GetPublicUserId() (value string) {
	if b == nil {
		return
	}
	return b.PublicUserId
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) GetNewName() (value string) {
	if b == nil {
		return
	}
	return b.NewName
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) GetBookmarked() (value *bool) {
	if b == nil {
		return
	}
	return b.Bookmarked
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) SetPublicUserId(publicUserId string) {
	b.PublicUserId = publicUserId
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) SetNewName(newName string) {
	b.NewName = newName
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) SetBookmarked(bookmarked *bool) {
	b.Bookmarked = bookmarked
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost to string"
	}
	return string(jsonData)
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost(tmp)
	return nil
}

func (b *BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost) MarshalJSON() ([]byte, error) {
	type embed BodyAddSharedVoiceV1VoicesAddPublicUserIdVoiceIdPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
