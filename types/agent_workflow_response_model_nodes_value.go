package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentWorkflowResponseModelNodesValue struct {
	Type            string
	End             *WorkflowEndNodeModelOutput
	OverrideAgent   *WorkflowOverrideAgentNodeModelOutput
	PhoneNumber     *WorkflowPhoneNumberNodeModelOutput
	Say             *WorkflowSayNodeModelOutput
	StandaloneAgent *WorkflowStandaloneAgentNodeModelOutput
	Start           *WorkflowStartNodeModelOutput
	Tool            *WorkflowToolNodeModelOutput
	UpdateState     *WorkflowUpdateStateNodeModelOutput
}

func (a *AgentWorkflowResponseModelNodesValue) GetEnd() *WorkflowEndNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.End == nil {
		return nil
	}
	return a.End
}

func (a *AgentWorkflowResponseModelNodesValue) GetOverrideAgent() *WorkflowOverrideAgentNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.OverrideAgent == nil {
		return nil
	}
	return a.OverrideAgent
}

func (a *AgentWorkflowResponseModelNodesValue) GetPhoneNumber() *WorkflowPhoneNumberNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.PhoneNumber == nil {
		return nil
	}
	return a.PhoneNumber
}

func (a *AgentWorkflowResponseModelNodesValue) GetSay() *WorkflowSayNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.Say == nil {
		return nil
	}
	return a.Say
}

func (a *AgentWorkflowResponseModelNodesValue) GetStandaloneAgent() *WorkflowStandaloneAgentNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.StandaloneAgent == nil {
		return nil
	}
	return a.StandaloneAgent
}

func (a *AgentWorkflowResponseModelNodesValue) GetStart() *WorkflowStartNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.Start == nil {
		return nil
	}
	return a.Start
}

func (a *AgentWorkflowResponseModelNodesValue) GetTool() *WorkflowToolNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.Tool == nil {
		return nil
	}
	return a.Tool
}

func (a *AgentWorkflowResponseModelNodesValue) GetUpdateState() *WorkflowUpdateStateNodeModelOutput {
	if a == nil {
		return nil
	}
	if a.UpdateState == nil {
		return nil
	}
	return a.UpdateState
}

func (a *AgentWorkflowResponseModelNodesValue) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AgentWorkflowResponseModelNodesValue) UnmarshalJSON(data []byte) error {
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
	case "end":
		value := new(WorkflowEndNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.End = value
	case "override_agent":
		value := new(WorkflowOverrideAgentNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.OverrideAgent = value
	case "phone_number":
		value := new(WorkflowPhoneNumberNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.PhoneNumber = value
	case "say":
		value := new(WorkflowSayNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Say = value
	case "standalone_agent":
		value := new(WorkflowStandaloneAgentNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.StandaloneAgent = value
	case "start":
		value := new(WorkflowStartNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Start = value
	case "tool":
		value := new(WorkflowToolNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Tool = value
	case "update_state":
		value := new(WorkflowUpdateStateNodeModelOutput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.UpdateState = value
	}
	return nil
}

func (a AgentWorkflowResponseModelNodesValue) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.End != nil {
		return internal.MarshalJSONWithDiscriminant(a.End, "type", "end")
	}
	if a.OverrideAgent != nil {
		return internal.MarshalJSONWithDiscriminant(a.OverrideAgent, "type", "override_agent")
	}
	if a.PhoneNumber != nil {
		return internal.MarshalJSONWithDiscriminant(a.PhoneNumber, "type", "phone_number")
	}
	if a.Say != nil {
		return internal.MarshalJSONWithDiscriminant(a.Say, "type", "say")
	}
	if a.StandaloneAgent != nil {
		return internal.MarshalJSONWithDiscriminant(a.StandaloneAgent, "type", "standalone_agent")
	}
	if a.Start != nil {
		return internal.MarshalJSONWithDiscriminant(a.Start, "type", "start")
	}
	if a.Tool != nil {
		return internal.MarshalJSONWithDiscriminant(a.Tool, "type", "tool")
	}
	if a.UpdateState != nil {
		return internal.MarshalJSONWithDiscriminant(a.UpdateState, "type", "update_state")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AgentWorkflowResponseModelNodesValueVisitor interface {
	VisitEnd(*WorkflowEndNodeModelOutput) error
	VisitOverrideAgent(*WorkflowOverrideAgentNodeModelOutput) error
	VisitPhoneNumber(*WorkflowPhoneNumberNodeModelOutput) error
	VisitSay(*WorkflowSayNodeModelOutput) error
	VisitStandaloneAgent(*WorkflowStandaloneAgentNodeModelOutput) error
	VisitStart(*WorkflowStartNodeModelOutput) error
	VisitTool(*WorkflowToolNodeModelOutput) error
	VisitUpdateState(*WorkflowUpdateStateNodeModelOutput) error
}

func (a *AgentWorkflowResponseModelNodesValue) Accept(visitor AgentWorkflowResponseModelNodesValueVisitor) error {
	if a.End != nil {
		return visitor.VisitEnd(a.End)
	}
	if a.OverrideAgent != nil {
		return visitor.VisitOverrideAgent(a.OverrideAgent)
	}
	if a.PhoneNumber != nil {
		return visitor.VisitPhoneNumber(a.PhoneNumber)
	}
	if a.Say != nil {
		return visitor.VisitSay(a.Say)
	}
	if a.StandaloneAgent != nil {
		return visitor.VisitStandaloneAgent(a.StandaloneAgent)
	}
	if a.Start != nil {
		return visitor.VisitStart(a.Start)
	}
	if a.Tool != nil {
		return visitor.VisitTool(a.Tool)
	}
	if a.UpdateState != nil {
		return visitor.VisitUpdateState(a.UpdateState)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AgentWorkflowResponseModelNodesValue) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.End != nil {
		fields = append(fields, "end")
	}
	if a.OverrideAgent != nil {
		fields = append(fields, "override_agent")
	}
	if a.PhoneNumber != nil {
		fields = append(fields, "phone_number")
	}
	if a.Say != nil {
		fields = append(fields, "say")
	}
	if a.StandaloneAgent != nil {
		fields = append(fields, "standalone_agent")
	}
	if a.Start != nil {
		fields = append(fields, "start")
	}
	if a.Tool != nil {
		fields = append(fields, "tool")
	}
	if a.UpdateState != nil {
		fields = append(fields, "update_state")
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
