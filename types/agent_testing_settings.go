package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Settings for agent testing configuration.
type AgentTestingSettings struct {
	// List of test IDs that should be run for this agent
	AttachedTests   []*AttachedTestModel `json:"attached_tests,omitempty" xml:"attached_tests,omitempty" maxItems:"5000"`
	extraProperties map[string]interface{}
}

func (a *AgentTestingSettings) GetAttachedTests() (value []*AttachedTestModel) {
	if a == nil {
		return
	}
	return a.AttachedTests
}

func (a *AgentTestingSettings) SetAttachedTests(attachedTests []*AttachedTestModel) {
	a.AttachedTests = attachedTests
}

func (a *AgentTestingSettings) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentTestingSettings) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentTestingSettings to string"
	}
	return string(jsonData)
}

func (a *AgentTestingSettings) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentTestingSettings
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentTestingSettings(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentTestingSettings) MarshalJSON() ([]byte, error) {
	type embed AgentTestingSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
