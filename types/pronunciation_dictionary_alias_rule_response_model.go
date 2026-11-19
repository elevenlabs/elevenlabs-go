package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryAliasRuleResponseModel struct {
	StringToReplace string `json:"string_to_replace" xml:"string_to_replace" required:"true"`
	// Whether the rule matches case-sensitively.
	CaseSensitive *bool `json:"case_sensitive,omitempty" xml:"case_sensitive,omitempty"`
	// Whether the rule only matches at word boundaries.
	WordBoundaries  *bool `json:"word_boundaries,omitempty" xml:"word_boundaries,omitempty"`
	type_           string
	Alias           string `json:"alias" xml:"alias" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PronunciationDictionaryAliasRuleResponseModel) GetStringToReplace() (value string) {
	if p == nil {
		return
	}
	return p.StringToReplace
}

func (p *PronunciationDictionaryAliasRuleResponseModel) GetCaseSensitive() (value *bool) {
	if p == nil {
		return
	}
	return p.CaseSensitive
}

func (p *PronunciationDictionaryAliasRuleResponseModel) GetWordBoundaries() (value *bool) {
	if p == nil {
		return
	}
	return p.WordBoundaries
}

func (p *PronunciationDictionaryAliasRuleResponseModel) GetAlias() (value string) {
	if p == nil {
		return
	}
	return p.Alias
}

func (p *PronunciationDictionaryAliasRuleResponseModel) SetStringToReplace(stringToReplace string) {
	p.StringToReplace = stringToReplace
}

func (p *PronunciationDictionaryAliasRuleResponseModel) SetCaseSensitive(caseSensitive *bool) {
	p.CaseSensitive = caseSensitive
}

func (p *PronunciationDictionaryAliasRuleResponseModel) SetWordBoundaries(wordBoundaries *bool) {
	p.WordBoundaries = wordBoundaries
}

func (p *PronunciationDictionaryAliasRuleResponseModel) SetAlias(alias string) {
	p.Alias = alias
}

func (p *PronunciationDictionaryAliasRuleResponseModel) Type() string {
	return p.type_
}

func (p *PronunciationDictionaryAliasRuleResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryAliasRuleResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryAliasRuleResponseModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryAliasRuleResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PronunciationDictionaryAliasRuleResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PronunciationDictionaryAliasRuleResponseModel(tmp)
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

func (p *PronunciationDictionaryAliasRuleResponseModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryAliasRuleResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "alias",
	}
	return json.Marshal(marshaler)
}
