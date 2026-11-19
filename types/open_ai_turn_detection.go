package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiTurnDetection struct {
	Type              *TurnDetectionType `json:"type,omitempty" xml:"type,omitempty"`
	Eagerness         *Eagerness         `json:"eagerness,omitempty" xml:"eagerness,omitempty"`
	CreateResponse    *bool              `json:"create_response,omitempty" xml:"create_response,omitempty"`
	InterruptResponse *bool              `json:"interrupt_response,omitempty" xml:"interrupt_response,omitempty"`
	Threshold         *float64           `json:"threshold,omitempty" xml:"threshold,omitempty"`
	PrefixPaddingMs   *int               `json:"prefix_padding_ms,omitempty" xml:"prefix_padding_ms,omitempty"`
	SilenceDurationMs *int               `json:"silence_duration_ms,omitempty" xml:"silence_duration_ms,omitempty"`
	TurnTimeout       *float64           `json:"turn_timeout,omitempty" xml:"turn_timeout,omitempty"`
	extraProperties   map[string]interface{}
}

func (o *OpenAiTurnDetection) GetType() (value *TurnDetectionType) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OpenAiTurnDetection) GetEagerness() (value *Eagerness) {
	if o == nil {
		return
	}
	return o.Eagerness
}

func (o *OpenAiTurnDetection) GetCreateResponse() (value *bool) {
	if o == nil {
		return
	}
	return o.CreateResponse
}

func (o *OpenAiTurnDetection) GetInterruptResponse() (value *bool) {
	if o == nil {
		return
	}
	return o.InterruptResponse
}

func (o *OpenAiTurnDetection) GetThreshold() (value *float64) {
	if o == nil {
		return
	}
	return o.Threshold
}

func (o *OpenAiTurnDetection) GetPrefixPaddingMs() (value *int) {
	if o == nil {
		return
	}
	return o.PrefixPaddingMs
}

func (o *OpenAiTurnDetection) GetSilenceDurationMs() (value *int) {
	if o == nil {
		return
	}
	return o.SilenceDurationMs
}

func (o *OpenAiTurnDetection) GetTurnTimeout() (value *float64) {
	if o == nil {
		return
	}
	return o.TurnTimeout
}

func (o *OpenAiTurnDetection) SetType(type_ *TurnDetectionType) {
	o.Type = type_
}

func (o *OpenAiTurnDetection) SetEagerness(eagerness *Eagerness) {
	o.Eagerness = eagerness
}

func (o *OpenAiTurnDetection) SetCreateResponse(createResponse *bool) {
	o.CreateResponse = createResponse
}

func (o *OpenAiTurnDetection) SetInterruptResponse(interruptResponse *bool) {
	o.InterruptResponse = interruptResponse
}

func (o *OpenAiTurnDetection) SetThreshold(threshold *float64) {
	o.Threshold = threshold
}

func (o *OpenAiTurnDetection) SetPrefixPaddingMs(prefixPaddingMs *int) {
	o.PrefixPaddingMs = prefixPaddingMs
}

func (o *OpenAiTurnDetection) SetSilenceDurationMs(silenceDurationMs *int) {
	o.SilenceDurationMs = silenceDurationMs
}

func (o *OpenAiTurnDetection) SetTurnTimeout(turnTimeout *float64) {
	o.TurnTimeout = turnTimeout
}

func (o *OpenAiTurnDetection) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiTurnDetection) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiTurnDetection to string"
	}
	return string(jsonData)
}

func (o *OpenAiTurnDetection) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiTurnDetection) MarshalJSON() ([]byte, error) {
	type embed OpenAiTurnDetection
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
