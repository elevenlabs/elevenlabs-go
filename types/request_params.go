package types

// ListHistoryRequest holds the optional parameters for the API request.
type ListHistoryRequest struct {
	PageSize                *int                             `json:"-" url:"page_size,omitempty"`
	StartAfterHistoryItemId *string                          `json:"-" url:"start_after_history_item_id,omitempty"`
	VoiceId                 *string                          `json:"-" url:"voice_id,omitempty"`
	ModelId                 *string                          `json:"-" url:"model_id,omitempty"`
	DateBeforeUnix          *int                             `json:"-" url:"date_before_unix,omitempty"`
	DateAfterUnix           *int                             `json:"-" url:"date_after_unix,omitempty"`
	SortDirection           *ListHistoryRequestSortDirection `json:"-" url:"sort_direction,omitempty"`
	Search                  *string                          `json:"-" url:"search,omitempty"`
	Source                  *ListHistoryRequestSource        `json:"-" url:"source,omitempty"`
}

func (p *ListHistoryRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListHistoryRequest) SetStartAfterHistoryItemId(startAfterHistoryItemId *string) {
	p.StartAfterHistoryItemId = startAfterHistoryItemId
}

func (p *ListHistoryRequest) SetVoiceId(voiceId *string) {
	p.VoiceId = voiceId
}

func (p *ListHistoryRequest) SetModelId(modelId *string) {
	p.ModelId = modelId
}

func (p *ListHistoryRequest) SetDateBeforeUnix(dateBeforeUnix *int) {
	p.DateBeforeUnix = dateBeforeUnix
}

func (p *ListHistoryRequest) SetDateAfterUnix(dateAfterUnix *int) {
	p.DateAfterUnix = dateAfterUnix
}

func (p *ListHistoryRequest) SetSortDirection(sortDirection *ListHistoryRequestSortDirection) {
	p.SortDirection = sortDirection
}

func (p *ListHistoryRequest) SetSearch(search *string) {
	p.Search = search
}

func (p *ListHistoryRequest) SetSource(source *ListHistoryRequestSource) {
	p.Source = source
}

// ListAudioIsolationRequest holds the optional parameters for the API request.
type ListAudioIsolationRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Page     *int    `json:"-" url:"page,omitempty"`
	Search   *string `json:"-" url:"search,omitempty"`
}

func (p *ListAudioIsolationRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListAudioIsolationRequest) SetPage(page *int) {
	p.Page = page
}

func (p *ListAudioIsolationRequest) SetSearch(search *string) {
	p.Search = search
}

// StreamTextToSpeechRequest holds the optional parameters for the API request.
type StreamTextToSpeechRequest struct {
	Text                            string                                               `json:"text" url:"-"`
	ModelId                         *string                                              `json:"model_id,omitempty" url:"-"`
	LanguageCode                    *string                                              `json:"language_code,omitempty" url:"-"`
	VoiceSettings                   *VoiceSettings                                       `json:"voice_settings,omitempty" url:"-"`
	PronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel `json:"pronunciation_dictionary_locators,omitempty" url:"-"`
	Seed                            *int                                                 `json:"seed,omitempty" url:"-"`
	PreviousText                    *string                                              `json:"previous_text,omitempty" url:"-"`
	NextText                        *string                                              `json:"next_text,omitempty" url:"-"`
	PreviousRequestIds              []string                                             `json:"previous_request_ids,omitempty" url:"-"`
	NextRequestIds                  []string                                             `json:"next_request_ids,omitempty" url:"-"`
	UsePvcAsIvc                     *bool                                                `json:"use_pvc_as_ivc,omitempty" url:"-"`
	ApplyTextNormalization          *BodyTextToSpeechStreamApplyTextNormalization        `json:"apply_text_normalization,omitempty" url:"-"`
	ApplyLanguageTextNormalization  *bool                                                `json:"apply_language_text_normalization,omitempty" url:"-"`
	AudioEffects                    *EffectsSpecInput                                    `json:"audio_effects,omitempty" url:"-"`
	EnableLogging                   *bool                                                `json:"-" url:"enable_logging,omitempty"`
	OptimizeStreamingLatency        *int                                                 `json:"-" url:"optimize_streaming_latency,omitempty"`
	OutputFormat                    *StreamTextToSpeechRequestOutputFormat               `json:"-" url:"output_format,omitempty"`
}

func (p *StreamTextToSpeechRequest) SetText(text string) {
	p.Text = text
}

func (p *StreamTextToSpeechRequest) SetModelId(modelId *string) {
	p.ModelId = modelId
}

func (p *StreamTextToSpeechRequest) SetLanguageCode(languageCode *string) {
	p.LanguageCode = languageCode
}

func (p *StreamTextToSpeechRequest) SetVoiceSettings(voiceSettings *VoiceSettings) {
	p.VoiceSettings = voiceSettings
}

func (p *StreamTextToSpeechRequest) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel) {
	p.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (p *StreamTextToSpeechRequest) SetSeed(seed *int) {
	p.Seed = seed
}

func (p *StreamTextToSpeechRequest) SetPreviousText(previousText *string) {
	p.PreviousText = previousText
}

func (p *StreamTextToSpeechRequest) SetNextText(nextText *string) {
	p.NextText = nextText
}

func (p *StreamTextToSpeechRequest) SetPreviousRequestIds(previousRequestIds []string) {
	p.PreviousRequestIds = previousRequestIds
}

func (p *StreamTextToSpeechRequest) SetNextRequestIds(nextRequestIds []string) {
	p.NextRequestIds = nextRequestIds
}

func (p *StreamTextToSpeechRequest) SetUsePvcAsIvc(usePvcAsIvc *bool) {
	p.UsePvcAsIvc = usePvcAsIvc
}

func (p *StreamTextToSpeechRequest) SetApplyTextNormalization(applyTextNormalization *BodyTextToSpeechStreamApplyTextNormalization) {
	p.ApplyTextNormalization = applyTextNormalization
}

func (p *StreamTextToSpeechRequest) SetApplyLanguageTextNormalization(applyLanguageTextNormalization *bool) {
	p.ApplyLanguageTextNormalization = applyLanguageTextNormalization
}

func (p *StreamTextToSpeechRequest) SetAudioEffects(audioEffects *EffectsSpecInput) {
	p.AudioEffects = audioEffects
}

func (p *StreamTextToSpeechRequest) SetEnableLogging(enableLogging *bool) {
	p.EnableLogging = enableLogging
}

func (p *StreamTextToSpeechRequest) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	p.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (p *StreamTextToSpeechRequest) SetOutputFormat(outputFormat *StreamTextToSpeechRequestOutputFormat) {
	p.OutputFormat = outputFormat
}

// StreamTextToSpeechWithTimestampsRequest holds the optional parameters for the API request.
type StreamTextToSpeechWithTimestampsRequest struct {
	Text                            string                                                      `json:"text" url:"-"`
	ModelId                         *string                                                     `json:"model_id,omitempty" url:"-"`
	LanguageCode                    *string                                                     `json:"language_code,omitempty" url:"-"`
	VoiceSettings                   *VoiceSettings                                              `json:"voice_settings,omitempty" url:"-"`
	PronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel        `json:"pronunciation_dictionary_locators,omitempty" url:"-"`
	Seed                            *int                                                        `json:"seed,omitempty" url:"-"`
	PreviousText                    *string                                                     `json:"previous_text,omitempty" url:"-"`
	NextText                        *string                                                     `json:"next_text,omitempty" url:"-"`
	PreviousRequestIds              []string                                                    `json:"previous_request_ids,omitempty" url:"-"`
	NextRequestIds                  []string                                                    `json:"next_request_ids,omitempty" url:"-"`
	UsePvcAsIvc                     *bool                                                       `json:"use_pvc_as_ivc,omitempty" url:"-"`
	ApplyTextNormalization          *BodyTextToSpeechStreamWithTimestampsApplyTextNormalization `json:"apply_text_normalization,omitempty" url:"-"`
	ApplyLanguageTextNormalization  *bool                                                       `json:"apply_language_text_normalization,omitempty" url:"-"`
	AudioEffects                    *EffectsSpecInput                                           `json:"audio_effects,omitempty" url:"-"`
	EnableLogging                   *bool                                                       `json:"-" url:"enable_logging,omitempty"`
	OptimizeStreamingLatency        *int                                                        `json:"-" url:"optimize_streaming_latency,omitempty"`
	OutputFormat                    *StreamWithTimestampsTextToSpeechRequestOutputFormat        `json:"-" url:"output_format,omitempty"`
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetText(text string) {
	p.Text = text
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetModelId(modelId *string) {
	p.ModelId = modelId
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetLanguageCode(languageCode *string) {
	p.LanguageCode = languageCode
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetVoiceSettings(voiceSettings *VoiceSettings) {
	p.VoiceSettings = voiceSettings
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel) {
	p.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetSeed(seed *int) {
	p.Seed = seed
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetPreviousText(previousText *string) {
	p.PreviousText = previousText
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetNextText(nextText *string) {
	p.NextText = nextText
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetPreviousRequestIds(previousRequestIds []string) {
	p.PreviousRequestIds = previousRequestIds
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetNextRequestIds(nextRequestIds []string) {
	p.NextRequestIds = nextRequestIds
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetUsePvcAsIvc(usePvcAsIvc *bool) {
	p.UsePvcAsIvc = usePvcAsIvc
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetApplyTextNormalization(applyTextNormalization *BodyTextToSpeechStreamWithTimestampsApplyTextNormalization) {
	p.ApplyTextNormalization = applyTextNormalization
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetApplyLanguageTextNormalization(applyLanguageTextNormalization *bool) {
	p.ApplyLanguageTextNormalization = applyLanguageTextNormalization
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetAudioEffects(audioEffects *EffectsSpecInput) {
	p.AudioEffects = audioEffects
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetEnableLogging(enableLogging *bool) {
	p.EnableLogging = enableLogging
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	p.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (p *StreamTextToSpeechWithTimestampsRequest) SetOutputFormat(outputFormat *StreamWithTimestampsTextToSpeechRequestOutputFormat) {
	p.OutputFormat = outputFormat
}

// GetVoicesRequest holds the optional parameters for the API request.
type GetVoicesRequest struct {
	WithSettings *bool `json:"-" url:"with_settings,omitempty"`
}

func (p *GetVoicesRequest) SetWithSettings(withSettings *bool) {
	p.WithSettings = withSettings
}

// SearchVoicesRequest holds the optional parameters for the API request.
type SearchVoicesRequest struct {
	NextPageToken        *string   `json:"-" url:"next_page_token,omitempty"`
	PageSize             *int      `json:"-" url:"page_size,omitempty"`
	Search               *string   `json:"-" url:"search,omitempty"`
	Sort                 *string   `json:"-" url:"sort,omitempty"`
	SortDirection        *string   `json:"-" url:"sort_direction,omitempty"`
	VoiceType            *string   `json:"-" url:"voice_type,omitempty"`
	Category             *string   `json:"-" url:"category,omitempty"`
	FineTuningState      *string   `json:"-" url:"fine_tuning_state,omitempty"`
	CollectionId         *string   `json:"-" url:"collection_id,omitempty"`
	Gender               *string   `json:"-" url:"gender,omitempty"`
	Age                  *string   `json:"-" url:"age,omitempty"`
	Language             []*string `json:"-" url:"language,omitempty"`
	Accent               *string   `json:"-" url:"accent,omitempty"`
	UseCases             []*string `json:"-" url:"use_cases,omitempty"`
	MinNoticePeriodDays  *int      `json:"-" url:"min_notice_period_days,omitempty"`
	IncludeCustomRates   *bool     `json:"-" url:"include_custom_rates,omitempty"`
	IncludeLiveModerated *bool     `json:"-" url:"include_live_moderated,omitempty"`
	HighQuality          *bool     `json:"-" url:"high_quality,omitempty"`
	IncludeTotalCount    *bool     `json:"-" url:"include_total_count,omitempty"`
	VoiceIds             []*string `json:"-" url:"voice_ids,omitempty"`
}

func (p *SearchVoicesRequest) SetNextPageToken(nextPageToken *string) {
	p.NextPageToken = nextPageToken
}

func (p *SearchVoicesRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *SearchVoicesRequest) SetSearch(search *string) {
	p.Search = search
}

func (p *SearchVoicesRequest) SetSort(sort *string) {
	p.Sort = sort
}

func (p *SearchVoicesRequest) SetSortDirection(sortDirection *string) {
	p.SortDirection = sortDirection
}

func (p *SearchVoicesRequest) SetVoiceType(voiceType *string) {
	p.VoiceType = voiceType
}

func (p *SearchVoicesRequest) SetCategory(category *string) {
	p.Category = category
}

func (p *SearchVoicesRequest) SetFineTuningState(fineTuningState *string) {
	p.FineTuningState = fineTuningState
}

func (p *SearchVoicesRequest) SetCollectionId(collectionId *string) {
	p.CollectionId = collectionId
}

func (p *SearchVoicesRequest) SetGender(gender *string) {
	p.Gender = gender
}

func (p *SearchVoicesRequest) SetAge(age *string) {
	p.Age = age
}

func (p *SearchVoicesRequest) SetLanguage(language []*string) {
	p.Language = language
}

func (p *SearchVoicesRequest) SetAccent(accent *string) {
	p.Accent = accent
}

func (p *SearchVoicesRequest) SetUseCases(useCases []*string) {
	p.UseCases = useCases
}

func (p *SearchVoicesRequest) SetMinNoticePeriodDays(minNoticePeriodDays *int) {
	p.MinNoticePeriodDays = minNoticePeriodDays
}

func (p *SearchVoicesRequest) SetIncludeCustomRates(includeCustomRates *bool) {
	p.IncludeCustomRates = includeCustomRates
}

func (p *SearchVoicesRequest) SetIncludeLiveModerated(includeLiveModerated *bool) {
	p.IncludeLiveModerated = includeLiveModerated
}

func (p *SearchVoicesRequest) SetHighQuality(highQuality *bool) {
	p.HighQuality = highQuality
}

func (p *SearchVoicesRequest) SetIncludeTotalCount(includeTotalCount *bool) {
	p.IncludeTotalCount = includeTotalCount
}

func (p *SearchVoicesRequest) SetVoiceIds(voiceIds []*string) {
	p.VoiceIds = voiceIds
}

// GetSharedVoicesRequest holds the optional parameters for the API request.
type GetSharedVoicesRequest struct {
	PageSize             *int                            `json:"-" url:"page_size,omitempty"`
	Category             *GetSharedVoicesRequestCategory `json:"-" url:"category,omitempty"`
	Gender               *string                         `json:"-" url:"gender,omitempty"`
	Age                  *string                         `json:"-" url:"age,omitempty"`
	Accent               *string                         `json:"-" url:"accent,omitempty"`
	Language             *string                         `json:"-" url:"language,omitempty"`
	Locale               *string                         `json:"-" url:"locale,omitempty"`
	Search               *string                         `json:"-" url:"search,omitempty"`
	UseCases             []*string                       `json:"-" url:"use_cases,omitempty"`
	Descriptives         []*string                       `json:"-" url:"descriptives,omitempty"`
	Featured             *bool                           `json:"-" url:"featured,omitempty"`
	MinNoticePeriodDays  *int                            `json:"-" url:"min_notice_period_days,omitempty"`
	IncludeCustomRates   *bool                           `json:"-" url:"include_custom_rates,omitempty"`
	IncludeLiveModerated *bool                           `json:"-" url:"include_live_moderated,omitempty"`
	ReaderAppEnabled     *bool                           `json:"-" url:"reader_app_enabled,omitempty"`
	OwnerId              *string                         `json:"-" url:"owner_id,omitempty"`
	Sort                 *GetSharedVoicesRequestSort     `json:"-" url:"sort,omitempty"`
	Page                 *int                            `json:"-" url:"page,omitempty"`
}

func (p *GetSharedVoicesRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *GetSharedVoicesRequest) SetCategory(category *GetSharedVoicesRequestCategory) {
	p.Category = category
}

func (p *GetSharedVoicesRequest) SetGender(gender *string) {
	p.Gender = gender
}

func (p *GetSharedVoicesRequest) SetAge(age *string) {
	p.Age = age
}

func (p *GetSharedVoicesRequest) SetAccent(accent *string) {
	p.Accent = accent
}

func (p *GetSharedVoicesRequest) SetLanguage(language *string) {
	p.Language = language
}

func (p *GetSharedVoicesRequest) SetLocale(locale *string) {
	p.Locale = locale
}

func (p *GetSharedVoicesRequest) SetSearch(search *string) {
	p.Search = search
}

func (p *GetSharedVoicesRequest) SetUseCases(useCases []*string) {
	p.UseCases = useCases
}

func (p *GetSharedVoicesRequest) SetDescriptives(descriptives []*string) {
	p.Descriptives = descriptives
}

func (p *GetSharedVoicesRequest) SetFeatured(featured *bool) {
	p.Featured = featured
}

func (p *GetSharedVoicesRequest) SetMinNoticePeriodDays(minNoticePeriodDays *int) {
	p.MinNoticePeriodDays = minNoticePeriodDays
}

func (p *GetSharedVoicesRequest) SetIncludeCustomRates(includeCustomRates *bool) {
	p.IncludeCustomRates = includeCustomRates
}

func (p *GetSharedVoicesRequest) SetIncludeLiveModerated(includeLiveModerated *bool) {
	p.IncludeLiveModerated = includeLiveModerated
}

func (p *GetSharedVoicesRequest) SetReaderAppEnabled(readerAppEnabled *bool) {
	p.ReaderAppEnabled = readerAppEnabled
}

func (p *GetSharedVoicesRequest) SetOwnerId(ownerId *string) {
	p.OwnerId = ownerId
}

func (p *GetSharedVoicesRequest) SetSort(sort *GetSharedVoicesRequestSort) {
	p.Sort = sort
}

func (p *GetSharedVoicesRequest) SetPage(page *int) {
	p.Page = page
}

// ListPronunciationDictionariesRequest holds the optional parameters for the API request.
type ListPronunciationDictionariesRequest struct {
	Cursor          *string                                   `json:"-" url:"cursor,omitempty"`
	PageSize        *int                                      `json:"-" url:"page_size,omitempty"`
	Sort            *ListPronunciationDictionariesRequestSort `json:"-" url:"sort,omitempty"`
	SortDirection   *string                                   `json:"-" url:"sort_direction,omitempty"`
	IncludeArchived *bool                                     `json:"-" url:"include_archived,omitempty"`
}

func (p *ListPronunciationDictionariesRequest) SetCursor(cursor *string) {
	p.Cursor = cursor
}

func (p *ListPronunciationDictionariesRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListPronunciationDictionariesRequest) SetSort(sort *ListPronunciationDictionariesRequestSort) {
	p.Sort = sort
}

func (p *ListPronunciationDictionariesRequest) SetSortDirection(sortDirection *string) {
	p.SortDirection = sortDirection
}

func (p *ListPronunciationDictionariesRequest) SetIncludeArchived(includeArchived *bool) {
	p.IncludeArchived = includeArchived
}

// ListDubbingRequest holds the optional parameters for the API request.
type ListDubbingRequest struct {
	Cursor              *string                                  `json:"-" url:"cursor,omitempty"`
	PageSize            *int                                     `json:"-" url:"page_size,omitempty"`
	DubbingStatus       *ListDubbingRequestDubbingStatus         `json:"-" url:"dubbing_status,omitempty"`
	DubbingStatuses     []*ListDubbingRequestDubbingStatusesItem `json:"-" url:"dubbing_statuses,omitempty"`
	DubbingModels       []*ListDubbingRequestDubbingModelsItem   `json:"-" url:"dubbing_models,omitempty"`
	TargetLanguageCodes []*string                                `json:"-" url:"target_language_codes,omitempty"`
	CreationSources     []*ListDubbingRequestCreationSourcesItem `json:"-" url:"creation_sources,omitempty"`
	FilterByCreator     *ListDubbingRequestFilterByCreator       `json:"-" url:"filter_by_creator,omitempty"`
	OrderBy             *ListDubbingRequestOrderBy               `json:"-" url:"order_by,omitempty"`
	OrderDirection      *ListDubbingRequestOrderDirection        `json:"-" url:"order_direction,omitempty"`
}

func (p *ListDubbingRequest) SetCursor(cursor *string) {
	p.Cursor = cursor
}

func (p *ListDubbingRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListDubbingRequest) SetDubbingStatus(dubbingStatus *ListDubbingRequestDubbingStatus) {
	p.DubbingStatus = dubbingStatus
}

func (p *ListDubbingRequest) SetDubbingStatuses(dubbingStatuses []*ListDubbingRequestDubbingStatusesItem) {
	p.DubbingStatuses = dubbingStatuses
}

func (p *ListDubbingRequest) SetDubbingModels(dubbingModels []*ListDubbingRequestDubbingModelsItem) {
	p.DubbingModels = dubbingModels
}

func (p *ListDubbingRequest) SetTargetLanguageCodes(targetLanguageCodes []*string) {
	p.TargetLanguageCodes = targetLanguageCodes
}

func (p *ListDubbingRequest) SetCreationSources(creationSources []*ListDubbingRequestCreationSourcesItem) {
	p.CreationSources = creationSources
}

func (p *ListDubbingRequest) SetFilterByCreator(filterByCreator *ListDubbingRequestFilterByCreator) {
	p.FilterByCreator = filterByCreator
}

func (p *ListDubbingRequest) SetOrderBy(orderBy *ListDubbingRequestOrderBy) {
	p.OrderBy = orderBy
}

func (p *ListDubbingRequest) SetOrderDirection(orderDirection *ListDubbingRequestOrderDirection) {
	p.OrderDirection = orderDirection
}

// ListWebhooksRequest holds the optional parameters for the API request.
type ListWebhooksRequest struct {
	IncludeUsages *bool `json:"-" url:"include_usages,omitempty"`
}

func (p *ListWebhooksRequest) SetIncludeUsages(includeUsages *bool) {
	p.IncludeUsages = includeUsages
}

// GetAgentsRequest holds the optional parameters for the API request.
type GetAgentsRequest struct {
	VersionId                                 *string `json:"-" url:"version_id,omitempty"`
	BranchId                                  *string `json:"-" url:"branch_id,omitempty"`
	IncludeProcedureGeneratedWorkflowElements *bool   `json:"-" url:"include_procedure_generated_workflow_elements,omitempty"`
}

func (p *GetAgentsRequest) SetVersionId(versionId *string) {
	p.VersionId = versionId
}

func (p *GetAgentsRequest) SetBranchId(branchId *string) {
	p.BranchId = branchId
}

func (p *GetAgentsRequest) SetIncludeProcedureGeneratedWorkflowElements(includeProcedureGeneratedWorkflowElements *bool) {
	p.IncludeProcedureGeneratedWorkflowElements = includeProcedureGeneratedWorkflowElements
}

// DeleteAgentsRequest holds the optional parameters for the API request.
type DeleteAgentsRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`
}

func (p *DeleteAgentsRequest) SetForce(force *bool) {
	p.Force = force
}

// ListAgentsRequest holds the optional parameters for the API request.
type ListAgentsRequest struct {
	PageSize            *int           `json:"-" url:"page_size,omitempty"`
	Search              *string        `json:"-" url:"search,omitempty"`
	Archived            *bool          `json:"-" url:"archived,omitempty"`
	ShowOnlyOwnedAgents *bool          `json:"-" url:"show_only_owned_agents,omitempty"`
	CreatedByUserId     *string        `json:"-" url:"created_by_user_id,omitempty"`
	Tags                []*string      `json:"-" url:"tags,omitempty"`
	SortDirection       *SortDirection `json:"-" url:"sort_direction,omitempty"`
	SortBy              *AgentSortBy   `json:"-" url:"sort_by,omitempty"`
	Cursor              *string        `json:"-" url:"cursor,omitempty"`
}

func (p *ListAgentsRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListAgentsRequest) SetSearch(search *string) {
	p.Search = search
}

func (p *ListAgentsRequest) SetArchived(archived *bool) {
	p.Archived = archived
}

func (p *ListAgentsRequest) SetShowOnlyOwnedAgents(showOnlyOwnedAgents *bool) {
	p.ShowOnlyOwnedAgents = showOnlyOwnedAgents
}

func (p *ListAgentsRequest) SetCreatedByUserId(createdByUserId *string) {
	p.CreatedByUserId = createdByUserId
}

func (p *ListAgentsRequest) SetTags(tags []*string) {
	p.Tags = tags
}

func (p *ListAgentsRequest) SetSortDirection(sortDirection *SortDirection) {
	p.SortDirection = sortDirection
}

func (p *ListAgentsRequest) SetSortBy(sortBy *AgentSortBy) {
	p.SortBy = sortBy
}

func (p *ListAgentsRequest) SetCursor(cursor *string) {
	p.Cursor = cursor
}

// ListSpeechEngineRequest holds the optional parameters for the API request.
type ListSpeechEngineRequest struct {
	PageSize      *int           `json:"-" url:"page_size,omitempty"`
	Search        *string        `json:"-" url:"search,omitempty"`
	SortDirection *SortDirection `json:"-" url:"sort_direction,omitempty"`
	SortBy        *AgentSortBy   `json:"-" url:"sort_by,omitempty"`
	Cursor        *string        `json:"-" url:"cursor,omitempty"`
}

func (p *ListSpeechEngineRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListSpeechEngineRequest) SetSearch(search *string) {
	p.Search = search
}

func (p *ListSpeechEngineRequest) SetSortDirection(sortDirection *SortDirection) {
	p.SortDirection = sortDirection
}

func (p *ListSpeechEngineRequest) SetSortBy(sortBy *AgentSortBy) {
	p.SortBy = sortBy
}

func (p *ListSpeechEngineRequest) SetCursor(cursor *string) {
	p.Cursor = cursor
}

// ListEnvironmentVariablesRequest holds the optional parameters for the API request.
type ListEnvironmentVariablesRequest struct {
	Cursor      *string                              `json:"-" url:"cursor,omitempty"`
	PageSize    *int                                 `json:"-" url:"page_size,omitempty"`
	Label       *string                              `json:"-" url:"label,omitempty"`
	Environment *string                              `json:"-" url:"environment,omitempty"`
	Type        *ListEnvironmentVariablesRequestType `json:"-" url:"type,omitempty"`
}

func (p *ListEnvironmentVariablesRequest) SetCursor(cursor *string) {
	p.Cursor = cursor
}

func (p *ListEnvironmentVariablesRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListEnvironmentVariablesRequest) SetLabel(label *string) {
	p.Label = label
}

func (p *ListEnvironmentVariablesRequest) SetEnvironment(environment *string) {
	p.Environment = environment
}

func (p *ListEnvironmentVariablesRequest) SetType(type_ *ListEnvironmentVariablesRequestType) {
	p.Type = type_
}

// ListAssetsRequest holds the optional parameters for the API request.
type ListAssetsRequest struct {
	PageSize *int    `json:"-" url:"page_size,omitempty"`
	Cursor   *string `json:"-" url:"cursor,omitempty"`
	Search   *string `json:"-" url:"search,omitempty"`
}

func (p *ListAssetsRequest) SetPageSize(pageSize *int) {
	p.PageSize = pageSize
}

func (p *ListAssetsRequest) SetCursor(cursor *string) {
	p.Cursor = cursor
}

func (p *ListAssetsRequest) SetSearch(search *string) {
	p.Search = search
}
