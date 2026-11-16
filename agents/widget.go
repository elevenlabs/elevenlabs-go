package agents

import (
	big "math/big"
)

var (
	getWidgetRequestFieldConversationSignature = big.NewInt(1 << 0)
)

type GetWidgetRequest struct {
	ConversationSignature *string `json:"-" url:"conversation_signature,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetWidgetRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetConversationSignature sets the ConversationSignature field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWidgetRequest) SetConversationSignature(conversationSignature *string) {
	g.ConversationSignature = conversationSignature
	g.require(getWidgetRequestFieldConversationSignature)
}
