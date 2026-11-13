package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type VideoReference struct {
	Type         string
	Asset        *StaticAssetReference
	Generation   *GenerationReference
	InlineBase64 *InlineVideoReference
}

func (v *VideoReference) GetAsset() *StaticAssetReference {
	if v == nil {
		return nil
	}
	if v.Asset == nil {
		return nil
	}
	return v.Asset
}

func (v *VideoReference) GetGeneration() *GenerationReference {
	if v == nil {
		return nil
	}
	if v.Generation == nil {
		return nil
	}
	return v.Generation
}

func (v *VideoReference) GetInlineBase64() *InlineVideoReference {
	if v == nil {
		return nil
	}
	if v.InlineBase64 == nil {
		return nil
	}
	return v.InlineBase64
}

func (v *VideoReference) GetType() string {
	if v == nil {
		return ""
	}
	return v.Type
}

func (v *VideoReference) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	v.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", v)
	}
	switch unmarshaler.Type {
	case "asset":
		value := new(StaticAssetReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.Asset = value
	case "generation":
		value := new(GenerationReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.Generation = value
	case "inline_base64":
		value := new(InlineVideoReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		v.InlineBase64 = value
	}
	return nil
}

func (v VideoReference) MarshalJSON() ([]byte, error) {
	if err := v.validate(); err != nil {
		return nil, err
	}
	if v.Asset != nil {
		return internal.MarshalJSONWithDiscriminant(v.Asset, "type", "asset")
	}
	if v.Generation != nil {
		return internal.MarshalJSONWithDiscriminant(v.Generation, "type", "generation")
	}
	if v.InlineBase64 != nil {
		return internal.MarshalJSONWithDiscriminant(v.InlineBase64, "type", "inline_base64")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", v)
}

type VideoReferenceVisitor interface {
	VisitAsset(*StaticAssetReference) error
	VisitGeneration(*GenerationReference) error
	VisitInlineBase64(*InlineVideoReference) error
}

func (v *VideoReference) Accept(visitor VideoReferenceVisitor) error {
	if v.Asset != nil {
		return visitor.VisitAsset(v.Asset)
	}
	if v.Generation != nil {
		return visitor.VisitGeneration(v.Generation)
	}
	if v.InlineBase64 != nil {
		return visitor.VisitInlineBase64(v.InlineBase64)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", v)
}

func (v *VideoReference) validate() error {
	if v == nil {
		return fmt.Errorf("type %T is nil", v)
	}
	var fields []string
	if v.Asset != nil {
		fields = append(fields, "asset")
	}
	if v.Generation != nil {
		fields = append(fields, "generation")
	}
	if v.InlineBase64 != nil {
		fields = append(fields, "inline_base64")
	}
	if len(fields) == 0 {
		if v.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", v, v.Type)
		}
		return fmt.Errorf("type %T is empty", v)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", v, fields)
	}
	if v.Type != "" {
		field := fields[0]
		if v.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				v,
				v.Type,
				v,
			)
		}
	}
	return nil
}
