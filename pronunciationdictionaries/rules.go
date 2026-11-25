package pronunciationdictionaries

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost struct {
	// The id of the pronunciation dictionary
	PronunciationDictionaryId string `json:"-" url:"-"`
	// List of pronunciation rules. Rule can be either:
	// an alias rule: {'string_to_replace': 'a', 'type': 'alias', 'alias': 'b', }
	// or a phoneme rule: {'string_to_replace': 'a', 'type': 'phoneme', 'phoneme': 'b', 'alphabet': 'ipa' }
	Rules []*BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem `json:"rules" xml:"rules" required:"true"`
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) GetPronunciationDictionaryId() (value string) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryId
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) GetRules() (value []*BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) {
	if b == nil {
		return
	}
	return b.Rules
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	b.PronunciationDictionaryId = pronunciationDictionaryId
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) SetRules(rules []*BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) {
	b.Rules = rules
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost to string"
	}
	return string(jsonData)
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost(tmp)
	return nil
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost) MarshalJSON() ([]byte, error) {
	type embed BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem struct {
	Type    string
	Alias   *types.PronunciationDictionaryAliasRuleRequestModel
	Phoneme *types.PronunciationDictionaryPhonemeRuleRequestModel
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) GetAlias() *types.PronunciationDictionaryAliasRuleRequestModel {
	if b == nil {
		return nil
	}
	if b.Alias == nil {
		return nil
	}
	return b.Alias
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) GetPhoneme() *types.PronunciationDictionaryPhonemeRuleRequestModel {
	if b == nil {
		return nil
	}
	if b.Phoneme == nil {
		return nil
	}
	return b.Phoneme
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) GetType() string {
	if b == nil {
		return ""
	}
	return b.Type
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) UnmarshalJSON(data []byte) error {
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
		value := new(types.PronunciationDictionaryAliasRuleRequestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Alias = value
	case "phoneme":
		value := new(types.PronunciationDictionaryPhonemeRuleRequestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Phoneme = value
	}
	return nil
}

func (b BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) MarshalJSON() ([]byte, error) {
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

type BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItemVisitor interface {
	VisitAlias(*types.PronunciationDictionaryAliasRuleRequestModel) error
	VisitPhoneme(*types.PronunciationDictionaryPhonemeRuleRequestModel) error
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) Accept(visitor BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItemVisitor) error {
	if b.Alias != nil {
		return visitor.VisitAlias(b.Alias)
	}
	if b.Phoneme != nil {
		return visitor.VisitPhoneme(b.Phoneme)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", b)
}

func (b *BodySetRulesOnThePronunciationDictionaryV1PronunciationDictionariesPronunciationDictionaryIdSetRulesPostRulesItem) validate() error {
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

type PronunciationDictionary struct {
	// The id of the pronunciation dictionary
	PronunciationDictionaryId string `json:"-" url:"-"`
	// List of pronunciation rules. Rule can be either:
	// an alias rule: {'string_to_replace': 'a', 'type': 'alias', 'alias': 'b', }
	// or a phoneme rule: {'string_to_replace': 'a', 'type': 'phoneme', 'phoneme': 'b', 'alphabet': 'ipa' }
	Rules []*PronunciationDictionaryRule `json:"rules" xml:"rules" required:"true"`
}

func (p *PronunciationDictionary) GetPronunciationDictionaryId() (value string) {
	if p == nil {
		return
	}
	return p.PronunciationDictionaryId
}

func (p *PronunciationDictionary) GetRules() (value []*PronunciationDictionaryRule) {
	if p == nil {
		return
	}
	return p.Rules
}

func (p *PronunciationDictionary) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	p.PronunciationDictionaryId = pronunciationDictionaryId
}

func (p *PronunciationDictionary) SetRules(rules []*PronunciationDictionaryRule) {
	p.Rules = rules
}

func (p *PronunciationDictionary) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PronunciationDictionary) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PronunciationDictionary to string"
	}
	return string(jsonData)
}

func (p *PronunciationDictionary) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PronunciationDictionary
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PronunciationDictionary(tmp)
	return nil
}

func (p *PronunciationDictionary) MarshalJSON() ([]byte, error) {
	type embed PronunciationDictionary
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}

type PronunciationDictionaryRule struct {
	Type    string
	Alias   *types.PronunciationDictionaryAliasRuleRequestModel
	Phoneme *types.PronunciationDictionaryPhonemeRuleRequestModel
}

func (p *PronunciationDictionaryRule) GetAlias() *types.PronunciationDictionaryAliasRuleRequestModel {
	if p == nil {
		return nil
	}
	if p.Alias == nil {
		return nil
	}
	return p.Alias
}

func (p *PronunciationDictionaryRule) GetPhoneme() *types.PronunciationDictionaryPhonemeRuleRequestModel {
	if p == nil {
		return nil
	}
	if p.Phoneme == nil {
		return nil
	}
	return p.Phoneme
}

func (p *PronunciationDictionaryRule) GetType() string {
	if p == nil {
		return ""
	}
	return p.Type
}

func (p *PronunciationDictionaryRule) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	p.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", p)
	}
	switch unmarshaler.Type {
	case "alias":
		value := new(types.PronunciationDictionaryAliasRuleRequestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Alias = value
	case "phoneme":
		value := new(types.PronunciationDictionaryPhonemeRuleRequestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		p.Phoneme = value
	}
	return nil
}

func (p PronunciationDictionaryRule) MarshalJSON() ([]byte, error) {
	if err := p.validate(); err != nil {
		return nil, err
	}
	if p.Alias != nil {
		return internal.MarshalJSONWithDiscriminant(p.Alias, "type", "alias")
	}
	if p.Phoneme != nil {
		return internal.MarshalJSONWithDiscriminant(p.Phoneme, "type", "phoneme")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", p)
}

type PronunciationDictionaryRuleVisitor interface {
	VisitAlias(*types.PronunciationDictionaryAliasRuleRequestModel) error
	VisitPhoneme(*types.PronunciationDictionaryPhonemeRuleRequestModel) error
}

func (p *PronunciationDictionaryRule) Accept(visitor PronunciationDictionaryRuleVisitor) error {
	if p.Alias != nil {
		return visitor.VisitAlias(p.Alias)
	}
	if p.Phoneme != nil {
		return visitor.VisitPhoneme(p.Phoneme)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", p)
}

func (p *PronunciationDictionaryRule) validate() error {
	if p == nil {
		return fmt.Errorf("type %T is nil", p)
	}
	var fields []string
	if p.Alias != nil {
		fields = append(fields, "alias")
	}
	if p.Phoneme != nil {
		fields = append(fields, "phoneme")
	}
	if len(fields) == 0 {
		if p.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", p, p.Type)
		}
		return fmt.Errorf("type %T is empty", p)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", p, fields)
	}
	if p.Type != "" {
		field := fields[0]
		if p.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				p,
				p.Type,
				p,
			)
		}
	}
	return nil
}

type RemovePronunciationDictionaryRulesRequest struct {
	// The id of the pronunciation dictionary
	PronunciationDictionaryId string `json:"-" url:"-"`
	// List of strings to remove from the pronunciation dictionary.
	RuleStrings []string `json:"rule_strings" xml:"rule_strings" required:"true"`
}

func (r *RemovePronunciationDictionaryRulesRequest) GetPronunciationDictionaryId() (value string) {
	if r == nil {
		return
	}
	return r.PronunciationDictionaryId
}

func (r *RemovePronunciationDictionaryRulesRequest) GetRuleStrings() (value []string) {
	if r == nil {
		return
	}
	return r.RuleStrings
}

func (r *RemovePronunciationDictionaryRulesRequest) SetPronunciationDictionaryId(pronunciationDictionaryId string) {
	r.PronunciationDictionaryId = pronunciationDictionaryId
}

func (r *RemovePronunciationDictionaryRulesRequest) SetRuleStrings(ruleStrings []string) {
	r.RuleStrings = ruleStrings
}

func (r *RemovePronunciationDictionaryRulesRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (r RemovePronunciationDictionaryRulesRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RemovePronunciationDictionaryRulesRequest to string"
	}
	return string(jsonData)
}

func (r *RemovePronunciationDictionaryRulesRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RemovePronunciationDictionaryRulesRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RemovePronunciationDictionaryRulesRequest(tmp)
	return nil
}

func (r *RemovePronunciationDictionaryRulesRequest) MarshalJSON() ([]byte, error) {
	type embed RemovePronunciationDictionaryRulesRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
