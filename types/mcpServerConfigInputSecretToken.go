package types

import (
	"encoding/json"
	"fmt"
)

// The secret token (Authorization header) stored as a workspace secret or in-place secret
type McpServerConfigInputSecretToken struct {
	ConvAiSecretLocator     *ConvAiSecretLocator
	ConvAiUserSecretDbModel *ConvAiUserSecretDbModel
	typ                     string
}

func (mcpServerConfigInputSecretToken *McpServerConfigInputSecretToken) GetConvAiSecretLocator() *ConvAiSecretLocator {
	if mcpServerConfigInputSecretToken == nil {
		return nil
	}
	if mcpServerConfigInputSecretToken.ConvAiSecretLocator == nil {
		return nil
	}
	return mcpServerConfigInputSecretToken.ConvAiSecretLocator
}

func (mcpServerConfigInputSecretToken *McpServerConfigInputSecretToken) GetConvAiUserSecretDbModel() *ConvAiUserSecretDbModel {
	if mcpServerConfigInputSecretToken == nil {
		return nil
	}
	if mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel == nil {
		return nil
	}
	return mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel
}

func (mcpServerConfigInputSecretToken *McpServerConfigInputSecretToken) UnmarshalJSON(data []byte) error {
	var valueConvAiSecretLocator *ConvAiSecretLocator
	if err := json.Unmarshal(data, &valueConvAiSecretLocator); err == nil {
		mcpServerConfigInputSecretToken.typ = "ConvAiSecretLocator"
		mcpServerConfigInputSecretToken.ConvAiSecretLocator = valueConvAiSecretLocator
		return nil
	}
	var valueConvAiUserSecretDbModel *ConvAiUserSecretDbModel
	if err := json.Unmarshal(data, &valueConvAiUserSecretDbModel); err == nil {
		mcpServerConfigInputSecretToken.typ = "ConvAiUserSecretDbModel"
		mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel = valueConvAiUserSecretDbModel
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, mcpServerConfigInputSecretToken)
}

func (mcpServerConfigInputSecretToken McpServerConfigInputSecretToken) MarshalJSON() ([]byte, error) {
	if mcpServerConfigInputSecretToken.typ == "ConvAiSecretLocator" || mcpServerConfigInputSecretToken.ConvAiSecretLocator != nil {
		return json.Marshal(mcpServerConfigInputSecretToken.ConvAiSecretLocator)
	}
	if mcpServerConfigInputSecretToken.typ == "ConvAiUserSecretDbModel" || mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel != nil {
		return json.Marshal(mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigInputSecretToken)
}

type McpServerConfigInputSecretTokenVisitor interface {
	VisitConvAiSecretLocator(*ConvAiSecretLocator) error
	VisitConvAiUserSecretDbModel(*ConvAiUserSecretDbModel) error
}

func (mcpServerConfigInputSecretToken *McpServerConfigInputSecretToken) Accept(visitor McpServerConfigInputSecretTokenVisitor) error {
	if mcpServerConfigInputSecretToken.typ == "ConvAiSecretLocator" || mcpServerConfigInputSecretToken.ConvAiSecretLocator != nil {
		return visitor.VisitConvAiSecretLocator(mcpServerConfigInputSecretToken.ConvAiSecretLocator)
	}
	if mcpServerConfigInputSecretToken.typ == "ConvAiUserSecretDbModel" || mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel != nil {
		return visitor.VisitConvAiUserSecretDbModel(mcpServerConfigInputSecretToken.ConvAiUserSecretDbModel)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", mcpServerConfigInputSecretToken)
}
