package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWorkspaceGroupPermissionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant WorkspaceGroupPermission
		wire    string
	}
	cases := []enumCase{
		{variant: WorkspaceGroupPermissionTextToSpeech, wire: "\"text_to_speech\""},
		{variant: WorkspaceGroupPermissionSpeechToSpeech, wire: "\"speech_to_speech\""},
		{variant: WorkspaceGroupPermissionSpeechToText, wire: "\"speech_to_text\""},
		{variant: WorkspaceGroupPermissionVoiceLab, wire: "\"voice_lab\""},
		{variant: WorkspaceGroupPermissionSoundEffects, wire: "\"sound_effects\""},
		{variant: WorkspaceGroupPermissionProjects, wire: "\"projects\""},
		{variant: WorkspaceGroupPermissionVoiceoverStudio, wire: "\"voiceover_studio\""},
		{variant: WorkspaceGroupPermissionDubbing, wire: "\"dubbing\""},
		{variant: WorkspaceGroupPermissionAudioNative, wire: "\"audio_native\""},
		{variant: WorkspaceGroupPermissionConversationalAi, wire: "\"conversational_ai\""},
		{variant: WorkspaceGroupPermissionConversationalAiRead, wire: "\"conversational_ai_read\""},
		{variant: WorkspaceGroupPermissionArchitect, wire: "\"architect\""},
		{variant: WorkspaceGroupPermissionVoiceIsolator, wire: "\"voice_isolator\""},
		{variant: WorkspaceGroupPermissionAiSpeechClassifier, wire: "\"ai_speech_classifier\""},
		{variant: WorkspaceGroupPermissionSynthidDetector, wire: "\"synthid_detector\""},
		{variant: WorkspaceGroupPermissionAddVoiceFromVoiceLibrary, wire: "\"add_voice_from_voice_library\""},
		{variant: WorkspaceGroupPermissionCreateInstantVoiceClone, wire: "\"create_instant_voice_clone\""},
		{variant: WorkspaceGroupPermissionCreateProfessionalVoiceClone, wire: "\"create_professional_voice_clone\""},
		{variant: WorkspaceGroupPermissionCreateUserApiKey, wire: "\"create_user_api_key\""},
		{variant: WorkspaceGroupPermissionPublishStudioProject, wire: "\"publish_studio_project\""},
		{variant: WorkspaceGroupPermissionMusic, wire: "\"music\""},
		{variant: WorkspaceGroupPermissionImageVideoGeneration, wire: "\"image_video_generation\""},
		{variant: WorkspaceGroupPermissionFlows, wire: "\"flows\""},
		{variant: WorkspaceGroupPermissionTemplates, wire: "\"templates\""},
		{variant: WorkspaceGroupPermissionShareVoiceExternally, wire: "\"share_voice_externally\""},
		{variant: WorkspaceGroupPermissionPublishVoiceToVoiceLibrary, wire: "\"publish_voice_to_voice_library\""},
		{variant: WorkspaceGroupPermissionViewFiatBalance, wire: "\"view_fiat_balance\""},
		{variant: WorkspaceGroupPermissionWorkspaceAnalyticsFullRead, wire: "\"workspace_analytics_full_read\""},
		{variant: WorkspaceGroupPermissionServiceAccountsManage, wire: "\"service_accounts_manage\""},
		{variant: WorkspaceGroupPermissionWebhooksManage, wire: "\"webhooks_manage\""},
		{variant: WorkspaceGroupPermissionGroupMembersManage, wire: "\"group_members_manage\""},
		{variant: WorkspaceGroupPermissionWorkspaceMembersInvite, wire: "\"workspace_members_invite\""},
		{variant: WorkspaceGroupPermissionWorkspaceMembersRemove, wire: "\"workspace_members_remove\""},
		{variant: WorkspaceGroupPermissionTermsOfServiceAccept, wire: "\"terms_of_service_accept\""},
		{variant: WorkspaceGroupPermissionAuditLogRead, wire: "\"audit_log_read\""},
		{variant: WorkspaceGroupPermissionConversationPrivacyManage, wire: "\"conversation_privacy_manage\""},
		{variant: WorkspaceGroupPermissionAllResourcesAccess, wire: "\"all_resources_access\""},
		{variant: WorkspaceGroupPermissionCopyResourcesCrossWorkspace, wire: "\"copy_resources_cross_workspace\""},
		{variant: WorkspaceGroupPermissionVoiceDesign, wire: "\"voice_design\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded WorkspaceGroupPermission
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
