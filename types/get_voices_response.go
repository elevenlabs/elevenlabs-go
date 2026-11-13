package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetVoicesResponse struct {
	// A list of available voices.
	Voices          []*Voice `json:"voices" xml:"voices" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetVoicesResponse) GetVoices() (value []*Voice) {
	if g == nil {
		return
	}
	return g.Voices
}

func (g *GetVoicesResponse) SetVoices(voices []*Voice) {
	g.Voices = voices
}

func (g *GetVoicesResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetVoicesResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetVoicesResponse to string"
	}
	return string(jsonData)
}

func (g *GetVoicesResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetVoicesResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetVoicesResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetVoicesResponse) MarshalJSON() ([]byte, error) {
	type embed GetVoicesResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
