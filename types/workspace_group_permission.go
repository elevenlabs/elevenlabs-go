package types

import (
	"fmt"
)

type WorkspaceGroupPermission string

const (
	WorkspaceGroupPermissionTextToSpeech                 WorkspaceGroupPermission = "text_to_speech"
	WorkspaceGroupPermissionSpeechToSpeech               WorkspaceGroupPermission = "speech_to_speech"
	WorkspaceGroupPermissionSpeechToText                 WorkspaceGroupPermission = "speech_to_text"
	WorkspaceGroupPermissionVoiceLab                     WorkspaceGroupPermission = "voice_lab"
	WorkspaceGroupPermissionSoundEffects                 WorkspaceGroupPermission = "sound_effects"
	WorkspaceGroupPermissionProjects                     WorkspaceGroupPermission = "projects"
	WorkspaceGroupPermissionVoiceoverStudio              WorkspaceGroupPermission = "voiceover_studio"
	WorkspaceGroupPermissionDubbing                      WorkspaceGroupPermission = "dubbing"
	WorkspaceGroupPermissionAudioNative                  WorkspaceGroupPermission = "audio_native"
	WorkspaceGroupPermissionConversationalAi             WorkspaceGroupPermission = "conversational_ai"
	WorkspaceGroupPermissionConversationalAiRead         WorkspaceGroupPermission = "conversational_ai_read"
	WorkspaceGroupPermissionArchitect                    WorkspaceGroupPermission = "architect"
	WorkspaceGroupPermissionVoiceIsolator                WorkspaceGroupPermission = "voice_isolator"
	WorkspaceGroupPermissionAiSpeechClassifier           WorkspaceGroupPermission = "ai_speech_classifier"
	WorkspaceGroupPermissionSynthidDetector              WorkspaceGroupPermission = "synthid_detector"
	WorkspaceGroupPermissionAddVoiceFromVoiceLibrary     WorkspaceGroupPermission = "add_voice_from_voice_library"
	WorkspaceGroupPermissionCreateInstantVoiceClone      WorkspaceGroupPermission = "create_instant_voice_clone"
	WorkspaceGroupPermissionCreateProfessionalVoiceClone WorkspaceGroupPermission = "create_professional_voice_clone"
	WorkspaceGroupPermissionCreateUserApiKey             WorkspaceGroupPermission = "create_user_api_key"
	WorkspaceGroupPermissionPublishStudioProject         WorkspaceGroupPermission = "publish_studio_project"
	WorkspaceGroupPermissionMusic                        WorkspaceGroupPermission = "music"
	WorkspaceGroupPermissionImageVideoGeneration         WorkspaceGroupPermission = "image_video_generation"
	WorkspaceGroupPermissionFlows                        WorkspaceGroupPermission = "flows"
	WorkspaceGroupPermissionTemplates                    WorkspaceGroupPermission = "templates"
	WorkspaceGroupPermissionShareVoiceExternally         WorkspaceGroupPermission = "share_voice_externally"
	WorkspaceGroupPermissionPublishVoiceToVoiceLibrary   WorkspaceGroupPermission = "publish_voice_to_voice_library"
	WorkspaceGroupPermissionViewFiatBalance              WorkspaceGroupPermission = "view_fiat_balance"
	WorkspaceGroupPermissionWorkspaceAnalyticsFullRead   WorkspaceGroupPermission = "workspace_analytics_full_read"
	WorkspaceGroupPermissionServiceAccountsManage        WorkspaceGroupPermission = "service_accounts_manage"
	WorkspaceGroupPermissionWebhooksManage               WorkspaceGroupPermission = "webhooks_manage"
	WorkspaceGroupPermissionGroupMembersManage           WorkspaceGroupPermission = "group_members_manage"
	WorkspaceGroupPermissionWorkspaceMembersInvite       WorkspaceGroupPermission = "workspace_members_invite"
	WorkspaceGroupPermissionWorkspaceMembersRemove       WorkspaceGroupPermission = "workspace_members_remove"
	WorkspaceGroupPermissionTermsOfServiceAccept         WorkspaceGroupPermission = "terms_of_service_accept"
	WorkspaceGroupPermissionAuditLogRead                 WorkspaceGroupPermission = "audit_log_read"
	WorkspaceGroupPermissionConversationPrivacyManage    WorkspaceGroupPermission = "conversation_privacy_manage"
	WorkspaceGroupPermissionAllResourcesAccess           WorkspaceGroupPermission = "all_resources_access"
	WorkspaceGroupPermissionCopyResourcesCrossWorkspace  WorkspaceGroupPermission = "copy_resources_cross_workspace"
	WorkspaceGroupPermissionVoiceDesign                  WorkspaceGroupPermission = "voice_design"
)

func NewWorkspaceGroupPermissionFromString(s string) (WorkspaceGroupPermission, error) {
	switch s {
	case "text_to_speech":
		return WorkspaceGroupPermissionTextToSpeech, nil
	case "speech_to_speech":
		return WorkspaceGroupPermissionSpeechToSpeech, nil
	case "speech_to_text":
		return WorkspaceGroupPermissionSpeechToText, nil
	case "voice_lab":
		return WorkspaceGroupPermissionVoiceLab, nil
	case "sound_effects":
		return WorkspaceGroupPermissionSoundEffects, nil
	case "projects":
		return WorkspaceGroupPermissionProjects, nil
	case "voiceover_studio":
		return WorkspaceGroupPermissionVoiceoverStudio, nil
	case "dubbing":
		return WorkspaceGroupPermissionDubbing, nil
	case "audio_native":
		return WorkspaceGroupPermissionAudioNative, nil
	case "conversational_ai":
		return WorkspaceGroupPermissionConversationalAi, nil
	case "conversational_ai_read":
		return WorkspaceGroupPermissionConversationalAiRead, nil
	case "architect":
		return WorkspaceGroupPermissionArchitect, nil
	case "voice_isolator":
		return WorkspaceGroupPermissionVoiceIsolator, nil
	case "ai_speech_classifier":
		return WorkspaceGroupPermissionAiSpeechClassifier, nil
	case "synthid_detector":
		return WorkspaceGroupPermissionSynthidDetector, nil
	case "add_voice_from_voice_library":
		return WorkspaceGroupPermissionAddVoiceFromVoiceLibrary, nil
	case "create_instant_voice_clone":
		return WorkspaceGroupPermissionCreateInstantVoiceClone, nil
	case "create_professional_voice_clone":
		return WorkspaceGroupPermissionCreateProfessionalVoiceClone, nil
	case "create_user_api_key":
		return WorkspaceGroupPermissionCreateUserApiKey, nil
	case "publish_studio_project":
		return WorkspaceGroupPermissionPublishStudioProject, nil
	case "music":
		return WorkspaceGroupPermissionMusic, nil
	case "image_video_generation":
		return WorkspaceGroupPermissionImageVideoGeneration, nil
	case "flows":
		return WorkspaceGroupPermissionFlows, nil
	case "templates":
		return WorkspaceGroupPermissionTemplates, nil
	case "share_voice_externally":
		return WorkspaceGroupPermissionShareVoiceExternally, nil
	case "publish_voice_to_voice_library":
		return WorkspaceGroupPermissionPublishVoiceToVoiceLibrary, nil
	case "view_fiat_balance":
		return WorkspaceGroupPermissionViewFiatBalance, nil
	case "workspace_analytics_full_read":
		return WorkspaceGroupPermissionWorkspaceAnalyticsFullRead, nil
	case "service_accounts_manage":
		return WorkspaceGroupPermissionServiceAccountsManage, nil
	case "webhooks_manage":
		return WorkspaceGroupPermissionWebhooksManage, nil
	case "group_members_manage":
		return WorkspaceGroupPermissionGroupMembersManage, nil
	case "workspace_members_invite":
		return WorkspaceGroupPermissionWorkspaceMembersInvite, nil
	case "workspace_members_remove":
		return WorkspaceGroupPermissionWorkspaceMembersRemove, nil
	case "terms_of_service_accept":
		return WorkspaceGroupPermissionTermsOfServiceAccept, nil
	case "audit_log_read":
		return WorkspaceGroupPermissionAuditLogRead, nil
	case "conversation_privacy_manage":
		return WorkspaceGroupPermissionConversationPrivacyManage, nil
	case "all_resources_access":
		return WorkspaceGroupPermissionAllResourcesAccess, nil
	case "copy_resources_cross_workspace":
		return WorkspaceGroupPermissionCopyResourcesCrossWorkspace, nil
	case "voice_design":
		return WorkspaceGroupPermissionVoiceDesign, nil
	}
	var t WorkspaceGroupPermission
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WorkspaceGroupPermission) Ptr() *WorkspaceGroupPermission {
	return &e
}
