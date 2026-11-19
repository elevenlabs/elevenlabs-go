package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Allows the agent to play DTMF tones during a phone call.
//
// This tool can be used to interact with automated phone systems, such as
// navigating phone menus, entering extensions, or inputting numeric codes.
type PlayDtmfToolConfig struct {
	// Send DTMF tones as out-of-band RTP events (RFC 4733) instead of in-band audio. Only effective for SIP trunk imported numbers.
	UseOutOfBandDtmf *bool `json:"use_out_of_band_dtmf,omitempty" xml:"use_out_of_band_dtmf,omitempty"`
	// If true, the agent will not generate further speech after playing DTMF tones. This prevents the agent's speech from interfering with IVR systems.
	SuppressTurnAfterDtmf *bool `json:"suppress_turn_after_dtmf,omitempty" xml:"suppress_turn_after_dtmf,omitempty"`
	extraProperties       map[string]interface{}
}

func (p *PlayDtmfToolConfig) GetUseOutOfBandDtmf() (value *bool) {
	if p == nil {
		return
	}
	return p.UseOutOfBandDtmf
}

func (p *PlayDtmfToolConfig) GetSuppressTurnAfterDtmf() (value *bool) {
	if p == nil {
		return
	}
	return p.SuppressTurnAfterDtmf
}

func (p *PlayDtmfToolConfig) SetUseOutOfBandDtmf(useOutOfBandDtmf *bool) {
	p.UseOutOfBandDtmf = useOutOfBandDtmf
}

func (p *PlayDtmfToolConfig) SetSuppressTurnAfterDtmf(suppressTurnAfterDtmf *bool) {
	p.SuppressTurnAfterDtmf = suppressTurnAfterDtmf
}

func (p *PlayDtmfToolConfig) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PlayDtmfToolConfig) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PlayDtmfToolConfig to string"
	}
	return string(jsonData)
}

func (p *PlayDtmfToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler PlayDtmfToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PlayDtmfToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "system_tool_type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PlayDtmfToolConfig) MarshalJSON() ([]byte, error) {
	type embed PlayDtmfToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
