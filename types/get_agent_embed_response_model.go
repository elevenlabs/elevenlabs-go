package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentEmbedResponseModel struct {
	AgentId         string                `json:"agent_id" xml:"agent_id" required:"true"`
	WidgetConfig    *WidgetConfigResponse `json:"widget_config" xml:"widget_config" required:"true"`
	extraProperties map[string]interface{}
}

func (g *GetAgentEmbedResponseModel) GetAgentId() (value string) {
	if g == nil {
		return
	}
	return g.AgentId
}

func (g *GetAgentEmbedResponseModel) GetWidgetConfig() (value *WidgetConfigResponse) {
	if g == nil {
		return
	}
	return g.WidgetConfig
}

func (g *GetAgentEmbedResponseModel) SetAgentId(agentId string) {
	g.AgentId = agentId
}

func (g *GetAgentEmbedResponseModel) SetWidgetConfig(widgetConfig *WidgetConfigResponse) {
	g.WidgetConfig = widgetConfig
}

func (g *GetAgentEmbedResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentEmbedResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentEmbedResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentEmbedResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, g); err != nil {
		return err
	}
	type alias GetAgentEmbedResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*g = GetAgentEmbedResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAgentEmbedResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentEmbedResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
