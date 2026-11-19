package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type WebhookTarget struct {
	Type string
	All  *WebhookTargetAll
	Ids  *WebhookTargetIds
}

func (w *WebhookTarget) GetAll() *WebhookTargetAll {
	if w == nil {
		return nil
	}
	if w.All == nil {
		return nil
	}
	return w.All
}

func (w *WebhookTarget) GetIds() *WebhookTargetIds {
	if w == nil {
		return nil
	}
	if w.Ids == nil {
		return nil
	}
	return w.Ids
}

func (w *WebhookTarget) GetType() string {
	if w == nil {
		return ""
	}
	return w.Type
}

func (w *WebhookTarget) UnmarshalJSON(data []byte) error {
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
	case "all":
		value := new(WebhookTargetAll)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.All = value
	case "ids":
		value := new(WebhookTargetIds)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		w.Ids = value
	}
	return nil
}

func (w WebhookTarget) MarshalJSON() ([]byte, error) {
	if err := w.validate(); err != nil {
		return nil, err
	}
	if w.All != nil {
		return internal.MarshalJSONWithDiscriminant(w.All, "type", "all")
	}
	if w.Ids != nil {
		return internal.MarshalJSONWithDiscriminant(w.Ids, "type", "ids")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", w)
}

type WebhookTargetVisitor interface {
	VisitAll(*WebhookTargetAll) error
	VisitIds(*WebhookTargetIds) error
}

func (w *WebhookTarget) Accept(visitor WebhookTargetVisitor) error {
	if w.All != nil {
		return visitor.VisitAll(w.All)
	}
	if w.Ids != nil {
		return visitor.VisitIds(w.Ids)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", w)
}

func (w *WebhookTarget) validate() error {
	if w == nil {
		return fmt.Errorf("type %T is nil", w)
	}
	var fields []string
	if w.All != nil {
		fields = append(fields, "all")
	}
	if w.Ids != nil {
		fields = append(fields, "ids")
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
