package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetLiveCountResponse struct {
	// The number of active ongoing conversations.
	Count           int `json:"count" xml:"count" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetLiveCountResponse) GetCount() (value int) {
	if g == nil {
		return
	}
	return g.Count
}

func (g *GetLiveCountResponse) SetCount(count int) {
	g.Count = count
}

func (g *GetLiveCountResponse) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetLiveCountResponse) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetLiveCountResponse to string"
	}
	return string(jsonData)
}

func (g *GetLiveCountResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetLiveCountResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetLiveCountResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetLiveCountResponse) MarshalJSON() ([]byte, error) {
	type embed GetLiveCountResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
