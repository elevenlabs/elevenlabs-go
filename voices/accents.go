package voices

import (
	big "math/big"
)

var (
	getAccentsRequestFieldLanguage = big.NewInt(1 << 0)
	getAccentsRequestFieldModelId  = big.NewInt(1 << 1)
)

type GetAccentsRequest struct {
	Language *string `json:"-" url:"language,omitempty"`
	ModelId  *string `json:"-" url:"model_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetAccentsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetLanguage sets the Language field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccentsRequest) SetLanguage(language *string) {
	g.Language = language
	g.require(getAccentsRequestFieldLanguage)
}

// SetModelId sets the ModelId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetAccentsRequest) SetModelId(modelId *string) {
	g.ModelId = modelId
	g.require(getAccentsRequestFieldModelId)
}
