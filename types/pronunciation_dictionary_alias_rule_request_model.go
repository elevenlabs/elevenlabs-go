package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryAliasRuleRequestModel struct {
	// The string to replace. Must be a non-empty string.
	StringToReplace string `json:"string_to_replace" xml:"string_to_replace" required:"true"`
	// Whether the rule should match case-sensitively.
	CaseSensitive *bool `json:"case_sensitive,omitempty" xml:"case_sensitive,omitempty"`
	// Whether the rule should only match at word boundaries.
	WordBoundaries *bool `json:"word_boundaries,omitempty" xml:"word_boundaries,omitempty"`
	// The type of the rule.
	type_ string
	// The alias for the string to be replaced.
	Alias           string `json:"alias" xml:"alias" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PronunciationDictionaryAliasRuleRequestModel) GetStringToReplace() (value string) {
	if p == nil {
		return
	}
	return p.StringToReplace
}

func (p *PronunciationDictionaryAliasRuleRequestModel) GetCaseSensitive() (value *bool) {
	if p == nil {
		return
	}
	return p.CaseSensitive
}

func (p *PronunciationDictionaryAliasRuleRequestModel) GetWordBoundaries() (value *bool) {
	if p == nil {
		return
	}
	return p.WordBoundaries
}

func (p *PronunciationDictionaryAliasRuleRequestModel) GetAlias() (value string) {
	if p == nil {
		return
	}
	return p.Alias
}

func (p *PronunciationDictionaryAliasRuleRequestModel) SetStringToReplace(stringToReplace string) {
	p.StringToReplace = stringToReplace
}

func (p *PronunciationDictionaryAliasRuleRequestModel) SetCaseSensitive(caseSensitive *bool) {
	p.CaseSensitive = caseSensitive
}

func (p *PronunciationDictionaryAliasRuleRequestModel) SetWordBoundaries(wordBoundaries *bool) {
	p.WordBoundaries = wordBoundaries
}

func (p *PronunciationDictionaryAliasRuleRequestModel) SetAlias(alias string) {
	p.Alias = alias
}

func (p *PronunciationDictionaryAliasRuleRequestModel) Type() string {
	return p.type_
}

func (p *PronunciationDictionaryAliasRuleRequestModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryAliasRuleRequestModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryAliasRuleRequestModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryAliasRuleRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PronunciationDictionaryAliasRuleRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PronunciationDictionaryAliasRuleRequestModel(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "alias" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "alias", *literals.Type)
	}
	p.type_ = "alias"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PronunciationDictionaryAliasRuleRequestModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryAliasRuleRequestModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "alias",
	}
	return json.Marshal(marshaler)
}
