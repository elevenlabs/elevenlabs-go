package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentKnowledgebaseSizeResponseModel struct {
	NumberOfPages   float64 `json:"number_of_pages" xml:"number_of_pages" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetAgentKnowledgebaseSizeResponseModel) GetNumberOfPages() (value float64) {
	if g == nil {
		return
	}
	return g.NumberOfPages
}

func (g *GetAgentKnowledgebaseSizeResponseModel) SetNumberOfPages(numberOfPages float64) {
	g.NumberOfPages = numberOfPages
}

func (g *GetAgentKnowledgebaseSizeResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentKnowledgebaseSizeResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentKnowledgebaseSizeResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentKnowledgebaseSizeResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetAgentKnowledgebaseSizeResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetAgentKnowledgebaseSizeResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAgentKnowledgebaseSizeResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentKnowledgebaseSizeResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
