package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VadScoreVadScoreEvent struct {
	VadScore        float64 `json:"vad_score" xml:"vad_score" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VadScoreVadScoreEvent) GetVadScore() (value float64) {
	if v == nil {
		return
	}
	return v.VadScore
}

func (v *VadScoreVadScoreEvent) SetVadScore(vadScore float64) {
	v.VadScore = vadScore
}

func (v *VadScoreVadScoreEvent) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VadScoreVadScoreEvent) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VadScoreVadScoreEvent to string"
	}
	return string(jsonData)
}

func (v *VadScoreVadScoreEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VadScoreVadScoreEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VadScoreVadScoreEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VadScoreVadScoreEvent) MarshalJSON() ([]byte, error) {
	type embed VadScoreVadScoreEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
