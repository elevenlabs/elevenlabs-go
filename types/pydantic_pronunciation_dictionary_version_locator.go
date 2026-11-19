package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A locator for other documents to be able to reference a specific dictionary and it's version.
// This is a pydantic version of PronunciationDictionaryVersionLocatorDBModel.
// Required to ensure compat with the rest of the agent data models.
type PydanticPronunciationDictionaryVersionLocator struct {
	// The ID of the pronunciation dictionary
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	// The ID of the version of the pronunciation dictionary
	VersionId       *string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PydanticPronunciationDictionaryVersionLocator) GetPronunciationDictionaryId() (value string) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryId
}

func (p *PydanticPronunciationDictionaryVersionLocator) GetVersionId() (value *string) {
	if p == nil {
		return
	}
	return p.VersionId
}

func (p *PydanticPronunciationDictionaryVersionLocator) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	p.PronunciationDictionaryId = pronunciationDictionaryId
}

func (p *PydanticPronunciationDictionaryVersionLocator) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *PydanticPronunciationDictionaryVersionLocator) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PydanticPronunciationDictionaryVersionLocator) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PydanticPronunciationDictionaryVersionLocator to string"
	}
	return string(jsonData)
}

func (p *PydanticPronunciationDictionaryVersionLocator) UnmarshalJSON(data []byte) error {
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

func (p *PydanticPronunciationDictionaryVersionLocator) MarshalJSON() ([]byte, error) {
	type embed PydanticPronunciationDictionaryVersionLocator
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
