package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RemoveVoiceFromUserVoiceCollectionResponseModel struct {
	// The ID of the voice collection.
	CollectionId string `json:"collection_id" xml:"collection_id" required:"true"`
	// The ID of the voice that was removed.
	VoiceId         string `json:"voice_id" xml:"voice_id" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) GetCollectionId() (value string) {
	if r == nil {
		return
	}
	return r.CollectionId
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) GetVoiceId() (value string) {
	if r == nil {
		return
	}
	return r.VoiceId
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) SetCollectionId(collectionId string) {
	r.CollectionId = collectionId
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) SetVoiceId(voiceId string) {
	r.VoiceId = voiceId
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RemoveVoiceFromUserVoiceCollectionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RemoveVoiceFromUserVoiceCollectionResponseModel to string"
	}
	return string(jsonData)
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RemoveVoiceFromUserVoiceCollectionResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RemoveVoiceFromUserVoiceCollectionResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RemoveVoiceFromUserVoiceCollectionResponseModel) MarshalJSON() ([]byte, error) {
	type embed RemoveVoiceFromUserVoiceCollectionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
