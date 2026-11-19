package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost struct {
	// URL of the page to extract content from.
	Url string `json:"url" xml:"url" required:"true"`
	// Author used in the player and inserted at the start of the uploaded article. If not provided, the default author set in the Player settings is used.
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// Title used in the player and inserted at the top of the uploaded article. If not provided, the default title set in the Player settings is used.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) GetUrl() (value string) {
	if b == nil {
		return
	}
	return b.Url
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) GetAuthor() (value *string) {
	if b == nil {
		return
	}
	return b.Author
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) SetUrl(url string) {
	b.Url = url
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) SetAuthor(author *string) {
	b.Author = author
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) SetTitle(title *string) {
	b.Title = title
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateAudioNativeContentFromUrlV1AudioNativeContentPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
