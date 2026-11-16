package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AuthConnectionDependenciesMcpServersItem struct {
	Type      string
	Available *DependentAvailableMcpServerIdentifier
	Unknown   *DependentUnknownMcpServerIdentifier
}

func (a *AuthConnectionDependenciesMcpServersItem) GetAvailable() *DependentAvailableMcpServerIdentifier {
	if a == nil {
		return nil
	}
	if a.Available == nil {
		return nil
	}
	return a.Available
}

func (a *AuthConnectionDependenciesMcpServersItem) GetUnknown() *DependentUnknownMcpServerIdentifier {
	if a == nil {
		return nil
	}
	if a.Unknown == nil {
		return nil
	}
	return a.Unknown
}

func (a *AuthConnectionDependenciesMcpServersItem) GetType() string {
	if a == nil {
		return ""
	}
	return a.Type
}

func (a *AuthConnectionDependenciesMcpServersItem) UnmarshalJSON(data []byte) error {
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
	case "available":
		value := new(DependentAvailableMcpServerIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Available = value
	case "unknown":
		value := new(DependentUnknownMcpServerIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		a.Unknown = value
	}
	return nil
}

func (a AuthConnectionDependenciesMcpServersItem) MarshalJSON() ([]byte, error) {
	if err := a.validate(); err != nil {
		return nil, err
	}
	if a.Available != nil {
		return internal.MarshalJSONWithDiscriminant(a.Available, "type", "available")
	}
	if a.Unknown != nil {
		return internal.MarshalJSONWithDiscriminant(a.Unknown, "type", "unknown")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", a)
}

type AuthConnectionDependenciesMcpServersItemVisitor interface {
	VisitAvailable(*DependentAvailableMcpServerIdentifier) error
	VisitUnknown(*DependentUnknownMcpServerIdentifier) error
}

func (a *AuthConnectionDependenciesMcpServersItem) Accept(visitor AuthConnectionDependenciesMcpServersItemVisitor) error {
	if a.Available != nil {
		return visitor.VisitAvailable(a.Available)
	}
	if a.Unknown != nil {
		return visitor.VisitUnknown(a.Unknown)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", a)
}

func (a *AuthConnectionDependenciesMcpServersItem) validate() error {
	if a == nil {
		return fmt.Errorf("type %T is nil", a)
	}
	var fields []string
	if a.Available != nil {
		fields = append(fields, "available")
	}
	if a.Unknown != nil {
		fields = append(fields, "unknown")
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
