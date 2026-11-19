package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentGuardrailOutputTriggerAction struct {
	Type    string
	EndCall *EndCallTriggerAction
	Retry   *RetryTriggerAction
}

func (contentGuardrailOutputTriggerAction *ContentGuardrailOutputTriggerAction) GetEndCall() *EndCallTriggerAction {
	if contentGuardrailOutputTriggerAction == nil {
		return nil
	}
	if contentGuardrailOutputTriggerAction.EndCall == nil {
		return nil
	}
	return contentGuardrailOutputTriggerAction.EndCall
}

func (contentGuardrailOutputTriggerAction *ContentGuardrailOutputTriggerAction) GetRetry() *RetryTriggerAction {
	if contentGuardrailOutputTriggerAction == nil {
		return nil
	}
	if contentGuardrailOutputTriggerAction.Retry == nil {
		return nil
	}
	return contentGuardrailOutputTriggerAction.Retry
}

func (contentGuardrailOutputTriggerAction *ContentGuardrailOutputTriggerAction) GetType() string {
	if contentGuardrailOutputTriggerAction == nil {
		return ""
	}
	return contentGuardrailOutputTriggerAction.Type
}

func (contentGuardrailOutputTriggerAction *ContentGuardrailOutputTriggerAction) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	contentGuardrailOutputTriggerAction.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", contentGuardrailOutputTriggerAction)
	}
	switch unmarshaler.Type {
	case "end_call":
		value := new(EndCallTriggerAction)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		contentGuardrailOutputTriggerAction.EndCall = value
	case "retry":
		value := new(RetryTriggerAction)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		contentGuardrailOutputTriggerAction.Retry = value
	}
	return nil
}

func (contentGuardrailOutputTriggerAction ContentGuardrailOutputTriggerAction) MarshalJSON() ([]byte, error) {
	if err := contentGuardrailOutputTriggerAction.validate(); err != nil {
		return nil, err
	}
	if contentGuardrailOutputTriggerAction.EndCall != nil {
		return internal.MarshalJSONWithDiscriminant(contentGuardrailOutputTriggerAction.EndCall, "type", "end_call")
	}
	if contentGuardrailOutputTriggerAction.Retry != nil {
		return internal.MarshalJSONWithDiscriminant(contentGuardrailOutputTriggerAction.Retry, "type", "retry")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", contentGuardrailOutputTriggerAction)
}

type ContentGuardrailOutputTriggerActionVisitor interface {
	VisitEndCall(*EndCallTriggerAction) error
	VisitRetry(*RetryTriggerAction) error
}

func (contentGuardrailOutputTriggerAction *ContentGuardrailOutputTriggerAction) Accept(visitor ContentGuardrailOutputTriggerActionVisitor) error {
	if contentGuardrailOutputTriggerAction.EndCall != nil {
		return visitor.VisitEndCall(contentGuardrailOutputTriggerAction.EndCall)
	}
	if contentGuardrailOutputTriggerAction.Retry != nil {
		return visitor.VisitRetry(contentGuardrailOutputTriggerAction.Retry)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", contentGuardrailOutputTriggerAction)
}

func (contentGuardrailOutputTriggerAction *ContentGuardrailOutputTriggerAction) validate() error {
	if contentGuardrailOutputTriggerAction == nil {
		return fmt.Errorf("type %T is nil", contentGuardrailOutputTriggerAction)
	}
	var fields []string
	if contentGuardrailOutputTriggerAction.EndCall != nil {
		fields = append(fields, "end_call")
	}
	if contentGuardrailOutputTriggerAction.Retry != nil {
		fields = append(fields, "retry")
	}
	if len(fields) == 0 {
		if contentGuardrailOutputTriggerAction.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", contentGuardrailOutputTriggerAction, contentGuardrailOutputTriggerAction.Type)
		}
		return fmt.Errorf("type %T is empty", contentGuardrailOutputTriggerAction)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", contentGuardrailOutputTriggerAction, fields)
	}
	if contentGuardrailOutputTriggerAction.Type != "" {
		field := fields[0]
		if contentGuardrailOutputTriggerAction.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				contentGuardrailOutputTriggerAction,
				contentGuardrailOutputTriggerAction.Type,
				contentGuardrailOutputTriggerAction,
			)
		}
	}
	return nil
}
