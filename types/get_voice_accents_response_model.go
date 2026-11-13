package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetVoiceAccentsResponseModel struct {
	// A list of available voice accents.
	Accents         []*VoiceAccentResponseModel `json:"accents" xml:"accents" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetVoiceAccentsResponseModel) GetAccents() (value []*VoiceAccentResponseModel) {
	if g == nil {
		return
	}
	return g.Accents
}

func (g *GetVoiceAccentsResponseModel) SetAccents(accents []*VoiceAccentResponseModel) {
	g.Accents = accents
}

func (g *GetVoiceAccentsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetVoiceAccentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetVoiceAccentsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetVoiceAccentsResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler GetVoiceAccentsResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetVoiceAccentsResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetVoiceAccentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetVoiceAccentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
