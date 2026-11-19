package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A receptionist.
type ReceptionCustomerFacingAgentResponse struct {
	// ID of the ElevenLabs agent behind this receptionist.
	AgentId *string                              `json:"agent_id,omitempty" xml:"agent_id,omitempty"`
	Config  *ReceptionCustomerFacingConfigOutput `json:"config,omitempty" xml:"config,omitempty"`
	// Unix timestamp when this receptionist was deactivated. None when active.
	DeactivatedAtUnix *int `json:"deactivated_at_unix,omitempty" xml:"deactivated_at_unix,omitempty"`
	// Why this receptionist was deactivated. None when active.
	DeactivationReason *ReceptionCustomerFacingAgentDeactivationReason `json:"deactivation_reason,omitempty" xml:"deactivation_reason,omitempty"`
	// Display name of the receptionist.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// False when the receptionist's underlying agent was deleted; it cannot take calls until it is set up again.
	AgentExists *bool `json:"agent_exists,omitempty" xml:"agent_exists,omitempty"`
	// Whether this receptionist can configure and use transfer rules.
	IsTransferAllowed *bool `json:"is_transfer_allowed,omitempty" xml:"is_transfer_allowed,omitempty"`
	// How the website widget looks. Null when the agent is missing. Unset labels follow the receptionist's language.
	Widget *ReceptionWidgetAppearance `json:"widget,omitempty" xml:"widget,omitempty"`
	// Whether the receptionist is active. Inactive receptionists take no live calls.
	IsActive        *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionCustomerFacingAgentResponse) GetAgentId() (value *string) {
	if r == nil {
		return
	}
	return r.AgentId
}

func (r *ReceptionCustomerFacingAgentResponse) GetConfig() (value *ReceptionCustomerFacingConfigOutput) {
	if r == nil {
		return
	}
	return r.Config
}

func (r *ReceptionCustomerFacingAgentResponse) GetDeactivatedAtUnix() (value *int) {
	if r == nil {
		return
	}
	return r.DeactivatedAtUnix
}

func (r *ReceptionCustomerFacingAgentResponse) GetDeactivationReason() (value *ReceptionCustomerFacingAgentDeactivationReason) {
	if r == nil {
		return
	}
	return r.DeactivationReason
}

func (r *ReceptionCustomerFacingAgentResponse) GetName() (value *string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionCustomerFacingAgentResponse) GetAgentExists() (value *bool) {
	if r == nil {
		return
	}
	return r.AgentExists
}

func (r *ReceptionCustomerFacingAgentResponse) GetIsTransferAllowed() (value *bool) {
	if r == nil {
		return
	}
	return r.IsTransferAllowed
}

func (r *ReceptionCustomerFacingAgentResponse) GetWidget() (value *ReceptionWidgetAppearance) {
	if r == nil {
		return
	}
	return r.Widget
}

func (r *ReceptionCustomerFacingAgentResponse) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionCustomerFacingAgentResponse) SetAgentId(agentId *string) {
	r.AgentId = agentId
}

func (r *ReceptionCustomerFacingAgentResponse) SetConfig(config *ReceptionCustomerFacingConfigOutput) {
	r.Config = config
}

func (r *ReceptionCustomerFacingAgentResponse) SetDeactivatedAtUnix(deactivatedAtUnix *int) {
	r.DeactivatedAtUnix = deactivatedAtUnix
}

func (r *ReceptionCustomerFacingAgentResponse) SetDeactivationReason(deactivationReason *ReceptionCustomerFacingAgentDeactivationReason) {
	r.DeactivationReason = deactivationReason
}

func (r *ReceptionCustomerFacingAgentResponse) SetName(name *string) {
	r.Name = name
}

func (r *ReceptionCustomerFacingAgentResponse) SetAgentExists(agentExists *bool) {
	r.AgentExists = agentExists
}

func (r *ReceptionCustomerFacingAgentResponse) SetIsTransferAllowed(isTransferAllowed *bool) {
	r.IsTransferAllowed = isTransferAllowed
}

func (r *ReceptionCustomerFacingAgentResponse) SetWidget(widget *ReceptionWidgetAppearance) {
	r.Widget = widget
}

func (r *ReceptionCustomerFacingAgentResponse) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionCustomerFacingAgentResponse) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionCustomerFacingAgentResponse) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionCustomerFacingAgentResponse to string"
	}
	return string(jsonData)
}

func (r *ReceptionCustomerFacingAgentResponse) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionCustomerFacingAgentResponse) MarshalJSON() ([]byte, error) {
	type embed ReceptionCustomerFacingAgentResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
