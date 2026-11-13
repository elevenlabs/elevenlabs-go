package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenAiSessionConfigOutputToolsItem struct {
	Type     string
	Function *OpenAiFunctionTool
	Mcp      *OpenAimcpTool
}

func (o *OpenAiSessionConfigOutputToolsItem) GetFunction() *OpenAiFunctionTool {
	if o == nil {
		return nil
	}
	if o.Function == nil {
		return nil
	}
	return o.Function
}

func (o *OpenAiSessionConfigOutputToolsItem) GetMcp() *OpenAimcpTool {
	if o == nil {
		return nil
	}
	if o.Mcp == nil {
		return nil
	}
	return o.Mcp
}

func (o *OpenAiSessionConfigOutputToolsItem) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *OpenAiSessionConfigOutputToolsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	o.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", o)
	}
	switch unmarshaler.Type {
	case "function":
		value := new(OpenAiFunctionTool)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		o.Function = value
	case "mcp":
		value := new(OpenAimcpTool)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		o.Mcp = value
	}
	return nil
}

func (o OpenAiSessionConfigOutputToolsItem) MarshalJSON() ([]byte, error) {
	if err := o.validate(); err != nil {
		return nil, err
	}
	if o.Function != nil {
		return internal.MarshalJSONWithDiscriminant(o.Function, "type", "function")
	}
	if o.Mcp != nil {
		return internal.MarshalJSONWithDiscriminant(o.Mcp, "type", "mcp")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", o)
}

type OpenAiSessionConfigOutputToolsItemVisitor interface {
	VisitFunction(*OpenAiFunctionTool) error
	VisitMcp(*OpenAimcpTool) error
}

func (o *OpenAiSessionConfigOutputToolsItem) Accept(visitor OpenAiSessionConfigOutputToolsItemVisitor) error {
	if o.Function != nil {
		return visitor.VisitFunction(o.Function)
	}
	if o.Mcp != nil {
		return visitor.VisitMcp(o.Mcp)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", o)
}

func (o *OpenAiSessionConfigOutputToolsItem) validate() error {
	if o == nil {
		return fmt.Errorf("type %T is nil", o)
	}
	var fields []string
	if o.Function != nil {
		fields = append(fields, "function")
	}
	if o.Mcp != nil {
		fields = append(fields, "mcp")
	}
	if len(fields) == 0 {
		if o.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", o, o.Type)
		}
		return fmt.Errorf("type %T is empty", o)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", o, fields)
	}
	if o.Type != "" {
		field := fields[0]
		if o.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				o,
				o.Type,
				o,
			)
		}
	}
	return nil
}
