package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryPhonemeRuleRequestModel struct {
	// The string to replace. Must be a non-empty string.
	StringToReplace string `json:"string_to_replace" xml:"string_to_replace" required:"true"`
	// Whether the rule should match case-sensitively.
	CaseSensitive *bool `json:"case_sensitive,omitempty" xml:"case_sensitive,omitempty"`
	// Whether the rule should only match at word boundaries.
	WordBoundaries *bool `json:"word_boundaries,omitempty" xml:"word_boundaries,omitempty"`
	// The type of the rule.
	type_ string
	// The phoneme rule.
	Phoneme string `json:"phoneme" xml:"phoneme" required:"true"`
	// The alphabet to use with the phoneme rule.
	Alphabet        string `json:"alphabet" xml:"alphabet" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) GetStringToReplace() (value string) {
	if p == nil {
		return
	}
	return p.StringToReplace
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) GetCaseSensitive() (value *bool) {
	if p == nil {
		return
	}
	return p.CaseSensitive
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) GetWordBoundaries() (value *bool) {
	if p == nil {
		return
	}
	return p.WordBoundaries
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) GetPhoneme() (value string) {
	if p == nil {
		return
	}
	return p.Phoneme
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) GetAlphabet() (value string) {
	if p == nil {
		return
	}
	return p.Alphabet
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) SetStringToReplace(stringToReplace string) {
	p.StringToReplace = stringToReplace
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) SetCaseSensitive(caseSensitive *bool) {
	p.CaseSensitive = caseSensitive
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) SetWordBoundaries(wordBoundaries *bool) {
	p.WordBoundaries = wordBoundaries
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) SetPhoneme(phoneme string) {
	p.Phoneme = phoneme
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) SetAlphabet(alphabet string) {
	p.Alphabet = alphabet
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) Type() string {
	return p.type_
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryPhonemeRuleRequestModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryPhonemeRuleRequestModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PronunciationDictionaryPhonemeRuleRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PronunciationDictionaryPhonemeRuleRequestModel(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "phoneme" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "phoneme", *literals.Type)
	}
	p.type_ = "phoneme"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PronunciationDictionaryPhonemeRuleRequestModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryPhonemeRuleRequestModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "phoneme",
	}
	return json.Marshal(marshaler)
}
