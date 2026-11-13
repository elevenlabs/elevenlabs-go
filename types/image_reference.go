package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ImageReference struct {
	Type         string
	Asset        *StaticAssetReference
	Generation   *GenerationReference
	InlineBase64 *InlineImageReference
}

func (i *ImageReference) GetAsset() *StaticAssetReference {
	if i == nil {
		return nil
	}
	if i.Asset == nil {
		return nil
	}
	return i.Asset
}

func (i *ImageReference) GetGeneration() *GenerationReference {
	if i == nil {
		return nil
	}
	if i.Generation == nil {
		return nil
	}
	return i.Generation
}

func (i *ImageReference) GetInlineBase64() *InlineImageReference {
	if i == nil {
		return nil
	}
	if i.InlineBase64 == nil {
		return nil
	}
	return i.InlineBase64
}

func (i *ImageReference) GetType() string {
	if i == nil {
		return ""
	}
	return i.Type
}

func (i *ImageReference) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	i.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", i)
	}
	switch unmarshaler.Type {
	case "asset":
		value := new(StaticAssetReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.Asset = value
	case "generation":
		value := new(GenerationReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.Generation = value
	case "inline_base64":
		value := new(InlineImageReference)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		i.InlineBase64 = value
	}
	return nil
}

func (i ImageReference) MarshalJSON() ([]byte, error) {
	if err := i.validate(); err != nil {
		return nil, err
	}
	if i.Asset != nil {
		return internal.MarshalJSONWithDiscriminant(i.Asset, "type", "asset")
	}
	if i.Generation != nil {
		return internal.MarshalJSONWithDiscriminant(i.Generation, "type", "generation")
	}
	if i.InlineBase64 != nil {
		return internal.MarshalJSONWithDiscriminant(i.InlineBase64, "type", "inline_base64")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", i)
}

type ImageReferenceVisitor interface {
	VisitAsset(*StaticAssetReference) error
	VisitGeneration(*GenerationReference) error
	VisitInlineBase64(*InlineImageReference) error
}

func (i *ImageReference) Accept(visitor ImageReferenceVisitor) error {
	if i.Asset != nil {
		return visitor.VisitAsset(i.Asset)
	}
	if i.Generation != nil {
		return visitor.VisitGeneration(i.Generation)
	}
	if i.InlineBase64 != nil {
		return visitor.VisitInlineBase64(i.InlineBase64)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", i)
}

func (i *ImageReference) validate() error {
	if i == nil {
		return fmt.Errorf("type %T is nil", i)
	}
	var fields []string
	if i.Asset != nil {
		fields = append(fields, "asset")
	}
	if i.Generation != nil {
		fields = append(fields, "generation")
	}
	if i.InlineBase64 != nil {
		fields = append(fields, "inline_base64")
	}
	if len(fields) == 0 {
		if i.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", i, i.Type)
		}
		return fmt.Errorf("type %T is empty", i)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", i, fields)
	}
	if i.Type != "" {
		field := fields[0]
		if i.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				i,
				i.Type,
				i,
			)
		}
	}
	return nil
}
