package configmanager

import (
	"github.com/elevenlabs/elevenlabs-go/apiconfig"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	history                    apiconfig.Config
	textToSoundEffects         apiconfig.Config
	audioIsolation             apiconfig.Config
	audioEffects               apiconfig.Config
	samples                    apiconfig.Config
	audio                      apiconfig.Config
	textToSpeech               apiconfig.Config
	textToDialogue             apiconfig.Config
	speechToSpeech             apiconfig.Config
	textToVoice                apiconfig.Config
	preview                    apiconfig.Config
	user                       apiconfig.Config
	subscription               apiconfig.Config
	collections                apiconfig.Config
	settings                   apiconfig.Config
	accents                    apiconfig.Config
	voices                     apiconfig.Config
	ivc                        apiconfig.Config
	studio                     apiconfig.Config
	pronunciationDictionaries  apiconfig.Config
	projects                   apiconfig.Config
	content                    apiconfig.Config
	snapshots                  apiconfig.Config
	chapters                   apiconfig.Config
	audioNative                apiconfig.Config
	project                    apiconfig.Config
	language                   apiconfig.Config
	transcript                 apiconfig.Config
	dubbing                    apiconfig.Config
	transcripts                apiconfig.Config
	models                     apiconfig.Config
	contentProvenance          apiconfig.Config
	rules                      apiconfig.Config
	apiKeys                    apiconfig.Config
	workspace                  apiconfig.Config
	auditLogs                  apiconfig.Config
	authConnections            apiconfig.Config
	serviceAccounts            apiconfig.Config
	groups                     apiconfig.Config
	members                    apiconfig.Config
	invites                    apiconfig.Config
	resources                  apiconfig.Config
	webhooks                   apiconfig.Config
	music                      apiconfig.Config
	speechToText               apiconfig.Config
	singleUse                  apiconfig.Config
	forcedAlignment            apiconfig.Config
	conversations              apiconfig.Config
	twilio                     apiconfig.Config
	exotel                     apiconfig.Config
	whatsapp                   apiconfig.Config
	agents                     apiconfig.Config
	summaries                  apiconfig.Config
	subagents                  apiconfig.Config
	widget                     apiconfig.Config
	link                       apiconfig.Config
	avatar                     apiconfig.Config
	holdAudio                  apiconfig.Config
	knowledgeBase              apiconfig.Config
	llmUsage                   apiconfig.Config
	tests                      apiconfig.Config
	folders                    apiconfig.Config
	invocations                apiconfig.Config
	users                      apiconfig.Config
	files                      apiconfig.Config
	feedback                   apiconfig.Config
	messages                   apiconfig.Config
	tags                       apiconfig.Config
	triageTickets              apiconfig.Config
	phoneNumbers               apiconfig.Config
	llm                        apiconfig.Config
	liveCount                  apiconfig.Config
	documents                  apiconfig.Config
	crawlJobs                  apiconfig.Config
	document                   apiconfig.Config
	chunk                      apiconfig.Config
	chunks                     apiconfig.Config
	topics                     apiconfig.Config
	tools                      apiconfig.Config
	executions                 apiconfig.Config
	analysisItems              apiconfig.Config
	dependentAgents            apiconfig.Config
	secrets                    apiconfig.Config
	batchCalls                 apiconfig.Config
	sipTrunk                   apiconfig.Config
	mcpServers                 apiconfig.Config
	toolApprovals              apiconfig.Config
	toolConfigs                apiconfig.Config
	whatsappAccounts           apiconfig.Config
	branches                   apiconfig.Config
	versions                   apiconfig.Config
	deployments                apiconfig.Config
	drafts                     apiconfig.Config
	mergeProposals             apiconfig.Config
	config                     apiconfig.Config
	outboundCampaigns          apiconfig.Config
	leads                      apiconfig.Config
	calls                      apiconfig.Config
	attention                  apiconfig.Config
	speechEngine               apiconfig.Config
	procedures                 apiconfig.Config
	analysis                   apiconfig.Config
	environmentVariables       apiconfig.Config
	environments               apiconfig.Config
	realtime                   apiconfig.Config
	compositionPlan            apiconfig.Config
	finetunes                  apiconfig.Config
	orders                     apiconfig.Config
	media                      apiconfig.Config
	items                      apiconfig.Config
	deliverables               apiconfig.Config
	languages                  apiconfig.Config
	pvc                        apiconfig.Config
	waveform                   apiconfig.Config
	speakers                   apiconfig.Config
	captcha                    apiconfig.Config
	verification               apiconfig.Config
	video                      apiconfig.Config
	image                      apiconfig.Config
	runs                       apiconfig.Config
	templates                  apiconfig.Config
	assets                     apiconfig.Config
	usage                      apiconfig.Config
	requests                   apiconfig.Config
	services                   apiconfig.Config
	products                   apiconfig.Config
	stock                      apiconfig.Config
	staff                      apiconfig.Config
	receptionists              apiconfig.Config
	transferRules              apiconfig.Config
	holidays                   apiconfig.Config
	locations                  apiconfig.Config
	businessProfile            apiconfig.Config
	knowledgeSources           apiconfig.Config
	workspaces                 apiconfig.Config
	subscriptions              apiconfig.Config
	tiers                      apiconfig.Config
	billing                    apiconfig.Config
	textToDialogueMultiContext apiconfig.Config
	textToSpeechMultiContext   apiconfig.Config
	translation                apiconfig.Config
	subtitling                 apiconfig.Config
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config apiconfig.Config) *ConfigManager {
	return &ConfigManager{
		history:                    config,
		textToSoundEffects:         config,
		audioIsolation:             config,
		audioEffects:               config,
		samples:                    config,
		audio:                      config,
		textToSpeech:               config,
		textToDialogue:             config,
		speechToSpeech:             config,
		textToVoice:                config,
		preview:                    config,
		user:                       config,
		subscription:               config,
		collections:                config,
		settings:                   config,
		accents:                    config,
		voices:                     config,
		ivc:                        config,
		studio:                     config,
		pronunciationDictionaries:  config,
		projects:                   config,
		content:                    config,
		snapshots:                  config,
		chapters:                   config,
		audioNative:                config,
		project:                    config,
		language:                   config,
		transcript:                 config,
		dubbing:                    config,
		transcripts:                config,
		models:                     config,
		contentProvenance:          config,
		rules:                      config,
		apiKeys:                    config,
		workspace:                  config,
		auditLogs:                  config,
		authConnections:            config,
		serviceAccounts:            config,
		groups:                     config,
		members:                    config,
		invites:                    config,
		resources:                  config,
		webhooks:                   config,
		music:                      config,
		speechToText:               config,
		singleUse:                  config,
		forcedAlignment:            config,
		conversations:              config,
		twilio:                     config,
		exotel:                     config,
		whatsapp:                   config,
		agents:                     config,
		summaries:                  config,
		subagents:                  config,
		widget:                     config,
		link:                       config,
		avatar:                     config,
		holdAudio:                  config,
		knowledgeBase:              config,
		llmUsage:                   config,
		tests:                      config,
		folders:                    config,
		invocations:                config,
		users:                      config,
		files:                      config,
		feedback:                   config,
		messages:                   config,
		tags:                       config,
		triageTickets:              config,
		phoneNumbers:               config,
		llm:                        config,
		liveCount:                  config,
		documents:                  config,
		crawlJobs:                  config,
		document:                   config,
		chunk:                      config,
		chunks:                     config,
		topics:                     config,
		tools:                      config,
		executions:                 config,
		analysisItems:              config,
		dependentAgents:            config,
		secrets:                    config,
		batchCalls:                 config,
		sipTrunk:                   config,
		mcpServers:                 config,
		toolApprovals:              config,
		toolConfigs:                config,
		whatsappAccounts:           config,
		branches:                   config,
		versions:                   config,
		deployments:                config,
		drafts:                     config,
		mergeProposals:             config,
		config:                     config,
		outboundCampaigns:          config,
		leads:                      config,
		calls:                      config,
		attention:                  config,
		speechEngine:               config,
		procedures:                 config,
		analysis:                   config,
		environmentVariables:       config,
		environments:               config,
		realtime:                   config,
		compositionPlan:            config,
		finetunes:                  config,
		orders:                     config,
		media:                      config,
		items:                      config,
		deliverables:               config,
		languages:                  config,
		pvc:                        config,
		waveform:                   config,
		speakers:                   config,
		captcha:                    config,
		verification:               config,
		video:                      config,
		image:                      config,
		runs:                       config,
		templates:                  config,
		assets:                     config,
		usage:                      config,
		requests:                   config,
		services:                   config,
		products:                   config,
		stock:                      config,
		staff:                      config,
		receptionists:              config,
		transferRules:              config,
		holidays:                   config,
		locations:                  config,
		businessProfile:            config,
		knowledgeSources:           config,
		workspaces:                 config,
		subscriptions:              config,
		tiers:                      config,
		billing:                    config,
		textToDialogueMultiContext: config,
		textToSpeechMultiContext:   config,
		translation:                config,
		subtitling:                 config,
	}
}

// SetBaseURL updates the BaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseURL(baseURL string) {
	c.history.SetBaseURL(baseURL)
	c.textToSoundEffects.SetBaseURL(baseURL)
	c.audioIsolation.SetBaseURL(baseURL)
	c.audioEffects.SetBaseURL(baseURL)
	c.samples.SetBaseURL(baseURL)
	c.audio.SetBaseURL(baseURL)
	c.textToSpeech.SetBaseURL(baseURL)
	c.textToDialogue.SetBaseURL(baseURL)
	c.speechToSpeech.SetBaseURL(baseURL)
	c.textToVoice.SetBaseURL(baseURL)
	c.preview.SetBaseURL(baseURL)
	c.user.SetBaseURL(baseURL)
	c.subscription.SetBaseURL(baseURL)
	c.collections.SetBaseURL(baseURL)
	c.settings.SetBaseURL(baseURL)
	c.accents.SetBaseURL(baseURL)
	c.voices.SetBaseURL(baseURL)
	c.ivc.SetBaseURL(baseURL)
	c.studio.SetBaseURL(baseURL)
	c.pronunciationDictionaries.SetBaseURL(baseURL)
	c.projects.SetBaseURL(baseURL)
	c.content.SetBaseURL(baseURL)
	c.snapshots.SetBaseURL(baseURL)
	c.chapters.SetBaseURL(baseURL)
	c.audioNative.SetBaseURL(baseURL)
	c.project.SetBaseURL(baseURL)
	c.language.SetBaseURL(baseURL)
	c.transcript.SetBaseURL(baseURL)
	c.dubbing.SetBaseURL(baseURL)
	c.transcripts.SetBaseURL(baseURL)
	c.models.SetBaseURL(baseURL)
	c.contentProvenance.SetBaseURL(baseURL)
	c.rules.SetBaseURL(baseURL)
	c.apiKeys.SetBaseURL(baseURL)
	c.workspace.SetBaseURL(baseURL)
	c.auditLogs.SetBaseURL(baseURL)
	c.authConnections.SetBaseURL(baseURL)
	c.serviceAccounts.SetBaseURL(baseURL)
	c.groups.SetBaseURL(baseURL)
	c.members.SetBaseURL(baseURL)
	c.invites.SetBaseURL(baseURL)
	c.resources.SetBaseURL(baseURL)
	c.webhooks.SetBaseURL(baseURL)
	c.music.SetBaseURL(baseURL)
	c.speechToText.SetBaseURL(baseURL)
	c.singleUse.SetBaseURL(baseURL)
	c.forcedAlignment.SetBaseURL(baseURL)
	c.conversations.SetBaseURL(baseURL)
	c.twilio.SetBaseURL(baseURL)
	c.exotel.SetBaseURL(baseURL)
	c.whatsapp.SetBaseURL(baseURL)
	c.agents.SetBaseURL(baseURL)
	c.summaries.SetBaseURL(baseURL)
	c.subagents.SetBaseURL(baseURL)
	c.widget.SetBaseURL(baseURL)
	c.link.SetBaseURL(baseURL)
	c.avatar.SetBaseURL(baseURL)
	c.holdAudio.SetBaseURL(baseURL)
	c.knowledgeBase.SetBaseURL(baseURL)
	c.llmUsage.SetBaseURL(baseURL)
	c.tests.SetBaseURL(baseURL)
	c.folders.SetBaseURL(baseURL)
	c.invocations.SetBaseURL(baseURL)
	c.users.SetBaseURL(baseURL)
	c.files.SetBaseURL(baseURL)
	c.feedback.SetBaseURL(baseURL)
	c.messages.SetBaseURL(baseURL)
	c.tags.SetBaseURL(baseURL)
	c.triageTickets.SetBaseURL(baseURL)
	c.phoneNumbers.SetBaseURL(baseURL)
	c.llm.SetBaseURL(baseURL)
	c.liveCount.SetBaseURL(baseURL)
	c.documents.SetBaseURL(baseURL)
	c.crawlJobs.SetBaseURL(baseURL)
	c.document.SetBaseURL(baseURL)
	c.chunk.SetBaseURL(baseURL)
	c.chunks.SetBaseURL(baseURL)
	c.topics.SetBaseURL(baseURL)
	c.tools.SetBaseURL(baseURL)
	c.executions.SetBaseURL(baseURL)
	c.analysisItems.SetBaseURL(baseURL)
	c.dependentAgents.SetBaseURL(baseURL)
	c.secrets.SetBaseURL(baseURL)
	c.batchCalls.SetBaseURL(baseURL)
	c.sipTrunk.SetBaseURL(baseURL)
	c.mcpServers.SetBaseURL(baseURL)
	c.toolApprovals.SetBaseURL(baseURL)
	c.toolConfigs.SetBaseURL(baseURL)
	c.whatsappAccounts.SetBaseURL(baseURL)
	c.branches.SetBaseURL(baseURL)
	c.versions.SetBaseURL(baseURL)
	c.deployments.SetBaseURL(baseURL)
	c.drafts.SetBaseURL(baseURL)
	c.mergeProposals.SetBaseURL(baseURL)
	c.config.SetBaseURL(baseURL)
	c.outboundCampaigns.SetBaseURL(baseURL)
	c.leads.SetBaseURL(baseURL)
	c.calls.SetBaseURL(baseURL)
	c.attention.SetBaseURL(baseURL)
	c.speechEngine.SetBaseURL(baseURL)
	c.procedures.SetBaseURL(baseURL)
	c.analysis.SetBaseURL(baseURL)
	c.environmentVariables.SetBaseURL(baseURL)
	c.environments.SetBaseURL(baseURL)
	c.realtime.SetBaseURL(baseURL)
	c.compositionPlan.SetBaseURL(baseURL)
	c.finetunes.SetBaseURL(baseURL)
	c.orders.SetBaseURL(baseURL)
	c.media.SetBaseURL(baseURL)
	c.items.SetBaseURL(baseURL)
	c.deliverables.SetBaseURL(baseURL)
	c.languages.SetBaseURL(baseURL)
	c.pvc.SetBaseURL(baseURL)
	c.waveform.SetBaseURL(baseURL)
	c.speakers.SetBaseURL(baseURL)
	c.captcha.SetBaseURL(baseURL)
	c.verification.SetBaseURL(baseURL)
	c.video.SetBaseURL(baseURL)
	c.image.SetBaseURL(baseURL)
	c.runs.SetBaseURL(baseURL)
	c.templates.SetBaseURL(baseURL)
	c.assets.SetBaseURL(baseURL)
	c.usage.SetBaseURL(baseURL)
	c.requests.SetBaseURL(baseURL)
	c.services.SetBaseURL(baseURL)
	c.products.SetBaseURL(baseURL)
	c.stock.SetBaseURL(baseURL)
	c.staff.SetBaseURL(baseURL)
	c.receptionists.SetBaseURL(baseURL)
	c.transferRules.SetBaseURL(baseURL)
	c.holidays.SetBaseURL(baseURL)
	c.locations.SetBaseURL(baseURL)
	c.businessProfile.SetBaseURL(baseURL)
	c.knowledgeSources.SetBaseURL(baseURL)
	c.workspaces.SetBaseURL(baseURL)
	c.subscriptions.SetBaseURL(baseURL)
	c.tiers.SetBaseURL(baseURL)
	c.billing.SetBaseURL(baseURL)
	c.textToDialogueMultiContext.SetBaseURL(baseURL)
	c.textToSpeechMultiContext.SetBaseURL(baseURL)
	c.translation.SetBaseURL(baseURL)
	c.subtitling.SetBaseURL(baseURL)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.history.SetTimeout(timeout)
	c.textToSoundEffects.SetTimeout(timeout)
	c.audioIsolation.SetTimeout(timeout)
	c.audioEffects.SetTimeout(timeout)
	c.samples.SetTimeout(timeout)
	c.audio.SetTimeout(timeout)
	c.textToSpeech.SetTimeout(timeout)
	c.textToDialogue.SetTimeout(timeout)
	c.speechToSpeech.SetTimeout(timeout)
	c.textToVoice.SetTimeout(timeout)
	c.preview.SetTimeout(timeout)
	c.user.SetTimeout(timeout)
	c.subscription.SetTimeout(timeout)
	c.collections.SetTimeout(timeout)
	c.settings.SetTimeout(timeout)
	c.accents.SetTimeout(timeout)
	c.voices.SetTimeout(timeout)
	c.ivc.SetTimeout(timeout)
	c.studio.SetTimeout(timeout)
	c.pronunciationDictionaries.SetTimeout(timeout)
	c.projects.SetTimeout(timeout)
	c.content.SetTimeout(timeout)
	c.snapshots.SetTimeout(timeout)
	c.chapters.SetTimeout(timeout)
	c.audioNative.SetTimeout(timeout)
	c.project.SetTimeout(timeout)
	c.language.SetTimeout(timeout)
	c.transcript.SetTimeout(timeout)
	c.dubbing.SetTimeout(timeout)
	c.transcripts.SetTimeout(timeout)
	c.models.SetTimeout(timeout)
	c.contentProvenance.SetTimeout(timeout)
	c.rules.SetTimeout(timeout)
	c.apiKeys.SetTimeout(timeout)
	c.workspace.SetTimeout(timeout)
	c.auditLogs.SetTimeout(timeout)
	c.authConnections.SetTimeout(timeout)
	c.serviceAccounts.SetTimeout(timeout)
	c.groups.SetTimeout(timeout)
	c.members.SetTimeout(timeout)
	c.invites.SetTimeout(timeout)
	c.resources.SetTimeout(timeout)
	c.webhooks.SetTimeout(timeout)
	c.music.SetTimeout(timeout)
	c.speechToText.SetTimeout(timeout)
	c.singleUse.SetTimeout(timeout)
	c.forcedAlignment.SetTimeout(timeout)
	c.conversations.SetTimeout(timeout)
	c.twilio.SetTimeout(timeout)
	c.exotel.SetTimeout(timeout)
	c.whatsapp.SetTimeout(timeout)
	c.agents.SetTimeout(timeout)
	c.summaries.SetTimeout(timeout)
	c.subagents.SetTimeout(timeout)
	c.widget.SetTimeout(timeout)
	c.link.SetTimeout(timeout)
	c.avatar.SetTimeout(timeout)
	c.holdAudio.SetTimeout(timeout)
	c.knowledgeBase.SetTimeout(timeout)
	c.llmUsage.SetTimeout(timeout)
	c.tests.SetTimeout(timeout)
	c.folders.SetTimeout(timeout)
	c.invocations.SetTimeout(timeout)
	c.users.SetTimeout(timeout)
	c.files.SetTimeout(timeout)
	c.feedback.SetTimeout(timeout)
	c.messages.SetTimeout(timeout)
	c.tags.SetTimeout(timeout)
	c.triageTickets.SetTimeout(timeout)
	c.phoneNumbers.SetTimeout(timeout)
	c.llm.SetTimeout(timeout)
	c.liveCount.SetTimeout(timeout)
	c.documents.SetTimeout(timeout)
	c.crawlJobs.SetTimeout(timeout)
	c.document.SetTimeout(timeout)
	c.chunk.SetTimeout(timeout)
	c.chunks.SetTimeout(timeout)
	c.topics.SetTimeout(timeout)
	c.tools.SetTimeout(timeout)
	c.executions.SetTimeout(timeout)
	c.analysisItems.SetTimeout(timeout)
	c.dependentAgents.SetTimeout(timeout)
	c.secrets.SetTimeout(timeout)
	c.batchCalls.SetTimeout(timeout)
	c.sipTrunk.SetTimeout(timeout)
	c.mcpServers.SetTimeout(timeout)
	c.toolApprovals.SetTimeout(timeout)
	c.toolConfigs.SetTimeout(timeout)
	c.whatsappAccounts.SetTimeout(timeout)
	c.branches.SetTimeout(timeout)
	c.versions.SetTimeout(timeout)
	c.deployments.SetTimeout(timeout)
	c.drafts.SetTimeout(timeout)
	c.mergeProposals.SetTimeout(timeout)
	c.config.SetTimeout(timeout)
	c.outboundCampaigns.SetTimeout(timeout)
	c.leads.SetTimeout(timeout)
	c.calls.SetTimeout(timeout)
	c.attention.SetTimeout(timeout)
	c.speechEngine.SetTimeout(timeout)
	c.procedures.SetTimeout(timeout)
	c.analysis.SetTimeout(timeout)
	c.environmentVariables.SetTimeout(timeout)
	c.environments.SetTimeout(timeout)
	c.realtime.SetTimeout(timeout)
	c.compositionPlan.SetTimeout(timeout)
	c.finetunes.SetTimeout(timeout)
	c.orders.SetTimeout(timeout)
	c.media.SetTimeout(timeout)
	c.items.SetTimeout(timeout)
	c.deliverables.SetTimeout(timeout)
	c.languages.SetTimeout(timeout)
	c.pvc.SetTimeout(timeout)
	c.waveform.SetTimeout(timeout)
	c.speakers.SetTimeout(timeout)
	c.captcha.SetTimeout(timeout)
	c.verification.SetTimeout(timeout)
	c.video.SetTimeout(timeout)
	c.image.SetTimeout(timeout)
	c.runs.SetTimeout(timeout)
	c.templates.SetTimeout(timeout)
	c.assets.SetTimeout(timeout)
	c.usage.SetTimeout(timeout)
	c.requests.SetTimeout(timeout)
	c.services.SetTimeout(timeout)
	c.products.SetTimeout(timeout)
	c.stock.SetTimeout(timeout)
	c.staff.SetTimeout(timeout)
	c.receptionists.SetTimeout(timeout)
	c.transferRules.SetTimeout(timeout)
	c.holidays.SetTimeout(timeout)
	c.locations.SetTimeout(timeout)
	c.businessProfile.SetTimeout(timeout)
	c.knowledgeSources.SetTimeout(timeout)
	c.workspaces.SetTimeout(timeout)
	c.subscriptions.SetTimeout(timeout)
	c.tiers.SetTimeout(timeout)
	c.billing.SetTimeout(timeout)
	c.textToDialogueMultiContext.SetTimeout(timeout)
	c.textToSpeechMultiContext.SetTimeout(timeout)
	c.translation.SetTimeout(timeout)
	c.subtitling.SetTimeout(timeout)
}

// SetRetryConfig updates the retry configuration across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetRetryConfig(retry apiconfig.RetryConfig) {
	c.history.SetRetryConfig(retry)
	c.textToSoundEffects.SetRetryConfig(retry)
	c.audioIsolation.SetRetryConfig(retry)
	c.audioEffects.SetRetryConfig(retry)
	c.samples.SetRetryConfig(retry)
	c.audio.SetRetryConfig(retry)
	c.textToSpeech.SetRetryConfig(retry)
	c.textToDialogue.SetRetryConfig(retry)
	c.speechToSpeech.SetRetryConfig(retry)
	c.textToVoice.SetRetryConfig(retry)
	c.preview.SetRetryConfig(retry)
	c.user.SetRetryConfig(retry)
	c.subscription.SetRetryConfig(retry)
	c.collections.SetRetryConfig(retry)
	c.settings.SetRetryConfig(retry)
	c.accents.SetRetryConfig(retry)
	c.voices.SetRetryConfig(retry)
	c.ivc.SetRetryConfig(retry)
	c.studio.SetRetryConfig(retry)
	c.pronunciationDictionaries.SetRetryConfig(retry)
	c.projects.SetRetryConfig(retry)
	c.content.SetRetryConfig(retry)
	c.snapshots.SetRetryConfig(retry)
	c.chapters.SetRetryConfig(retry)
	c.audioNative.SetRetryConfig(retry)
	c.project.SetRetryConfig(retry)
	c.language.SetRetryConfig(retry)
	c.transcript.SetRetryConfig(retry)
	c.dubbing.SetRetryConfig(retry)
	c.transcripts.SetRetryConfig(retry)
	c.models.SetRetryConfig(retry)
	c.contentProvenance.SetRetryConfig(retry)
	c.rules.SetRetryConfig(retry)
	c.apiKeys.SetRetryConfig(retry)
	c.workspace.SetRetryConfig(retry)
	c.auditLogs.SetRetryConfig(retry)
	c.authConnections.SetRetryConfig(retry)
	c.serviceAccounts.SetRetryConfig(retry)
	c.groups.SetRetryConfig(retry)
	c.members.SetRetryConfig(retry)
	c.invites.SetRetryConfig(retry)
	c.resources.SetRetryConfig(retry)
	c.webhooks.SetRetryConfig(retry)
	c.music.SetRetryConfig(retry)
	c.speechToText.SetRetryConfig(retry)
	c.singleUse.SetRetryConfig(retry)
	c.forcedAlignment.SetRetryConfig(retry)
	c.conversations.SetRetryConfig(retry)
	c.twilio.SetRetryConfig(retry)
	c.exotel.SetRetryConfig(retry)
	c.whatsapp.SetRetryConfig(retry)
	c.agents.SetRetryConfig(retry)
	c.summaries.SetRetryConfig(retry)
	c.subagents.SetRetryConfig(retry)
	c.widget.SetRetryConfig(retry)
	c.link.SetRetryConfig(retry)
	c.avatar.SetRetryConfig(retry)
	c.holdAudio.SetRetryConfig(retry)
	c.knowledgeBase.SetRetryConfig(retry)
	c.llmUsage.SetRetryConfig(retry)
	c.tests.SetRetryConfig(retry)
	c.folders.SetRetryConfig(retry)
	c.invocations.SetRetryConfig(retry)
	c.users.SetRetryConfig(retry)
	c.files.SetRetryConfig(retry)
	c.feedback.SetRetryConfig(retry)
	c.messages.SetRetryConfig(retry)
	c.tags.SetRetryConfig(retry)
	c.triageTickets.SetRetryConfig(retry)
	c.phoneNumbers.SetRetryConfig(retry)
	c.llm.SetRetryConfig(retry)
	c.liveCount.SetRetryConfig(retry)
	c.documents.SetRetryConfig(retry)
	c.crawlJobs.SetRetryConfig(retry)
	c.document.SetRetryConfig(retry)
	c.chunk.SetRetryConfig(retry)
	c.chunks.SetRetryConfig(retry)
	c.topics.SetRetryConfig(retry)
	c.tools.SetRetryConfig(retry)
	c.executions.SetRetryConfig(retry)
	c.analysisItems.SetRetryConfig(retry)
	c.dependentAgents.SetRetryConfig(retry)
	c.secrets.SetRetryConfig(retry)
	c.batchCalls.SetRetryConfig(retry)
	c.sipTrunk.SetRetryConfig(retry)
	c.mcpServers.SetRetryConfig(retry)
	c.toolApprovals.SetRetryConfig(retry)
	c.toolConfigs.SetRetryConfig(retry)
	c.whatsappAccounts.SetRetryConfig(retry)
	c.branches.SetRetryConfig(retry)
	c.versions.SetRetryConfig(retry)
	c.deployments.SetRetryConfig(retry)
	c.drafts.SetRetryConfig(retry)
	c.mergeProposals.SetRetryConfig(retry)
	c.config.SetRetryConfig(retry)
	c.outboundCampaigns.SetRetryConfig(retry)
	c.leads.SetRetryConfig(retry)
	c.calls.SetRetryConfig(retry)
	c.attention.SetRetryConfig(retry)
	c.speechEngine.SetRetryConfig(retry)
	c.procedures.SetRetryConfig(retry)
	c.analysis.SetRetryConfig(retry)
	c.environmentVariables.SetRetryConfig(retry)
	c.environments.SetRetryConfig(retry)
	c.realtime.SetRetryConfig(retry)
	c.compositionPlan.SetRetryConfig(retry)
	c.finetunes.SetRetryConfig(retry)
	c.orders.SetRetryConfig(retry)
	c.media.SetRetryConfig(retry)
	c.items.SetRetryConfig(retry)
	c.deliverables.SetRetryConfig(retry)
	c.languages.SetRetryConfig(retry)
	c.pvc.SetRetryConfig(retry)
	c.waveform.SetRetryConfig(retry)
	c.speakers.SetRetryConfig(retry)
	c.captcha.SetRetryConfig(retry)
	c.verification.SetRetryConfig(retry)
	c.video.SetRetryConfig(retry)
	c.image.SetRetryConfig(retry)
	c.runs.SetRetryConfig(retry)
	c.templates.SetRetryConfig(retry)
	c.assets.SetRetryConfig(retry)
	c.usage.SetRetryConfig(retry)
	c.requests.SetRetryConfig(retry)
	c.services.SetRetryConfig(retry)
	c.products.SetRetryConfig(retry)
	c.stock.SetRetryConfig(retry)
	c.staff.SetRetryConfig(retry)
	c.receptionists.SetRetryConfig(retry)
	c.transferRules.SetRetryConfig(retry)
	c.holidays.SetRetryConfig(retry)
	c.locations.SetRetryConfig(retry)
	c.businessProfile.SetRetryConfig(retry)
	c.knowledgeSources.SetRetryConfig(retry)
	c.workspaces.SetRetryConfig(retry)
	c.subscriptions.SetRetryConfig(retry)
	c.tiers.SetRetryConfig(retry)
	c.billing.SetRetryConfig(retry)
	c.textToDialogueMultiContext.SetRetryConfig(retry)
	c.textToSpeechMultiContext.SetRetryConfig(retry)
	c.translation.SetRetryConfig(retry)
	c.subtitling.SetRetryConfig(retry)
}

// GetHistory returns the configuration for the History service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetHistory() *apiconfig.Config {
	return &c.history
}

// GetTextToSoundEffects returns the configuration for the TextToSoundEffects service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTextToSoundEffects() *apiconfig.Config {
	return &c.textToSoundEffects
}

// GetAudioIsolation returns the configuration for the AudioIsolation service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAudioIsolation() *apiconfig.Config {
	return &c.audioIsolation
}

// GetAudioEffects returns the configuration for the AudioEffects service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAudioEffects() *apiconfig.Config {
	return &c.audioEffects
}

// GetSamples returns the configuration for the Samples service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSamples() *apiconfig.Config {
	return &c.samples
}

// GetAudio returns the configuration for the Audio service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAudio() *apiconfig.Config {
	return &c.audio
}

// GetTextToSpeech returns the configuration for the TextToSpeech service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTextToSpeech() *apiconfig.Config {
	return &c.textToSpeech
}

// GetTextToDialogue returns the configuration for the TextToDialogue service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTextToDialogue() *apiconfig.Config {
	return &c.textToDialogue
}

// GetSpeechToSpeech returns the configuration for the SpeechToSpeech service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSpeechToSpeech() *apiconfig.Config {
	return &c.speechToSpeech
}

// GetTextToVoice returns the configuration for the TextToVoice service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTextToVoice() *apiconfig.Config {
	return &c.textToVoice
}

// GetPreview returns the configuration for the Preview service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPreview() *apiconfig.Config {
	return &c.preview
}

// GetUser returns the configuration for the User service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetUser() *apiconfig.Config {
	return &c.user
}

// GetSubscription returns the configuration for the Subscription service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSubscription() *apiconfig.Config {
	return &c.subscription
}

// GetCollections returns the configuration for the Collections service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetCollections() *apiconfig.Config {
	return &c.collections
}

// GetSettings returns the configuration for the Settings service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSettings() *apiconfig.Config {
	return &c.settings
}

// GetAccents returns the configuration for the Accents service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAccents() *apiconfig.Config {
	return &c.accents
}

// GetVoices returns the configuration for the Voices service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetVoices() *apiconfig.Config {
	return &c.voices
}

// GetIvc returns the configuration for the Ivc service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetIvc() *apiconfig.Config {
	return &c.ivc
}

// GetStudio returns the configuration for the Studio service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetStudio() *apiconfig.Config {
	return &c.studio
}

// GetPronunciationDictionaries returns the configuration for the PronunciationDictionaries service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPronunciationDictionaries() *apiconfig.Config {
	return &c.pronunciationDictionaries
}

// GetProjects returns the configuration for the Projects service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetProjects() *apiconfig.Config {
	return &c.projects
}

// GetContent returns the configuration for the Content service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetContent() *apiconfig.Config {
	return &c.content
}

// GetSnapshots returns the configuration for the Snapshots service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSnapshots() *apiconfig.Config {
	return &c.snapshots
}

// GetChapters returns the configuration for the Chapters service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetChapters() *apiconfig.Config {
	return &c.chapters
}

// GetAudioNative returns the configuration for the AudioNative service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAudioNative() *apiconfig.Config {
	return &c.audioNative
}

// GetProject returns the configuration for the Project service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetProject() *apiconfig.Config {
	return &c.project
}

// GetLanguage returns the configuration for the Language service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLanguage() *apiconfig.Config {
	return &c.language
}

// GetTranscript returns the configuration for the Transcript service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTranscript() *apiconfig.Config {
	return &c.transcript
}

// GetDubbing returns the configuration for the Dubbing service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDubbing() *apiconfig.Config {
	return &c.dubbing
}

// GetTranscripts returns the configuration for the Transcripts service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTranscripts() *apiconfig.Config {
	return &c.transcripts
}

// GetModels returns the configuration for the Models service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetModels() *apiconfig.Config {
	return &c.models
}

// GetContentProvenance returns the configuration for the ContentProvenance service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetContentProvenance() *apiconfig.Config {
	return &c.contentProvenance
}

// GetRules returns the configuration for the Rules service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetRules() *apiconfig.Config {
	return &c.rules
}

// GetApiKeys returns the configuration for the ApiKeys service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetApiKeys() *apiconfig.Config {
	return &c.apiKeys
}

// GetWorkspace returns the configuration for the Workspace service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWorkspace() *apiconfig.Config {
	return &c.workspace
}

// GetAuditLogs returns the configuration for the AuditLogs service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAuditLogs() *apiconfig.Config {
	return &c.auditLogs
}

// GetAuthConnections returns the configuration for the AuthConnections service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAuthConnections() *apiconfig.Config {
	return &c.authConnections
}

// GetServiceAccounts returns the configuration for the ServiceAccounts service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetServiceAccounts() *apiconfig.Config {
	return &c.serviceAccounts
}

// GetGroups returns the configuration for the Groups service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetGroups() *apiconfig.Config {
	return &c.groups
}

// GetMembers returns the configuration for the Members service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMembers() *apiconfig.Config {
	return &c.members
}

// GetInvites returns the configuration for the Invites service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetInvites() *apiconfig.Config {
	return &c.invites
}

// GetResources returns the configuration for the Resources service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetResources() *apiconfig.Config {
	return &c.resources
}

// GetWebhooks returns the configuration for the Webhooks service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWebhooks() *apiconfig.Config {
	return &c.webhooks
}

// GetMusic returns the configuration for the Music service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMusic() *apiconfig.Config {
	return &c.music
}

// GetSpeechToText returns the configuration for the SpeechToText service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSpeechToText() *apiconfig.Config {
	return &c.speechToText
}

// GetSingleUse returns the configuration for the SingleUse service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSingleUse() *apiconfig.Config {
	return &c.singleUse
}

// GetForcedAlignment returns the configuration for the ForcedAlignment service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetForcedAlignment() *apiconfig.Config {
	return &c.forcedAlignment
}

// GetConversations returns the configuration for the Conversations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetConversations() *apiconfig.Config {
	return &c.conversations
}

// GetTwilio returns the configuration for the Twilio service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTwilio() *apiconfig.Config {
	return &c.twilio
}

// GetExotel returns the configuration for the Exotel service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetExotel() *apiconfig.Config {
	return &c.exotel
}

// GetWhatsapp returns the configuration for the Whatsapp service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWhatsapp() *apiconfig.Config {
	return &c.whatsapp
}

// GetAgents returns the configuration for the Agents service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAgents() *apiconfig.Config {
	return &c.agents
}

// GetSummaries returns the configuration for the Summaries service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSummaries() *apiconfig.Config {
	return &c.summaries
}

// GetSubagents returns the configuration for the Subagents service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSubagents() *apiconfig.Config {
	return &c.subagents
}

// GetWidget returns the configuration for the Widget service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWidget() *apiconfig.Config {
	return &c.widget
}

// GetLink returns the configuration for the Link service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLink() *apiconfig.Config {
	return &c.link
}

// GetAvatar returns the configuration for the Avatar service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAvatar() *apiconfig.Config {
	return &c.avatar
}

// GetHoldAudio returns the configuration for the HoldAudio service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetHoldAudio() *apiconfig.Config {
	return &c.holdAudio
}

// GetKnowledgeBase returns the configuration for the KnowledgeBase service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetKnowledgeBase() *apiconfig.Config {
	return &c.knowledgeBase
}

// GetLlmUsage returns the configuration for the LlmUsage service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLlmUsage() *apiconfig.Config {
	return &c.llmUsage
}

// GetTests returns the configuration for the Tests service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTests() *apiconfig.Config {
	return &c.tests
}

// GetFolders returns the configuration for the Folders service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetFolders() *apiconfig.Config {
	return &c.folders
}

// GetInvocations returns the configuration for the Invocations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetInvocations() *apiconfig.Config {
	return &c.invocations
}

// GetUsers returns the configuration for the Users service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetUsers() *apiconfig.Config {
	return &c.users
}

// GetFiles returns the configuration for the Files service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetFiles() *apiconfig.Config {
	return &c.files
}

// GetFeedback returns the configuration for the Feedback service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetFeedback() *apiconfig.Config {
	return &c.feedback
}

// GetMessages returns the configuration for the Messages service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMessages() *apiconfig.Config {
	return &c.messages
}

// GetTags returns the configuration for the Tags service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTags() *apiconfig.Config {
	return &c.tags
}

// GetTriageTickets returns the configuration for the TriageTickets service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTriageTickets() *apiconfig.Config {
	return &c.triageTickets
}

// GetPhoneNumbers returns the configuration for the PhoneNumbers service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPhoneNumbers() *apiconfig.Config {
	return &c.phoneNumbers
}

// GetLlm returns the configuration for the Llm service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLlm() *apiconfig.Config {
	return &c.llm
}

// GetLiveCount returns the configuration for the LiveCount service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLiveCount() *apiconfig.Config {
	return &c.liveCount
}

// GetDocuments returns the configuration for the Documents service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDocuments() *apiconfig.Config {
	return &c.documents
}

// GetCrawlJobs returns the configuration for the CrawlJobs service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetCrawlJobs() *apiconfig.Config {
	return &c.crawlJobs
}

// GetDocument returns the configuration for the Document service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDocument() *apiconfig.Config {
	return &c.document
}

// GetChunk returns the configuration for the Chunk service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetChunk() *apiconfig.Config {
	return &c.chunk
}

// GetChunks returns the configuration for the Chunks service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetChunks() *apiconfig.Config {
	return &c.chunks
}

// GetTopics returns the configuration for the Topics service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTopics() *apiconfig.Config {
	return &c.topics
}

// GetTools returns the configuration for the Tools service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTools() *apiconfig.Config {
	return &c.tools
}

// GetExecutions returns the configuration for the Executions service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetExecutions() *apiconfig.Config {
	return &c.executions
}

// GetAnalysisItems returns the configuration for the AnalysisItems service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAnalysisItems() *apiconfig.Config {
	return &c.analysisItems
}

// GetDependentAgents returns the configuration for the DependentAgents service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDependentAgents() *apiconfig.Config {
	return &c.dependentAgents
}

// GetSecrets returns the configuration for the Secrets service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSecrets() *apiconfig.Config {
	return &c.secrets
}

// GetBatchCalls returns the configuration for the BatchCalls service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetBatchCalls() *apiconfig.Config {
	return &c.batchCalls
}

// GetSipTrunk returns the configuration for the SipTrunk service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSipTrunk() *apiconfig.Config {
	return &c.sipTrunk
}

// GetMcpServers returns the configuration for the McpServers service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMcpServers() *apiconfig.Config {
	return &c.mcpServers
}

// GetToolApprovals returns the configuration for the ToolApprovals service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetToolApprovals() *apiconfig.Config {
	return &c.toolApprovals
}

// GetToolConfigs returns the configuration for the ToolConfigs service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetToolConfigs() *apiconfig.Config {
	return &c.toolConfigs
}

// GetWhatsappAccounts returns the configuration for the WhatsappAccounts service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWhatsappAccounts() *apiconfig.Config {
	return &c.whatsappAccounts
}

// GetBranches returns the configuration for the Branches service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetBranches() *apiconfig.Config {
	return &c.branches
}

// GetVersions returns the configuration for the Versions service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetVersions() *apiconfig.Config {
	return &c.versions
}

// GetDeployments returns the configuration for the Deployments service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDeployments() *apiconfig.Config {
	return &c.deployments
}

// GetDrafts returns the configuration for the Drafts service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDrafts() *apiconfig.Config {
	return &c.drafts
}

// GetMergeProposals returns the configuration for the MergeProposals service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMergeProposals() *apiconfig.Config {
	return &c.mergeProposals
}

// GetConfig returns the configuration for the Config service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetConfig() *apiconfig.Config {
	return &c.config
}

// GetOutboundCampaigns returns the configuration for the OutboundCampaigns service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetOutboundCampaigns() *apiconfig.Config {
	return &c.outboundCampaigns
}

// GetLeads returns the configuration for the Leads service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLeads() *apiconfig.Config {
	return &c.leads
}

// GetCalls returns the configuration for the Calls service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetCalls() *apiconfig.Config {
	return &c.calls
}

// GetAttention returns the configuration for the Attention service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAttention() *apiconfig.Config {
	return &c.attention
}

// GetSpeechEngine returns the configuration for the SpeechEngine service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSpeechEngine() *apiconfig.Config {
	return &c.speechEngine
}

// GetProcedures returns the configuration for the Procedures service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetProcedures() *apiconfig.Config {
	return &c.procedures
}

// GetAnalysis returns the configuration for the Analysis service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAnalysis() *apiconfig.Config {
	return &c.analysis
}

// GetEnvironmentVariables returns the configuration for the EnvironmentVariables service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetEnvironmentVariables() *apiconfig.Config {
	return &c.environmentVariables
}

// GetEnvironments returns the configuration for the Environments service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetEnvironments() *apiconfig.Config {
	return &c.environments
}

// GetRealtime returns the configuration for the Realtime service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetRealtime() *apiconfig.Config {
	return &c.realtime
}

// GetCompositionPlan returns the configuration for the CompositionPlan service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetCompositionPlan() *apiconfig.Config {
	return &c.compositionPlan
}

// GetFinetunes returns the configuration for the Finetunes service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetFinetunes() *apiconfig.Config {
	return &c.finetunes
}

// GetOrders returns the configuration for the Orders service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetOrders() *apiconfig.Config {
	return &c.orders
}

// GetMedia returns the configuration for the Media service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetMedia() *apiconfig.Config {
	return &c.media
}

// GetItems returns the configuration for the Items service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetItems() *apiconfig.Config {
	return &c.items
}

// GetDeliverables returns the configuration for the Deliverables service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDeliverables() *apiconfig.Config {
	return &c.deliverables
}

// GetLanguages returns the configuration for the Languages service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLanguages() *apiconfig.Config {
	return &c.languages
}

// GetPvc returns the configuration for the Pvc service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPvc() *apiconfig.Config {
	return &c.pvc
}

// GetWaveform returns the configuration for the Waveform service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWaveform() *apiconfig.Config {
	return &c.waveform
}

// GetSpeakers returns the configuration for the Speakers service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSpeakers() *apiconfig.Config {
	return &c.speakers
}

// GetCaptcha returns the configuration for the Captcha service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetCaptcha() *apiconfig.Config {
	return &c.captcha
}

// GetVerification returns the configuration for the Verification service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetVerification() *apiconfig.Config {
	return &c.verification
}

// GetVideo returns the configuration for the Video service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetVideo() *apiconfig.Config {
	return &c.video
}

// GetImage returns the configuration for the Image service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetImage() *apiconfig.Config {
	return &c.image
}

// GetRuns returns the configuration for the Runs service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetRuns() *apiconfig.Config {
	return &c.runs
}

// GetTemplates returns the configuration for the Templates service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplates() *apiconfig.Config {
	return &c.templates
}

// GetAssets returns the configuration for the Assets service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAssets() *apiconfig.Config {
	return &c.assets
}

// GetUsage returns the configuration for the Usage service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetUsage() *apiconfig.Config {
	return &c.usage
}

// GetRequests returns the configuration for the Requests service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetRequests() *apiconfig.Config {
	return &c.requests
}

// GetServices returns the configuration for the Services service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetServices() *apiconfig.Config {
	return &c.services
}

// GetProducts returns the configuration for the Products service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetProducts() *apiconfig.Config {
	return &c.products
}

// GetStock returns the configuration for the Stock service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetStock() *apiconfig.Config {
	return &c.stock
}

// GetStaff returns the configuration for the Staff service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetStaff() *apiconfig.Config {
	return &c.staff
}

// GetReceptionists returns the configuration for the Receptionists service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetReceptionists() *apiconfig.Config {
	return &c.receptionists
}

// GetTransferRules returns the configuration for the TransferRules service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTransferRules() *apiconfig.Config {
	return &c.transferRules
}

// GetHolidays returns the configuration for the Holidays service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetHolidays() *apiconfig.Config {
	return &c.holidays
}

// GetLocations returns the configuration for the Locations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetLocations() *apiconfig.Config {
	return &c.locations
}

// GetBusinessProfile returns the configuration for the BusinessProfile service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetBusinessProfile() *apiconfig.Config {
	return &c.businessProfile
}

// GetKnowledgeSources returns the configuration for the KnowledgeSources service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetKnowledgeSources() *apiconfig.Config {
	return &c.knowledgeSources
}

// GetWorkspaces returns the configuration for the Workspaces service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWorkspaces() *apiconfig.Config {
	return &c.workspaces
}

// GetSubscriptions returns the configuration for the Subscriptions service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSubscriptions() *apiconfig.Config {
	return &c.subscriptions
}

// GetTiers returns the configuration for the Tiers service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTiers() *apiconfig.Config {
	return &c.tiers
}

// GetBilling returns the configuration for the Billing service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetBilling() *apiconfig.Config {
	return &c.billing
}

// GetTextToDialogueMultiContext returns the configuration for the TextToDialogueMultiContext service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTextToDialogueMultiContext() *apiconfig.Config {
	return &c.textToDialogueMultiContext
}

// GetTextToSpeechMultiContext returns the configuration for the TextToSpeechMultiContext service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTextToSpeechMultiContext() *apiconfig.Config {
	return &c.textToSpeechMultiContext
}

// GetTranslation returns the configuration for the Translation service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTranslation() *apiconfig.Config {
	return &c.translation
}

// GetSubtitling returns the configuration for the Subtitling service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSubtitling() *apiconfig.Config {
	return &c.subtitling
}

// GetBaseURL returns the currently configured base URL.
// All services share the same base URL; this reads it from the first service's config.
func (c *ConfigManager) GetBaseURL() string {
	return c.history.BaseURL
}
