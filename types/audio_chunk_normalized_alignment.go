package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioChunkNormalizedAlignment struct {
	Chars            []string `json:"chars" xml:"chars" required:"true"`
	CharStartTimesMs []int    `json:"char_start_times_ms" xml:"char_start_times_ms" required:"true"`
	CharDurationsMs  []int    `json:"char_durations_ms" xml:"char_durations_ms" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AudioChunkNormalizedAlignment) GetChars() (value []string) {
	if a == nil {
		return
	}
	return a.Chars
}

func (a *AudioChunkNormalizedAlignment) GetCharStartTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharStartTimesMs
}

func (a *AudioChunkNormalizedAlignment) GetCharDurationsMs() (value []int) {
	if a == nil {
		return
	}
	return a.CharDurationsMs
}

func (a *AudioChunkNormalizedAlignment) SetChars(chars []string) {
	a.Chars = chars
}

func (a *AudioChunkNormalizedAlignment) SetCharStartTimesMs(charStartTimesMs []int) {
	a.CharStartTimesMs = charStartTimesMs
}

func (a *AudioChunkNormalizedAlignment) SetCharDurationsMs(charDurationsMs []int) {
	a.CharDurationsMs = charDurationsMs
}

func (a *AudioChunkNormalizedAlignment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioChunkNormalizedAlignment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioChunkNormalizedAlignment to string"
	}
	return string(jsonData)
}

func (a *AudioChunkNormalizedAlignment) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioChunkNormalizedAlignment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioChunkNormalizedAlignment(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioChunkNormalizedAlignment) MarshalJSON() ([]byte, error) {
	type embed AudioChunkNormalizedAlignment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
