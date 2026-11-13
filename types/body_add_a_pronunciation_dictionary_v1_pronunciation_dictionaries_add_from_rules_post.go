package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost struct {
	// List of pronunciation rules. Rule can be either:
	// an alias rule: {'string_to_replace': 'a', 'type': 'alias', 'alias': 'b', }
	// or a phoneme rule: {'string_to_replace': 'a', 'type': 'phoneme', 'phoneme': 'b', 'alphabet': 'ipa' }
	Rules []*BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem `json:"rules" xml:"rules" required:"true"`
	// The name of the pronunciation dictionary, used for identification only.
	Name string `json:"name" xml:"name" required:"true"`
	// A description of the pronunciation dictionary, used for identification only.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Should be one of 'admin', 'editor' or 'viewer'. If not provided, defaults to no access.
	WorkspaceAccess *AddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostWorkspaceAccess `json:"workspace_access,omitempty" xml:"workspace_access,omitempty"`
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) GetRules() (value []*BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) {
	if b == nil {
		return
	}
	return b.Rules
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) GetWorkspaceAccess() (value *AddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostWorkspaceAccess) {
	if b == nil {
		return
	}
	return b.WorkspaceAccess
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) SetRules(rules []*BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) {
	b.Rules = rules
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) SetName(name string) {
	b.Name = name
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) SetWorkspaceAccess(workspaceAccess *AddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostWorkspaceAccess) {
	b.WorkspaceAccess = workspaceAccess
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost to string"
	}
	return string(jsonData)
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost) MarshalJSON() ([]byte, error) {
	type embed BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
