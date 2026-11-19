package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UnitTestRunResponseModelTestInfo struct {
	Type       string
	Llm        *ResponseUnitTestModel
	Simulation *SimulationTestModel
	Tool       *ToolCallUnitTestModel
}

func (u *UnitTestRunResponseModelTestInfo) GetLlm() *ResponseUnitTestModel {
	if u == nil {
		return nil
	}
	if u.Llm == nil {
		return nil
	}
	return u.Llm
}

func (u *UnitTestRunResponseModelTestInfo) GetSimulation() *SimulationTestModel {
	if u == nil {
		return nil
	}
	if u.Simulation == nil {
		return nil
	}
	return u.Simulation
}

func (u *UnitTestRunResponseModelTestInfo) GetTool() *ToolCallUnitTestModel {
	if u == nil {
		return nil
	}
	if u.Tool == nil {
		return nil
	}
	return u.Tool
}

func (u *UnitTestRunResponseModelTestInfo) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UnitTestRunResponseModelTestInfo) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", u)
	}
	switch unmarshaler.Type {
	case "llm":
		value := new(ResponseUnitTestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Llm = value
	case "simulation":
		value := new(SimulationTestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Simulation = value
	case "tool":
		value := new(ToolCallUnitTestModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Tool = value
	}
	return nil
}

func (u UnitTestRunResponseModelTestInfo) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(u.Llm, "type", "llm")
	}
	if u.Simulation != nil {
		return internal.MarshalJSONWithDiscriminant(u.Simulation, "type", "simulation")
	}
	if u.Tool != nil {
		return internal.MarshalJSONWithDiscriminant(u.Tool, "type", "tool")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UnitTestRunResponseModelTestInfoVisitor interface {
	VisitLlm(*ResponseUnitTestModel) error
	VisitSimulation(*SimulationTestModel) error
	VisitTool(*ToolCallUnitTestModel) error
}

func (u *UnitTestRunResponseModelTestInfo) Accept(visitor UnitTestRunResponseModelTestInfoVisitor) error {
	if u.Llm != nil {
		return visitor.VisitLlm(u.Llm)
	}
	if u.Simulation != nil {
		return visitor.VisitSimulation(u.Simulation)
	}
	if u.Tool != nil {
		return visitor.VisitTool(u.Tool)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UnitTestRunResponseModelTestInfo) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Llm != nil {
		fields = append(fields, "llm")
	}
	if u.Simulation != nil {
		fields = append(fields, "simulation")
	}
	if u.Tool != nil {
		fields = append(fields, "tool")
	}
	if len(fields) == 0 {
		if u.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Type)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Type != "" {
		field := fields[0]
		if u.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Type,
				u,
			)
		}
	}
	return nil
}
