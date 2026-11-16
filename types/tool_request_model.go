package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolRequestModel struct {
	// Configuration for the tool
	ToolConfig *ToolRequestModelToolConfig `json:"tool_config" xml:"tool_config" required:"true"`
	// Mock responses with optional parameter conditions. Evaluated top-to-bottom; first match wins.
	ResponseMocks []*ToolResponseMockConfigInput `json:"response_mocks,omitempty" xml:"response_mocks,omitempty"`
}

func (t *ToolRequestModel) GetToolConfig() (value *ToolRequestModelToolConfig) {
	if t == nil {
		return
	}
	return t.ToolConfig
}

func (t *ToolRequestModel) GetResponseMocks() (value []*ToolResponseMockConfigInput) {
	if t == nil {
		return
	}
	return t.ResponseMocks
}

func (t *ToolRequestModel) SetToolConfig(toolConfig *ToolRequestModelToolConfig) {
	t.ToolConfig = toolConfig
}

func (t *ToolRequestModel) SetResponseMocks(responseMocks []*ToolResponseMockConfigInput) {
	t.ResponseMocks = responseMocks
}

func (t *ToolRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (t ToolRequestModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolRequestModel to string"
	}
	return string(jsonData)
}

func (t *ToolRequestModel) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, t)
}

func (t *ToolRequestModel) MarshalJSON() ([]byte, error) {
	type embed ToolRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
