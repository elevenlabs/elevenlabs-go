package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryLocatorResponseModel struct {
	PronunciationDictionaryId string  `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	VersionId                 *string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties           map[string]interface{}
}

func (p *PronunciationDictionaryLocatorResponseModel) GetPronunciationDictionaryId() (value string) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryId
}

func (p *PronunciationDictionaryLocatorResponseModel) GetVersionId() (value *string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *PronunciationDictionaryLocatorResponseModel) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	p.PronunciationDictionaryId = pronunciationDictionaryId
}

func (p *PronunciationDictionaryLocatorResponseModel) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *PronunciationDictionaryLocatorResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryLocatorResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryLocatorResponseModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryLocatorResponseModel) UnmarshalJSON(data []byte) error {
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

func (p *PronunciationDictionaryLocatorResponseModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryLocatorResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
