package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	big "math/big"
)

type UpdateWhatsAppAccountRequest struct {
	PhoneNumberId              string  `json:"-" url:"-"`
	AssignedAgentId            *string `json:"assigned_agent_id,omitempty" xml:"assigned_agent_id,omitempty"`
	EnableMessaging            *bool   `json:"enable_messaging,omitempty" xml:"enable_messaging,omitempty"`
	EnableAudioMessageResponse *bool   `json:"enable_audio_message_response,omitempty" xml:"enable_audio_message_response,omitempty"`
	EnableTypingIndicator      *bool   `json:"enable_typing_indicator,omitempty" xml:"enable_typing_indicator,omitempty"`
}

func (u *UpdateWhatsAppAccountRequest) GetPhoneNumberId() (value string) {
	if u == nil {
		return
	}
	return u.PhoneNumberId
}

func (u *UpdateWhatsAppAccountRequest) GetAssignedAgentId() (value *string) {
	if u == nil {
		return
	}
	return u.AssignedAgentId
}

func (u *UpdateWhatsAppAccountRequest) GetEnableMessaging() (value *bool) {
	if u == nil {
		return
	}
	return u.EnableMessaging
}

func (u *UpdateWhatsAppAccountRequest) GetEnableAudioMessageResponse() (value *bool) {
	if u == nil {
		return
	}
	return u.EnableAudioMessageResponse
}

func (u *UpdateWhatsAppAccountRequest) GetEnableTypingIndicator() (value *bool) {
	if u == nil {
		return
	}
	return u.EnableTypingIndicator
}

func (u *UpdateWhatsAppAccountRequest) SetPhoneNumberId(phoneNumberId string) {
	u.PhoneNumberId = phoneNumberId
}

func (u *UpdateWhatsAppAccountRequest) SetAssignedAgentId(assignedAgentId *string) {
	u.AssignedAgentId = assignedAgentId
}

func (u *UpdateWhatsAppAccountRequest) SetEnableMessaging(enableMessaging *bool) {
	u.EnableMessaging = enableMessaging
}

func (u *UpdateWhatsAppAccountRequest) SetEnableAudioMessageResponse(enableAudioMessageResponse *bool) {
	u.EnableAudioMessageResponse = enableAudioMessageResponse
}

func (u *UpdateWhatsAppAccountRequest) SetEnableTypingIndicator(enableTypingIndicator *bool) {
	u.EnableTypingIndicator = enableTypingIndicator
}

func (u *UpdateWhatsAppAccountRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateWhatsAppAccountRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateWhatsAppAccountRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateWhatsAppAccountRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateWhatsAppAccountRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateWhatsAppAccountRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

var (
	listWhatsappAccountsRequestFieldAgentId = big.NewInt(1 << 0)
)

type ListWhatsappAccountsRequest struct {
	AgentId *string `json:"-" url:"agent_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListWhatsappAccountsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListWhatsappAccountsRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listWhatsappAccountsRequestFieldAgentId)
}
