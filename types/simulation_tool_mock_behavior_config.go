package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Simulation/preview-side config: tools are identified by IDs, resolved to names at runtime.
type SimulationToolMockBehaviorConfig struct {
	MockingStrategy  *MockingStrategy     `json:"mocking_strategy,omitempty" xml:"mocking_strategy,omitempty"`
	FallbackStrategy *MockNoMatchBehavior `json:"fallback_strategy,omitempty" xml:"fallback_strategy,omitempty"`
	// Tool IDs to mock. Resolved to tool names before being passed to the orchestrator.
	MockedToolIds   []string `json:"mocked_tool_ids,omitempty" xml:"mocked_tool_ids,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SimulationToolMockBehaviorConfig) GetMockingStrategy() (value *MockingStrategy) {
	if s == nil {
		return
	}
	return s.MockingStrategy
}

func (s *SimulationToolMockBehaviorConfig) GetFallbackStrategy() (value *MockNoMatchBehavior) {
	if s == nil {
		return
	}
	return s.FallbackStrategy
}

func (s *SimulationToolMockBehaviorConfig) GetMockedToolIds() (value []string) {
	if s == nil {
		return
	}
	return s.MockedToolIds
}

func (s *SimulationToolMockBehaviorConfig) SetMockingStrategy(mockingStrategy *MockingStrategy) {
	s.MockingStrategy = mockingStrategy
}

func (s *SimulationToolMockBehaviorConfig) SetFallbackStrategy(fallbackStrategy *MockNoMatchBehavior) {
	s.FallbackStrategy = fallbackStrategy
}

func (s *SimulationToolMockBehaviorConfig) SetMockedToolIds(mockedToolIds []string) {
	s.MockedToolIds = mockedToolIds
}

func (s *SimulationToolMockBehaviorConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SimulationToolMockBehaviorConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SimulationToolMockBehaviorConfig to string"
	}
	return string(jsonData)
}

func (s *SimulationToolMockBehaviorConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SimulationToolMockBehaviorConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SimulationToolMockBehaviorConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SimulationToolMockBehaviorConfig) MarshalJSON() ([]byte, error) {
	type embed SimulationToolMockBehaviorConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
