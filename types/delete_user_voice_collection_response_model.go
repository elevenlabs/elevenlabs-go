package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DeleteUserVoiceCollectionResponseModel struct {
	// The ID of the deleted voice collection.
	CollectionId    string `json:"collection_id" xml:"collection_id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DeleteUserVoiceCollectionResponseModel) GetCollectionId() (value string) {
	if d == nil {
		return
	}
	return d.CollectionId
}

func (d *DeleteUserVoiceCollectionResponseModel) SetCollectionId(collectionId string) {
	d.CollectionId = collectionId
}

func (d *DeleteUserVoiceCollectionResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteUserVoiceCollectionResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteUserVoiceCollectionResponseModel to string"
	}
	return string(jsonData)
}

func (d *DeleteUserVoiceCollectionResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DeleteUserVoiceCollectionResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DeleteUserVoiceCollectionResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteUserVoiceCollectionResponseModel) MarshalJSON() ([]byte, error) {
	type embed DeleteUserVoiceCollectionResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
