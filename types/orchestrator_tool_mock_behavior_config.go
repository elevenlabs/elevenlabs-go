package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Orchestrator-side config: tools are identified by resolved names.
type OrchestratorToolMockBehaviorConfig struct {
	MockingStrategy  *MockingStrategy     `json:"mocking_strategy,omitempty" xml:"mocking_strategy,omitempty"`
	FallbackStrategy *MockNoMatchBehavior `json:"fallback_strategy,omitempty" xml:"fallback_strategy,omitempty"`
	// Tool names to mock. Only used when mocking_strategy is 'selected'.
	MockedToolNames []string `json:"mocked_tool_names,omitempty" xml:"mocked_tool_names,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OrchestratorToolMockBehaviorConfig) GetMockingStrategy() (value *MockingStrategy) {
	if o == nil {
		return
	}
	return o.MockingStrategy
}

func (o *OrchestratorToolMockBehaviorConfig) GetFallbackStrategy() (value *MockNoMatchBehavior) {
	if o == nil {
		return
	}
	return o.FallbackStrategy
}

func (o *OrchestratorToolMockBehaviorConfig) GetMockedToolNames() (value []string) {
	if o == nil {
		return
	}
	return o.MockedToolNames
}

func (o *OrchestratorToolMockBehaviorConfig) SetMockingStrategy(mockingStrategy *MockingStrategy) {
	o.MockingStrategy = mockingStrategy
}

func (o *OrchestratorToolMockBehaviorConfig) SetFallbackStrategy(fallbackStrategy *MockNoMatchBehavior) {
	o.FallbackStrategy = fallbackStrategy
}

func (o *OrchestratorToolMockBehaviorConfig) SetMockedToolNames(mockedToolNames []string) {
	o.MockedToolNames = mockedToolNames
}

func (o *OrchestratorToolMockBehaviorConfig) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrchestratorToolMockBehaviorConfig) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrchestratorToolMockBehaviorConfig to string"
	}
	return string(jsonData)
}

func (o *OrchestratorToolMockBehaviorConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler OrchestratorToolMockBehaviorConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrchestratorToolMockBehaviorConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OrchestratorToolMockBehaviorConfig) MarshalJSON() ([]byte, error) {
	type embed OrchestratorToolMockBehaviorConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
