package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PronunciationDictionaryPhonemeRuleResponseModel struct {
	StringToReplace string `json:"string_to_replace" xml:"string_to_replace" required:"true"`
	// Whether the rule matches case-sensitively.
	CaseSensitive *bool `json:"case_sensitive,omitempty" xml:"case_sensitive,omitempty"`
	// Whether the rule only matches at word boundaries.
	WordBoundaries  *bool `json:"word_boundaries,omitempty" xml:"word_boundaries,omitempty"`
	type_           string
	Phoneme         string `json:"phoneme" xml:"phoneme" required:"true"`
	Alphabet        string `json:"alphabet" xml:"alphabet" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) GetStringToReplace() (value string) {
	if p == nil {
		return
	}
	return p.StringToReplace
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) GetCaseSensitive() (value *bool) {
	if p == nil {
		return
	}
	return p.CaseSensitive
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) GetWordBoundaries() (value *bool) {
	if p == nil {
		return
	}
	return p.WordBoundaries
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) GetPhoneme() (value string) {
	if p == nil {
		return
	}
	return p.Phoneme
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) GetAlphabet() (value string) {
	if p == nil {
		return
	}
	return p.Alphabet
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) SetStringToReplace(stringToReplace string) {
	p.StringToReplace = stringToReplace
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) SetCaseSensitive(caseSensitive *bool) {
	p.CaseSensitive = caseSensitive
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) SetWordBoundaries(wordBoundaries *bool) {
	p.WordBoundaries = wordBoundaries
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) SetPhoneme(phoneme string) {
	p.Phoneme = phoneme
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) SetAlphabet(alphabet string) {
	p.Alphabet = alphabet
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) Type() string {
	return p.type_
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PronunciationDictionaryPhonemeRuleResponseModel) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionaryPhonemeRuleResponseModel to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionaryPhonemeRuleResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PronunciationDictionaryPhonemeRuleResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PronunciationDictionaryPhonemeRuleResponseModel(tmp)
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

func (p *PronunciationDictionaryPhonemeRuleResponseModel) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionaryPhonemeRuleResponseModel
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "phoneme",
	}
	return json.Marshal(marshaler)
}
