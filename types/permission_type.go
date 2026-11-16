package types

import (
	"fmt"
)

type PermissionType string

const (
	PermissionTypeTextToSpeech                    PermissionType = "text_to_speech"
	PermissionTypeSpeechToSpeech                  PermissionType = "speech_to_speech"
	PermissionTypeSpeechToText                    PermissionType = "speech_to_text"
	PermissionTypeModelsRead                      PermissionType = "models_read"
	PermissionTypeModelsWrite                     PermissionType = "models_write"
	PermissionTypeVoicesRead                      PermissionType = "voices_read"
	PermissionTypeVoicesWrite                     PermissionType = "voices_write"
	PermissionTypeVoiceCollectionsRead            PermissionType = "voice_collections_read"
	PermissionTypeVoiceCollectionsWrite           PermissionType = "voice_collections_write"
	PermissionTypeSpeechHistoryRead               PermissionType = "speech_history_read"
	PermissionTypeSpeechHistoryWrite              PermissionType = "speech_history_write"
	PermissionTypeSoundGeneration                 PermissionType = "sound_generation"
	PermissionTypeAudioIsolation                  PermissionType = "audio_isolation"
	PermissionTypeVoiceGeneration                 PermissionType = "voice_generation"
	PermissionTypeDubbingRead                     PermissionType = "dubbing_read"
	PermissionTypeDubbingWrite                    PermissionType = "dubbing_write"
	PermissionTypePronunciationDictionariesRead   PermissionType = "pronunciation_dictionaries_read"
	PermissionTypePronunciationDictionariesWrite  PermissionType = "pronunciation_dictionaries_write"
	PermissionTypeUserRead                        PermissionType = "user_read"
	PermissionTypeUserWrite                       PermissionType = "user_write"
	PermissionTypeProjectsRead                    PermissionType = "projects_read"
	PermissionTypeProjectsWrite                   PermissionType = "projects_write"
	PermissionTypeAudioNativeRead                 PermissionType = "audio_native_read"
	PermissionTypeAudioNativeWrite                PermissionType = "audio_native_write"
	PermissionTypeWorkspaceRead                   PermissionType = "workspace_read"
	PermissionTypeWorkspaceWrite                  PermissionType = "workspace_write"
	PermissionTypeForcedAlignment                 PermissionType = "forced_alignment"
	PermissionTypeConvaiRead                      PermissionType = "convai_read"
	PermissionTypeConvaiWrite                     PermissionType = "convai_write"
	PermissionTypeAgentsArchitect                 PermissionType = "agents_architect"
	PermissionTypeAgentsEnvironmentVariablesWrite PermissionType = "agents_environment_variables_write"
	PermissionTypeMusicGeneration                 PermissionType = "music_generation"
	PermissionTypeImageVideoGeneration            PermissionType = "image_video_generation"
	PermissionTypeFlows                           PermissionType = "flows"
	PermissionTypeTemplates                       PermissionType = "templates"
	PermissionTypeAddVoiceFromVoiceLibrary        PermissionType = "add_voice_from_voice_library"
	PermissionTypeCreateInstantVoiceClone         PermissionType = "create_instant_voice_clone"
	PermissionTypeCreateProfessionalVoiceClone    PermissionType = "create_professional_voice_clone"
	PermissionTypePublishVoiceToVoiceLibrary      PermissionType = "publish_voice_to_voice_library"
	PermissionTypeShareVoiceExternally            PermissionType = "share_voice_externally"
	PermissionTypeCreateUserApiKey                PermissionType = "create_user_api_key"
	PermissionTypeWorkspaceAnalyticsFullRead      PermissionType = "workspace_analytics_full_read"
	PermissionTypeWebhooksWrite                   PermissionType = "webhooks_write"
	PermissionTypeServiceAccountWrite             PermissionType = "service_account_write"
	PermissionTypeGroupMembersManage              PermissionType = "group_members_manage"
	PermissionTypeWorkspaceMembersRead            PermissionType = "workspace_members_read"
	PermissionTypeWorkspaceMembersInvite          PermissionType = "workspace_members_invite"
	PermissionTypeWorkspaceMembersRemove          PermissionType = "workspace_members_remove"
	PermissionTypeTermsOfServiceAccept            PermissionType = "terms_of_service_accept"
	PermissionTypeAuditLogRead                    PermissionType = "audit_log_read"
	PermissionTypeConversationPrivacyManage       PermissionType = "conversation_privacy_manage"
	PermissionTypeCopyResourcesCrossWorkspace     PermissionType = "copy_resources_cross_workspace"
	PermissionTypeSynthidDetector                 PermissionType = "synthid_detector"
	PermissionTypeWatermark                       PermissionType = "watermark"
)

func NewPermissionTypeFromString(s string) (PermissionType, error) {
	switch s {
	case "text_to_speech":
		return PermissionTypeTextToSpeech, nil
	case "speech_to_speech":
		return PermissionTypeSpeechToSpeech, nil
	case "speech_to_text":
		return PermissionTypeSpeechToText, nil
	case "models_read":
		return PermissionTypeModelsRead, nil
	case "models_write":
		return PermissionTypeModelsWrite, nil
	case "voices_read":
		return PermissionTypeVoicesRead, nil
	case "voices_write":
		return PermissionTypeVoicesWrite, nil
	case "voice_collections_read":
		return PermissionTypeVoiceCollectionsRead, nil
	case "voice_collections_write":
		return PermissionTypeVoiceCollectionsWrite, nil
	case "speech_history_read":
		return PermissionTypeSpeechHistoryRead, nil
	case "speech_history_write":
		return PermissionTypeSpeechHistoryWrite, nil
	case "sound_generation":
		return PermissionTypeSoundGeneration, nil
	case "audio_isolation":
		return PermissionTypeAudioIsolation, nil
	case "voice_generation":
		return PermissionTypeVoiceGeneration, nil
	case "dubbing_read":
		return PermissionTypeDubbingRead, nil
	case "dubbing_write":
		return PermissionTypeDubbingWrite, nil
	case "pronunciation_dictionaries_read":
		return PermissionTypePronunciationDictionariesRead, nil
	case "pronunciation_dictionaries_write":
		return PermissionTypePronunciationDictionariesWrite, nil
	case "user_read":
		return PermissionTypeUserRead, nil
	case "user_write":
		return PermissionTypeUserWrite, nil
	case "projects_read":
		return PermissionTypeProjectsRead, nil
	case "projects_write":
		return PermissionTypeProjectsWrite, nil
	case "audio_native_read":
		return PermissionTypeAudioNativeRead, nil
	case "audio_native_write":
		return PermissionTypeAudioNativeWrite, nil
	case "workspace_read":
		return PermissionTypeWorkspaceRead, nil
	case "workspace_write":
		return PermissionTypeWorkspaceWrite, nil
	case "forced_alignment":
		return PermissionTypeForcedAlignment, nil
	case "convai_read":
		return PermissionTypeConvaiRead, nil
	case "convai_write":
		return PermissionTypeConvaiWrite, nil
	case "agents_architect":
		return PermissionTypeAgentsArchitect, nil
	case "agents_environment_variables_write":
		return PermissionTypeAgentsEnvironmentVariablesWrite, nil
	case "music_generation":
		return PermissionTypeMusicGeneration, nil
	case "image_video_generation":
		return PermissionTypeImageVideoGeneration, nil
	case "flows":
		return PermissionTypeFlows, nil
	case "templates":
		return PermissionTypeTemplates, nil
	case "add_voice_from_voice_library":
		return PermissionTypeAddVoiceFromVoiceLibrary, nil
	case "create_instant_voice_clone":
		return PermissionTypeCreateInstantVoiceClone, nil
	case "create_professional_voice_clone":
		return PermissionTypeCreateProfessionalVoiceClone, nil
	case "publish_voice_to_voice_library":
		return PermissionTypePublishVoiceToVoiceLibrary, nil
	case "share_voice_externally":
		return PermissionTypeShareVoiceExternally, nil
	case "create_user_api_key":
		return PermissionTypeCreateUserApiKey, nil
	case "workspace_analytics_full_read":
		return PermissionTypeWorkspaceAnalyticsFullRead, nil
	case "webhooks_write":
		return PermissionTypeWebhooksWrite, nil
	case "service_account_write":
		return PermissionTypeServiceAccountWrite, nil
	case "group_members_manage":
		return PermissionTypeGroupMembersManage, nil
	case "workspace_members_read":
		return PermissionTypeWorkspaceMembersRead, nil
	case "workspace_members_invite":
		return PermissionTypeWorkspaceMembersInvite, nil
	case "workspace_members_remove":
		return PermissionTypeWorkspaceMembersRemove, nil
	case "terms_of_service_accept":
		return PermissionTypeTermsOfServiceAccept, nil
	case "audit_log_read":
		return PermissionTypeAuditLogRead, nil
	case "conversation_privacy_manage":
		return PermissionTypeConversationPrivacyManage, nil
	case "copy_resources_cross_workspace":
		return PermissionTypeCopyResourcesCrossWorkspace, nil
	case "synthid_detector":
		return PermissionTypeSynthidDetector, nil
	case "watermark":
		return PermissionTypeWatermark, nil
	}
	var t PermissionType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PermissionType) Ptr() *PermissionType {
	return &e
}
