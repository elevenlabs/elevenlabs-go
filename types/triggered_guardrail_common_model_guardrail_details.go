package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TriggeredGuardrailCommonModelGuardrailDetails struct {
	Type       string
	Classifier *GuardrailClassifierVerdict
	Llm        *GuardrailLlmVerdict
	Unknown    *GuardrailUnknownVerdict
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) GetClassifier() *GuardrailClassifierVerdict {
	if t == nil {
		return nil
	}
	if t.Classifier == nil {
		return nil
	}
	return t.Classifier
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) GetLlm() *GuardrailLlmVerdict {
	if t == nil {
		return nil
	}
	if t.Llm == nil {
		return nil
	}
	return t.Llm
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) GetUnknown() *GuardrailUnknownVerdict {
	if t == nil {
		return nil
	}
	if t.Unknown == nil {
		return nil
	}
	return t.Unknown
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) GetType() string {
	if t == nil {
		return ""
	}
	return t.Type
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) UnmarshalJSON(data []byte) error {
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
	case "classifier":
		value := new(GuardrailClassifierVerdict)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Classifier = value
	case "llm":
		value := new(GuardrailLlmVerdict)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Llm = value
	case "unknown":
		value := new(GuardrailUnknownVerdict)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		t.Unknown = value
	}
	return nil
}

func (t TriggeredGuardrailCommonModelGuardrailDetails) MarshalJSON() ([]byte, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	if t.Classifier != nil {
		return internal.MarshalJSONWithDiscriminant(t.Classifier, "type", "classifier")
	}
	if t.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(t.Llm, "type", "llm")
	}
	if t.Unknown != nil {
		return internal.MarshalJSONWithDiscriminant(t.Unknown, "type", "unknown")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", t)
}

type TriggeredGuardrailCommonModelGuardrailDetailsVisitor interface {
	VisitClassifier(*GuardrailClassifierVerdict) error
	VisitLlm(*GuardrailLlmVerdict) error
	VisitUnknown(*GuardrailUnknownVerdict) error
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) Accept(visitor TriggeredGuardrailCommonModelGuardrailDetailsVisitor) error {
	if t.Classifier != nil {
		return visitor.VisitClassifier(t.Classifier)
	}
	if t.Llm != nil {
		return visitor.VisitLlm(t.Llm)
	}
	if t.Unknown != nil {
		return visitor.VisitUnknown(t.Unknown)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", t)
}

func (t *TriggeredGuardrailCommonModelGuardrailDetails) validate() error {
	if t == nil {
		return fmt.Errorf("type %T is nil", t)
	}
	var fields []string
	if t.Classifier != nil {
		fields = append(fields, "classifier")
	}
	if t.Llm != nil {
		fields = append(fields, "llm")
	}
	if t.Unknown != nil {
		fields = append(fields, "unknown")
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
