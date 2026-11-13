package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConvAiStoredSecretDependenciesToolsItem struct {
	Type      string
	Available *DependentAvailableToolIdentifier
	Unknown   *DependentUnknownToolIdentifier
}

func (c *ConvAiStoredSecretDependenciesToolsItem) GetAvailable() *DependentAvailableToolIdentifier {
	if c == nil {
		return nil
	}
	if c.Available == nil {
		return nil
	}
	return c.Available
}

func (c *ConvAiStoredSecretDependenciesToolsItem) GetUnknown() *DependentUnknownToolIdentifier {
	if c == nil {
		return nil
	}
	if c.Unknown == nil {
		return nil
	}
	return c.Unknown
}

func (c *ConvAiStoredSecretDependenciesToolsItem) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *ConvAiStoredSecretDependenciesToolsItem) UnmarshalJSON(data []byte) error {
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
	case "available":
		value := new(DependentAvailableToolIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Available = value
	case "unknown":
		value := new(DependentUnknownToolIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Unknown = value
	}
	return nil
}

func (c ConvAiStoredSecretDependenciesToolsItem) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Available != nil {
		return internal.MarshalJSONWithDiscriminant(c.Available, "type", "available")
	}
	if c.Unknown != nil {
		return internal.MarshalJSONWithDiscriminant(c.Unknown, "type", "unknown")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type ConvAiStoredSecretDependenciesToolsItemVisitor interface {
	VisitAvailable(*DependentAvailableToolIdentifier) error
	VisitUnknown(*DependentUnknownToolIdentifier) error
}

func (c *ConvAiStoredSecretDependenciesToolsItem) Accept(visitor ConvAiStoredSecretDependenciesToolsItemVisitor) error {
	if c.Available != nil {
		return visitor.VisitAvailable(c.Available)
	}
	if c.Unknown != nil {
		return visitor.VisitUnknown(c.Unknown)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *ConvAiStoredSecretDependenciesToolsItem) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Available != nil {
		fields = append(fields, "available")
	}
	if c.Unknown != nil {
		fields = append(fields, "unknown")
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
