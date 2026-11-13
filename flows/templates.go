package flows

import (
	big "math/big"
)

var (
	listTemplatesRequestFieldCursor              = big.NewInt(1 << 0)
	listTemplatesRequestFieldPageSize            = big.NewInt(1 << 1)
	listTemplatesRequestFieldVersionsPerTemplate = big.NewInt(1 << 2)
	listTemplatesRequestFieldSearch              = big.NewInt(1 << 3)
)

type ListTemplatesRequest struct {
	Cursor              *string `json:"-" url:"cursor,omitempty"`
	PageSize            *int    `json:"-" url:"page_size,omitempty"`
	VersionsPerTemplate *int    `json:"-" url:"versions_per_template,omitempty"`
	Search              *string `json:"-" url:"search,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListTemplatesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTemplatesRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listTemplatesRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTemplatesRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listTemplatesRequestFieldPageSize)
}

// SetVersionsPerTemplate sets the VersionsPerTemplate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTemplatesRequest) SetVersionsPerTemplate(versionsPerTemplate *int) {
	l.VersionsPerTemplate = versionsPerTemplate
	l.require(listTemplatesRequestFieldVersionsPerTemplate)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTemplatesRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listTemplatesRequestFieldSearch)
}

var (
	getTemplatesRequestFieldVersionsPerTemplate = big.NewInt(1 << 0)
)

type GetTemplatesRequest struct {
	VersionsPerTemplate *int `json:"-" url:"versions_per_template,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetTemplatesRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetVersionsPerTemplate sets the VersionsPerTemplate field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetTemplatesRequest) SetVersionsPerTemplate(versionsPerTemplate *int) {
	g.VersionsPerTemplate = versionsPerTemplate
	g.require(getTemplatesRequestFieldVersionsPerTemplate)
}
