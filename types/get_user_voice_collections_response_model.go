package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetUserVoiceCollectionsResponseModel struct {
	// The voice collections the caller can access, sorted by title.
	Collections     []*UserVoiceCollectionResponseModel `json:"collections" xml:"collections" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetUserVoiceCollectionsResponseModel) GetCollections() (value []*UserVoiceCollectionResponseModel) {
	if g == nil {
		return
	}
	return g.Collections
}

func (g *GetUserVoiceCollectionsResponseModel) SetCollections(collections []*UserVoiceCollectionResponseModel) {
	g.Collections = collections
}

func (g *GetUserVoiceCollectionsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetUserVoiceCollectionsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetUserVoiceCollectionsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetUserVoiceCollectionsResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler GetUserVoiceCollectionsResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetUserVoiceCollectionsResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetUserVoiceCollectionsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetUserVoiceCollectionsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
