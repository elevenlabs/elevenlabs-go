package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWorkspaceResourceTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant WorkspaceResourceType
		wire    string
	}
	cases := []enumCase{
		{variant: WorkspaceResourceTypeVoice, wire: "\"voice\""},
		{variant: WorkspaceResourceTypeVoiceCollection, wire: "\"voice_collection\""},
		{variant: WorkspaceResourceTypePronunciationDictionary, wire: "\"pronunciation_dictionary\""},
		{variant: WorkspaceResourceTypeDubbing, wire: "\"dubbing\""},
		{variant: WorkspaceResourceTypeDubbingProject, wire: "\"dubbing_project\""},
		{variant: WorkspaceResourceTypeProject, wire: "\"project\""},
		{variant: WorkspaceResourceTypeConvaiAgents, wire: "\"convai_agents\""},
		{variant: WorkspaceResourceTypeConvaiKnowledgeBaseDocuments, wire: "\"convai_knowledge_base_documents\""},
		{variant: WorkspaceResourceTypeConvaiTools, wire: "\"convai_tools\""},
		{variant: WorkspaceResourceTypeConvaiSettings, wire: "\"convai_settings\""},
		{variant: WorkspaceResourceTypeConvaiSecrets, wire: "\"convai_secrets\""},
		{variant: WorkspaceResourceTypeWorkspaceAuthConnections, wire: "\"workspace_auth_connections\""},
		{variant: WorkspaceResourceTypeConvaiPhoneNumbers, wire: "\"convai_phone_numbers\""},
		{variant: WorkspaceResourceTypeConvaiMcpServers, wire: "\"convai_mcp_servers\""},
		{variant: WorkspaceResourceTypeConvaiApiIntegrationConnections, wire: "\"convai_api_integration_connections\""},
		{variant: WorkspaceResourceTypeConvaiApiIntegrationTriggerConnections, wire: "\"convai_api_integration_trigger_connections\""},
		{variant: WorkspaceResourceTypeConvaiBatchCalls, wire: "\"convai_batch_calls\""},
		{variant: WorkspaceResourceTypeConvaiAgentResponseTests, wire: "\"convai_agent_response_tests\""},
		{variant: WorkspaceResourceTypeConvaiTestSuiteInvocations, wire: "\"convai_test_suite_invocations\""},
		{variant: WorkspaceResourceTypeConvaiCrawlJobs, wire: "\"convai_crawl_jobs\""},
		{variant: WorkspaceResourceTypeConvaiCrawlTasks, wire: "\"convai_crawl_tasks\""},
		{variant: WorkspaceResourceTypeConvaiKbExternalSyncJobs, wire: "\"convai_kb_external_sync_jobs\""},
		{variant: WorkspaceResourceTypeConvaiWhatsappAccounts, wire: "\"convai_whatsapp_accounts\""},
		{variant: WorkspaceResourceTypeConvaiAgentVersions, wire: "\"convai_agent_versions\""},
		{variant: WorkspaceResourceTypeConvaiAgentBranches, wire: "\"convai_agent_branches\""},
		{variant: WorkspaceResourceTypeConvaiAgentVersionsDeployments, wire: "\"convai_agent_versions_deployments\""},
		{variant: WorkspaceResourceTypeConvaiAgentExperiments, wire: "\"convai_agent_experiments\""},
		{variant: WorkspaceResourceTypeConvaiMemoryEntries, wire: "\"convai_memory_entries\""},
		{variant: WorkspaceResourceTypeConvaiTemplates, wire: "\"convai_templates\""},
		{variant: WorkspaceResourceTypeDashboard, wire: "\"dashboard\""},
		{variant: WorkspaceResourceTypeDashboardConfiguration, wire: "\"dashboard_configuration\""},
		{variant: WorkspaceResourceTypeConvaiAgentDrafts, wire: "\"convai_agent_drafts\""},
		{variant: WorkspaceResourceTypeResourceLocators, wire: "\"resource_locators\""},
		{variant: WorkspaceResourceTypeAssets, wire: "\"assets\""},
		{variant: WorkspaceResourceTypeContentGenerations, wire: "\"content_generations\""},
		{variant: WorkspaceResourceTypeContentTemplates, wire: "\"content_templates\""},
		{variant: WorkspaceResourceTypeContentSkills, wire: "\"content_skills\""},
		{variant: WorkspaceResourceTypeSongs, wire: "\"songs\""},
		{variant: WorkspaceResourceTypeTranscriptionTasks, wire: "\"transcription_tasks\""},
		{variant: WorkspaceResourceTypeAvatars, wire: "\"avatars\""},
		{variant: WorkspaceResourceTypeAvatarVideoGenerations, wire: "\"avatar_video_generations\""},
		{variant: WorkspaceResourceTypeResourceCollection, wire: "\"resource_collection\""},
		{variant: WorkspaceResourceTypeStudioProjects, wire: "\"studio_projects\""},
		{variant: WorkspaceResourceTypeConvaiAnalysisItems, wire: "\"convai_analysis_items\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded WorkspaceResourceType
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
