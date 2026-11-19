package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InitializeConnectionPronunciationDictionaryLocatorsItem struct {
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	VersionId                 string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties           map[string]interface{}
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryId
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) GetVersionId() (value string) {
	if i == nil {
		return
	}
	return i.VersionId
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	i.PronunciationDictionaryId = pronunciationDictionaryId
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) SetVersionId(versionId string) {
	i.VersionId = versionId
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitializeConnectionPronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitializeConnectionPronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitializeConnectionPronunciationDictionaryLocatorsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitializeConnectionPronunciationDictionaryLocatorsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitializeConnectionPronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed InitializeConnectionPronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
