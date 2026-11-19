package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentQueueingConfig struct {
	// Hold callers in a wait queue when the agent is at its concurrency limit, instead of rejecting them immediately. New agents are created with queueing enabled unless this field is set explicitly.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Maximum time a caller can wait in the queue before being rejected
	WaitTimeoutSeconds *int `json:"wait_timeout_seconds,omitempty" xml:"wait_timeout_seconds,omitempty" max:"1800"`
	// Custom hold audio played to queued callers; when unset, callers hear the default hold tone. Read-only: set it by uploading a file through the agent hold-audio endpoint.
	HoldAudio       *AgentHoldAudioConfig `json:"hold_audio,omitempty" xml:"hold_audio,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentQueueingConfig) GetEnabled() (value *bool) {
	if a == nil {
		return
	}
	return a.Enabled
}

func (a *AgentQueueingConfig) GetWaitTimeoutSeconds() (value *int) {
	if a == nil {
		return
	}
	return a.WaitTimeoutSeconds
}

func (a *AgentQueueingConfig) GetHoldAudio() (value *AgentHoldAudioConfig) {
	if a == nil {
		return
	}
	return a.HoldAudio
}

func (a *AgentQueueingConfig) SetEnabled(enabled *bool) {
	a.Enabled = enabled
}

func (a *AgentQueueingConfig) SetWaitTimeoutSeconds(waitTimeoutSeconds *int) {
	a.WaitTimeoutSeconds = waitTimeoutSeconds
}

func (a *AgentQueueingConfig) SetHoldAudio(holdAudio *AgentHoldAudioConfig) {
	a.HoldAudio = holdAudio
}

func (a *AgentQueueingConfig) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentQueueingConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentQueueingConfig to string"
	}
	return string(jsonData)
}

func (a *AgentQueueingConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentQueueingConfig) MarshalJSON() ([]byte, error) {
	type embed AgentQueueingConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
