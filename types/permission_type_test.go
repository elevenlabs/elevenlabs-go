package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPermissionTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant PermissionType
		wire    string
	}
	cases := []enumCase{
		{variant: PermissionTypeTextToSpeech, wire: "\"text_to_speech\""},
		{variant: PermissionTypeSpeechToSpeech, wire: "\"speech_to_speech\""},
		{variant: PermissionTypeSpeechToText, wire: "\"speech_to_text\""},
		{variant: PermissionTypeModelsRead, wire: "\"models_read\""},
		{variant: PermissionTypeModelsWrite, wire: "\"models_write\""},
		{variant: PermissionTypeVoicesRead, wire: "\"voices_read\""},
		{variant: PermissionTypeVoicesWrite, wire: "\"voices_write\""},
		{variant: PermissionTypeVoiceCollectionsRead, wire: "\"voice_collections_read\""},
		{variant: PermissionTypeVoiceCollectionsWrite, wire: "\"voice_collections_write\""},
		{variant: PermissionTypeSpeechHistoryRead, wire: "\"speech_history_read\""},
		{variant: PermissionTypeSpeechHistoryWrite, wire: "\"speech_history_write\""},
		{variant: PermissionTypeSoundGeneration, wire: "\"sound_generation\""},
		{variant: PermissionTypeAudioIsolation, wire: "\"audio_isolation\""},
		{variant: PermissionTypeVoiceGeneration, wire: "\"voice_generation\""},
		{variant: PermissionTypeDubbingRead, wire: "\"dubbing_read\""},
		{variant: PermissionTypeDubbingWrite, wire: "\"dubbing_write\""},
		{variant: PermissionTypePronunciationDictionariesRead, wire: "\"pronunciation_dictionaries_read\""},
		{variant: PermissionTypePronunciationDictionariesWrite, wire: "\"pronunciation_dictionaries_write\""},
		{variant: PermissionTypeUserRead, wire: "\"user_read\""},
		{variant: PermissionTypeUserWrite, wire: "\"user_write\""},
		{variant: PermissionTypeProjectsRead, wire: "\"projects_read\""},
		{variant: PermissionTypeProjectsWrite, wire: "\"projects_write\""},
		{variant: PermissionTypeAudioNativeRead, wire: "\"audio_native_read\""},
		{variant: PermissionTypeAudioNativeWrite, wire: "\"audio_native_write\""},
		{variant: PermissionTypeWorkspaceRead, wire: "\"workspace_read\""},
		{variant: PermissionTypeWorkspaceWrite, wire: "\"workspace_write\""},
		{variant: PermissionTypeForcedAlignment, wire: "\"forced_alignment\""},
		{variant: PermissionTypeConvaiRead, wire: "\"convai_read\""},
		{variant: PermissionTypeConvaiWrite, wire: "\"convai_write\""},
		{variant: PermissionTypeAgentsArchitect, wire: "\"agents_architect\""},
		{variant: PermissionTypeAgentsEnvironmentVariablesWrite, wire: "\"agents_environment_variables_write\""},
		{variant: PermissionTypeMusicGeneration, wire: "\"music_generation\""},
		{variant: PermissionTypeImageVideoGeneration, wire: "\"image_video_generation\""},
		{variant: PermissionTypeFlows, wire: "\"flows\""},
		{variant: PermissionTypeTemplates, wire: "\"templates\""},
		{variant: PermissionTypeAddVoiceFromVoiceLibrary, wire: "\"add_voice_from_voice_library\""},
		{variant: PermissionTypeCreateInstantVoiceClone, wire: "\"create_instant_voice_clone\""},
		{variant: PermissionTypeCreateProfessionalVoiceClone, wire: "\"create_professional_voice_clone\""},
		{variant: PermissionTypePublishVoiceToVoiceLibrary, wire: "\"publish_voice_to_voice_library\""},
		{variant: PermissionTypeShareVoiceExternally, wire: "\"share_voice_externally\""},
		{variant: PermissionTypeCreateUserApiKey, wire: "\"create_user_api_key\""},
		{variant: PermissionTypeWorkspaceAnalyticsFullRead, wire: "\"workspace_analytics_full_read\""},
		{variant: PermissionTypeWebhooksWrite, wire: "\"webhooks_write\""},
		{variant: PermissionTypeServiceAccountWrite, wire: "\"service_account_write\""},
		{variant: PermissionTypeGroupMembersManage, wire: "\"group_members_manage\""},
		{variant: PermissionTypeWorkspaceMembersRead, wire: "\"workspace_members_read\""},
		{variant: PermissionTypeWorkspaceMembersInvite, wire: "\"workspace_members_invite\""},
		{variant: PermissionTypeWorkspaceMembersRemove, wire: "\"workspace_members_remove\""},
		{variant: PermissionTypeTermsOfServiceAccept, wire: "\"terms_of_service_accept\""},
		{variant: PermissionTypeAuditLogRead, wire: "\"audit_log_read\""},
		{variant: PermissionTypeConversationPrivacyManage, wire: "\"conversation_privacy_manage\""},
		{variant: PermissionTypeCopyResourcesCrossWorkspace, wire: "\"copy_resources_cross_workspace\""},
		{variant: PermissionTypeSynthidDetector, wire: "\"synthid_detector\""},
		{variant: PermissionTypeWatermark, wire: "\"watermark\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded PermissionType
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
