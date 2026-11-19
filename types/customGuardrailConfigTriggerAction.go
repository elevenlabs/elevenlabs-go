package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CustomGuardrailConfigTriggerAction struct {
	Type    string
	EndCall *EndCallTriggerAction
	Retry   *RetryTriggerAction
}

func (customGuardrailConfigTriggerAction *CustomGuardrailConfigTriggerAction) GetEndCall() *EndCallTriggerAction {
	if customGuardrailConfigTriggerAction == nil {
		return nil
	}
	if customGuardrailConfigTriggerAction.EndCall == nil {
		return nil
	}
	return customGuardrailConfigTriggerAction.EndCall
}

func (customGuardrailConfigTriggerAction *CustomGuardrailConfigTriggerAction) GetRetry() *RetryTriggerAction {
	if customGuardrailConfigTriggerAction == nil {
		return nil
	}
	if customGuardrailConfigTriggerAction.Retry == nil {
		return nil
	}
	return customGuardrailConfigTriggerAction.Retry
}

func (customGuardrailConfigTriggerAction *CustomGuardrailConfigTriggerAction) GetType() string {
	if customGuardrailConfigTriggerAction == nil {
		return ""
	}
	return customGuardrailConfigTriggerAction.Type
}

func (customGuardrailConfigTriggerAction *CustomGuardrailConfigTriggerAction) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	customGuardrailConfigTriggerAction.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", customGuardrailConfigTriggerAction)
	}
	switch unmarshaler.Type {
	case "end_call":
		value := new(EndCallTriggerAction)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		customGuardrailConfigTriggerAction.EndCall = value
	case "retry":
		value := new(RetryTriggerAction)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		customGuardrailConfigTriggerAction.Retry = value
	}
	return nil
}

func (customGuardrailConfigTriggerAction CustomGuardrailConfigTriggerAction) MarshalJSON() ([]byte, error) {
	if err := customGuardrailConfigTriggerAction.validate(); err != nil {
		return nil, err
	}
	if customGuardrailConfigTriggerAction.EndCall != nil {
		return internal.MarshalJSONWithDiscriminant(customGuardrailConfigTriggerAction.EndCall, "type", "end_call")
	}
	if customGuardrailConfigTriggerAction.Retry != nil {
		return internal.MarshalJSONWithDiscriminant(customGuardrailConfigTriggerAction.Retry, "type", "retry")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", customGuardrailConfigTriggerAction)
}

type CustomGuardrailConfigTriggerActionVisitor interface {
	VisitEndCall(*EndCallTriggerAction) error
	VisitRetry(*RetryTriggerAction) error
}

func (customGuardrailConfigTriggerAction *CustomGuardrailConfigTriggerAction) Accept(visitor CustomGuardrailConfigTriggerActionVisitor) error {
	if customGuardrailConfigTriggerAction.EndCall != nil {
		return visitor.VisitEndCall(customGuardrailConfigTriggerAction.EndCall)
	}
	if customGuardrailConfigTriggerAction.Retry != nil {
		return visitor.VisitRetry(customGuardrailConfigTriggerAction.Retry)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", customGuardrailConfigTriggerAction)
}

func (customGuardrailConfigTriggerAction *CustomGuardrailConfigTriggerAction) validate() error {
	if customGuardrailConfigTriggerAction == nil {
		return fmt.Errorf("type %T is nil", customGuardrailConfigTriggerAction)
	}
	var fields []string
	if customGuardrailConfigTriggerAction.EndCall != nil {
		fields = append(fields, "end_call")
	}
	if customGuardrailConfigTriggerAction.Retry != nil {
		fields = append(fields, "retry")
	}
	if len(fields) == 0 {
		if customGuardrailConfigTriggerAction.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", customGuardrailConfigTriggerAction, customGuardrailConfigTriggerAction.Type)
		}
		return fmt.Errorf("type %T is empty", customGuardrailConfigTriggerAction)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", customGuardrailConfigTriggerAction, fields)
	}
	if customGuardrailConfigTriggerAction.Type != "" {
		field := fields[0]
		if customGuardrailConfigTriggerAction.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				customGuardrailConfigTriggerAction,
				customGuardrailConfigTriggerAction.Type,
				customGuardrailConfigTriggerAction,
			)
		}
	}
	return nil
}
