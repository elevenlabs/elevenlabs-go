package conversations

import (
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

var (
	getTopicsRequestFieldPageSize                  = big.NewInt(1 << 0)
	getTopicsRequestFieldSortBy                    = big.NewInt(1 << 1)
	getTopicsRequestFieldSortDirection             = big.NewInt(1 << 2)
	getTopicsRequestFieldFromUnixSecs              = big.NewInt(1 << 3)
	getTopicsRequestFieldToUnixSecs                = big.NewInt(1 << 4)
	getTopicsRequestFieldIncludeEvaluationCriteria = big.NewInt(1 << 5)
	getTopicsRequestFieldCursor                    = big.NewInt(1 << 6)
)

type GetTopicsRequest struct {
	PageSize                  *int                 `json:"-" url:"page_size,omitempty"`
	SortBy                    *types.TopicSortBy   `json:"-" url:"sort_by,omitempty"`
	SortDirection             *types.SortDirection `json:"-" url:"sort_direction,omitempty"`
	FromUnixSecs              *int                 `json:"-" url:"from_unix_secs,omitempty"`
	ToUnixSecs                *int                 `json:"-" url:"to_unix_secs,omitempty"`
	IncludeEvaluationCriteria *bool                `json:"-" url:"include_evaluation_criteria,omitempty"`
	Cursor                    *string              `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetTopicsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getTopicsRequestFieldPageSize)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetSortBy(sortBy *types.TopicSortBy) {
	g.SortBy = sortBy
	g.require(getTopicsRequestFieldSortBy)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetSortDirection(sortDirection *types.SortDirection) {
	g.SortDirection = sortDirection
	g.require(getTopicsRequestFieldSortDirection)
}

// SetFromUnixSecs sets the FromUnixSecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetFromUnixSecs(fromUnixSecs *int) {
	g.FromUnixSecs = fromUnixSecs
	g.require(getTopicsRequestFieldFromUnixSecs)
}

// SetToUnixSecs sets the ToUnixSecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetToUnixSecs(toUnixSecs *int) {
	g.ToUnixSecs = toUnixSecs
	g.require(getTopicsRequestFieldToUnixSecs)
}

// SetIncludeEvaluationCriteria sets the IncludeEvaluationCriteria field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetIncludeEvaluationCriteria(includeEvaluationCriteria *bool) {
	g.IncludeEvaluationCriteria = includeEvaluationCriteria
	g.require(getTopicsRequestFieldIncludeEvaluationCriteria)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTopicsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getTopicsRequestFieldCursor)
}
