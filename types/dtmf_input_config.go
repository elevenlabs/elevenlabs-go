package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for DTMF (keypad) input collection during phone calls.
type DtmfInputConfig struct {
	// Timeout in seconds to wait for additional DTMF digits
	DtmfInputTimeout *float64 `json:"dtmf_input_timeout,omitempty" xml:"dtmf_input_timeout,omitempty" min:"0.5" max:"10"`
	// If true, pressing # immediately completes DTMF input
	HashTerminator *bool `json:"hash_terminator,omitempty" xml:"hash_terminator,omitempty"`
	// If true, replace the caller's DTMF (keypad) entries with a redaction marker in the transcript, conversation log and analysis. Digits the agent repeats back or passes to a tool are not affected.
	RedactInput     *bool `json:"redact_input,omitempty" xml:"redact_input,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DtmfInputConfig) GetDtmfInputTimeout() (value *float64) {
	if d == nil {
		return
	}
	return d.DtmfInputTimeout
}

func (d *DtmfInputConfig) GetHashTerminator() (value *bool) {
	if d == nil {
		return
	}
	return d.HashTerminator
}

func (d *DtmfInputConfig) GetRedactInput() (value *bool) {
	if d == nil {
		return
	}
	return d.RedactInput
}

func (d *DtmfInputConfig) SetDtmfInputTimeout(dtmfInputTimeout *float64) {
	d.DtmfInputTimeout = dtmfInputTimeout
}

func (d *DtmfInputConfig) SetHashTerminator(hashTerminator *bool) {
	d.HashTerminator = hashTerminator
}

func (d *DtmfInputConfig) SetRedactInput(redactInput *bool) {
	d.RedactInput = redactInput
}

func (d *DtmfInputConfig) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DtmfInputConfig) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DtmfInputConfig to string"
	}
	return string(jsonData)
}

func (d *DtmfInputConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler DtmfInputConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DtmfInputConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DtmfInputConfig) MarshalJSON() ([]byte, error) {
	type embed DtmfInputConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
