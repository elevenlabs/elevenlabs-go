package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem struct {
	Type    string
	Alias   *PronunciationDictionaryAliasRuleRequestModel
	Phoneme *PronunciationDictionaryPhonemeRuleRequestModel
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) GetAlias() *PronunciationDictionaryAliasRuleRequestModel {
	if b == nil {
		return nil
	}
	if b.Alias == nil {
		return nil
	}
	return b.Alias
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) GetPhoneme() *PronunciationDictionaryPhonemeRuleRequestModel {
	if b == nil {
		return nil
	}
	if b.Phoneme == nil {
		return nil
	}
	return b.Phoneme
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) GetType() string {
	if b == nil {
		return ""
	}
	return b.Type
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	b.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", b)
	}
	switch unmarshaler.Type {
	case "alias":
		value := new(PronunciationDictionaryAliasRuleRequestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Alias = value
	case "phoneme":
		value := new(PronunciationDictionaryPhonemeRuleRequestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Phoneme = value
	}
	return nil
}

func (b BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) MarshalJSON() ([]byte, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	if b.Alias != nil {
		return internal.MarshalJSONWithDiscriminant(b.Alias, "type", "alias")
	}
	if b.Phoneme != nil {
		return internal.MarshalJSONWithDiscriminant(b.Phoneme, "type", "phoneme")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", b)
}

type BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItemVisitor interface {
	VisitAlias(*PronunciationDictionaryAliasRuleRequestModel) error
	VisitPhoneme(*PronunciationDictionaryPhonemeRuleRequestModel) error
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) Accept(visitor BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItemVisitor) error {
	if b.Alias != nil {
		return visitor.VisitAlias(b.Alias)
	}
	if b.Phoneme != nil {
		return visitor.VisitPhoneme(b.Phoneme)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", b)
}

func (b *BodyAddAPronunciationDictionaryV1PronunciationDictionariesAddFromRulesPostRulesItem) validate() error {
	if b == nil {
		return fmt.Errorf("type %T is nil", b)
	}
	var fields []string
	if b.Alias != nil {
		fields = append(fields, "alias")
	}
	if b.Phoneme != nil {
		fields = append(fields, "phoneme")
	}
	if len(fields) == 0 {
		if b.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", b, b.Type)
		}
		return fmt.Errorf("type %T is empty", b)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", b, fields)
	}
	if b.Type != "" {
		field := fields[0]
		if b.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				b,
				b.Type,
				b,
			)
		}
	}
	return nil
}
