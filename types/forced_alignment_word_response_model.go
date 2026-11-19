package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Model representing a single word with its timing information from the aligner.
type ForcedAlignmentWordResponseModel struct {
	// The word that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the word in seconds.
	Start float64 `json:"start" xml:"start" required:"true"`
	// The end time of the word in seconds.
	End float64 `json:"end" xml:"end" required:"true"`
	// The average alignment loss/confidence score for this word, calculated from its constituent characters.
	Loss            float64 `json:"loss" xml:"loss" required:"true"`
	extraProperties map[string]interface{}
}

func (f *ForcedAlignmentWordResponseModel) GetText() (value string) {
	if f == nil {
		return
	}
	return f.Text
}

func (f *ForcedAlignmentWordResponseModel) GetStart() (value float64) {
	if f == nil {
		return
	}
	return f.Start
}

func (f *ForcedAlignmentWordResponseModel) GetEnd() (value float64) {
	if f == nil {
		return
	}
	return f.End
}

func (f *ForcedAlignmentWordResponseModel) GetLoss() (value float64) {
	if f == nil {
		return
	}
	return f.Loss
}

func (f *ForcedAlignmentWordResponseModel) SetText(text string) {
	f.Text = text
}

func (f *ForcedAlignmentWordResponseModel) SetStart(start float64) {
	f.Start = start
}

func (f *ForcedAlignmentWordResponseModel) SetEnd(end float64) {
	f.End = end
}

func (f *ForcedAlignmentWordResponseModel) SetLoss(loss float64) {
	f.Loss = loss
}

func (f *ForcedAlignmentWordResponseModel) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f ForcedAlignmentWordResponseModel) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: ForcedAlignmentWordResponseModel to string"
	}
	return string(jsonData)
}

func (f *ForcedAlignmentWordResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, f); err != nil {
		return err
	}
	type alias ForcedAlignmentWordResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*f = ForcedAlignmentWordResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *ForcedAlignmentWordResponseModel) MarshalJSON() ([]byte, error) {
	type embed ForcedAlignmentWordResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
