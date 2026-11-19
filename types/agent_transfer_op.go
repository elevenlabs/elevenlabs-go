package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentTransferOp struct {
	Type    string
	Pop     *AgentTransferOpPop
	Push    *AgentTransferOpPush
	Replace *AgentTransferOpReplace
}

func (a *AgentTransferOp) GetPop() *AgentTransferOpPop {
	if a == nil {
		return nil
	}
	if a.Pop == nil {
		return nil
	}
	return a.Pop
}

func (a *AgentTransferOp) GetPush() *AgentTransferOpPush {
	if a == nil {
		return nil
	}
	if a.Push == nil {
		return nil
	}
	return a.Push
}

func (a *AgentTransferOp) GetReplace() *AgentTransferOpReplace {
	if a == nil {
		return nil
	}
	if a.Replace == nil {
		return nil
	}
	return a.Replace
}

func (a *AgentTransferOp) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AgentTransferOp) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	a.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", a)
	}
	switch unmarshaler.Type {
	case "pop":
		value := new(AgentTransferOpPop)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Pop = value
	case "push":
		value := new(AgentTransferOpPush)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Push = value
	case "replace":
		value := new(AgentTransferOpReplace)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Replace = value
	}
	return nil
}

func (a AgentTransferOp) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.Pop != nil {
		return internal.MarshalJSONWithDiscriminant(a.Pop, "type", "pop")
	}
	if a.Push != nil {
		return internal.MarshalJSONWithDiscriminant(a.Push, "type", "push")
	}
	if a.Replace != nil {
		return internal.MarshalJSONWithDiscriminant(a.Replace, "type", "replace")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AgentTransferOpVisitor interface {
	VisitPop(*AgentTransferOpPop) error
	VisitPush(*AgentTransferOpPush) error
	VisitReplace(*AgentTransferOpReplace) error
}

func (a *AgentTransferOp) Accept(visitor AgentTransferOpVisitor) error {
	if a.Pop != nil {
		return visitor.VisitPop(a.Pop)
	}
	if a.Push != nil {
		return visitor.VisitPush(a.Push)
	}
	if a.Replace != nil {
		return visitor.VisitReplace(a.Replace)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AgentTransferOp) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.Pop != nil {
		fields = append(fields, "pop")
	}
	if a.Push != nil {
		fields = append(fields, "push")
	}
	if a.Replace != nil {
		fields = append(fields, "replace")
	}
	if len(fields) == 0 {
		if a.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", a, a.Type)
		}
		return fmt.Errorf("type %T is empty", a)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", a, fields)
	}
	if a.Type != "" {
		field := fields[0]
		if a.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				a,
				a.Type,
				a,
			)
		}
	}
	return nil
}
