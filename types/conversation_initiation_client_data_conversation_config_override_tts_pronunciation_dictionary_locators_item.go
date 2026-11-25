package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A locator for other documents to be able to reference a specific dictionary and it's version.
// This is a pydantic version of PronunciationDictionaryVersionLocatorDBModel.
// Required to ensure compat with the rest of the agent data models.
type ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem struct {
	// The ID of the pronunciation dictionary
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	// The ID of the version of the pronunciation dictionary
	VersionId       string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if c == nil {
		return
	}
	return c.PronunciationDictionaryId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) GetVersionId() (value string) {
	if c == nil {
		return
	}
	return c.VersionId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	c.PronunciationDictionaryId = pronunciationDictionaryId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) SetVersionId(versionId string) {
	c.VersionId = versionId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
