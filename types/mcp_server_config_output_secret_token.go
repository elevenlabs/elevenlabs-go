package types

import (
	"encoding/json"
	"fmt"
)

// The secret token (Authorization header) stored as a workspace secret or in-place secret
type McpServerConfigOutputSecretToken struct {
	ConvAiSecretLocator     *ConvAiSecretLocator
	ConvAiUserSecretDbModel *ConvAiUserSecretDbModel
	typ                     string
}

func (m *McpServerConfigOutputSecretToken) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if m == nil {
		return nil
	}
	if m.ConvAiSecretLocator == nil {
		return nil
	}
	return m.ConvAiSecretLocator
}

func (m *McpServerConfigOutputSecretToken) GetConvAiUserSecretDbModel() *ConvAiUserSecretDbModel {
	if m == nil {
		return nil
	}
	if m.ConvAiUserSecretDbModel == nil {
		return nil
	}
	return m.ConvAiUserSecretDbModel
}

func (m *McpServerConfigOutputSecretToken) UnmarshalJSON(data []byte) error {
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		m.typ = "ConvAiSecretLocator"
		m.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiUserSecretDbModel *ConvAiUserSecretDbModel
	if err := json.Unmarshal(data, &valueConvAiUserSecretDbModel); err == nil {
		m.typ = "ConvAiUserSecretDbModel"
		m.ConvAiUserSecretDbModel = valueConvAiUserSecretDbModel
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m McpServerConfigOutputSecretToken) MarshalJSON() ([]byte, error) {
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return json.Marshal(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiUserSecretDbModel" || m.ConvAiUserSecretDbModel != nil {
		return json.Marshal(m.ConvAiUserSecretDbModel)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type McpServerConfigOutputSecretTokenVisitor interface {
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiUserSecretDbModel(*ConvAiUserSecretDbModel) error
}

func (m *McpServerConfigOutputSecretToken) Accept(visitor McpServerConfigOutputSecretTokenVisitor) error {
	if m.typ == "ConvAiSecretLocator" || m.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(m.ConvAiSecretLocator)
	}
	if m.typ == "ConvAiUserSecretDbModel" || m.ConvAiUserSecretDbModel != nil {
		return visitor.VisitConvAiUserSecretDbModel(m.ConvAiUserSecretDbModel)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
