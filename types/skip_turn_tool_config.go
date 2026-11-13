package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Allows the agent to explicitly skip its turn.
//
// This tool should be invoked by the LLM when the user indicates they would like
// to think or take a short pause before continuing the conversation—e.g. when
// they say: "Give me a second", "Let me think", or "One moment please".  After
// calling this tool, the assistant should not speak until the user speaks
// again, or if wait_timeout_secs is set, until that wait elapses and the
// agent generates a check-in.
type SkipTurnToolConfig struct {
	// Seconds to wait after skip_turn before the agent generates a contextual check-in. The "End conversation after silence" timer is paused during the wait. -1 disables the wait: after skip_turn the agent stays silent until the caller speaks, and that timer keeps running. Applies to voice conversations only.
	WaitTimeoutSecs *float64 `json:"wait_timeout_secs,omitempty" xml:"wait_timeout_secs,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SkipTurnToolConfig) GetWaitTimeoutSecs() (value *float64) {
	if s == nil {
		return
	}
	return s.WaitTimeoutSecs
}

func (s *SkipTurnToolConfig) SetWaitTimeoutSecs(waitTimeoutSecs *float64) {
	s.WaitTimeoutSecs = waitTimeoutSecs
}

func (s *SkipTurnToolConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SkipTurnToolConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SkipTurnToolConfig to string"
	}
	return string(jsonData)
}

func (s *SkipTurnToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler SkipTurnToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SkipTurnToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "system_tool_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SkipTurnToolConfig) MarshalJSON() ([]byte, error) {
	type embed SkipTurnToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
