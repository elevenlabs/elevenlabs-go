package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ToolResponseModel struct {
	Id string `json:"id" xml:"id" required:"true"`
	// The type of tool
	ToolConfig *ToolResponseModelToolConfig `json:"tool_config" xml:"tool_config" required:"true"`
	AccessInfo *ResourceAccessInfo          `json:"access_info" xml:"access_info" required:"true"`
	UsageStats *ToolUsageStatsResponseModel `json:"usage_stats" xml:"usage_stats" required:"true"`
	// Mock responses with optional parameter conditions. Evaluated top-to-bottom; first match wins.
	ResponseMocks   []*ToolResponseMockConfigOutput `json:"response_mocks,omitempty" xml:"response_mocks,omitempty"`
	extraProperties map[string]interface{}
}

func (t *ToolResponseModel) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *ToolResponseModel) GetToolConfig() (value *ToolResponseModelToolConfig) {
	if t == nil {
		return
	}
	return t.ToolConfig
}

func (t *ToolResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if t == nil {
		return
	}
	return t.AccessInfo
}

func (t *ToolResponseModel) GetUsageStats() (value *ToolUsageStatsResponseModel) {
	if t == nil {
		return
	}
	return t.UsageStats
}

func (t *ToolResponseModel) GetResponseMocks() (value []*ToolResponseMockConfigOutput) {
	if t == nil {
		return
	}
	return t.ResponseMocks
}

func (t *ToolResponseModel) SetId(id string) {
	t.Id = id
}

func (t *ToolResponseModel) SetToolConfig(toolConfig *ToolResponseModelToolConfig) {
	t.ToolConfig = toolConfig
}

func (t *ToolResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	t.AccessInfo = accessInfo
}

func (t *ToolResponseModel) SetUsageStats(usageStats *ToolUsageStatsResponseModel) {
	t.UsageStats = usageStats
}

func (t *ToolResponseModel) SetResponseMocks(responseMocks []*ToolResponseMockConfigOutput) {
	t.ResponseMocks = responseMocks
}

func (t *ToolResponseModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolResponseModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolResponseModel to string"
	}
	return string(jsonData)
}

func (t *ToolResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolResponseModel) MarshalJSON() ([]byte, error) {
	type embed ToolResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
