package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTranscriptResponseModel struct {
	Language        string                        `json:"language" xml:"language" required:"true"`
	Utterances      []*DubbingTranscriptUtterance `json:"utterances" xml:"utterances" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriptResponseModel) GetLanguage() (value string) {
	if d == nil {
		return
	}
	return d.Language
}

func (d *DubbingTranscriptResponseModel) GetUtterances() (value []*DubbingTranscriptUtterance) {
	if d == nil {
		return
	}
	return d.Utterances
}

func (d *DubbingTranscriptResponseModel) SetLanguage(language string) {
	d.Language = language
}

func (d *DubbingTranscriptResponseModel) SetUtterances(utterances []*DubbingTranscriptUtterance) {
	d.Utterances = utterances
}

func (d *DubbingTranscriptResponseModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptResponseModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptResponseModel to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTranscriptResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTranscriptResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscriptResponseModel) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
