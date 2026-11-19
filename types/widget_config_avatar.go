package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The avatar of the widget
type WidgetConfigAvatar struct {
	Type  string
	Image *ImageAvatar
	Orb   *OrbAvatar
	Url   *UrlAvatar
}

func (w *WidgetConfigAvatar) GetImage() *ImageAvatar {
	if w == nil {
		return nil
	}
	if w.Image == nil {
		return nil
	}
	return w.Image
}

func (w *WidgetConfigAvatar) GetOrb() *OrbAvatar {
	if w == nil {
		return nil
	}
	if w.Orb == nil {
		return nil
	}
	return w.Orb
}

func (w *WidgetConfigAvatar) GetUrl() *UrlAvatar {
	if w == nil {
		return nil
	}
	if w.Url == nil {
		return nil
	}
	return w.Url
}

func (w *WidgetConfigAvatar) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WidgetConfigAvatar) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	w.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", w)
	}
	switch unmarshaler.Type {
	case "image":
		value := new(ImageAvatar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Image = value
	case "orb":
		value := new(OrbAvatar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Orb = value
	case "url":
		value := new(UrlAvatar)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Url = value
	}
	return nil
}

func (w WidgetConfigAvatar) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Image != nil {
		return internal.MarshalJSONWithDiscriminant(w.Image, "type", "image")
	}
	if w.Orb != nil {
		return internal.MarshalJSONWithDiscriminant(w.Orb, "type", "orb")
	}
	if w.Url != nil {
		return internal.MarshalJSONWithDiscriminant(w.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WidgetConfigAvatarVisitor interface {
	VisitImage(*ImageAvatar) error
	VisitOrb(*OrbAvatar) error
	VisitUrl(*UrlAvatar) error
}

func (w *WidgetConfigAvatar) Accept(visitor WidgetConfigAvatarVisitor) error {
	if w.Image != nil {
		return visitor.VisitImage(w.Image)
	}
	if w.Orb != nil {
		return visitor.VisitOrb(w.Orb)
	}
	if w.Url != nil {
		return visitor.VisitUrl(w.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WidgetConfigAvatar) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Image != nil {
		fields = append(fields, "image")
	}
	if w.Orb != nil {
		fields = append(fields, "orb")
	}
	if w.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if w.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", w, w.Type)
		}
		return fmt.Errorf("type %T is empty", w)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", w, fields)
	}
	if w.Type != "" {
		field := fields[0]
		if w.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				w,
				w.Type,
				w,
			)
		}
	}
	return nil
}
