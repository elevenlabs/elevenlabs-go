package documents

import (
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

var (
	getChunkRequestFieldEmbeddingModel = big.NewInt(1 << 0)
)

type GetChunkRequest struct {
	EmbeddingModel *types.EmbeddingModelEnum `json:"-" url:"embedding_model,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetChunkRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetEmbeddingModel sets the EmbeddingModel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetChunkRequest) SetEmbeddingModel(embeddingModel *types.EmbeddingModelEnum) {
	g.EmbeddingModel = embeddingModel
	g.require(getChunkRequestFieldEmbeddingModel)
}
