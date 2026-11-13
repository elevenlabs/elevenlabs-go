package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A locator for other documents to be able to reference a specific dictionary and it's version.
// This is a pydantic version of PronunciationDictionaryVersionLocatorDBModel.
// Required to ensure compat with the rest of the agent data models.
type TextToDialogueClientMessagePronunciationDictionaryLocatorsItem struct {
	// The ID of the pronunciation dictionary
	PronunciationDictionaryId string `json:"pronunciation_dictionary_id" xml:"pronunciation_dictionary_id" required:"true"`
	// The ID of the version of the pronunciation dictionary
	VersionId       string `json:"version_id" xml:"version_id" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) GetPronunciationDictionaryId() (value string) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryId
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) GetVersionId() (value string) {
	if t == nil {
		return
	}
	return t.VersionId
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	t.PronunciationDictionaryId = pronunciationDictionaryId
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) SetVersionId(versionId string) {
	t.VersionId = versionId
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextToDialogueClientMessagePronunciationDictionaryLocatorsItem to string"
	}
	return string(jsonData)
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TextToDialogueClientMessagePronunciationDictionaryLocatorsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TextToDialogueClientMessagePronunciationDictionaryLocatorsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TextToDialogueClientMessagePronunciationDictionaryLocatorsItem) MarshalJSON() ([]byte, error) {
	type embed TextToDialogueClientMessagePronunciationDictionaryLocatorsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
