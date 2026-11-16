package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A single customer-facing receptionist agent.
//
// A workspace can have multiple of these. Each entry owns its own ConvAI
// agent id and full “CustomerFacingConfig“ (voice, language, first message,
// system tool toggles, agent rules, transfer rules).
type ReceptionCustomerFacingAgent struct {
	// ID of the ElevenLabs agent behind this receptionist.
	AgentId *string                             `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Config  *ReceptionCustomerFacingConfigInput `json:"config,omitempty" xml:"config,omitempty"`
	// Unix timestamp when this receptionist was deactivated. None when active.
	DeactivatedAtUnix *int `json:"deactivated_at_unix,omitempty" xml:"deactivated_at_unix,omitempty"`
	// Why this receptionist was deactivated. None when active.
	DeactivationReason *ReceptionCustomerFacingAgentDeactivationReason `json:"deactivation_reason,omitempty" xml:"deactivation_reason,omitempty"`
	extraProperties    map[string]interface{}
}

func (r *ReceptionCustomerFacingAgent) GetAgentId() (value *string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionCustomerFacingAgent) GetConfig() (value *ReceptionCustomerFacingConfigInput) {
	if r == nil {
		return
	}
	return r.Config
}

func (r *ReceptionCustomerFacingAgent) GetDeactivatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.DeactivatedAtUnix
}

func (r *ReceptionCustomerFacingAgent) GetDeactivationReason() (value *ReceptionCustomerFacingAgentDeactivationReason) {
	if r == nil {
		return
	}
	return r.DeactivationReason
}

func (r *ReceptionCustomerFacingAgent) SetAgentId(agentId *string) {
	r.AgentId = agentId
}

func (r *ReceptionCustomerFacingAgent) SetConfig(config *ReceptionCustomerFacingConfigInput) {
	r.Config = config
}

func (r *ReceptionCustomerFacingAgent) SetDeactivatedAtUnix(deactivatedAtUnix *int) {
	r.DeactivatedAtUnix = deactivatedAtUnix
}

func (r *ReceptionCustomerFacingAgent) SetDeactivationReason(deactivationReason *ReceptionCustomerFacingAgentDeactivationReason) {
	r.DeactivationReason = deactivationReason
}

func (r *ReceptionCustomerFacingAgent) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionCustomerFacingAgent) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCustomerFacingAgent to string"
	}
	return string(jsonData)
}

func (r *ReceptionCustomerFacingAgent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionCustomerFacingAgent) MarshalJSON() ([]byte, error) {
	type embed ReceptionCustomerFacingAgent
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
