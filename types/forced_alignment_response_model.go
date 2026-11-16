package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Model representing the response from the aligner service.
type ForcedAlignmentResponseModel struct {
	// List of characters with their timing information.
	Characters []*ForcedAlignmentCharacterResponseModel `json:"characters" xml:"characters" required:"true"`
	// List of words with their timing information.
	Words []*ForcedAlignmentWordResponseModel `json:"words" xml:"words" required:"true"`
	// The average alignment loss/confidence score for the entire transcript, calculated from all characters.
	Loss            float64 `json:"loss" xml:"loss" required:"true"`
	extraProperties map[string]interface{}
}

func (f *ForcedAlignmentResponseModel) GetCharacters() (value []*ForcedAlignmentCharacterResponseModel) {
	if f == nil {
		return
	}
	return f.Characters
}

func (f *ForcedAlignmentResponseModel) GetWords() (value []*ForcedAlignmentWordResponseModel) {
	if f == nil {
		return
	}
	return f.Words
}

func (f *ForcedAlignmentResponseModel) GetLoss() (value float64) {
	if f == nil {
		return
	}
	return f.Loss
}

func (f *ForcedAlignmentResponseModel) SetCharacters(characters []*ForcedAlignmentCharacterResponseModel) {
	f.Characters = characters
}

func (f *ForcedAlignmentResponseModel) SetWords(words []*ForcedAlignmentWordResponseModel) {
	f.Words = words
}

func (f *ForcedAlignmentResponseModel) SetLoss(loss float64) {
	f.Loss = loss
}

func (f *ForcedAlignmentResponseModel) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f ForcedAlignmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: ForcedAlignmentResponseModel to string"
	}
	return string(jsonData)
}

func (f *ForcedAlignmentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias ForcedAlignmentResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = ForcedAlignmentResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *ForcedAlignmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed ForcedAlignmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
