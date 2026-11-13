package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneChannelConfigInput struct {
	// The phone number to call from.
	AgentPhoneNumberId string                 `json:"agent_phone_number_id" xml:"agent_phone_number_id" required:"true" maxLength:"256" minLength:"1"`
	ModeConfig         *PhoneVoiceConfigInput `json:"mode_config,omitempty" xml:"mode_config,omitempty"`
	extraProperties    map[string]interface{}
}

func (p *PhoneChannelConfigInput) GetAgentPhoneNumberId() (value string) {
	if p == nil {
		return
	}
	return p.AgentPhoneNumberId
}

func (p *PhoneChannelConfigInput) GetModeConfig() (value *PhoneVoiceConfigInput) {
	if p == nil {
		return
	}
	return p.ModeConfig
}

func (p *PhoneChannelConfigInput) SetAgentPhoneNumberId(agentPhoneNumberId string) {
	p.AgentPhoneNumberId = agentPhoneNumberId
}

func (p *PhoneChannelConfigInput) SetModeConfig(modeConfig *PhoneVoiceConfigInput) {
	p.ModeConfig = modeConfig
}

func (p *PhoneChannelConfigInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneChannelConfigInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneChannelConfigInput to string"
	}
	return string(jsonData)
}

func (p *PhoneChannelConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneChannelConfigInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneChannelConfigInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "channel")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneChannelConfigInput) MarshalJSON() ([]byte, error) {
	type embed PhoneChannelConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
