package conversations

import (
	"fmt"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type TextSearchMessagesRequestExcludeStatusesItem string

const (
	TextSearchMessagesRequestExcludeStatusesItemInitiated  TextSearchMessagesRequestExcludeStatusesItem = "initiated"
	TextSearchMessagesRequestExcludeStatusesItemInProgress TextSearchMessagesRequestExcludeStatusesItem = "in-progress"
	TextSearchMessagesRequestExcludeStatusesItemProcessing TextSearchMessagesRequestExcludeStatusesItem = "processing"
	TextSearchMessagesRequestExcludeStatusesItemDone       TextSearchMessagesRequestExcludeStatusesItem = "done"
	TextSearchMessagesRequestExcludeStatusesItemFailed     TextSearchMessagesRequestExcludeStatusesItem = "failed"
)

func NewTextSearchMessagesRequestExcludeStatusesItemFromString(s string) (TextSearchMessagesRequestExcludeStatusesItem, error) {
	switch s {
	case "initiated":
		return TextSearchMessagesRequestExcludeStatusesItemInitiated, nil
	case "in-progress":
		return TextSearchMessagesRequestExcludeStatusesItemInProgress, nil
	case "processing":
		return TextSearchMessagesRequestExcludeStatusesItemProcessing, nil
	case "done":
		return TextSearchMessagesRequestExcludeStatusesItemDone, nil
	case "failed":
		return TextSearchMessagesRequestExcludeStatusesItemFailed, nil
	}
	var t TextSearchMessagesRequestExcludeStatusesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextSearchMessagesRequestExcludeStatusesItem) Ptr() *TextSearchMessagesRequestExcludeStatusesItem {
	return &e
}

type TextSearchMessagesRequestSummaryMode string

const (
	TextSearchMessagesRequestSummaryModeExclude TextSearchMessagesRequestSummaryMode = "exclude"
	TextSearchMessagesRequestSummaryModeInclude TextSearchMessagesRequestSummaryMode = "include"
)

func NewTextSearchMessagesRequestSummaryModeFromString(s string) (TextSearchMessagesRequestSummaryMode, error) {
	switch s {
	case "exclude":
		return TextSearchMessagesRequestSummaryModeExclude, nil
	case "include":
		return TextSearchMessagesRequestSummaryModeInclude, nil
	}
	var t TextSearchMessagesRequestSummaryMode
	return t, fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextSearchMessagesRequestSummaryMode) Ptr() *TextSearchMessagesRequestSummaryMode {
	return &e
}

var (
	textSearchMessagesRequestFieldTextQuery                    = big.NewInt(1 << 0)
	textSearchMessagesRequestFieldAgentId                      = big.NewInt(1 << 1)
	textSearchMessagesRequestFieldVisitedAgentIds              = big.NewInt(1 << 2)
	textSearchMessagesRequestFieldVisitedAgentBranchIds        = big.NewInt(1 << 3)
	textSearchMessagesRequestFieldTriggeredProcedureIds        = big.NewInt(1 << 4)
	textSearchMessagesRequestFieldCallSuccessful               = big.NewInt(1 << 5)
	textSearchMessagesRequestFieldCallStartBeforeUnix          = big.NewInt(1 << 6)
	textSearchMessagesRequestFieldCallStartAfterUnix           = big.NewInt(1 << 7)
	textSearchMessagesRequestFieldCallDurationMinSecs          = big.NewInt(1 << 8)
	textSearchMessagesRequestFieldCallDurationMaxSecs          = big.NewInt(1 << 9)
	textSearchMessagesRequestFieldRatingMax                    = big.NewInt(1 << 10)
	textSearchMessagesRequestFieldRatingMin                    = big.NewInt(1 << 11)
	textSearchMessagesRequestFieldHasFeedbackComment           = big.NewInt(1 << 12)
	textSearchMessagesRequestFieldUserId                       = big.NewInt(1 << 13)
	textSearchMessagesRequestFieldEvaluationParams             = big.NewInt(1 << 14)
	textSearchMessagesRequestFieldDataCollectionParams         = big.NewInt(1 << 15)
	textSearchMessagesRequestFieldDynamicVariableParams        = big.NewInt(1 << 16)
	textSearchMessagesRequestFieldToolNames                    = big.NewInt(1 << 17)
	textSearchMessagesRequestFieldToolNamesSuccessful          = big.NewInt(1 << 18)
	textSearchMessagesRequestFieldToolNamesErrored             = big.NewInt(1 << 19)
	textSearchMessagesRequestFieldToolIds                      = big.NewInt(1 << 20)
	textSearchMessagesRequestFieldToolIdsSuccessful            = big.NewInt(1 << 21)
	textSearchMessagesRequestFieldToolIdsErrored               = big.NewInt(1 << 22)
	textSearchMessagesRequestFieldIncludeInvalidToolCalls      = big.NewInt(1 << 23)
	textSearchMessagesRequestFieldMainLanguages                = big.NewInt(1 << 24)
	textSearchMessagesRequestFieldPhoneNumberIds               = big.NewInt(1 << 25)
	textSearchMessagesRequestFieldWhatsappPhoneNumberIds       = big.NewInt(1 << 26)
	textSearchMessagesRequestFieldExcludeStatuses              = big.NewInt(1 << 27)
	textSearchMessagesRequestFieldTerminationReasons           = big.NewInt(1 << 28)
	textSearchMessagesRequestFieldPageSize                     = big.NewInt(1 << 29)
	textSearchMessagesRequestFieldSummaryMode                  = big.NewInt(1 << 30)
	textSearchMessagesRequestFieldConversationInitiationSource = big.NewInt(1 << 31)
	textSearchMessagesRequestFieldTextOnly                     = big.NewInt(1 << 32)
	textSearchMessagesRequestFieldConversationProductType      = big.NewInt(1 << 33)
	textSearchMessagesRequestFieldBranchId                     = big.NewInt(1 << 34)
	textSearchMessagesRequestFieldVersionId                    = big.NewInt(1 << 35)
	textSearchMessagesRequestFieldTopicIds                     = big.NewInt(1 << 36)
	textSearchMessagesRequestFieldSortBy                       = big.NewInt(1 << 37)
	textSearchMessagesRequestFieldCursor                       = big.NewInt(1 << 38)
)

type TextSearchMessagesRequest struct {
	TextQuery                    string                                          `json:"-" url:"text_query,omitempty"`
	AgentId                      *string                                         `json:"-" url:"agent_id,omitempty"`
	VisitedAgentIds              []*string                                       `json:"-" url:"visited_agent_ids,omitempty"`
	VisitedAgentBranchIds        []*string                                       `json:"-" url:"visited_agent_branch_ids,omitempty"`
	TriggeredProcedureIds        []*string                                       `json:"-" url:"triggered_procedure_ids,omitempty"`
	CallSuccessful               *types.EvaluationResultFilter                   `json:"-" url:"call_successful,omitempty"`
	CallStartBeforeUnix          *int                                            `json:"-" url:"call_start_before_unix,omitempty"`
	CallStartAfterUnix           *int                                            `json:"-" url:"call_start_after_unix,omitempty"`
	CallDurationMinSecs          *int                                            `json:"-" url:"call_duration_min_secs,omitempty"`
	CallDurationMaxSecs          *int                                            `json:"-" url:"call_duration_max_secs,omitempty"`
	RatingMax                    *int                                            `json:"-" url:"rating_max,omitempty"`
	RatingMin                    *int                                            `json:"-" url:"rating_min,omitempty"`
	HasFeedbackComment           *bool                                           `json:"-" url:"has_feedback_comment,omitempty"`
	UserId                       *string                                         `json:"-" url:"user_id,omitempty"`
	EvaluationParams             []*string                                       `json:"-" url:"evaluation_params,omitempty"`
	DataCollectionParams         []*string                                       `json:"-" url:"data_collection_params,omitempty"`
	DynamicVariableParams        []*string                                       `json:"-" url:"dynamic_variable_params,omitempty"`
	ToolNames                    []*string                                       `json:"-" url:"tool_names,omitempty"`
	ToolNamesSuccessful          []*string                                       `json:"-" url:"tool_names_successful,omitempty"`
	ToolNamesErrored             []*string                                       `json:"-" url:"tool_names_errored,omitempty"`
	ToolIds                      []*string                                       `json:"-" url:"tool_ids,omitempty"`
	ToolIdsSuccessful            []*string                                       `json:"-" url:"tool_ids_successful,omitempty"`
	ToolIdsErrored               []*string                                       `json:"-" url:"tool_ids_errored,omitempty"`
	IncludeInvalidToolCalls      *bool                                           `json:"-" url:"include_invalid_tool_calls,omitempty"`
	MainLanguages                []*string                                       `json:"-" url:"main_languages,omitempty"`
	PhoneNumberIds               []*string                                       `json:"-" url:"phone_number_ids,omitempty"`
	WhatsappPhoneNumberIds       []*string                                       `json:"-" url:"whatsapp_phone_number_ids,omitempty"`
	ExcludeStatuses              []*TextSearchMessagesRequestExcludeStatusesItem `json:"-" url:"exclude_statuses,omitempty"`
	TerminationReasons           []*string                                       `json:"-" url:"termination_reasons,omitempty"`
	PageSize                     *int                                            `json:"-" url:"page_size,omitempty"`
	SummaryMode                  *TextSearchMessagesRequestSummaryMode           `json:"-" url:"summary_mode,omitempty"`
	ConversationInitiationSource *types.ConversationInitiationSource             `json:"-" url:"conversation_initiation_source,omitempty"`
	TextOnly                     *bool                                           `json:"-" url:"text_only,omitempty"`
	ConversationProductType      *types.ConversationProduct                      `json:"-" url:"conversation_product_type,omitempty"`
	BranchId                     *string                                         `json:"-" url:"branch_id,omitempty"`
	VersionId                    *string                                         `json:"-" url:"version_id,omitempty"`
	TopicIds                     []*string                                       `json:"-" url:"topic_ids,omitempty"`
	SortBy                       *types.MessageSearchSortBy                      `json:"-" url:"sort_by,omitempty"`
	Cursor                       *string                                         `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (t *TextSearchMessagesRequest) require(field *big.Int) {
	if t.explicitFields == nil {
		t.explicitFields = big.NewInt(0)
	}
	t.explicitFields.Or(t.explicitFields, field)
}

// SetTextQuery sets the TextQuery field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetTextQuery(textQuery string) {
	t.TextQuery = textQuery
	t.require(textSearchMessagesRequestFieldTextQuery)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetAgentId(agentId *string) {
	t.AgentId = agentId
	t.require(textSearchMessagesRequestFieldAgentId)
}

// SetVisitedAgentIds sets the VisitedAgentIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetVisitedAgentIds(visitedAgentIds []*string) {
	t.VisitedAgentIds = visitedAgentIds
	t.require(textSearchMessagesRequestFieldVisitedAgentIds)
}

// SetVisitedAgentBranchIds sets the VisitedAgentBranchIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetVisitedAgentBranchIds(visitedAgentBranchIds []*string) {
	t.VisitedAgentBranchIds = visitedAgentBranchIds
	t.require(textSearchMessagesRequestFieldVisitedAgentBranchIds)
}

// SetTriggeredProcedureIds sets the TriggeredProcedureIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetTriggeredProcedureIds(triggeredProcedureIds []*string) {
	t.TriggeredProcedureIds = triggeredProcedureIds
	t.require(textSearchMessagesRequestFieldTriggeredProcedureIds)
}

// SetCallSuccessful sets the CallSuccessful field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetCallSuccessful(callSuccessful *types.EvaluationResultFilter) {
	t.CallSuccessful = callSuccessful
	t.require(textSearchMessagesRequestFieldCallSuccessful)
}

// SetCallStartBeforeUnix sets the CallStartBeforeUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetCallStartBeforeUnix(callStartBeforeUnix *int) {
	t.CallStartBeforeUnix = callStartBeforeUnix
	t.require(textSearchMessagesRequestFieldCallStartBeforeUnix)
}

// SetCallStartAfterUnix sets the CallStartAfterUnix field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetCallStartAfterUnix(callStartAfterUnix *int) {
	t.CallStartAfterUnix = callStartAfterUnix
	t.require(textSearchMessagesRequestFieldCallStartAfterUnix)
}

// SetCallDurationMinSecs sets the CallDurationMinSecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetCallDurationMinSecs(callDurationMinSecs *int) {
	t.CallDurationMinSecs = callDurationMinSecs
	t.require(textSearchMessagesRequestFieldCallDurationMinSecs)
}

// SetCallDurationMaxSecs sets the CallDurationMaxSecs field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetCallDurationMaxSecs(callDurationMaxSecs *int) {
	t.CallDurationMaxSecs = callDurationMaxSecs
	t.require(textSearchMessagesRequestFieldCallDurationMaxSecs)
}

// SetRatingMax sets the RatingMax field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetRatingMax(ratingMax *int) {
	t.RatingMax = ratingMax
	t.require(textSearchMessagesRequestFieldRatingMax)
}

// SetRatingMin sets the RatingMin field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetRatingMin(ratingMin *int) {
	t.RatingMin = ratingMin
	t.require(textSearchMessagesRequestFieldRatingMin)
}

// SetHasFeedbackComment sets the HasFeedbackComment field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetHasFeedbackComment(hasFeedbackComment *bool) {
	t.HasFeedbackComment = hasFeedbackComment
	t.require(textSearchMessagesRequestFieldHasFeedbackComment)
}

// SetUserId sets the UserId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetUserId(userId *string) {
	t.UserId = userId
	t.require(textSearchMessagesRequestFieldUserId)
}

// SetEvaluationParams sets the EvaluationParams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetEvaluationParams(evaluationParams []*string) {
	t.EvaluationParams = evaluationParams
	t.require(textSearchMessagesRequestFieldEvaluationParams)
}

// SetDataCollectionParams sets the DataCollectionParams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetDataCollectionParams(dataCollectionParams []*string) {
	t.DataCollectionParams = dataCollectionParams
	t.require(textSearchMessagesRequestFieldDataCollectionParams)
}

// SetDynamicVariableParams sets the DynamicVariableParams field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetDynamicVariableParams(dynamicVariableParams []*string) {
	t.DynamicVariableParams = dynamicVariableParams
	t.require(textSearchMessagesRequestFieldDynamicVariableParams)
}

// SetToolNames sets the ToolNames field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetToolNames(toolNames []*string) {
	t.ToolNames = toolNames
	t.require(textSearchMessagesRequestFieldToolNames)
}

// SetToolNamesSuccessful sets the ToolNamesSuccessful field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetToolNamesSuccessful(toolNamesSuccessful []*string) {
	t.ToolNamesSuccessful = toolNamesSuccessful
	t.require(textSearchMessagesRequestFieldToolNamesSuccessful)
}

// SetToolNamesErrored sets the ToolNamesErrored field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetToolNamesErrored(toolNamesErrored []*string) {
	t.ToolNamesErrored = toolNamesErrored
	t.require(textSearchMessagesRequestFieldToolNamesErrored)
}

// SetToolIds sets the ToolIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetToolIds(toolIds []*string) {
	t.ToolIds = toolIds
	t.require(textSearchMessagesRequestFieldToolIds)
}

// SetToolIdsSuccessful sets the ToolIdsSuccessful field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetToolIdsSuccessful(toolIdsSuccessful []*string) {
	t.ToolIdsSuccessful = toolIdsSuccessful
	t.require(textSearchMessagesRequestFieldToolIdsSuccessful)
}

// SetToolIdsErrored sets the ToolIdsErrored field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetToolIdsErrored(toolIdsErrored []*string) {
	t.ToolIdsErrored = toolIdsErrored
	t.require(textSearchMessagesRequestFieldToolIdsErrored)
}

// SetIncludeInvalidToolCalls sets the IncludeInvalidToolCalls field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetIncludeInvalidToolCalls(includeInvalidToolCalls *bool) {
	t.IncludeInvalidToolCalls = includeInvalidToolCalls
	t.require(textSearchMessagesRequestFieldIncludeInvalidToolCalls)
}

// SetMainLanguages sets the MainLanguages field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetMainLanguages(mainLanguages []*string) {
	t.MainLanguages = mainLanguages
	t.require(textSearchMessagesRequestFieldMainLanguages)
}

// SetPhoneNumberIds sets the PhoneNumberIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetPhoneNumberIds(phoneNumberIds []*string) {
	t.PhoneNumberIds = phoneNumberIds
	t.require(textSearchMessagesRequestFieldPhoneNumberIds)
}

// SetWhatsappPhoneNumberIds sets the WhatsappPhoneNumberIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetWhatsappPhoneNumberIds(whatsappPhoneNumberIds []*string) {
	t.WhatsappPhoneNumberIds = whatsappPhoneNumberIds
	t.require(textSearchMessagesRequestFieldWhatsappPhoneNumberIds)
}

// SetExcludeStatuses sets the ExcludeStatuses field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetExcludeStatuses(excludeStatuses []*TextSearchMessagesRequestExcludeStatusesItem) {
	t.ExcludeStatuses = excludeStatuses
	t.require(textSearchMessagesRequestFieldExcludeStatuses)
}

// SetTerminationReasons sets the TerminationReasons field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetTerminationReasons(terminationReasons []*string) {
	t.TerminationReasons = terminationReasons
	t.require(textSearchMessagesRequestFieldTerminationReasons)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetPageSize(pageSize *int) {
	t.PageSize = pageSize
	t.require(textSearchMessagesRequestFieldPageSize)
}

// SetSummaryMode sets the SummaryMode field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetSummaryMode(summaryMode *TextSearchMessagesRequestSummaryMode) {
	t.SummaryMode = summaryMode
	t.require(textSearchMessagesRequestFieldSummaryMode)
}

// SetConversationInitiationSource sets the ConversationInitiationSource field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetConversationInitiationSource(conversationInitiationSource *types.ConversationInitiationSource) {
	t.ConversationInitiationSource = conversationInitiationSource
	t.require(textSearchMessagesRequestFieldConversationInitiationSource)
}

// SetTextOnly sets the TextOnly field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetTextOnly(textOnly *bool) {
	t.TextOnly = textOnly
	t.require(textSearchMessagesRequestFieldTextOnly)
}

// SetConversationProductType sets the ConversationProductType field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetConversationProductType(conversationProductType *types.ConversationProduct) {
	t.ConversationProductType = conversationProductType
	t.require(textSearchMessagesRequestFieldConversationProductType)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetBranchId(branchId *string) {
	t.BranchId = branchId
	t.require(textSearchMessagesRequestFieldBranchId)
}

// SetVersionId sets the VersionId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetVersionId(versionId *string) {
	t.VersionId = versionId
	t.require(textSearchMessagesRequestFieldVersionId)
}

// SetTopicIds sets the TopicIds field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetTopicIds(topicIds []*string) {
	t.TopicIds = topicIds
	t.require(textSearchMessagesRequestFieldTopicIds)
}

// SetSortBy sets the SortBy field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetSortBy(sortBy *types.MessageSearchSortBy) {
	t.SortBy = sortBy
	t.require(textSearchMessagesRequestFieldSortBy)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (t *TextSearchMessagesRequest) SetCursor(cursor *string) {
	t.Cursor = cursor
	t.require(textSearchMessagesRequestFieldCursor)
}

var (
	searchMessagesRequestFieldTextQuery = big.NewInt(1 << 0)
	searchMessagesRequestFieldAgentId   = big.NewInt(1 << 1)
	searchMessagesRequestFieldBranchId  = big.NewInt(1 << 2)
	searchMessagesRequestFieldPageSize  = big.NewInt(1 << 3)
	searchMessagesRequestFieldCursor    = big.NewInt(1 << 4)
)

type SearchMessagesRequest struct {
	TextQuery string  `json:"-" url:"text_query,omitempty"`
	AgentId   *string `json:"-" url:"agent_id,omitempty"`
	BranchId  *string `json:"-" url:"branch_id,omitempty"`
	PageSize  *int    `json:"-" url:"page_size,omitempty"`
	Cursor    *string `json:"-" url:"cursor,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (s *SearchMessagesRequest) require(field *big.Int) {
	if s.explicitFields == nil {
		s.explicitFields = big.NewInt(0)
	}
	s.explicitFields.Or(s.explicitFields, field)
}

// SetTextQuery sets the TextQuery field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchMessagesRequest) SetTextQuery(textQuery string) {
	s.TextQuery = textQuery
	s.require(searchMessagesRequestFieldTextQuery)
}

// SetAgentId sets the AgentId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchMessagesRequest) SetAgentId(agentId *string) {
	s.AgentId = agentId
	s.require(searchMessagesRequestFieldAgentId)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchMessagesRequest) SetBranchId(branchId *string) {
	s.BranchId = branchId
	s.require(searchMessagesRequestFieldBranchId)
}

// SetPageSize sets the PageSize field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchMessagesRequest) SetPageSize(pageSize *int) {
	s.PageSize = pageSize
	s.require(searchMessagesRequestFieldPageSize)
}

// SetCursor sets the Cursor field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (s *SearchMessagesRequest) SetCursor(cursor *string) {
	s.Cursor = cursor
	s.require(searchMessagesRequestFieldCursor)
}
