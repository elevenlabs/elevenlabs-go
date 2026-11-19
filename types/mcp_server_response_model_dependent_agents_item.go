package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type McpServerResponseModelDependentAgentsItem struct {
	Type      string
	Available *DependentAvailableAgentIdentifier
	Unknown   *DependentUnknownAgentIdentifier
}

func (m *McpServerResponseModelDependentAgentsItem) GetAvailable() *DependentAvailableAgentIdentifier {
	if m == nil {
		return nil
	}
	if m.Available == nil {
		return nil
	}
	return m.Available
}

func (m *McpServerResponseModelDependentAgentsItem) GetUnknown() *DependentUnknownAgentIdentifier {
	if m == nil {
		return nil
	}
	if m.Unknown == nil {
		return nil
	}
	return m.Unknown
}

func (m *McpServerResponseModelDependentAgentsItem) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}

func (m *McpServerResponseModelDependentAgentsItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	m.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", m)
	}
	switch unmarshaler.Type {
	case "available":
		value := new(DependentAvailableAgentIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Available = value
	case "unknown":
		value := new(DependentUnknownAgentIdentifier)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Unknown = value
	}
	return nil
}

func (m McpServerResponseModelDependentAgentsItem) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Available != nil {
		return internal.MarshalJSONWithDiscriminant(m.Available, "type", "available")
	}
	if m.Unknown != nil {
		return internal.MarshalJSONWithDiscriminant(m.Unknown, "type", "unknown")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type McpServerResponseModelDependentAgentsItemVisitor interface {
	VisitAvailable(*DependentAvailableAgentIdentifier) error
	VisitUnknown(*DependentUnknownAgentIdentifier) error
}

func (m *McpServerResponseModelDependentAgentsItem) Accept(visitor McpServerResponseModelDependentAgentsItemVisitor) error {
	if m.Available != nil {
		return visitor.VisitAvailable(m.Available)
	}
	if m.Unknown != nil {
		return visitor.VisitUnknown(m.Unknown)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *McpServerResponseModelDependentAgentsItem) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Available != nil {
		fields = append(fields, "available")
	}
	if m.Unknown != nil {
		fields = append(fields, "unknown")
	}
	if len(fields) == 0 {
		if m.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", m, m.Type)
		}
		return fmt.Errorf("type %T is empty", m)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", m, fields)
	}
	if m.Type != "" {
		field := fields[0]
		if m.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				m,
				m.Type,
				m,
			)
		}
	}
	return nil
}
