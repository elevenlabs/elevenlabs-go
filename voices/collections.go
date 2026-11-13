package voices

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreateVoiceCollectionV1VoicesCollectionsPost struct {
	// Title of the collection to create/update.
	Title string `json:"title" xml:"title" required:"true" maxLength:"1000"`
	// Icon of the collection to create/update.
	Icon string `json:"icon" xml:"icon" required:"true" maxLength:"1000"`
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) GetTitle() (value string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) GetIcon() (value string) {
	if b == nil {
		return
	}
	return b.Icon
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) SetTitle(title string) {
	b.Title = title
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) SetIcon(icon string) {
	b.Icon = icon
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateVoiceCollectionV1VoicesCollectionsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateVoiceCollectionV1VoicesCollectionsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyCreateVoiceCollectionV1VoicesCollectionsPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyCreateVoiceCollectionV1VoicesCollectionsPost(tmp)
	return nil
}

func (b *BodyCreateVoiceCollectionV1VoicesCollectionsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateVoiceCollectionV1VoicesCollectionsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch struct {
	// Collection ID.
	CollectionId string `json:"-" url:"-"`
	// Title of the collection to create/update.
	Title *string `json:"title,omitempty" xml:"title,omitempty" maxLength:"1000"`
	// Icon of the collection to create/update.
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty" maxLength:"1000"`
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) GetCollectionId() (value string) {
	if b == nil {
		return
	}
	return b.CollectionId
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) GetIcon() (value *string) {
	if b == nil {
		return
	}
	return b.Icon
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) SetCollectionId(collectionId string) {
	b.CollectionId = collectionId
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) SetTitle(title *string) {
	b.Title = title
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) SetIcon(icon *string) {
	b.Icon = icon
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateVoiceCollectionV1VoicesCollectionsCollectionIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
