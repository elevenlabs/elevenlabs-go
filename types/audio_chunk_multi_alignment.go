package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioChunkMultiAlignment struct {
	Chars            []string `json:"chars" xml:"chars" required:"true"`
	CharStartTimesMs []int    `json:"char_start_times_ms" xml:"char_start_times_ms" required:"true"`
	CharDurationsMs  []int    `json:"char_durations_ms" xml:"char_durations_ms" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AudioChunkMultiAlignment) GetChars() (value []string) {
	if a == nil {
		return
	}
	return a.Chars
}

func (a *AudioChunkMultiAlignment) GetCharStartTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharStartTimesMs
}

func (a *AudioChunkMultiAlignment) GetCharDurationsMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharDurationsMs
}

func (a *AudioChunkMultiAlignment) SetChars(chars []string) {
	a.Chars = chars
}

func (a *AudioChunkMultiAlignment) SetCharStartTimesMs(charStartTimesMs []int) {
	a.CharStartTimesMs = charStartTimesMs
}

func (a *AudioChunkMultiAlignment) SetCharDurationsMs(charDurationsMs []int) {
	a.CharDurationsMs = charDurationsMs
}

func (a *AudioChunkMultiAlignment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioChunkMultiAlignment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioChunkMultiAlignment to string"
	}
	return string(jsonData)
}

func (a *AudioChunkMultiAlignment) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioChunkMultiAlignment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioChunkMultiAlignment(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioChunkMultiAlignment) MarshalJSON() ([]byte, error) {
	type embed AudioChunkMultiAlignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
