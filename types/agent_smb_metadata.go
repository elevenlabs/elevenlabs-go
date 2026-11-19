package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Identifies an agent as part of the SMB product.
type AgentSmbMetadata struct {
	AgentType SmbAgentType `json:"agent_type" xml:"agent_type" required:"true"`
	// True for pre-signup onboarding agents that have not yet been transferred to a user workspace.
	IsEphemeral *bool `json:"is_ephemeral,omitempty" xml:"is_ephemeral,omitempty"`
	// Hash of the SMB data+code inputs this agent was last regenerated against; the conversation-start drift gate. None (unstamped) regenerates once on the next conversation.
	SourceFingerprint *string `json:"source_fingerprint,omitempty" xml:"source_fingerprint,omitempty"`
	extraProperties   map[string]interface{}
}

func (a *AgentSmbMetadata) GetAgentType() (value SmbAgentType) {
	if a == nil {
		return
	}
	return a.AgentType
}

func (a *AgentSmbMetadata) GetIsEphemeral() (value *bool) {
	if a == nil {
		return
	}
	return a.IsEphemeral
}

func (a *AgentSmbMetadata) GetSourceFingerprint() (value *string) {
	if a == nil {
		return
	}
	return a.SourceFingerprint
}

func (a *AgentSmbMetadata) SetAgentType(agentType SmbAgentType) {
	a.AgentType = agentType
}

func (a *AgentSmbMetadata) SetIsEphemeral(isEphemeral *bool) {
	a.IsEphemeral = isEphemeral
}

func (a *AgentSmbMetadata) SetSourceFingerprint(sourceFingerprint *string) {
	a.SourceFingerprint = sourceFingerprint
}

func (a *AgentSmbMetadata) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentSmbMetadata) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentSmbMetadata to string"
	}
	return string(jsonData)
}

func (a *AgentSmbMetadata) UnmarshalJSON(data []byte) error {
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

func (a *AgentSmbMetadata) MarshalJSON() ([]byte, error) {
	type embed AgentSmbMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
