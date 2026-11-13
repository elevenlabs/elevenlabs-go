package agents

import (
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

var (
	listToolsRequestFieldSearch                 = big.NewInt(1 << 0)
	listToolsRequestFieldPageSize               = big.NewInt(1 << 1)
	listToolsRequestFieldShowOnlyOwnedDocuments = big.NewInt(1 << 2)
	listToolsRequestFieldCreatedByUserId        = big.NewInt(1 << 3)
	listToolsRequestFieldTypes                  = big.NewInt(1 << 4)
	listToolsRequestFieldSortDirection          = big.NewInt(1 << 5)
	listToolsRequestFieldSortBy                 = big.NewInt(1 << 6)
	listToolsRequestFieldUsedByAgentId          = big.NewInt(1 << 7)
	listToolsRequestFieldCursor                 = big.NewInt(1 << 8)
)

type ListToolsRequest struct {
	Search                 *string                 `json:"-" url:"search,omitempty"`
	PageSize               *int                    `json:"-" url:"page_size,omitempty"`
	ShowOnlyOwnedDocuments *bool                   `json:"-" url:"show_only_owned_documents,omitempty"`
	CreatedByUserId        *string                 `json:"-" url:"created_by_user_id,omitempty"`
	Types                  []*types.ToolTypeFilter `json:"-" url:"types,omitempty"`
	SortDirection          *types.SortDirection    `json:"-" url:"sort_direction,omitempty"`
	SortBy                 *types.ToolSortBy       `json:"-" url:"sort_by,omitempty"`
	UsedByAgentId          *string                 `json:"-" url:"used_by_agent_id,omitempty"`
	Cursor                 *string                 `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListToolsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listToolsRequestFieldSearch)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listToolsRequestFieldPageSize)
}

// SetShowOnlyOwnedDocuments sets the ShowOnlyOwnedDocuments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetShowOnlyOwnedDocuments(showOnlyOwnedDocuments *bool) {
	l.ShowOnlyOwnedDocuments = showOnlyOwnedDocuments
	l.require(listToolsRequestFieldShowOnlyOwnedDocuments)
}

// SetCreatedByUserId sets the CreatedByUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetCreatedByUserId(createdByUserId *string) {
	l.CreatedByUserId = createdByUserId
	l.require(listToolsRequestFieldCreatedByUserId)
}

// SetTypes sets the Types field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetTypes(types []*types.ToolTypeFilter) {
	l.Types = types
	l.require(listToolsRequestFieldTypes)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetSortDirection(sortDirection *types.SortDirection) {
	l.SortDirection = sortDirection
	l.require(listToolsRequestFieldSortDirection)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetSortBy(sortBy *types.ToolSortBy) {
	l.SortBy = sortBy
	l.require(listToolsRequestFieldSortBy)
}

// SetUsedByAgentId sets the UsedByAgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetUsedByAgentId(usedByAgentId *string) {
	l.UsedByAgentId = usedByAgentId
	l.require(listToolsRequestFieldUsedByAgentId)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListToolsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listToolsRequestFieldCursor)
}

var (
	getToolsRequestFieldEnvironment = big.NewInt(1 << 0)
)

type GetToolsRequest struct {
	Environment *string `json:"-" url:"environment,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetToolsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetEnvironment sets the Environment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetToolsRequest) SetEnvironment(environment *string) {
	g.Environment = environment
	g.require(getToolsRequestFieldEnvironment)
}

var (
	deleteToolsRequestFieldForce = big.NewInt(1 << 0)
)

type DeleteToolsRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteToolsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetForce sets the Force field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteToolsRequest) SetForce(force *bool) {
	d.Force = force
	d.require(deleteToolsRequestFieldForce)
}

var (
	getDependentAgentsToolsRequestFieldCursor   = big.NewInt(1 << 0)
	getDependentAgentsToolsRequestFieldPageSize = big.NewInt(1 << 1)
)

type GetDependentAgentsToolsRequest struct {
	Cursor   *string `json:"-" url:"cursor,omitempty"`
	PageSize *int    `json:"-" url:"page_size,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetDependentAgentsToolsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDependentAgentsToolsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getDependentAgentsToolsRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetDependentAgentsToolsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getDependentAgentsToolsRequestFieldPageSize)
}
