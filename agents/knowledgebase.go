package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type GetOrCreateRagIndexesKnowledgeBaseResponse struct {
	Status  string
	Failure *types.BatchFailureResponseModel
	Success *types.RagIndexBatchSuccessfulResponseModel
}

func (g *GetOrCreateRagIndexesKnowledgeBaseResponse) GetFailure() *types.BatchFailureResponseModel {
	if g == nil {
		return nil
	}
	if g.Failure == nil {
		return nil
	}
	return g.Failure
}

func (g *GetOrCreateRagIndexesKnowledgeBaseResponse) GetSuccess() *types.RagIndexBatchSuccessfulResponseModel {
	if g == nil {
		return nil
	}
	if g.Success == nil {
		return nil
	}
	return g.Success
}

func (g *GetOrCreateRagIndexesKnowledgeBaseResponse) GetStatus() string {
	if g == nil {
		return ""
	}
	return g.Status
}

func (g *GetOrCreateRagIndexesKnowledgeBaseResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Status = unmarshaler.Status
	if unmarshaler.Status == "" {
		return fmt.Errorf("%T did not include discriminant status", g)
	}
	switch unmarshaler.Status {
	case "failure":
		value := new(types.BatchFailureResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Failure = value
	case "success":
		value := new(types.RagIndexBatchSuccessfulResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Success = value
	}
	return nil
}

func (g GetOrCreateRagIndexesKnowledgeBaseResponse) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Failure != nil {
		return internal.MarshalJSONWithDiscriminant(g.Failure, "status", "failure")
	}
	if g.Success != nil {
		return internal.MarshalJSONWithDiscriminant(g.Success, "status", "success")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetOrCreateRagIndexesKnowledgeBaseResponseVisitor interface {
	VisitFailure(*types.BatchFailureResponseModel) error
	VisitSuccess(*types.RagIndexBatchSuccessfulResponseModel) error
}

func (g *GetOrCreateRagIndexesKnowledgeBaseResponse) Accept(visitor GetOrCreateRagIndexesKnowledgeBaseResponseVisitor) error {
	if g.Failure != nil {
		return visitor.VisitFailure(g.Failure)
	}
	if g.Success != nil {
		return visitor.VisitSuccess(g.Success)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetOrCreateRagIndexesKnowledgeBaseResponse) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Failure != nil {
		fields = append(fields, "failure")
	}
	if g.Success != nil {
		fields = append(fields, "success")
	}
	if len(fields) == 0 {
		if g.Status != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Status)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Status != "" {
		field := fields[0]
		if g.Status != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Status,
				g,
			)
		}
	}
	return nil
}

type BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost struct {
	// List of requested RAG indexes. Minimum 1, maximum 100 items.
	Items []*types.GetOrCreateRagIndexRequestModel `json:"items" xml:"items" required:"true" minItems:"1" maxItems:"100"`
}

func (b *BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost) GetItems() (value []*types.GetOrCreateRagIndexRequestModel) {
	if b == nil {
		return
	}
	return b.Items
}

func (b *BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost) SetItems(items []*types.GetOrCreateRagIndexRequestModel) {
	b.Items = items
}

func (b *BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost to string"
	}
	return string(jsonData)
}

func (b *BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost) UnmarshalJSON(data []byte) error {
	type unmarshaler BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost(value)
	return nil
}

func (b *BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost) MarshalJSON() ([]byte, error) {
	type embed BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listKnowledgeBaseRequestFieldPageSize               = big.NewInt(1 << 0)
	listKnowledgeBaseRequestFieldSearch                 = big.NewInt(1 << 1)
	listKnowledgeBaseRequestFieldShowOnlyOwnedDocuments = big.NewInt(1 << 2)
	listKnowledgeBaseRequestFieldCreatedByUserId        = big.NewInt(1 << 3)
	listKnowledgeBaseRequestFieldTypes                  = big.NewInt(1 << 4)
	listKnowledgeBaseRequestFieldParentFolderId         = big.NewInt(1 << 5)
	listKnowledgeBaseRequestFieldAncestorFolderId       = big.NewInt(1 << 6)
	listKnowledgeBaseRequestFieldFoldersFirst           = big.NewInt(1 << 7)
	listKnowledgeBaseRequestFieldSortDirection          = big.NewInt(1 << 8)
	listKnowledgeBaseRequestFieldSortBy                 = big.NewInt(1 << 9)
	listKnowledgeBaseRequestFieldCursor                 = big.NewInt(1 << 10)
)

type ListKnowledgeBaseRequest struct {
	PageSize               *int                               `json:"-" url:"page_size,omitempty"`
	Search                 *string                            `json:"-" url:"search,omitempty"`
	ShowOnlyOwnedDocuments *bool                              `json:"-" url:"show_only_owned_documents,omitempty"`
	CreatedByUserId        *string                            `json:"-" url:"created_by_user_id,omitempty"`
	Types                  []*types.KnowledgeBaseDocumentType `json:"-" url:"types,omitempty"`
	ParentFolderId         *string                            `json:"-" url:"parent_folder_id,omitempty"`
	AncestorFolderId       *string                            `json:"-" url:"ancestor_folder_id,omitempty"`
	FoldersFirst           *bool                              `json:"-" url:"folders_first,omitempty"`
	SortDirection          *types.SortDirection               `json:"-" url:"sort_direction,omitempty"`
	SortBy                 *types.KnowledgeBaseSortBy         `json:"-" url:"sort_by,omitempty"`
	Cursor                 *string                            `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListKnowledgeBaseRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listKnowledgeBaseRequestFieldPageSize)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listKnowledgeBaseRequestFieldSearch)
}

// SetShowOnlyOwnedDocuments sets the ShowOnlyOwnedDocuments field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetShowOnlyOwnedDocuments(showOnlyOwnedDocuments *bool) {
	l.ShowOnlyOwnedDocuments = showOnlyOwnedDocuments
	l.require(listKnowledgeBaseRequestFieldShowOnlyOwnedDocuments)
}

// SetCreatedByUserId sets the CreatedByUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetCreatedByUserId(createdByUserId *string) {
	l.CreatedByUserId = createdByUserId
	l.require(listKnowledgeBaseRequestFieldCreatedByUserId)
}

// SetTypes sets the Types field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetTypes(types []*types.KnowledgeBaseDocumentType) {
	l.Types = types
	l.require(listKnowledgeBaseRequestFieldTypes)
}

// SetParentFolderId sets the ParentFolderId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetParentFolderId(parentFolderId *string) {
	l.ParentFolderId = parentFolderId
	l.require(listKnowledgeBaseRequestFieldParentFolderId)
}

// SetAncestorFolderId sets the AncestorFolderId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetAncestorFolderId(ancestorFolderId *string) {
	l.AncestorFolderId = ancestorFolderId
	l.require(listKnowledgeBaseRequestFieldAncestorFolderId)
}

// SetFoldersFirst sets the FoldersFirst field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetFoldersFirst(foldersFirst *bool) {
	l.FoldersFirst = foldersFirst
	l.require(listKnowledgeBaseRequestFieldFoldersFirst)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetSortDirection(sortDirection *types.SortDirection) {
	l.SortDirection = sortDirection
	l.require(listKnowledgeBaseRequestFieldSortDirection)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetSortBy(sortBy *types.KnowledgeBaseSortBy) {
	l.SortBy = sortBy
	l.require(listKnowledgeBaseRequestFieldSortBy)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListKnowledgeBaseRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listKnowledgeBaseRequestFieldCursor)
}

var (
	searchKnowledgeBaseRequestFieldQuery    = big.NewInt(1 << 0)
	searchKnowledgeBaseRequestFieldPageSize = big.NewInt(1 << 1)
	searchKnowledgeBaseRequestFieldTypes    = big.NewInt(1 << 2)
	searchKnowledgeBaseRequestFieldCursor   = big.NewInt(1 << 3)
)

type SearchKnowledgeBaseRequest struct {
	Query    string                             `json:"-" url:"query,omitempty"`
	PageSize *int                               `json:"-" url:"page_size,omitempty"`
	Types    []*types.KnowledgeBaseDocumentType `json:"-" url:"types,omitempty"`
	Cursor   *string                            `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (s *SearchKnowledgeBaseRequest) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetQuery sets the Query field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchKnowledgeBaseRequest) SetQuery(query string) {
	s.Query = query
	s.require(searchKnowledgeBaseRequestFieldQuery)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchKnowledgeBaseRequest) SetPageSize(pageSize *int) {
	s.PageSize = pageSize
	s.require(searchKnowledgeBaseRequestFieldPageSize)
}

// SetTypes sets the Types field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchKnowledgeBaseRequest) SetTypes(types []*types.KnowledgeBaseDocumentType) {
	s.Types = types
	s.require(searchKnowledgeBaseRequestFieldTypes)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchKnowledgeBaseRequest) SetCursor(cursor *string) {
	s.Cursor = cursor
	s.require(searchKnowledgeBaseRequestFieldCursor)
}
