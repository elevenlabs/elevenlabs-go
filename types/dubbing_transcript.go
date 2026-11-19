package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTranscript struct {
	Language        string                        `json:"language" xml:"language" required:"true"`
	Utterances      []*DubbingTranscriptUtterance `json:"utterances" xml:"utterances" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscript) GetLanguage() (value string) {
	if d == nil {
		return
	}
	return d.Language
}

func (d *DubbingTranscript) GetUtterances() (value []*DubbingTranscriptUtterance) {
	if d == nil {
		return
	}
	return d.Utterances
}

func (d *DubbingTranscript) SetLanguage(language string) {
	d.Language = language
}

func (d *DubbingTranscript) SetUtterances(utterances []*DubbingTranscriptUtterance) {
	d.Utterances = utterances
}

func (d *DubbingTranscript) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscript) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscript to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTranscript(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscript) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
