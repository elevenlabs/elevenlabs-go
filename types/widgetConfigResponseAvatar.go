package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The avatar of the widget
type WidgetConfigResponseAvatar struct {
	Type  string
	Image *ImageAvatar
	Orb   *OrbAvatar
	Url   *UrlAvatar
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) GetImage() *ImageAvatar {
	if widgetConfigResponseAvatar == nil {
		return nil
	}
	if widgetConfigResponseAvatar.Image == nil {
		return nil
	}
	return widgetConfigResponseAvatar.Image
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) GetOrb() *OrbAvatar {
	if widgetConfigResponseAvatar == nil {
		return nil
	}
	if widgetConfigResponseAvatar.Orb == nil {
		return nil
	}
	return widgetConfigResponseAvatar.Orb
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) GetUrl() *UrlAvatar {
	if widgetConfigResponseAvatar == nil {
		return nil
	}
	if widgetConfigResponseAvatar.Url == nil {
		return nil
	}
	return widgetConfigResponseAvatar.Url
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) GetType() string {
	if widgetConfigResponseAvatar == nil {
		return ""
	}
	return widgetConfigResponseAvatar.Type
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	widgetConfigResponseAvatar.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", widgetConfigResponseAvatar)
	}
	switch unmarshaler.Type {
	case "image":
		value := new(ImageAvatar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		widgetConfigResponseAvatar.Image = value
	case "orb":
		value := new(OrbAvatar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		widgetConfigResponseAvatar.Orb = value
	case "url":
		value := new(UrlAvatar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		widgetConfigResponseAvatar.Url = value
	}
	return nil
}

func (widgetConfigResponseAvatar WidgetConfigResponseAvatar) MarshalJSON() ([]byte, error) {
	if err := widgetConfigResponseAvatar.validate(); err != nil {
		return nil, err
	}
	if widgetConfigResponseAvatar.Image != nil {
		return internal.MarshalJSONWithDiscriminant(widgetConfigResponseAvatar.Image, "type", "image")
	}
	if widgetConfigResponseAvatar.Orb != nil {
		return internal.MarshalJSONWithDiscriminant(widgetConfigResponseAvatar.Orb, "type", "orb")
	}
	if widgetConfigResponseAvatar.Url != nil {
		return internal.MarshalJSONWithDiscriminant(widgetConfigResponseAvatar.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", widgetConfigResponseAvatar)
}

type WidgetConfigResponseAvatarVisitor interface {
	VisitImage(*ImageAvatar) error
	VisitOrb(*OrbAvatar) error
	VisitUrl(*UrlAvatar) error
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) Accept(visitor WidgetConfigResponseAvatarVisitor) error {
	if widgetConfigResponseAvatar.Image != nil {
		return visitor.VisitImage(widgetConfigResponseAvatar.Image)
	}
	if widgetConfigResponseAvatar.Orb != nil {
		return visitor.VisitOrb(widgetConfigResponseAvatar.Orb)
	}
	if widgetConfigResponseAvatar.Url != nil {
		return visitor.VisitUrl(widgetConfigResponseAvatar.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", widgetConfigResponseAvatar)
}

func (widgetConfigResponseAvatar *WidgetConfigResponseAvatar) validate() error {
	if widgetConfigResponseAvatar == nil {
		return fmt.Errorf("type %T is nil", widgetConfigResponseAvatar)
	}
	var fields []string
	if widgetConfigResponseAvatar.Image != nil {
		fields = append(fields, "image")
	}
	if widgetConfigResponseAvatar.Orb != nil {
		fields = append(fields, "orb")
	}
	if widgetConfigResponseAvatar.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if widgetConfigResponseAvatar.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", widgetConfigResponseAvatar, widgetConfigResponseAvatar.Type)
		}
		return fmt.Errorf("type %T is empty", widgetConfigResponseAvatar)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", widgetConfigResponseAvatar, fields)
	}
	if widgetConfigResponseAvatar.Type != "" {
		field := fields[0]
		if widgetConfigResponseAvatar.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				widgetConfigResponseAvatar,
				widgetConfigResponseAvatar.Type,
				widgetConfigResponseAvatar,
			)
		}
	}
	return nil
}
