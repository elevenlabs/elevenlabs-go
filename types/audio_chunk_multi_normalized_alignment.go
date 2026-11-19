package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioChunkMultiNormalizedAlignment struct {
	Chars            []string `json:"chars" xml:"chars" required:"true"`
	CharStartTimesMs []int    `json:"char_start_times_ms" xml:"char_start_times_ms" required:"true"`
	CharDurationsMs  []int    `json:"char_durations_ms" xml:"char_durations_ms" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AudioChunkMultiNormalizedAlignment) GetChars() (value []string) {
	if a == nil {
		return
	}
	return a.Chars
}

func (a *AudioChunkMultiNormalizedAlignment) GetCharStartTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharStartTimesMs
}

func (a *AudioChunkMultiNormalizedAlignment) GetCharDurationsMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharDurationsMs
}

func (a *AudioChunkMultiNormalizedAlignment) SetChars(chars []string) {
	a.Chars = chars
}

func (a *AudioChunkMultiNormalizedAlignment) SetCharStartTimesMs(charStartTimesMs []int) {
	a.CharStartTimesMs = charStartTimesMs
}

func (a *AudioChunkMultiNormalizedAlignment) SetCharDurationsMs(charDurationsMs []int) {
	a.CharDurationsMs = charDurationsMs
}

func (a *AudioChunkMultiNormalizedAlignment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioChunkMultiNormalizedAlignment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioChunkMultiNormalizedAlignment to string"
	}
	return string(jsonData)
}

func (a *AudioChunkMultiNormalizedAlignment) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioChunkMultiNormalizedAlignment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioChunkMultiNormalizedAlignment(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioChunkMultiNormalizedAlignment) MarshalJSON() ([]byte, error) {
	type embed AudioChunkMultiNormalizedAlignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
