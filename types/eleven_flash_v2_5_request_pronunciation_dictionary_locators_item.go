package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A pronunciation dictionary to apply during speech synthesis.
type ElevenFlashV25RequestPronunciationDictionaryLocatorsItem struct {
	// The ID of a pronunciation dictionary created via `POST /v1/pronunciation-dictionaries/add-from-file` or `POST /v1/pronunciation-dictionaries/add-from-rules`.
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	// The version of the dictionary to use. Omit to use the latest version.
	VersionId       *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if e == nil {
		return
	}
	return e.PronunciationDictionaryId
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) GetVersionId() (value *string) {
	if e == nil {
		return
	}
	return e.VersionId
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	e.PronunciationDictionaryId = pronunciationDictionaryId
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) SetVersionId(versionId *string) {
	e.VersionId = versionId
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ElevenFlashV25RequestPronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed ElevenFlashV25RequestPronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
