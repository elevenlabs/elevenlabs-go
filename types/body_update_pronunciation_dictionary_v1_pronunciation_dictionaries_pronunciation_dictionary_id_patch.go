package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch struct {
	// The id of the pronunciation dictionary
	PronunciationDictionaryId string `json:"-" url:"-"`
	// Whether to archive the pronunciation dictionary.
	Archived *bool `json:"archived,omitempty" xml:"archived,omitempty"`
	// The name of the pronunciation dictionary, used for identification only.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) GetPronunciationDictionaryId() (value string) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryId
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) GetArchived() (value *bool) {
	if b == nil {
		return
	}
	return b.Archived
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	b.PronunciationDictionaryId = pronunciationDictionaryId
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) SetArchived(archived *bool) {
	b.Archived = archived
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) SetName(name *string) {
	b.Name = name
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch(tmp)
	return nil
}

func (b *BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdatePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
