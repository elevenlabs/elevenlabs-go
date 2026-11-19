package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InitialiseContextPronunciationDictionaryLocatorsItem struct {
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	VersionId                 string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties           map[string]interface{}
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryId
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) GetVersionId() (value string) {
	if i == nil {
		return
	}
	return i.VersionId
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	i.PronunciationDictionaryId = pronunciationDictionaryId
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) SetVersionId(versionId string) {
	i.VersionId = versionId
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitialiseContextPronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitialiseContextPronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitialiseContextPronunciationDictionaryLocatorsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitialiseContextPronunciationDictionaryLocatorsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitialiseContextPronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed InitialiseContextPronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
