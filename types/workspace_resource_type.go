package types

import (
	"fmt"
)

// Resource types that can be shared in the workspace. The name always need to match the collection names
type WorkspaceResourceType string

const (
	WorkspaceResourceTypeVoice                                  WorkspaceResourceType = "voice"
	WorkspaceResourceTypeVoiceCollection                        WorkspaceResourceType = "voice_collection"
	WorkspaceResourceTypePronunciationDictionary                WorkspaceResourceType = "pronunciation_dictionary"
	WorkspaceResourceTypeDubbing                                WorkspaceResourceType = "dubbing"
	WorkspaceResourceTypeDubbingProject                         WorkspaceResourceType = "dubbing_project"
	WorkspaceResourceTypeProject                                WorkspaceResourceType = "project"
	WorkspaceResourceTypeConvaiAgents                           WorkspaceResourceType = "convai_agents"
	WorkspaceResourceTypeConvaiKnowledgeBaseDocuments           WorkspaceResourceType = "convai_knowledge_base_documents"
	WorkspaceResourceTypeConvaiTools                            WorkspaceResourceType = "convai_tools"
	WorkspaceResourceTypeConvaiSettings                         WorkspaceResourceType = "convai_settings"
	WorkspaceResourceTypeConvaiSecrets                          WorkspaceResourceType = "convai_secrets"
	WorkspaceResourceTypeWorkspaceAuthConnections               WorkspaceResourceType = "workspace_auth_connections"
	WorkspaceResourceTypeConvaiPhoneNumbers                     WorkspaceResourceType = "convai_phone_numbers"
	WorkspaceResourceTypeConvaiMcpServers                       WorkspaceResourceType = "convai_mcp_servers"
	WorkspaceResourceTypeConvaiApiIntegrationConnections        WorkspaceResourceType = "convai_api_integration_connections"
	WorkspaceResourceTypeConvaiApiIntegrationTriggerConnections WorkspaceResourceType = "convai_api_integration_trigger_connections"
	WorkspaceResourceTypeConvaiBatchCalls                       WorkspaceResourceType = "convai_batch_calls"
	WorkspaceResourceTypeConvaiAgentResponseTests               WorkspaceResourceType = "convai_agent_response_tests"
	WorkspaceResourceTypeConvaiTestSuiteInvocations             WorkspaceResourceType = "convai_test_suite_invocations"
	WorkspaceResourceTypeConvaiCrawlJobs                        WorkspaceResourceType = "convai_crawl_jobs"
	WorkspaceResourceTypeConvaiCrawlTasks                       WorkspaceResourceType = "convai_crawl_tasks"
	WorkspaceResourceTypeConvaiKbExternalSyncJobs               WorkspaceResourceType = "convai_kb_external_sync_jobs"
	WorkspaceResourceTypeConvaiWhatsappAccounts                 WorkspaceResourceType = "convai_whatsapp_accounts"
	WorkspaceResourceTypeConvaiAgentVersions                    WorkspaceResourceType = "convai_agent_versions"
	WorkspaceResourceTypeConvaiAgentBranches                    WorkspaceResourceType = "convai_agent_branches"
	WorkspaceResourceTypeConvaiAgentVersionsDeployments         WorkspaceResourceType = "convai_agent_versions_deployments"
	WorkspaceResourceTypeConvaiAgentExperiments                 WorkspaceResourceType = "convai_agent_experiments"
	WorkspaceResourceTypeConvaiMemoryEntries                    WorkspaceResourceType = "convai_memory_entries"
	WorkspaceResourceTypeConvaiTemplates                        WorkspaceResourceType = "convai_templates"
	WorkspaceResourceTypeDashboard                              WorkspaceResourceType = "dashboard"
	WorkspaceResourceTypeDashboardConfiguration                 WorkspaceResourceType = "dashboard_configuration"
	WorkspaceResourceTypeConvaiAgentDrafts                      WorkspaceResourceType = "convai_agent_drafts"
	WorkspaceResourceTypeResourceLocators                       WorkspaceResourceType = "resource_locators"
	WorkspaceResourceTypeAssets                                 WorkspaceResourceType = "assets"
	WorkspaceResourceTypeContentGenerations                     WorkspaceResourceType = "content_generations"
	WorkspaceResourceTypeContentTemplates                       WorkspaceResourceType = "content_templates"
	WorkspaceResourceTypeContentSkills                          WorkspaceResourceType = "content_skills"
	WorkspaceResourceTypeSongs                                  WorkspaceResourceType = "songs"
	WorkspaceResourceTypeTranscriptionTasks                     WorkspaceResourceType = "transcription_tasks"
	WorkspaceResourceTypeAvatars                                WorkspaceResourceType = "avatars"
	WorkspaceResourceTypeAvatarVideoGenerations                 WorkspaceResourceType = "avatar_video_generations"
	WorkspaceResourceTypeResourceCollection                     WorkspaceResourceType = "resource_collection"
	WorkspaceResourceTypeStudioProjects                         WorkspaceResourceType = "studio_projects"
	WorkspaceResourceTypeConvaiAnalysisItems                    WorkspaceResourceType = "convai_analysis_items"
)

func NewWorkspaceResourceTypeFromString(s string) (WorkspaceResourceType, error) {
	switch s {
	case "voice":
		return WorkspaceResourceTypeVoice, nil
	case "voice_collection":
		return WorkspaceResourceTypeVoiceCollection, nil
	case "pronunciation_dictionary":
		return WorkspaceResourceTypePronunciationDictionary, nil
	case "dubbing":
		return WorkspaceResourceTypeDubbing, nil
	case "dubbing_project":
		return WorkspaceResourceTypeDubbingProject, nil
	case "project":
		return WorkspaceResourceTypeProject, nil
	case "convai_agents":
		return WorkspaceResourceTypeConvaiAgents, nil
	case "convai_knowledge_base_documents":
		return WorkspaceResourceTypeConvaiKnowledgeBaseDocuments, nil
	case "convai_tools":
		return WorkspaceResourceTypeConvaiTools, nil
	case "convai_settings":
		return WorkspaceResourceTypeConvaiSettings, nil
	case "convai_secrets":
		return WorkspaceResourceTypeConvaiSecrets, nil
	case "workspace_auth_connections":
		return WorkspaceResourceTypeWorkspaceAuthConnections, nil
	case "convai_phone_numbers":
		return WorkspaceResourceTypeConvaiPhoneNumbers, nil
	case "convai_mcp_servers":
		return WorkspaceResourceTypeConvaiMcpServers, nil
	case "convai_api_integration_connections":
		return WorkspaceResourceTypeConvaiApiIntegrationConnections, nil
	case "convai_api_integration_trigger_connections":
		return WorkspaceResourceTypeConvaiApiIntegrationTriggerConnections, nil
	case "convai_batch_calls":
		return WorkspaceResourceTypeConvaiBatchCalls, nil
	case "convai_agent_response_tests":
		return WorkspaceResourceTypeConvaiAgentResponseTests, nil
	case "convai_test_suite_invocations":
		return WorkspaceResourceTypeConvaiTestSuiteInvocations, nil
	case "convai_crawl_jobs":
		return WorkspaceResourceTypeConvaiCrawlJobs, nil
	case "convai_crawl_tasks":
		return WorkspaceResourceTypeConvaiCrawlTasks, nil
	case "convai_kb_external_sync_jobs":
		return WorkspaceResourceTypeConvaiKbExternalSyncJobs, nil
	case "convai_whatsapp_accounts":
		return WorkspaceResourceTypeConvaiWhatsappAccounts, nil
	case "convai_agent_versions":
		return WorkspaceResourceTypeConvaiAgentVersions, nil
	case "convai_agent_branches":
		return WorkspaceResourceTypeConvaiAgentBranches, nil
	case "convai_agent_versions_deployments":
		return WorkspaceResourceTypeConvaiAgentVersionsDeployments, nil
	case "convai_agent_experiments":
		return WorkspaceResourceTypeConvaiAgentExperiments, nil
	case "convai_memory_entries":
		return WorkspaceResourceTypeConvaiMemoryEntries, nil
	case "convai_templates":
		return WorkspaceResourceTypeConvaiTemplates, nil
	case "dashboard":
		return WorkspaceResourceTypeDashboard, nil
	case "dashboard_configuration":
		return WorkspaceResourceTypeDashboardConfiguration, nil
	case "convai_agent_drafts":
		return WorkspaceResourceTypeConvaiAgentDrafts, nil
	case "resource_locators":
		return WorkspaceResourceTypeResourceLocators, nil
	case "assets":
		return WorkspaceResourceTypeAssets, nil
	case "content_generations":
		return WorkspaceResourceTypeContentGenerations, nil
	case "content_templates":
		return WorkspaceResourceTypeContentTemplates, nil
	case "content_skills":
		return WorkspaceResourceTypeContentSkills, nil
	case "songs":
		return WorkspaceResourceTypeSongs, nil
	case "transcription_tasks":
		return WorkspaceResourceTypeTranscriptionTasks, nil
	case "avatars":
		return WorkspaceResourceTypeAvatars, nil
	case "avatar_video_generations":
		return WorkspaceResourceTypeAvatarVideoGenerations, nil
	case "resource_collection":
		return WorkspaceResourceTypeResourceCollection, nil
	case "studio_projects":
		return WorkspaceResourceTypeStudioProjects, nil
	case "convai_analysis_items":
		return WorkspaceResourceTypeConvaiAnalysisItems, nil
	}
	var t WorkspaceResourceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WorkspaceResourceType) Ptr() *WorkspaceResourceType {
	return &e
}
