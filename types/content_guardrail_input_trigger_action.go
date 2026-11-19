package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentGuardrailInputTriggerAction struct {
	Type    string
	EndCall *EndCallTriggerAction
	Retry   *RetryTriggerAction
}

func (c *ContentGuardrailInputTriggerAction) GetEndCall() *EndCallTriggerAction {
	if c == nil {
		return nil
	}
	if c.EndCall == nil {
		return nil
	}
	return c.EndCall
}

func (c *ContentGuardrailInputTriggerAction) GetRetry() *RetryTriggerAction {
	if c == nil {
		return nil
	}
	if c.Retry == nil {
		return nil
	}
	return c.Retry
}

func (c *ContentGuardrailInputTriggerAction) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ContentGuardrailInputTriggerAction) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", c)
	}
	switch unmarshaler.Type {
	case "end_call":
		value := new(EndCallTriggerAction)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.EndCall = value
	case "retry":
		value := new(RetryTriggerAction)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Retry = value
	}
	return nil
}

func (c ContentGuardrailInputTriggerAction) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.EndCall != nil {
		return internal.MarshalJSONWithDiscriminant(c.EndCall, "type", "end_call")
	}
	if c.Retry != nil {
		return internal.MarshalJSONWithDiscriminant(c.Retry, "type", "retry")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ContentGuardrailInputTriggerActionVisitor interface {
	VisitEndCall(*EndCallTriggerAction) error
	VisitRetry(*RetryTriggerAction) error
}

func (c *ContentGuardrailInputTriggerAction) Accept(visitor ContentGuardrailInputTriggerActionVisitor) error {
	if c.EndCall != nil {
		return visitor.VisitEndCall(c.EndCall)
	}
	if c.Retry != nil {
		return visitor.VisitRetry(c.Retry)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ContentGuardrailInputTriggerAction) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.EndCall != nil {
		fields = append(fields, "end_call")
	}
	if c.Retry != nil {
		fields = append(fields, "retry")
	}
	if len(fields) == 0 {
		if c.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Type)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Type != "" {
		field := fields[0]
		if c.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Type,
				c,
			)
		}
	}
	return nil
}
