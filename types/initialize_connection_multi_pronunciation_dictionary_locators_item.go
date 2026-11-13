package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InitializeConnectionMultiPronunciationDictionaryLocatorsItem struct {
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	VersionId                 string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties           map[string]interface{}
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryId
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) GetVersionId() (value string) {
	if i == nil {
		return
	}
	return i.VersionId
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	i.PronunciationDictionaryId = pronunciationDictionaryId
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) SetVersionId(versionId string) {
	i.VersionId = versionId
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionMultiPronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionMultiPronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitializeConnectionMultiPronunciationDictionaryLocatorsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitializeConnectionMultiPronunciationDictionaryLocatorsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnectionMultiPronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionMultiPronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
