package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Definition of the analysis item (evaluation or data collection).
type CreateAnalysisItemsRequestBody struct {
	Kind           string
	DataCollection *types.UserDataCollectionAnalysisItemInput
	Evaluation     *types.UserEvaluationAnalysisItem
}

func (c *CreateAnalysisItemsRequestBody) GetDataCollection() *types.UserDataCollectionAnalysisItemInput {
	if c == nil {
		return nil
	}
	if c.DataCollection == nil {
		return nil
	}
	return c.DataCollection
}

func (c *CreateAnalysisItemsRequestBody) GetEvaluation() *types.UserEvaluationAnalysisItem {
	if c == nil {
		return nil
	}
	if c.Evaluation == nil {
		return nil
	}
	return c.Evaluation
}

func (c *CreateAnalysisItemsRequestBody) GetKind() string {
	if c == nil {
		return ""
	}
	return c.Kind
}

func (c *CreateAnalysisItemsRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", c)
	}
	switch unmarshaler.Kind {
	case "data_collection":
		value := new(types.UserDataCollectionAnalysisItemInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.DataCollection = value
	case "evaluation":
		value := new(types.UserEvaluationAnalysisItem)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Evaluation = value
	}
	return nil
}

func (c CreateAnalysisItemsRequestBody) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.DataCollection != nil {
		return internal.MarshalJSONWithDiscriminant(c.DataCollection, "kind", "data_collection")
	}
	if c.Evaluation != nil {
		return internal.MarshalJSONWithDiscriminant(c.Evaluation, "kind", "evaluation")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CreateAnalysisItemsRequestBodyVisitor interface {
	VisitDataCollection(*types.UserDataCollectionAnalysisItemInput) error
	VisitEvaluation(*types.UserEvaluationAnalysisItem) error
}

func (c *CreateAnalysisItemsRequestBody) Accept(visitor CreateAnalysisItemsRequestBodyVisitor) error {
	if c.DataCollection != nil {
		return visitor.VisitDataCollection(c.DataCollection)
	}
	if c.Evaluation != nil {
		return visitor.VisitEvaluation(c.Evaluation)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CreateAnalysisItemsRequestBody) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.DataCollection != nil {
		fields = append(fields, "data_collection")
	}
	if c.Evaluation != nil {
		fields = append(fields, "evaluation")
	}
	if len(fields) == 0 {
		if c.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Kind)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Kind != "" {
		field := fields[0]
		if c.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Kind,
				c,
			)
		}
	}
	return nil
}

type UpdateAnalysisItemAttachmentsRequestModel struct {
	// ID of the requested analysis item.
	AnalysisItemId string                                     `json:"-" url:"-"`
	Source         types.AnalysisItemSource                   `json:"source" xml:"source" required:"true"`
	Attach         []*types.AnalysisItemAttachmentDestination `json:"attach,omitempty" xml:"attach,omitempty"`
	Detach         []*types.AnalysisItemAttachmentDestination `json:"detach,omitempty" xml:"detach,omitempty"`
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) GetAnalysisItemId() (value string) {
	if u == nil {
		return
	}
	return u.AnalysisItemId
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) GetSource() (value types.AnalysisItemSource) {
	if u == nil {
		return
	}
	return u.Source
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) GetAttach() (value []*types.AnalysisItemAttachmentDestination) {
	if u == nil {
		return
	}
	return u.Attach
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) GetDetach() (value []*types.AnalysisItemAttachmentDestination) {
	if u == nil {
		return
	}
	return u.Detach
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) SetAnalysisItemId(analysisItemId string) {
	u.AnalysisItemId = analysisItemId
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) SetSource(source types.AnalysisItemSource) {
	u.Source = source
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) SetAttach(attach []*types.AnalysisItemAttachmentDestination) {
	u.Attach = attach
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) SetDetach(detach []*types.AnalysisItemAttachmentDestination) {
	u.Detach = detach
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateAnalysisItemAttachmentsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAnalysisItemAttachmentsRequestModel to string"
	}
	return string(jsonData)
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UpdateAnalysisItemAttachmentsRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UpdateAnalysisItemAttachmentsRequestModel(tmp)
	return nil
}

func (u *UpdateAnalysisItemAttachmentsRequestModel) MarshalJSON() ([]byte, error) {
	type embed UpdateAnalysisItemAttachmentsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}

// Definition of the analysis item (evaluation or data collection).
type UpdateAnalysisItemsRequestBody struct {
	Kind           string
	DataCollection *types.UserDataCollectionAnalysisItemInput
	Evaluation     *types.UserEvaluationAnalysisItem
}

func (u *UpdateAnalysisItemsRequestBody) GetDataCollection() *types.UserDataCollectionAnalysisItemInput {
	if u == nil {
		return nil
	}
	if u.DataCollection == nil {
		return nil
	}
	return u.DataCollection
}

func (u *UpdateAnalysisItemsRequestBody) GetEvaluation() *types.UserEvaluationAnalysisItem {
	if u == nil {
		return nil
	}
	if u.Evaluation == nil {
		return nil
	}
	return u.Evaluation
}

func (u *UpdateAnalysisItemsRequestBody) GetKind() string {
	if u == nil {
		return ""
	}
	return u.Kind
}

func (u *UpdateAnalysisItemsRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Kind string `json:"kind"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Kind = unmarshaler.Kind
	if unmarshaler.Kind == "" {
		return fmt.Errorf("%T did not include discriminant kind", u)
	}
	switch unmarshaler.Kind {
	case "data_collection":
		value := new(types.UserDataCollectionAnalysisItemInput)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.DataCollection = value
	case "evaluation":
		value := new(types.UserEvaluationAnalysisItem)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Evaluation = value
	}
	return nil
}

func (u UpdateAnalysisItemsRequestBody) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.DataCollection != nil {
		return internal.MarshalJSONWithDiscriminant(u.DataCollection, "kind", "data_collection")
	}
	if u.Evaluation != nil {
		return internal.MarshalJSONWithDiscriminant(u.Evaluation, "kind", "evaluation")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateAnalysisItemsRequestBodyVisitor interface {
	VisitDataCollection(*types.UserDataCollectionAnalysisItemInput) error
	VisitEvaluation(*types.UserEvaluationAnalysisItem) error
}

func (u *UpdateAnalysisItemsRequestBody) Accept(visitor UpdateAnalysisItemsRequestBodyVisitor) error {
	if u.DataCollection != nil {
		return visitor.VisitDataCollection(u.DataCollection)
	}
	if u.Evaluation != nil {
		return visitor.VisitEvaluation(u.Evaluation)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateAnalysisItemsRequestBody) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.DataCollection != nil {
		fields = append(fields, "data_collection")
	}
	if u.Evaluation != nil {
		fields = append(fields, "evaluation")
	}
	if len(fields) == 0 {
		if u.Kind != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Kind)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Kind != "" {
		field := fields[0]
		if u.Kind != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Kind,
				u,
			)
		}
	}
	return nil
}

var (
	listAnalysisItemsRequestFieldSearch          = big.NewInt(1 << 0)
	listAnalysisItemsRequestFieldPageSize        = big.NewInt(1 << 1)
	listAnalysisItemsRequestFieldKind            = big.NewInt(1 << 2)
	listAnalysisItemsRequestFieldAnalysisItemIds = big.NewInt(1 << 3)
	listAnalysisItemsRequestFieldCreatedByUserId = big.NewInt(1 << 4)
	listAnalysisItemsRequestFieldSortDirection   = big.NewInt(1 << 5)
	listAnalysisItemsRequestFieldCursor          = big.NewInt(1 << 6)
)

type ListAnalysisItemsRequest struct {
	Search          *string                 `json:"-" url:"search,omitempty"`
	PageSize        *int                    `json:"-" url:"page_size,omitempty"`
	Kind            *types.AnalysisItemKind `json:"-" url:"kind,omitempty"`
	AnalysisItemIds []*string               `json:"-" url:"analysis_item_ids,omitempty"`
	CreatedByUserId *string                 `json:"-" url:"created_by_user_id,omitempty"`
	SortDirection   *types.SortDirection    `json:"-" url:"sort_direction,omitempty"`
	Cursor          *string                 `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListAnalysisItemsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listAnalysisItemsRequestFieldSearch)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listAnalysisItemsRequestFieldPageSize)
}

// SetKind sets the Kind field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetKind(kind *types.AnalysisItemKind) {
	l.Kind = kind
	l.require(listAnalysisItemsRequestFieldKind)
}

// SetAnalysisItemIds sets the AnalysisItemIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetAnalysisItemIds(analysisItemIds []*string) {
	l.AnalysisItemIds = analysisItemIds
	l.require(listAnalysisItemsRequestFieldAnalysisItemIds)
}

// SetCreatedByUserId sets the CreatedByUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetCreatedByUserId(createdByUserId *string) {
	l.CreatedByUserId = createdByUserId
	l.require(listAnalysisItemsRequestFieldCreatedByUserId)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetSortDirection(sortDirection *types.SortDirection) {
	l.SortDirection = sortDirection
	l.require(listAnalysisItemsRequestFieldSortDirection)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListAnalysisItemsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listAnalysisItemsRequestFieldCursor)
}

var (
	deleteAnalysisItemsRequestFieldForce = big.NewInt(1 << 0)
)

type DeleteAnalysisItemsRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteAnalysisItemsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetForce sets the Force field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteAnalysisItemsRequest) SetForce(force *bool) {
	d.Force = force
	d.require(deleteAnalysisItemsRequestFieldForce)
}

var (
	listVersionsAnalysisItemsRequestFieldPageSize = big.NewInt(1 << 0)
	listVersionsAnalysisItemsRequestFieldCursor   = big.NewInt(1 << 1)
)

type ListVersionsAnalysisItemsRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListVersionsAnalysisItemsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVersionsAnalysisItemsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listVersionsAnalysisItemsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListVersionsAnalysisItemsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listVersionsAnalysisItemsRequestFieldCursor)
}
