package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AddVoiceToUserVoiceCollectionResponseModel struct {
	// The ID of the voice collection.
	CollectionId string `json:"collection_id" xml:"collection_id" required:"true"`
	// The ID of the voice that was added.
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) GetCollectionId() (value string) {
	if a == nil {
		return
	}
	return a.CollectionId
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) GetVoiceId() (value string) {
	if a == nil {
		return
	}
	return a.VoiceId
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) SetCollectionId(collectionId string) {
	a.CollectionId = collectionId
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) SetVoiceId(voiceId string) {
	a.VoiceId = voiceId
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AddVoiceToUserVoiceCollectionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddVoiceToUserVoiceCollectionResponseModel to string"
	}
	return string(jsonData)
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddVoiceToUserVoiceCollectionResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddVoiceToUserVoiceCollectionResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AddVoiceToUserVoiceCollectionResponseModel) MarshalJSON() ([]byte, error) {
	type embed AddVoiceToUserVoiceCollectionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
