package agents

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

// Create Chat Response Test Request Information
type CreateTestsRequestBody struct {
	Type       string
	Llm        *types.CreateResponseUnitTestRequest
	Simulation *types.CreateSimulationTestRequest
	Tool       *types.CreateToolCallUnitTestRequest
}

func (c *CreateTestsRequestBody) GetLlm() *types.CreateResponseUnitTestRequest {
	if c == nil {
		return nil
	}
	if c.Llm == nil {
		return nil
	}
	return c.Llm
}

func (c *CreateTestsRequestBody) GetSimulation() *types.CreateSimulationTestRequest {
	if c == nil {
		return nil
	}
	if c.Simulation == nil {
		return nil
	}
	return c.Simulation
}

func (c *CreateTestsRequestBody) GetTool() *types.CreateToolCallUnitTestRequest {
	if c == nil {
		return nil
	}
	if c.Tool == nil {
		return nil
	}
	return c.Tool
}

func (c *CreateTestsRequestBody) GetType() string {
	if c == nil {
		return ""
	}
	return c.Type
}

func (c *CreateTestsRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	c.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", c)
	}
	switch unmarshaler.Type {
	case "llm":
		value := new(types.CreateResponseUnitTestRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Llm = value
	case "simulation":
		value := new(types.CreateSimulationTestRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Simulation = value
	case "tool":
		value := new(types.CreateToolCallUnitTestRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		c.Tool = value
	}
	return nil
}

func (c CreateTestsRequestBody) MarshalJSON() ([]byte, error) {
	if err := c.validate(); err != nil {
		return nil, err
	}
	if c.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(c.Llm, "type", "llm")
	}
	if c.Simulation != nil {
		return internal.MarshalJSONWithDiscriminant(c.Simulation, "type", "simulation")
	}
	if c.Tool != nil {
		return internal.MarshalJSONWithDiscriminant(c.Tool, "type", "tool")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", c)
}

type CreateTestsRequestBodyVisitor interface {
	VisitLlm(*types.CreateResponseUnitTestRequest) error
	VisitSimulation(*types.CreateSimulationTestRequest) error
	VisitTool(*types.CreateToolCallUnitTestRequest) error
}

func (c *CreateTestsRequestBody) Accept(visitor CreateTestsRequestBodyVisitor) error {
	if c.Llm != nil {
		return visitor.VisitLlm(c.Llm)
	}
	if c.Simulation != nil {
		return visitor.VisitSimulation(c.Simulation)
	}
	if c.Tool != nil {
		return visitor.VisitTool(c.Tool)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", c)
}

func (c *CreateTestsRequestBody) validate() error {
	if c == nil {
		return fmt.Errorf("type %T is nil", c)
	}
	var fields []string
	if c.Llm != nil {
		fields = append(fields, "llm")
	}
	if c.Simulation != nil {
		fields = append(fields, "simulation")
	}
	if c.Tool != nil {
		fields = append(fields, "tool")
	}
	if len(fields) == 0 {
		if c.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", c, c.Type)
		}
		return fmt.Errorf("type %T is empty", c)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", c, fields)
	}
	if c.Type != "" {
		field := fields[0]
		if c.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				c,
				c.Type,
				c,
			)
		}
	}
	return nil
}

type BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost struct {
	// The IDs of tests or folders to move.
	EntityIds []string `json:"entity_ids" xml:"entity_ids" required:"true" minItems:"1" uniqueItems:"true"`
	// The folder to move the entities to. If not set, the entities will be moved to the root folder.
	MoveTo *string `json:"move_to,omitempty" xml:"move_to,omitempty"`
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) GetEntityIds() (value []string) {
	if b == nil {
		return
	}
	return b.EntityIds
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) GetMoveTo() (value *string) {
	if b == nil {
		return
	}
	return b.MoveTo
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) SetEntityIds(entityIds []string) {
	b.EntityIds = entityIds
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) SetMoveTo(moveTo *string) {
	b.MoveTo = moveTo
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost to string"
	}
	return string(jsonData)
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost) MarshalJSON() ([]byte, error) {
	type embed BodyBulkMoveTestsToFolderV1ConvaiAgentTestingBulkMovePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type GetTestsResponse struct {
	Type       string
	Llm        *types.GetResponseUnitTestResponseModel
	Simulation *types.GetSimulationTestResponseModel
	Tool       *types.GetToolCallUnitTestResponseModel
}

func (g *GetTestsResponse) GetLlm() *types.GetResponseUnitTestResponseModel {
	if g == nil {
		return nil
	}
	if g.Llm == nil {
		return nil
	}
	return g.Llm
}

func (g *GetTestsResponse) GetSimulation() *types.GetSimulationTestResponseModel {
	if g == nil {
		return nil
	}
	if g.Simulation == nil {
		return nil
	}
	return g.Simulation
}

func (g *GetTestsResponse) GetTool() *types.GetToolCallUnitTestResponseModel {
	if g == nil {
		return nil
	}
	if g.Tool == nil {
		return nil
	}
	return g.Tool
}

func (g *GetTestsResponse) GetType() string {
	if g == nil {
		return ""
	}
	return g.Type
}

func (g *GetTestsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	g.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", g)
	}
	switch unmarshaler.Type {
	case "llm":
		value := new(types.GetResponseUnitTestResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Llm = value
	case "simulation":
		value := new(types.GetSimulationTestResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Simulation = value
	case "tool":
		value := new(types.GetToolCallUnitTestResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		g.Tool = value
	}
	return nil
}

func (g GetTestsResponse) MarshalJSON() ([]byte, error) {
	if err := g.validate(); err != nil {
		return nil, err
	}
	if g.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(g.Llm, "type", "llm")
	}
	if g.Simulation != nil {
		return internal.MarshalJSONWithDiscriminant(g.Simulation, "type", "simulation")
	}
	if g.Tool != nil {
		return internal.MarshalJSONWithDiscriminant(g.Tool, "type", "tool")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", g)
}

type GetTestsResponseVisitor interface {
	VisitLlm(*types.GetResponseUnitTestResponseModel) error
	VisitSimulation(*types.GetSimulationTestResponseModel) error
	VisitTool(*types.GetToolCallUnitTestResponseModel) error
}

func (g *GetTestsResponse) Accept(visitor GetTestsResponseVisitor) error {
	if g.Llm != nil {
		return visitor.VisitLlm(g.Llm)
	}
	if g.Simulation != nil {
		return visitor.VisitSimulation(g.Simulation)
	}
	if g.Tool != nil {
		return visitor.VisitTool(g.Tool)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", g)
}

func (g *GetTestsResponse) validate() error {
	if g == nil {
		return fmt.Errorf("type %T is nil", g)
	}
	var fields []string
	if g.Llm != nil {
		fields = append(fields, "llm")
	}
	if g.Simulation != nil {
		fields = append(fields, "simulation")
	}
	if g.Tool != nil {
		fields = append(fields, "tool")
	}
	if len(fields) == 0 {
		if g.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", g, g.Type)
		}
		return fmt.Errorf("type %T is empty", g)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", g, fields)
	}
	if g.Type != "" {
		field := fields[0]
		if g.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				g,
				g.Type,
				g,
			)
		}
	}
	return nil
}

type UpdateTestsResponse struct {
	Type       string
	Llm        *types.GetResponseUnitTestResponseModel
	Simulation *types.GetSimulationTestResponseModel
	Tool       *types.GetToolCallUnitTestResponseModel
}

func (u *UpdateTestsResponse) GetLlm() *types.GetResponseUnitTestResponseModel {
	if u == nil {
		return nil
	}
	if u.Llm == nil {
		return nil
	}
	return u.Llm
}

func (u *UpdateTestsResponse) GetSimulation() *types.GetSimulationTestResponseModel {
	if u == nil {
		return nil
	}
	if u.Simulation == nil {
		return nil
	}
	return u.Simulation
}

func (u *UpdateTestsResponse) GetTool() *types.GetToolCallUnitTestResponseModel {
	if u == nil {
		return nil
	}
	if u.Tool == nil {
		return nil
	}
	return u.Tool
}

func (u *UpdateTestsResponse) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UpdateTestsResponse) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", u)
	}
	switch unmarshaler.Type {
	case "llm":
		value := new(types.GetResponseUnitTestResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Llm = value
	case "simulation":
		value := new(types.GetSimulationTestResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Simulation = value
	case "tool":
		value := new(types.GetToolCallUnitTestResponseModel)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Tool = value
	}
	return nil
}

func (u UpdateTestsResponse) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(u.Llm, "type", "llm")
	}
	if u.Simulation != nil {
		return internal.MarshalJSONWithDiscriminant(u.Simulation, "type", "simulation")
	}
	if u.Tool != nil {
		return internal.MarshalJSONWithDiscriminant(u.Tool, "type", "tool")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateTestsResponseVisitor interface {
	VisitLlm(*types.GetResponseUnitTestResponseModel) error
	VisitSimulation(*types.GetSimulationTestResponseModel) error
	VisitTool(*types.GetToolCallUnitTestResponseModel) error
}

func (u *UpdateTestsResponse) Accept(visitor UpdateTestsResponseVisitor) error {
	if u.Llm != nil {
		return visitor.VisitLlm(u.Llm)
	}
	if u.Simulation != nil {
		return visitor.VisitSimulation(u.Simulation)
	}
	if u.Tool != nil {
		return visitor.VisitTool(u.Tool)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateTestsResponse) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Llm != nil {
		fields = append(fields, "llm")
	}
	if u.Simulation != nil {
		fields = append(fields, "simulation")
	}
	if u.Tool != nil {
		fields = append(fields, "tool")
	}
	if len(fields) == 0 {
		if u.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Type)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Type != "" {
		field := fields[0]
		if u.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Type,
				u,
			)
		}
	}
	return nil
}

// Agent test to update
type UpdateTestsRequestBody struct {
	Type       string
	Llm        *types.UpdateResponseUnitTestRequest
	Simulation *types.UpdateSimulationTestRequest
	Tool       *types.UpdateToolCallUnitTestRequest
}

func (u *UpdateTestsRequestBody) GetLlm() *types.UpdateResponseUnitTestRequest {
	if u == nil {
		return nil
	}
	if u.Llm == nil {
		return nil
	}
	return u.Llm
}

func (u *UpdateTestsRequestBody) GetSimulation() *types.UpdateSimulationTestRequest {
	if u == nil {
		return nil
	}
	if u.Simulation == nil {
		return nil
	}
	return u.Simulation
}

func (u *UpdateTestsRequestBody) GetTool() *types.UpdateToolCallUnitTestRequest {
	if u == nil {
		return nil
	}
	if u.Tool == nil {
		return nil
	}
	return u.Tool
}

func (u *UpdateTestsRequestBody) GetType() string {
	if u == nil {
		return ""
	}
	return u.Type
}

func (u *UpdateTestsRequestBody) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	u.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", u)
	}
	switch unmarshaler.Type {
	case "llm":
		value := new(types.UpdateResponseUnitTestRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Llm = value
	case "simulation":
		value := new(types.UpdateSimulationTestRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Simulation = value
	case "tool":
		value := new(types.UpdateToolCallUnitTestRequest)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		u.Tool = value
	}
	return nil
}

func (u UpdateTestsRequestBody) MarshalJSON() ([]byte, error) {
	if err := u.validate(); err != nil {
		return nil, err
	}
	if u.Llm != nil {
		return internal.MarshalJSONWithDiscriminant(u.Llm, "type", "llm")
	}
	if u.Simulation != nil {
		return internal.MarshalJSONWithDiscriminant(u.Simulation, "type", "simulation")
	}
	if u.Tool != nil {
		return internal.MarshalJSONWithDiscriminant(u.Tool, "type", "tool")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", u)
}

type UpdateTestsRequestBodyVisitor interface {
	VisitLlm(*types.UpdateResponseUnitTestRequest) error
	VisitSimulation(*types.UpdateSimulationTestRequest) error
	VisitTool(*types.UpdateToolCallUnitTestRequest) error
}

func (u *UpdateTestsRequestBody) Accept(visitor UpdateTestsRequestBodyVisitor) error {
	if u.Llm != nil {
		return visitor.VisitLlm(u.Llm)
	}
	if u.Simulation != nil {
		return visitor.VisitSimulation(u.Simulation)
	}
	if u.Tool != nil {
		return visitor.VisitTool(u.Tool)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", u)
}

func (u *UpdateTestsRequestBody) validate() error {
	if u == nil {
		return fmt.Errorf("type %T is nil", u)
	}
	var fields []string
	if u.Llm != nil {
		fields = append(fields, "llm")
	}
	if u.Simulation != nil {
		fields = append(fields, "simulation")
	}
	if u.Tool != nil {
		fields = append(fields, "tool")
	}
	if len(fields) == 0 {
		if u.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", u, u.Type)
		}
		return fmt.Errorf("type %T is empty", u)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", u, fields)
	}
	if u.Type != "" {
		field := fields[0]
		if u.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				u,
				u.Type,
				u,
			)
		}
	}
	return nil
}

type ListTestsByIdsRequestModel struct {
	// List of test IDs to fetch. No duplicates allowed. Prefer at most 1000 IDs per request.
	TestIds []string `json:"test_ids" xml:"test_ids" required:"true"`
}

func (l *ListTestsByIdsRequestModel) GetTestIds() (value []string) {
	if l == nil {
		return
	}
	return l.TestIds
}

func (l *ListTestsByIdsRequestModel) SetTestIds(testIds []string) {
	l.TestIds = testIds
}

func (l *ListTestsByIdsRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (l ListTestsByIdsRequestModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTestsByIdsRequestModel to string"
	}
	return string(jsonData)
}

func (l *ListTestsByIdsRequestModel) UnmarshalJSON(data []byte) error {
	type unmarshaler ListTestsByIdsRequestModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListTestsByIdsRequestModel(value)
	return nil
}

func (l *ListTestsByIdsRequestModel) MarshalJSON() ([]byte, error) {
	type embed ListTestsByIdsRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}

type ListTestsRequestSortMode string

const (
	ListTestsRequestSortModeDefault      ListTestsRequestSortMode = "default"
	ListTestsRequestSortModeFoldersFirst ListTestsRequestSortMode = "folders_first"
)

func NewListTestsRequestSortModeFromString(s string) (ListTestsRequestSortMode, error) {
	switch s {
	case "default":
		return ListTestsRequestSortModeDefault, nil
	case "folders_first":
		return ListTestsRequestSortModeFoldersFirst, nil
	}
	var t ListTestsRequestSortMode
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListTestsRequestSortMode) Ptr() *ListTestsRequestSortMode {
	return &e
}

var (
	listTestsRequestFieldCursor          = big.NewInt(1 << 0)
	listTestsRequestFieldPageSize        = big.NewInt(1 << 1)
	listTestsRequestFieldSearch          = big.NewInt(1 << 2)
	listTestsRequestFieldParentFolderId  = big.NewInt(1 << 3)
	listTestsRequestFieldTypes           = big.NewInt(1 << 4)
	listTestsRequestFieldIncludeFolders  = big.NewInt(1 << 5)
	listTestsRequestFieldSortMode        = big.NewInt(1 << 6)
	listTestsRequestFieldSharingMode     = big.NewInt(1 << 7)
	listTestsRequestFieldCreatedByUserId = big.NewInt(1 << 8)
)

type ListTestsRequest struct {
	Cursor          *string                   `json:"-" url:"cursor,omitempty"`
	PageSize        *int                      `json:"-" url:"page_size,omitempty"`
	Search          *string                   `json:"-" url:"search,omitempty"`
	ParentFolderId  *string                   `json:"-" url:"parent_folder_id,omitempty"`
	Types           []*types.TestType         `json:"-" url:"types,omitempty"`
	IncludeFolders  *bool                     `json:"-" url:"include_folders,omitempty"`
	SortMode        *ListTestsRequestSortMode `json:"-" url:"sort_mode,omitempty"`
	SharingMode     *types.TestSharingMode    `json:"-" url:"sharing_mode,omitempty"`
	CreatedByUserId *string                   `json:"-" url:"created_by_user_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListTestsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listTestsRequestFieldCursor)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listTestsRequestFieldPageSize)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listTestsRequestFieldSearch)
}

// SetParentFolderId sets the ParentFolderId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetParentFolderId(parentFolderId *string) {
	l.ParentFolderId = parentFolderId
	l.require(listTestsRequestFieldParentFolderId)
}

// SetTypes sets the Types field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetTypes(types []*types.TestType) {
	l.Types = types
	l.require(listTestsRequestFieldTypes)
}

// SetIncludeFolders sets the IncludeFolders field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetIncludeFolders(includeFolders *bool) {
	l.IncludeFolders = includeFolders
	l.require(listTestsRequestFieldIncludeFolders)
}

// SetSortMode sets the SortMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetSortMode(sortMode *ListTestsRequestSortMode) {
	l.SortMode = sortMode
	l.require(listTestsRequestFieldSortMode)
}

// SetSharingMode sets the SharingMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetSharingMode(sharingMode *types.TestSharingMode) {
	l.SharingMode = sharingMode
	l.require(listTestsRequestFieldSharingMode)
}

// SetCreatedByUserId sets the CreatedByUserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListTestsRequest) SetCreatedByUserId(createdByUserId *string) {
	l.CreatedByUserId = createdByUserId
	l.require(listTestsRequestFieldCreatedByUserId)
}
