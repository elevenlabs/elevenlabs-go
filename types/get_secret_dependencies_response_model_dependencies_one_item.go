package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type GetSecretDependenciesResponseModelDependenciesOneItem struct {
	Type      string
	Available *DependentAvailableAgentIdentifier
	Unknown   *DependentUnknownAgentIdentifier
}

func (g *GetSecretDependenciesResponseModelDependenciesOneItem) GetAvailable() *DependentAvailableAgentIdentifier {
	if g == nil {
		return nil
	}
	if g.Available == nil {
		return nil
	}
	return g.Available
}

func (g *GetSecretDependenciesResponseModelDependenciesOneItem) GetUnknown() *DependentUnknownAgentIdentifier {
	if g == nil {
		return nil
	}
	if g.Unknown == nil {
		return nil
	}
	return g.Unknown
}

func (g *GetSecretDependenciesResponseModelDependenciesOneItem) GetType() string {
	if g == nil {
		return ""
	}
	return g.Type
}

func (g *GetSecretDependenciesResponseModelDependenciesOneItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", g)
	}
	switch unmarshaler.Type {
	case "available":
		value := new(DependentAvailableAgentIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Available = value
	case "unknown":
		value := new(DependentUnknownAgentIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Unknown = value
	}
	return nil
}

func (g GetSecretDependenciesResponseModelDependenciesOneItem) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Available != nil {
		return internal.MarshalJSONWithDiscriminant(g.Available, "type", "available")
	}
	if g.Unknown != nil {
		return internal.MarshalJSONWithDiscriminant(g.Unknown, "type", "unknown")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetSecretDependenciesResponseModelDependenciesOneItemVisitor interface {
	VisitAvailable(*DependentAvailableAgentIdentifier) error
	VisitUnknown(*DependentUnknownAgentIdentifier) error
}

func (g *GetSecretDependenciesResponseModelDependenciesOneItem) Accept(visitor GetSecretDependenciesResponseModelDependenciesOneItemVisitor) error {
	if g.Available != nil {
		return visitor.VisitAvailable(g.Available)
	}
	if g.Unknown != nil {
		return visitor.VisitUnknown(g.Unknown)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetSecretDependenciesResponseModelDependenciesOneItem) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Available != nil {
		fields = append(fields, "available")
	}
	if g.Unknown != nil {
		fields = append(fields, "unknown")
	}
	if len(fields) == 0 {
		if g.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Type)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Type != "" {
		field := fields[0]
		if g.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Type,
				g,
			)
		}
	}
	return nil
}
