package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PhoneChannelConfigOutput struct {
	// The phone number to call from.
	AgentPhoneNumberId string                  `json:"agent_phone_number_id" xml:"agent_phone_number_id" required:"true" maxLength:"256" minLength:"1"`
	ModeConfig         *PhoneVoiceConfigOutput `json:"mode_config,omitempty" xml:"mode_config,omitempty"`
	extraProperties    map[string]interface{}
}

func (p *PhoneChannelConfigOutput) GetAgentPhoneNumberId() (value string) {
	if p == nil {
		return
	}
	return p.AgentPhoneNumberId
}

func (p *PhoneChannelConfigOutput) GetModeConfig() (value *PhoneVoiceConfigOutput) {
	if p == nil {
		return
	}
	return p.ModeConfig
}

func (p *PhoneChannelConfigOutput) SetAgentPhoneNumberId(agentPhoneNumberId string) {
	p.AgentPhoneNumberId = agentPhoneNumberId
}

func (p *PhoneChannelConfigOutput) SetModeConfig(modeConfig *PhoneVoiceConfigOutput) {
	p.ModeConfig = modeConfig
}

func (p *PhoneChannelConfigOutput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PhoneChannelConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PhoneChannelConfigOutput to string"
	}
	return string(jsonData)
}

func (p *PhoneChannelConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PhoneChannelConfigOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PhoneChannelConfigOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "channel")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PhoneChannelConfigOutput) MarshalJSON() ([]byte, error) {
	type embed PhoneChannelConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
