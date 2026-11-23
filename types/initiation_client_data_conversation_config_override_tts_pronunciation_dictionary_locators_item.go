package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A locator for other documents to be able to reference a specific dictionary and it's version.
// This is a pydantic version of PronunciationDictionaryVersionLocatorDBModel.
// Required to ensure compat with the rest of the agent data models.
type InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem struct {
	// The ID of the pronunciation dictionary
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	// The ID of the version of the pronunciation dictionary
	VersionId       string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties map[string]interface{}
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if i == nil {
		return
	}
	return i.PronunciationDictionaryId
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) GetVersionId() (value string) {
	if i == nil {
		return
	}
	return i.VersionId
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	i.PronunciationDictionaryId = pronunciationDictionaryId
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) SetVersionId(versionId string) {
	i.VersionId = versionId
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, i); err != nil {
		return err
	}
	type alias InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*i = InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
