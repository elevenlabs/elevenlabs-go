package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type MemoryEntrySearchResultSource struct {
	Type         string
	Conversation *ConversationSource
	Manual       *ManualSource
}

func (m *MemoryEntrySearchResultSource) GetConversation() *ConversationSource {
	if m == nil {
		return nil
	}
	if m.Conversation == nil {
		return nil
	}
	return m.Conversation
}

func (m *MemoryEntrySearchResultSource) GetManual() *ManualSource {
	if m == nil {
		return nil
	}
	if m.Manual == nil {
		return nil
	}
	return m.Manual
}

func (m *MemoryEntrySearchResultSource) GetType() string {
	if m == nil {
		return ""
	}
	return m.Type
}

func (m *MemoryEntrySearchResultSource) UnmarshalJSON(data []byte) error {
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
	case "conversation":
		value := new(ConversationSource)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Conversation = value
	case "manual":
		value := new(ManualSource)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		m.Manual = value
	}
	return nil
}

func (m MemoryEntrySearchResultSource) MarshalJSON() ([]byte, error) {
	if err := m.validate(); err != nil {
		return nil, err
	}
	if m.Conversation != nil {
		return internal.MarshalJSONWithDiscriminant(m.Conversation, "type", "conversation")
	}
	if m.Manual != nil {
		return internal.MarshalJSONWithDiscriminant(m.Manual, "type", "manual")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", m)
}

type MemoryEntrySearchResultSourceVisitor interface {
	VisitConversation(*ConversationSource) error
	VisitManual(*ManualSource) error
}

func (m *MemoryEntrySearchResultSource) Accept(visitor MemoryEntrySearchResultSourceVisitor) error {
	if m.Conversation != nil {
		return visitor.VisitConversation(m.Conversation)
	}
	if m.Manual != nil {
		return visitor.VisitManual(m.Manual)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", m)
}

func (m *MemoryEntrySearchResultSource) validate() error {
	if m == nil {
		return fmt.Errorf("type %T is nil", m)
	}
	var fields []string
	if m.Conversation != nil {
		fields = append(fields, "conversation")
	}
	if m.Manual != nil {
		fields = append(fields, "manual")
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
