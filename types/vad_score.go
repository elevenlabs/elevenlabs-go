package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VadScore struct {
	VadScoreEvent   *VadScoreVadScoreEvent `json:"vad_score_event" xml:"vad_score_event" required:"true"`
	Type            *string                `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VadScore) GetVadScoreEvent() (value *VadScoreVadScoreEvent) {
	if v == nil {
		return
	}
	return v.VadScoreEvent
}

func (v *VadScore) GetType() (value *string) {
	if v == nil {
		return
	}
	return v.Type
}

func (v *VadScore) SetVadScoreEvent(vadScoreEvent *VadScoreVadScoreEvent) {
	v.VadScoreEvent = vadScoreEvent
}

func (v *VadScore) SetType(type_ *string) {
	v.Type = type_
}

func (v *VadScore) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VadScore) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VadScore to string"
	}
	return string(jsonData)
}

func (v *VadScore) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VadScore
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VadScore(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VadScore) MarshalJSON() ([]byte, error) {
	type embed VadScore
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
