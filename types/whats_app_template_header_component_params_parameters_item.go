package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WhatsAppTemplateHeaderComponentParamsParametersItem struct {
	Type     string
	Document *WhatsAppTemplateDocumentParam
	Image    *WhatsAppTemplateImageParam
	Location *WhatsAppTemplateLocationParam
	Text     *WhatsAppTemplateTextParam
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) GetDocument() *WhatsAppTemplateDocumentParam {
	if w == nil {
		return nil
	}
	if w.Document == nil {
		return nil
	}
	return w.Document
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) GetImage() *WhatsAppTemplateImageParam {
	if w == nil {
		return nil
	}
	if w.Image == nil {
		return nil
	}
	return w.Image
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) GetLocation() *WhatsAppTemplateLocationParam {
	if w == nil {
		return nil
	}
	if w.Location == nil {
		return nil
	}
	return w.Location
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) GetText() *WhatsAppTemplateTextParam {
	if w == nil {
		return nil
	}
	if w.Text == nil {
		return nil
	}
	return w.Text
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) UnmarshalJSON(data []byte) error {
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
	case "document":
		value := new(WhatsAppTemplateDocumentParam)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Document = value
	case "image":
		value := new(WhatsAppTemplateImageParam)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Image = value
	case "location":
		value := new(WhatsAppTemplateLocationParam)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Location = value
	case "text":
		value := new(WhatsAppTemplateTextParam)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Text = value
	}
	return nil
}

func (w WhatsAppTemplateHeaderComponentParamsParametersItem) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.Document != nil {
		return internal.MarshalJSONWithDiscriminant(w.Document, "type", "document")
	}
	if w.Image != nil {
		return internal.MarshalJSONWithDiscriminant(w.Image, "type", "image")
	}
	if w.Location != nil {
		return internal.MarshalJSONWithDiscriminant(w.Location, "type", "location")
	}
	if w.Text != nil {
		return internal.MarshalJSONWithDiscriminant(w.Text, "type", "text")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WhatsAppTemplateHeaderComponentParamsParametersItemVisitor interface {
	VisitDocument(*WhatsAppTemplateDocumentParam) error
	VisitImage(*WhatsAppTemplateImageParam) error
	VisitLocation(*WhatsAppTemplateLocationParam) error
	VisitText(*WhatsAppTemplateTextParam) error
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) Accept(visitor WhatsAppTemplateHeaderComponentParamsParametersItemVisitor) error {
	if w.Document != nil {
		return visitor.VisitDocument(w.Document)
	}
	if w.Image != nil {
		return visitor.VisitImage(w.Image)
	}
	if w.Location != nil {
		return visitor.VisitLocation(w.Location)
	}
	if w.Text != nil {
		return visitor.VisitText(w.Text)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WhatsAppTemplateHeaderComponentParamsParametersItem) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.Document != nil {
		fields = append(fields, "document")
	}
	if w.Image != nil {
		fields = append(fields, "image")
	}
	if w.Location != nil {
		fields = append(fields, "location")
	}
	if w.Text != nil {
		fields = append(fields, "text")
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
