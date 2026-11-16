package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioReference struct {
	Type         string
	Asset        *StaticAssetReference
	Generation   *GenerationReference
	InlineBase64 *InlineAudioReference
}

func (a *AudioReference) GetAsset() *StaticAssetReference {
	if a == nil {
		return nil
	}
	if a.Asset == nil {
		return nil
	}
	return a.Asset
}

func (a *AudioReference) GetGeneration() *GenerationReference {
	if a == nil {
		return nil
	}
	if a.Generation == nil {
		return nil
	}
	return a.Generation
}

func (a *AudioReference) GetInlineBase64() *InlineAudioReference {
	if a == nil {
		return nil
	}
	if a.InlineBase64 == nil {
		return nil
	}
	return a.InlineBase64
}

func (a *AudioReference) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AudioReference) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", a)
	}
	switch unmarshaler.Type {
	case "asset":
		value := new(StaticAssetReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Asset = value
	case "generation":
		value := new(GenerationReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Generation = value
	case "inline_base64":
		value := new(InlineAudioReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.InlineBase64 = value
	}
	return nil
}

func (a AudioReference) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.Asset != nil {
		return internal.MarshalJSONWithDiscriminant(a.Asset, "type", "asset")
	}
	if a.Generation != nil {
		return internal.MarshalJSONWithDiscriminant(a.Generation, "type", "generation")
	}
	if a.InlineBase64 != nil {
		return internal.MarshalJSONWithDiscriminant(a.InlineBase64, "type", "inline_base64")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AudioReferenceVisitor interface {
	VisitAsset(*StaticAssetReference) error
	VisitGeneration(*GenerationReference) error
	VisitInlineBase64(*InlineAudioReference) error
}

func (a *AudioReference) Accept(visitor AudioReferenceVisitor) error {
	if a.Asset != nil {
		return visitor.VisitAsset(a.Asset)
	}
	if a.Generation != nil {
		return visitor.VisitGeneration(a.Generation)
	}
	if a.InlineBase64 != nil {
		return visitor.VisitInlineBase64(a.InlineBase64)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AudioReference) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.Asset != nil {
		fields = append(fields, "asset")
	}
	if a.Generation != nil {
		fields = append(fields, "generation")
	}
	if a.InlineBase64 != nil {
		fields = append(fields, "inline_base64")
	}
	if len(fields) == 0 {
		if a.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Type)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Type != "" {
		field := fields[0]
		if a.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Type,
				a,
			)
		}
	}
	return nil
}
