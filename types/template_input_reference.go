package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TemplateInputReference struct {
	Type         string
	Asset        *StaticAssetReference
	Generation   *GenerationReference
	InlineBase64 *InlineBase64Reference
	Voice        *VoiceReference
}

func (t *TemplateInputReference) GetAsset() *StaticAssetReference {
	if t == nil {
		return nil
	}
	if t.Asset == nil {
		return nil
	}
	return t.Asset
}

func (t *TemplateInputReference) GetGeneration() *GenerationReference {
	if t == nil {
		return nil
	}
	if t.Generation == nil {
		return nil
	}
	return t.Generation
}

func (t *TemplateInputReference) GetInlineBase64() *InlineBase64Reference {
	if t == nil {
		return nil
	}
	if t.InlineBase64 == nil {
		return nil
	}
	return t.InlineBase64
}

func (t *TemplateInputReference) GetVoice() *VoiceReference {
	if t == nil {
		return nil
	}
	if t.Voice == nil {
		return nil
	}
	return t.Voice
}

func (t *TemplateInputReference) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TemplateInputReference) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	t.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", t)
	}
	switch unmarshaler.Type {
	case "asset":
		value := new(StaticAssetReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Asset = value
	case "generation":
		value := new(GenerationReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Generation = value
	case "inline_base64":
		value := new(InlineBase64Reference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.InlineBase64 = value
	case "voice":
		value := new(VoiceReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Voice = value
	}
	return nil
}

func (t TemplateInputReference) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Asset != nil {
		return internal.MarshalJSONWithDiscriminant(t.Asset, "type", "asset")
	}
	if t.Generation != nil {
		return internal.MarshalJSONWithDiscriminant(t.Generation, "type", "generation")
	}
	if t.InlineBase64 != nil {
		return internal.MarshalJSONWithDiscriminant(t.InlineBase64, "type", "inline_base64")
	}
	if t.Voice != nil {
		return internal.MarshalJSONWithDiscriminant(t.Voice, "type", "voice")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TemplateInputReferenceVisitor interface {
	VisitAsset(*StaticAssetReference) error
	VisitGeneration(*GenerationReference) error
	VisitInlineBase64(*InlineBase64Reference) error
	VisitVoice(*VoiceReference) error
}

func (t *TemplateInputReference) Accept(visitor TemplateInputReferenceVisitor) error {
	if t.Asset != nil {
		return visitor.VisitAsset(t.Asset)
	}
	if t.Generation != nil {
		return visitor.VisitGeneration(t.Generation)
	}
	if t.InlineBase64 != nil {
		return visitor.VisitInlineBase64(t.InlineBase64)
	}
	if t.Voice != nil {
		return visitor.VisitVoice(t.Voice)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TemplateInputReference) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Asset != nil {
		fields = append(fields, "asset")
	}
	if t.Generation != nil {
		fields = append(fields, "generation")
	}
	if t.InlineBase64 != nil {
		fields = append(fields, "inline_base64")
	}
	if t.Voice != nil {
		fields = append(fields, "voice")
	}
	if len(fields) == 0 {
		if t.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", t, t.Type)
		}
		return fmt.Errorf("type %T is empty", t)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", t, fields)
	}
	if t.Type != "" {
		field := fields[0]
		if t.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				t,
				t.Type,
				t,
			)
		}
	}
	return nil
}
