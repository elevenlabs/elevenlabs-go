package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioOutputNormalizedAlignment struct {
	Chars            []string `json:"chars" xml:"chars" required:"true"`
	CharStartTimesMs []int    `json:"char_start_times_ms" xml:"char_start_times_ms" required:"true"`
	CharDurationsMs  []int    `json:"char_durations_ms" xml:"char_durations_ms" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AudioOutputNormalizedAlignment) GetChars() (value []string) {
	if a == nil {
		return
	}
	return a.Chars
}

func (a *AudioOutputNormalizedAlignment) GetCharStartTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharStartTimesMs
}

func (a *AudioOutputNormalizedAlignment) GetCharDurationsMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharDurationsMs
}

func (a *AudioOutputNormalizedAlignment) SetChars(chars []string) {
	a.Chars = chars
}

func (a *AudioOutputNormalizedAlignment) SetCharStartTimesMs(charStartTimesMs []int) {
	a.CharStartTimesMs = charStartTimesMs
}

func (a *AudioOutputNormalizedAlignment) SetCharDurationsMs(charDurationsMs []int) {
	a.CharDurationsMs = charDurationsMs
}

func (a *AudioOutputNormalizedAlignment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioOutputNormalizedAlignment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioOutputNormalizedAlignment to string"
	}
	return string(jsonData)
}

func (a *AudioOutputNormalizedAlignment) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioOutputNormalizedAlignment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioOutputNormalizedAlignment(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioOutputNormalizedAlignment) MarshalJSON() ([]byte, error) {
	type embed AudioOutputNormalizedAlignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
