package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetPronunciationDictionaryWithRulesResponseModelRulesItem struct {
	Type    string
	Alias   *PronunciationDictionaryAliasRuleResponseModel
	Phoneme *PronunciationDictionaryPhonemeRuleResponseModel
}

func (g *GetPronunciationDictionaryWithRulesResponseModelRulesItem) GetAlias() *PronunciationDictionaryAliasRuleResponseModel {
	if g == nil {
		return nil
	}
	if g.Alias == nil {
		return nil
	}
	return g.Alias
}

func (g *GetPronunciationDictionaryWithRulesResponseModelRulesItem) GetPhoneme() *PronunciationDictionaryPhonemeRuleResponseModel {
	if g == nil {
		return nil
	}
	if g.Phoneme == nil {
		return nil
	}
	return g.Phoneme
}

func (g *GetPronunciationDictionaryWithRulesResponseModelRulesItem) GetType() string {
	if g == nil {
		return ""
	}
	return g.Type
}

func (g *GetPronunciationDictionaryWithRulesResponseModelRulesItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", g)
	}
	switch unmarshaler.Type {
	case "alias":
		value := new(PronunciationDictionaryAliasRuleResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Alias = value
	case "phoneme":
		value := new(PronunciationDictionaryPhonemeRuleResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Phoneme = value
	}
	return nil
}

func (g GetPronunciationDictionaryWithRulesResponseModelRulesItem) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Alias != nil {
		return internal.MarshalJSONWithDiscriminant(g.Alias, "type", "alias")
	}
	if g.Phoneme != nil {
		return internal.MarshalJSONWithDiscriminant(g.Phoneme, "type", "phoneme")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetPronunciationDictionaryWithRulesResponseModelRulesItemVisitor interface {
	VisitAlias(*PronunciationDictionaryAliasRuleResponseModel) error
	VisitPhoneme(*PronunciationDictionaryPhonemeRuleResponseModel) error
}

func (g *GetPronunciationDictionaryWithRulesResponseModelRulesItem) Accept(visitor GetPronunciationDictionaryWithRulesResponseModelRulesItemVisitor) error {
	if g.Alias != nil {
		return visitor.VisitAlias(g.Alias)
	}
	if g.Phoneme != nil {
		return visitor.VisitPhoneme(g.Phoneme)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetPronunciationDictionaryWithRulesResponseModelRulesItem) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Alias != nil {
		fields = append(fields, "alias")
	}
	if g.Phoneme != nil {
		fields = append(fields, "phoneme")
	}
	if len(fields) == 0 {
		if g.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Type)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Type != "" {
		field := fields[0]
		if g.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Type,
				g,
			)
		}
	}
	return nil
}
