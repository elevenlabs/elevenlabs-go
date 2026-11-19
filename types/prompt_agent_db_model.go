package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PromptAgentDbModel struct {
	Tools           any `json:"tools,omitempty" xml:"tools,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PromptAgentDbModel) GetTools() (value any) {
	if p == nil {
		return
	}
	return p.Tools
}

func (p *PromptAgentDbModel) SetTools(tools any) {
	p.Tools = tools
}

func (p *PromptAgentDbModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PromptAgentDbModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PromptAgentDbModel to string"
	}
	return string(jsonData)
}

func (p *PromptAgentDbModel) UnmarshalJSON(data []byte) error {
	type unmarshaler PromptAgentDbModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PromptAgentDbModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PromptAgentDbModel) MarshalJSON() ([]byte, error) {
	type embed PromptAgentDbModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
