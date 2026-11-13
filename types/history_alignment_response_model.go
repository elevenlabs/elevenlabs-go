package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type HistoryAlignmentResponseModel struct {
	// The characters in the alignment.
	Characters []string `json:"characters" xml:"characters" required:"true"`
	// The start times of the characters in seconds.
	CharacterStartTimesSeconds []float64 `json:"character_start_times_seconds" xml:"character_start_times_seconds" required:"true"`
	// The end times of the characters in seconds.
	CharacterEndTimesSeconds []float64 `json:"character_end_times_seconds" xml:"character_end_times_seconds" required:"true"`
	extraProperties          map[string]interface{}
}

func (h *HistoryAlignmentResponseModel) GetCharacters() (value []string) {
	if h == nil {
		return
	}
	return h.Characters
}

func (h *HistoryAlignmentResponseModel) GetCharacterStartTimesSeconds() (value []float64) {
	if h == nil {
		return
	}
	return h.CharacterStartTimesSeconds
}

func (h *HistoryAlignmentResponseModel) GetCharacterEndTimesSeconds() (value []float64) {
	if h == nil {
		return
	}
	return h.CharacterEndTimesSeconds
}

func (h *HistoryAlignmentResponseModel) SetCharacters(characters []string) {
	h.Characters = characters
}

func (h *HistoryAlignmentResponseModel) SetCharacterStartTimesSeconds(characterStartTimesSeconds []float64) {
	h.CharacterStartTimesSeconds = characterStartTimesSeconds
}

func (h *HistoryAlignmentResponseModel) SetCharacterEndTimesSeconds(characterEndTimesSeconds []float64) {
	h.CharacterEndTimesSeconds = characterEndTimesSeconds
}

func (h *HistoryAlignmentResponseModel) GetExtraProperties() map[string]interface{} {
	if h == nil {
		return nil
	}
	return h.extraProperties
}

func (h HistoryAlignmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "error converting struct: HistoryAlignmentResponseModel to string"
	}
	return string(jsonData)
}

func (h *HistoryAlignmentResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler HistoryAlignmentResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*h = HistoryAlignmentResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *h)
	if err != nil {
		return err
	}
	h.extraProperties = extraProperties
	return nil
}

func (h *HistoryAlignmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed HistoryAlignmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*h),
	}
	return json.Marshal(marshaler)
}
