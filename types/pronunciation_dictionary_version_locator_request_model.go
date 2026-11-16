package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryVersionLocatorRequestModel struct {
	// The ID of the pronunciation dictionary.
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	// The ID of the version of the pronunciation dictionary. If not provided, the latest version will be used.
	VersionId       *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) GetPronunciationDictionaryId() (value string) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryId
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) GetVersionId() (value *string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	p.PronunciationDictionaryId = pronunciationDictionaryId
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryVersionLocatorRequestModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryVersionLocatorRequestModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PronunciationDictionaryVersionLocatorRequestModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryVersionLocatorRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
