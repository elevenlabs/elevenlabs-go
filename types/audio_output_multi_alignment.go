package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioOutputMultiAlignment struct {
	Chars            []string `json:"chars" xml:"chars" required:"true"`
	CharStartTimesMs []int    `json:"char_start_times_ms" xml:"char_start_times_ms" required:"true"`
	CharDurationsMs  []int    `json:"char_durations_ms" xml:"char_durations_ms" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AudioOutputMultiAlignment) GetChars() (value []string) {
	if a == nil {
		return
	}
	return a.Chars
}

func (a *AudioOutputMultiAlignment) GetCharStartTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharStartTimesMs
}

func (a *AudioOutputMultiAlignment) GetCharDurationsMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharDurationsMs
}

func (a *AudioOutputMultiAlignment) SetChars(chars []string) {
	a.Chars = chars
}

func (a *AudioOutputMultiAlignment) SetCharStartTimesMs(charStartTimesMs []int) {
	a.CharStartTimesMs = charStartTimesMs
}

func (a *AudioOutputMultiAlignment) SetCharDurationsMs(charDurationsMs []int) {
	a.CharDurationsMs = charDurationsMs
}

func (a *AudioOutputMultiAlignment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioOutputMultiAlignment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioOutputMultiAlignment to string"
	}
	return string(jsonData)
}

func (a *AudioOutputMultiAlignment) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioOutputMultiAlignment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioOutputMultiAlignment(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioOutputMultiAlignment) MarshalJSON() ([]byte, error) {
	type embed AudioOutputMultiAlignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
