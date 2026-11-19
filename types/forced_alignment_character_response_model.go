package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Model representing a single character with its timing information from the aligner.
type ForcedAlignmentCharacterResponseModel struct {
	// The character that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the character in seconds.
	Start float64 `json:"start" xml:"start" required:"true"`
	// The end time of the character in seconds.
	End             float64 `json:"end" xml:"end" required:"true"`
	extraProperties map[string]interface{}
}

func (f *ForcedAlignmentCharacterResponseModel) GetText() (value string) {
	if f == nil {
		return
	}
	return f.Text
}

func (f *ForcedAlignmentCharacterResponseModel) GetStart() (value float64) {
	if f == nil {
		return
	}
	return f.Start
}

func (f *ForcedAlignmentCharacterResponseModel) GetEnd() (value float64) {
	if f == nil {
		return
	}
	return f.End
}

func (f *ForcedAlignmentCharacterResponseModel) SetText(text string) {
	f.Text = text
}

func (f *ForcedAlignmentCharacterResponseModel) SetStart(start float64) {
	f.Start = start
}

func (f *ForcedAlignmentCharacterResponseModel) SetEnd(end float64) {
	f.End = end
}

func (f *ForcedAlignmentCharacterResponseModel) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f ForcedAlignmentCharacterResponseModel) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: ForcedAlignmentCharacterResponseModel to string"
	}
	return string(jsonData)
}

func (f *ForcedAlignmentCharacterResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias ForcedAlignmentCharacterResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = ForcedAlignmentCharacterResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *ForcedAlignmentCharacterResponseModel) MarshalJSON() ([]byte, error) {
	type embed ForcedAlignmentCharacterResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
