package agents

import (
	"fmt"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type ListConversationsRequestExcludeStatusesItem string

const (
	ListConversationsRequestExcludeStatusesItemInitiated  ListConversationsRequestExcludeStatusesItem = "initiated"
	ListConversationsRequestExcludeStatusesItemInProgress ListConversationsRequestExcludeStatusesItem = "in-progress"
	ListConversationsRequestExcludeStatusesItemProcessing ListConversationsRequestExcludeStatusesItem = "processing"
	ListConversationsRequestExcludeStatusesItemDone       ListConversationsRequestExcludeStatusesItem = "done"
	ListConversationsRequestExcludeStatusesItemFailed     ListConversationsRequestExcludeStatusesItem = "failed"
)

func NewListConversationsRequestExcludeStatusesItemFromString(s string) (ListConversationsRequestExcludeStatusesItem, error) {
	switch s {
	case "initiated":
		return ListConversationsRequestExcludeStatusesItemInitiated, nil
	case "in-progress":
		return ListConversationsRequestExcludeStatusesItemInProgress, nil
	case "processing":
		return ListConversationsRequestExcludeStatusesItemProcessing, nil
	case "done":
		return ListConversationsRequestExcludeStatusesItemDone, nil
	case "failed":
		return ListConversationsRequestExcludeStatusesItemFailed, nil
	}
	var t ListConversationsRequestExcludeStatusesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListConversationsRequestExcludeStatusesItem) Ptr() *ListConversationsRequestExcludeStatusesItem {
	return &e
}

type ListConversationsRequestSummaryMode string

const (
	ListConversationsRequestSummaryModeExclude ListConversationsRequestSummaryMode = "exclude"
	ListConversationsRequestSummaryModeInclude ListConversationsRequestSummaryMode = "include"
)

func NewListConversationsRequestSummaryModeFromString(s string) (ListConversationsRequestSummaryMode, error) {
	switch s {
	case "exclude":
		return ListConversationsRequestSummaryModeExclude, nil
	case "include":
		return ListConversationsRequestSummaryModeInclude, nil
	}
	var t ListConversationsRequestSummaryMode
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ListConversationsRequestSummaryMode) Ptr() *ListConversationsRequestSummaryMode {
	return &e
}

type GetConversationsRequestFormat string

const (
	GetConversationsRequestFormatJson          GetConversationsRequestFormat = "json"
	GetConversationsRequestFormatOpentelemetry GetConversationsRequestFormat = "opentelemetry"
)

func NewGetConversationsRequestFormatFromString(s string) (GetConversationsRequestFormat, error) {
	switch s {
	case "json":
		return GetConversationsRequestFormatJson, nil
	case "opentelemetry":
		return GetConversationsRequestFormatOpentelemetry, nil
	}
	var t GetConversationsRequestFormat
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GetConversationsRequestFormat) Ptr() *GetConversationsRequestFormat {
	return &e
}

var (
	getSignedUrlConversationsRequestFieldAgentId               = big.NewInt(1 << 0)
	getSignedUrlConversationsRequestFieldIncludeConversationId = big.NewInt(1 << 1)
	getSignedUrlConversationsRequestFieldBranchId              = big.NewInt(1 << 2)
	getSignedUrlConversationsRequestFieldVersionId             = big.NewInt(1 << 3)
	getSignedUrlConversationsRequestFieldEnvironment           = big.NewInt(1 << 4)
	getSignedUrlConversationsRequestFieldDebugEventsRequest    = big.NewInt(1 << 5)
)

type GetSignedUrlConversationsRequest struct {
	AgentId               string  `json:"-" url:"agent_id,omitempty"`
	IncludeConversationId *bool   `json:"-" url:"include_conversation_id,omitempty"`
	BranchId              *string `json:"-" url:"branch_id,omitempty"`
	VersionId             *string `json:"-" url:"version_id,omitempty"`
	Environment           *string `json:"-" url:"environment,omitempty"`
	DebugEventsRequest    *bool   `json:"-" url:"debug_events_request,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSignedUrlConversationsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSignedUrlConversationsRequest) SetAgentId(agentId string) {
	g.AgentId = agentId
	g.require(getSignedUrlConversationsRequestFieldAgentId)
}

// SetIncludeConversationId sets the IncludeConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSignedUrlConversationsRequest) SetIncludeConversationId(includeConversationId *bool) {
	g.IncludeConversationId = includeConversationId
	g.require(getSignedUrlConversationsRequestFieldIncludeConversationId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSignedUrlConversationsRequest) SetBranchId(branchId *string) {
	g.BranchId = branchId
	g.require(getSignedUrlConversationsRequestFieldBranchId)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSignedUrlConversationsRequest) SetVersionId(versionId *string) {
	g.VersionId = versionId
	g.require(getSignedUrlConversationsRequestFieldVersionId)
}

// SetEnvironment sets the Environment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSignedUrlConversationsRequest) SetEnvironment(environment *string) {
	g.Environment = environment
	g.require(getSignedUrlConversationsRequestFieldEnvironment)
}

// SetDebugEventsRequest sets the DebugEventsRequest field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSignedUrlConversationsRequest) SetDebugEventsRequest(debugEventsRequest *bool) {
	g.DebugEventsRequest = debugEventsRequest
	g.require(getSignedUrlConversationsRequestFieldDebugEventsRequest)
}

var (
	getWebrtcTokenConversationsRequestFieldAgentId            = big.NewInt(1 << 0)
	getWebrtcTokenConversationsRequestFieldParticipantName    = big.NewInt(1 << 1)
	getWebrtcTokenConversationsRequestFieldBranchId           = big.NewInt(1 << 2)
	getWebrtcTokenConversationsRequestFieldVersionId          = big.NewInt(1 << 3)
	getWebrtcTokenConversationsRequestFieldEnvironment        = big.NewInt(1 << 4)
	getWebrtcTokenConversationsRequestFieldDebugEventsRequest = big.NewInt(1 << 5)
)

type GetWebrtcTokenConversationsRequest struct {
	AgentId            string  `json:"-" url:"agent_id,omitempty"`
	ParticipantName    *string `json:"-" url:"participant_name,omitempty"`
	BranchId           *string `json:"-" url:"branch_id,omitempty"`
	VersionId          *string `json:"-" url:"version_id,omitempty"`
	Environment        *string `json:"-" url:"environment,omitempty"`
	DebugEventsRequest *bool   `json:"-" url:"debug_events_request,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetWebrtcTokenConversationsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWebrtcTokenConversationsRequest) SetAgentId(agentId string) {
	g.AgentId = agentId
	g.require(getWebrtcTokenConversationsRequestFieldAgentId)
}

// SetParticipantName sets the ParticipantName field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWebrtcTokenConversationsRequest) SetParticipantName(participantName *string) {
	g.ParticipantName = participantName
	g.require(getWebrtcTokenConversationsRequestFieldParticipantName)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWebrtcTokenConversationsRequest) SetBranchId(branchId *string) {
	g.BranchId = branchId
	g.require(getWebrtcTokenConversationsRequestFieldBranchId)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWebrtcTokenConversationsRequest) SetVersionId(versionId *string) {
	g.VersionId = versionId
	g.require(getWebrtcTokenConversationsRequestFieldVersionId)
}

// SetEnvironment sets the Environment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWebrtcTokenConversationsRequest) SetEnvironment(environment *string) {
	g.Environment = environment
	g.require(getWebrtcTokenConversationsRequestFieldEnvironment)
}

// SetDebugEventsRequest sets the DebugEventsRequest field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetWebrtcTokenConversationsRequest) SetDebugEventsRequest(debugEventsRequest *bool) {
	g.DebugEventsRequest = debugEventsRequest
	g.require(getWebrtcTokenConversationsRequestFieldDebugEventsRequest)
}

var (
	listConversationsRequestFieldCursor                       = big.NewInt(1 << 0)
	listConversationsRequestFieldAgentId                      = big.NewInt(1 << 1)
	listConversationsRequestFieldVisitedAgentIds              = big.NewInt(1 << 2)
	listConversationsRequestFieldVisitedAgentBranchIds        = big.NewInt(1 << 3)
	listConversationsRequestFieldTriggeredProcedureIds        = big.NewInt(1 << 4)
	listConversationsRequestFieldCallSuccessful               = big.NewInt(1 << 5)
	listConversationsRequestFieldCallStartBeforeUnix          = big.NewInt(1 << 6)
	listConversationsRequestFieldCallStartAfterUnix           = big.NewInt(1 << 7)
	listConversationsRequestFieldCallDurationMinSecs          = big.NewInt(1 << 8)
	listConversationsRequestFieldCallDurationMaxSecs          = big.NewInt(1 << 9)
	listConversationsRequestFieldRatingMax                    = big.NewInt(1 << 10)
	listConversationsRequestFieldRatingMin                    = big.NewInt(1 << 11)
	listConversationsRequestFieldHasFeedbackComment           = big.NewInt(1 << 12)
	listConversationsRequestFieldUserId                       = big.NewInt(1 << 13)
	listConversationsRequestFieldEvaluationParams             = big.NewInt(1 << 14)
	listConversationsRequestFieldDataCollectionParams         = big.NewInt(1 << 15)
	listConversationsRequestFieldDynamicVariableParams        = big.NewInt(1 << 16)
	listConversationsRequestFieldDataCollectionFilter         = big.NewInt(1 << 17)
	listConversationsRequestFieldDynamicVariableFilter        = big.NewInt(1 << 18)
	listConversationsRequestFieldDataCollectionIds            = big.NewInt(1 << 19)
	listConversationsRequestFieldEvaluationCriteriaIds        = big.NewInt(1 << 20)
	listConversationsRequestFieldToolNames                    = big.NewInt(1 << 21)
	listConversationsRequestFieldToolNamesSuccessful          = big.NewInt(1 << 22)
	listConversationsRequestFieldToolNamesErrored             = big.NewInt(1 << 23)
	listConversationsRequestFieldToolIds                      = big.NewInt(1 << 24)
	listConversationsRequestFieldToolIdsSuccessful            = big.NewInt(1 << 25)
	listConversationsRequestFieldToolIdsErrored               = big.NewInt(1 << 26)
	listConversationsRequestFieldIncludeInvalidToolCalls      = big.NewInt(1 << 27)
	listConversationsRequestFieldMainLanguages                = big.NewInt(1 << 28)
	listConversationsRequestFieldPhoneNumberIds               = big.NewInt(1 << 29)
	listConversationsRequestFieldWhatsappPhoneNumberIds       = big.NewInt(1 << 30)
	listConversationsRequestFieldPageSize                     = big.NewInt(1 << 31)
	listConversationsRequestFieldSummaryMode                  = big.NewInt(1 << 32)
	listConversationsRequestFieldSearch                       = big.NewInt(1 << 33)
	listConversationsRequestFieldConversationInitiationSource = big.NewInt(1 << 34)
	listConversationsRequestFieldTextOnly                     = big.NewInt(1 << 35)
	listConversationsRequestFieldConversationProductType      = big.NewInt(1 << 36)
	listConversationsRequestFieldBranchId                     = big.NewInt(1 << 37)
	listConversationsRequestFieldVersionId                    = big.NewInt(1 << 38)
	listConversationsRequestFieldParentConversationId         = big.NewInt(1 << 39)
	listConversationsRequestFieldTopicIds                     = big.NewInt(1 << 40)
	listConversationsRequestFieldExcludeStatuses              = big.NewInt(1 << 41)
	listConversationsRequestFieldTagIds                       = big.NewInt(1 << 42)
	listConversationsRequestFieldWorkflowNodeEnteredId        = big.NewInt(1 << 43)
	listConversationsRequestFieldTerminationReasons           = big.NewInt(1 << 44)
	listConversationsRequestFieldGuardrailTypes               = big.NewInt(1 << 45)
	listConversationsRequestFieldCustomGuardrailNames         = big.NewInt(1 << 46)
	listConversationsRequestFieldSortDirection                = big.NewInt(1 << 47)
)

type ListConversationsRequest struct {
	Cursor                       *string                                        `json:"-" url:"cursor,omitempty"`
	AgentId                      *string                                        `json:"-" url:"agent_id,omitempty"`
	VisitedAgentIds              []*string                                      `json:"-" url:"visited_agent_ids,omitempty"`
	VisitedAgentBranchIds        []*string                                      `json:"-" url:"visited_agent_branch_ids,omitempty"`
	TriggeredProcedureIds        []*string                                      `json:"-" url:"triggered_procedure_ids,omitempty"`
	CallSuccessful               *types.EvaluationResultFilter                  `json:"-" url:"call_successful,omitempty"`
	CallStartBeforeUnix          *int                                           `json:"-" url:"call_start_before_unix,omitempty"`
	CallStartAfterUnix           *int                                           `json:"-" url:"call_start_after_unix,omitempty"`
	CallDurationMinSecs          *int                                           `json:"-" url:"call_duration_min_secs,omitempty"`
	CallDurationMaxSecs          *int                                           `json:"-" url:"call_duration_max_secs,omitempty"`
	RatingMax                    *int                                           `json:"-" url:"rating_max,omitempty"`
	RatingMin                    *int                                           `json:"-" url:"rating_min,omitempty"`
	HasFeedbackComment           *bool                                          `json:"-" url:"has_feedback_comment,omitempty"`
	UserId                       *string                                        `json:"-" url:"user_id,omitempty"`
	EvaluationParams             []*string                                      `json:"-" url:"evaluation_params,omitempty"`
	DataCollectionParams         []*string                                      `json:"-" url:"data_collection_params,omitempty"`
	DynamicVariableParams        []*string                                      `json:"-" url:"dynamic_variable_params,omitempty"`
	DataCollectionFilter         *string                                        `json:"-" url:"data_collection_filter,omitempty"`
	DynamicVariableFilter        *string                                        `json:"-" url:"dynamic_variable_filter,omitempty"`
	DataCollectionIds            []*string                                      `json:"-" url:"data_collection_ids,omitempty"`
	EvaluationCriteriaIds        []*string                                      `json:"-" url:"evaluation_criteria_ids,omitempty"`
	ToolNames                    []*string                                      `json:"-" url:"tool_names,omitempty"`
	ToolNamesSuccessful          []*string                                      `json:"-" url:"tool_names_successful,omitempty"`
	ToolNamesErrored             []*string                                      `json:"-" url:"tool_names_errored,omitempty"`
	ToolIds                      []*string                                      `json:"-" url:"tool_ids,omitempty"`
	ToolIdsSuccessful            []*string                                      `json:"-" url:"tool_ids_successful,omitempty"`
	ToolIdsErrored               []*string                                      `json:"-" url:"tool_ids_errored,omitempty"`
	IncludeInvalidToolCalls      *bool                                          `json:"-" url:"include_invalid_tool_calls,omitempty"`
	MainLanguages                []*string                                      `json:"-" url:"main_languages,omitempty"`
	PhoneNumberIds               []*string                                      `json:"-" url:"phone_number_ids,omitempty"`
	WhatsappPhoneNumberIds       []*string                                      `json:"-" url:"whatsapp_phone_number_ids,omitempty"`
	PageSize                     *int                                           `json:"-" url:"page_size,omitempty"`
	SummaryMode                  *ListConversationsRequestSummaryMode           `json:"-" url:"summary_mode,omitempty"`
	Search                       *string                                        `json:"-" url:"search,omitempty"`
	ConversationInitiationSource *types.ConversationInitiationSource            `json:"-" url:"conversation_initiation_source,omitempty"`
	TextOnly                     *bool                                          `json:"-" url:"text_only,omitempty"`
	ConversationProductType      *types.ConversationProduct                     `json:"-" url:"conversation_product_type,omitempty"`
	BranchId                     *string                                        `json:"-" url:"branch_id,omitempty"`
	VersionId                    *string                                        `json:"-" url:"version_id,omitempty"`
	ParentConversationId         *string                                        `json:"-" url:"parent_conversation_id,omitempty"`
	TopicIds                     []*string                                      `json:"-" url:"topic_ids,omitempty"`
	ExcludeStatuses              []*ListConversationsRequestExcludeStatusesItem `json:"-" url:"exclude_statuses,omitempty"`
	TagIds                       []*string                                      `json:"-" url:"tag_ids,omitempty"`
	WorkflowNodeEnteredId        *string                                        `json:"-" url:"workflow_node_entered_id,omitempty"`
	TerminationReasons           []*string                                      `json:"-" url:"termination_reasons,omitempty"`
	GuardrailTypes               []*types.GuardrailType                         `json:"-" url:"guardrail_types,omitempty"`
	CustomGuardrailNames         []*string                                      `json:"-" url:"custom_guardrail_names,omitempty"`
	SortDirection                *types.SortDirection                           `json:"-" url:"sort_direction,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListConversationsRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCursor(cursor *string) {
	l.Cursor = cursor
	l.require(listConversationsRequestFieldCursor)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetAgentId(agentId *string) {
	l.AgentId = agentId
	l.require(listConversationsRequestFieldAgentId)
}

// SetVisitedAgentIds sets the VisitedAgentIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetVisitedAgentIds(visitedAgentIds []*string) {
	l.VisitedAgentIds = visitedAgentIds
	l.require(listConversationsRequestFieldVisitedAgentIds)
}

// SetVisitedAgentBranchIds sets the VisitedAgentBranchIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetVisitedAgentBranchIds(visitedAgentBranchIds []*string) {
	l.VisitedAgentBranchIds = visitedAgentBranchIds
	l.require(listConversationsRequestFieldVisitedAgentBranchIds)
}

// SetTriggeredProcedureIds sets the TriggeredProcedureIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetTriggeredProcedureIds(triggeredProcedureIds []*string) {
	l.TriggeredProcedureIds = triggeredProcedureIds
	l.require(listConversationsRequestFieldTriggeredProcedureIds)
}

// SetCallSuccessful sets the CallSuccessful field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCallSuccessful(callSuccessful *types.EvaluationResultFilter) {
	l.CallSuccessful = callSuccessful
	l.require(listConversationsRequestFieldCallSuccessful)
}

// SetCallStartBeforeUnix sets the CallStartBeforeUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCallStartBeforeUnix(callStartBeforeUnix *int) {
	l.CallStartBeforeUnix = callStartBeforeUnix
	l.require(listConversationsRequestFieldCallStartBeforeUnix)
}

// SetCallStartAfterUnix sets the CallStartAfterUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCallStartAfterUnix(callStartAfterUnix *int) {
	l.CallStartAfterUnix = callStartAfterUnix
	l.require(listConversationsRequestFieldCallStartAfterUnix)
}

// SetCallDurationMinSecs sets the CallDurationMinSecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCallDurationMinSecs(callDurationMinSecs *int) {
	l.CallDurationMinSecs = callDurationMinSecs
	l.require(listConversationsRequestFieldCallDurationMinSecs)
}

// SetCallDurationMaxSecs sets the CallDurationMaxSecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCallDurationMaxSecs(callDurationMaxSecs *int) {
	l.CallDurationMaxSecs = callDurationMaxSecs
	l.require(listConversationsRequestFieldCallDurationMaxSecs)
}

// SetRatingMax sets the RatingMax field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetRatingMax(ratingMax *int) {
	l.RatingMax = ratingMax
	l.require(listConversationsRequestFieldRatingMax)
}

// SetRatingMin sets the RatingMin field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetRatingMin(ratingMin *int) {
	l.RatingMin = ratingMin
	l.require(listConversationsRequestFieldRatingMin)
}

// SetHasFeedbackComment sets the HasFeedbackComment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetHasFeedbackComment(hasFeedbackComment *bool) {
	l.HasFeedbackComment = hasFeedbackComment
	l.require(listConversationsRequestFieldHasFeedbackComment)
}

// SetUserId sets the UserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetUserId(userId *string) {
	l.UserId = userId
	l.require(listConversationsRequestFieldUserId)
}

// SetEvaluationParams sets the EvaluationParams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetEvaluationParams(evaluationParams []*string) {
	l.EvaluationParams = evaluationParams
	l.require(listConversationsRequestFieldEvaluationParams)
}

// SetDataCollectionParams sets the DataCollectionParams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetDataCollectionParams(dataCollectionParams []*string) {
	l.DataCollectionParams = dataCollectionParams
	l.require(listConversationsRequestFieldDataCollectionParams)
}

// SetDynamicVariableParams sets the DynamicVariableParams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetDynamicVariableParams(dynamicVariableParams []*string) {
	l.DynamicVariableParams = dynamicVariableParams
	l.require(listConversationsRequestFieldDynamicVariableParams)
}

// SetDataCollectionFilter sets the DataCollectionFilter field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetDataCollectionFilter(dataCollectionFilter *string) {
	l.DataCollectionFilter = dataCollectionFilter
	l.require(listConversationsRequestFieldDataCollectionFilter)
}

// SetDynamicVariableFilter sets the DynamicVariableFilter field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetDynamicVariableFilter(dynamicVariableFilter *string) {
	l.DynamicVariableFilter = dynamicVariableFilter
	l.require(listConversationsRequestFieldDynamicVariableFilter)
}

// SetDataCollectionIds sets the DataCollectionIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetDataCollectionIds(dataCollectionIds []*string) {
	l.DataCollectionIds = dataCollectionIds
	l.require(listConversationsRequestFieldDataCollectionIds)
}

// SetEvaluationCriteriaIds sets the EvaluationCriteriaIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetEvaluationCriteriaIds(evaluationCriteriaIds []*string) {
	l.EvaluationCriteriaIds = evaluationCriteriaIds
	l.require(listConversationsRequestFieldEvaluationCriteriaIds)
}

// SetToolNames sets the ToolNames field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetToolNames(toolNames []*string) {
	l.ToolNames = toolNames
	l.require(listConversationsRequestFieldToolNames)
}

// SetToolNamesSuccessful sets the ToolNamesSuccessful field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetToolNamesSuccessful(toolNamesSuccessful []*string) {
	l.ToolNamesSuccessful = toolNamesSuccessful
	l.require(listConversationsRequestFieldToolNamesSuccessful)
}

// SetToolNamesErrored sets the ToolNamesErrored field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetToolNamesErrored(toolNamesErrored []*string) {
	l.ToolNamesErrored = toolNamesErrored
	l.require(listConversationsRequestFieldToolNamesErrored)
}

// SetToolIds sets the ToolIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetToolIds(toolIds []*string) {
	l.ToolIds = toolIds
	l.require(listConversationsRequestFieldToolIds)
}

// SetToolIdsSuccessful sets the ToolIdsSuccessful field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetToolIdsSuccessful(toolIdsSuccessful []*string) {
	l.ToolIdsSuccessful = toolIdsSuccessful
	l.require(listConversationsRequestFieldToolIdsSuccessful)
}

// SetToolIdsErrored sets the ToolIdsErrored field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetToolIdsErrored(toolIdsErrored []*string) {
	l.ToolIdsErrored = toolIdsErrored
	l.require(listConversationsRequestFieldToolIdsErrored)
}

// SetIncludeInvalidToolCalls sets the IncludeInvalidToolCalls field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetIncludeInvalidToolCalls(includeInvalidToolCalls *bool) {
	l.IncludeInvalidToolCalls = includeInvalidToolCalls
	l.require(listConversationsRequestFieldIncludeInvalidToolCalls)
}

// SetMainLanguages sets the MainLanguages field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetMainLanguages(mainLanguages []*string) {
	l.MainLanguages = mainLanguages
	l.require(listConversationsRequestFieldMainLanguages)
}

// SetPhoneNumberIds sets the PhoneNumberIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetPhoneNumberIds(phoneNumberIds []*string) {
	l.PhoneNumberIds = phoneNumberIds
	l.require(listConversationsRequestFieldPhoneNumberIds)
}

// SetWhatsappPhoneNumberIds sets the WhatsappPhoneNumberIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetWhatsappPhoneNumberIds(whatsappPhoneNumberIds []*string) {
	l.WhatsappPhoneNumberIds = whatsappPhoneNumberIds
	l.require(listConversationsRequestFieldWhatsappPhoneNumberIds)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetPageSize(pageSize *int) {
	l.PageSize = pageSize
	l.require(listConversationsRequestFieldPageSize)
}

// SetSummaryMode sets the SummaryMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetSummaryMode(summaryMode *ListConversationsRequestSummaryMode) {
	l.SummaryMode = summaryMode
	l.require(listConversationsRequestFieldSummaryMode)
}

// SetSearch sets the Search field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetSearch(search *string) {
	l.Search = search
	l.require(listConversationsRequestFieldSearch)
}

// SetConversationInitiationSource sets the ConversationInitiationSource field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetConversationInitiationSource(conversationInitiationSource *types.ConversationInitiationSource) {
	l.ConversationInitiationSource = conversationInitiationSource
	l.require(listConversationsRequestFieldConversationInitiationSource)
}

// SetTextOnly sets the TextOnly field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetTextOnly(textOnly *bool) {
	l.TextOnly = textOnly
	l.require(listConversationsRequestFieldTextOnly)
}

// SetConversationProductType sets the ConversationProductType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetConversationProductType(conversationProductType *types.ConversationProduct) {
	l.ConversationProductType = conversationProductType
	l.require(listConversationsRequestFieldConversationProductType)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetBranchId(branchId *string) {
	l.BranchId = branchId
	l.require(listConversationsRequestFieldBranchId)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetVersionId(versionId *string) {
	l.VersionId = versionId
	l.require(listConversationsRequestFieldVersionId)
}

// SetParentConversationId sets the ParentConversationId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetParentConversationId(parentConversationId *string) {
	l.ParentConversationId = parentConversationId
	l.require(listConversationsRequestFieldParentConversationId)
}

// SetTopicIds sets the TopicIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetTopicIds(topicIds []*string) {
	l.TopicIds = topicIds
	l.require(listConversationsRequestFieldTopicIds)
}

// SetExcludeStatuses sets the ExcludeStatuses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetExcludeStatuses(excludeStatuses []*ListConversationsRequestExcludeStatusesItem) {
	l.ExcludeStatuses = excludeStatuses
	l.require(listConversationsRequestFieldExcludeStatuses)
}

// SetTagIds sets the TagIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetTagIds(tagIds []*string) {
	l.TagIds = tagIds
	l.require(listConversationsRequestFieldTagIds)
}

// SetWorkflowNodeEnteredId sets the WorkflowNodeEnteredId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetWorkflowNodeEnteredId(workflowNodeEnteredId *string) {
	l.WorkflowNodeEnteredId = workflowNodeEnteredId
	l.require(listConversationsRequestFieldWorkflowNodeEnteredId)
}

// SetTerminationReasons sets the TerminationReasons field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetTerminationReasons(terminationReasons []*string) {
	l.TerminationReasons = terminationReasons
	l.require(listConversationsRequestFieldTerminationReasons)
}

// SetGuardrailTypes sets the GuardrailTypes field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetGuardrailTypes(guardrailTypes []*types.GuardrailType) {
	l.GuardrailTypes = guardrailTypes
	l.require(listConversationsRequestFieldGuardrailTypes)
}

// SetCustomGuardrailNames sets the CustomGuardrailNames field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetCustomGuardrailNames(customGuardrailNames []*string) {
	l.CustomGuardrailNames = customGuardrailNames
	l.require(listConversationsRequestFieldCustomGuardrailNames)
}

// SetSortDirection sets the SortDirection field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListConversationsRequest) SetSortDirection(sortDirection *types.SortDirection) {
	l.SortDirection = sortDirection
	l.require(listConversationsRequestFieldSortDirection)
}

var (
	resolveConversationsRequestFieldAgentId   = big.NewInt(1 << 0)
	resolveConversationsRequestFieldReference = big.NewInt(1 << 1)
)

type ResolveConversationsRequest struct {
	AgentId   string `json:"-" url:"agent_id,omitempty"`
	Reference string `json:"-" url:"reference,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (r *ResolveConversationsRequest) require(field *big.Int) {
	if r.explicitFields == nil {
		r.explicitFields = big.NewInt(0)
	}
	r.explicitFields.Or(r.explicitFields, field)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResolveConversationsRequest) SetAgentId(agentId string) {
	r.AgentId = agentId
	r.require(resolveConversationsRequestFieldAgentId)
}

// SetReference sets the Reference field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (r *ResolveConversationsRequest) SetReference(reference string) {
	r.Reference = reference
	r.require(resolveConversationsRequestFieldReference)
}

var (
	getConversationsRequestFieldFormat = big.NewInt(1 << 0)
)

type GetConversationsRequest struct {
	Format *GetConversationsRequestFormat `json:"-" url:"format,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetConversationsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetFormat sets the Format field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetConversationsRequest) SetFormat(format *GetConversationsRequestFormat) {
	g.Format = format
	g.require(getConversationsRequestFieldFormat)
}

var (
	getSummaryConversationsRequestFieldMaxMessages = big.NewInt(1 << 0)
)

type GetSummaryConversationsRequest struct {
	MaxMessages *int `json:"-" url:"max_messages,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSummaryConversationsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetMaxMessages sets the MaxMessages field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSummaryConversationsRequest) SetMaxMessages(maxMessages *int) {
	g.MaxMessages = maxMessages
	g.require(getSummaryConversationsRequestFieldMaxMessages)
}

var (
	getSipMessagesConversationsRequestFieldPageSize = big.NewInt(1 << 0)
	getSipMessagesConversationsRequestFieldCursor   = big.NewInt(1 << 1)
)

type GetSipMessagesConversationsRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (g *GetSipMessagesConversationsRequest) require(field *big.Int) {
	if g.explicitFields == nil {
		g.explicitFields = big.NewInt(0)
	}
	g.explicitFields.Or(g.explicitFields, field)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSipMessagesConversationsRequest) SetPageSize(pageSize *int) {
	g.PageSize = pageSize
	g.require(getSipMessagesConversationsRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (g *GetSipMessagesConversationsRequest) SetCursor(cursor *string) {
	g.Cursor = cursor
	g.require(getSipMessagesConversationsRequestFieldCursor)
}
