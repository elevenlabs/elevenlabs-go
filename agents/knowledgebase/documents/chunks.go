package documents

import (
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

var (
	listChunksRequestFieldEmbeddingModel = big.NewInt(1 << 0)
	listChunksRequestFieldPageSize       = big.NewInt(1 << 1)
	listChunksRequestFieldCursor         = big.NewInt(1 << 2)
)

type ListChunksRequest struct {
	EmbeddingModel types.EmbeddingModelEnum `json:"-" url:"embedding_model,omitempty"`
	PageSize       *int                     `json:"-" url:"page_size,omitempty"`
	Cursor         *string                  `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListChunksRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetEmbeddingModel sets the EmbeddingModel field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChunksRequest) SetEmbeddingModel(embeddingModel types.EmbeddingModelEnum) {
	l.EmbeddingModel = embeddingModel
	l.require(listChunksRequestFieldEmbeddingModel)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChunksRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listChunksRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListChunksRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listChunksRequestFieldCursor)
}
