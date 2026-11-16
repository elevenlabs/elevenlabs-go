package types

import (
	"fmt"
)

type WebhookTypes string

const (
	WebhookTypesProjectConversionStatus             WebhookTypes = "project_conversion_status"
	WebhookTypesChapterConversionStatus             WebhookTypes = "chapter_conversion_status"
	WebhookTypesPostCallTranscription               WebhookTypes = "post_call_transcription"
	WebhookTypesPostCallTranscriptionOtel           WebhookTypes = "post_call_transcription_otel"
	WebhookTypesPostCallAudio                       WebhookTypes = "post_call_audio"
	WebhookTypesPostCallUnredactedTranscription     WebhookTypes = "post_call_unredacted_transcription"
	WebhookTypesPostCallUnredactedTranscriptionOtel WebhookTypes = "post_call_unredacted_transcription_otel"
	WebhookTypesPostCallUnredactedAudio             WebhookTypes = "post_call_unredacted_audio"
	WebhookTypesPostCallAnalysis                    WebhookTypes = "post_call_analysis"
	WebhookTypesCallInitiationFailure               WebhookTypes = "call_initiation_failure"
	WebhookTypesConversationInitiationClientData    WebhookTypes = "conversation_initiation_client_data"
	WebhookTypesAnsweringMachineDetection           WebhookTypes = "answering_machine_detection"
	WebhookTypesVoiceRemovalNotice                  WebhookTypes = "voice_removal_notice"
	WebhookTypesVoiceRemovalNoticeWithdrawn         WebhookTypes = "voice_removal_notice_withdrawn"
	WebhookTypesVoiceRemoved                        WebhookTypes = "voice_removed"
	WebhookTypesSpeechToTextTranscription           WebhookTypes = "speech_to_text_transcription"
	WebhookTypesWebhookError                        WebhookTypes = "webhook_error"
	WebhookTypesSpeechToTextBatch                   WebhookTypes = "speech_to_text_batch"
	WebhookTypesAgentQaEvaluationCompleted          WebhookTypes = "agent_qa_evaluation_completed"
	WebhookTypesConvaiAlertingNotification          WebhookTypes = "convai_alerting_notification"
	WebhookTypesFlowsGeneration                     WebhookTypes = "flows_generation"
	WebhookTypesFlowsTemplateRun                    WebhookTypes = "flows_template_run"
	WebhookTypesDubbingProjectReady                 WebhookTypes = "dubbing_project_ready"
	WebhookTypesDubbingProjectFailed                WebhookTypes = "dubbing_project_failed"
	WebhookTypesDubbingLanguageCompleted            WebhookTypes = "dubbing_language_completed"
	WebhookTypesDubbingLanguageFailed               WebhookTypes = "dubbing_language_failed"
	WebhookTypesSupportTicketMessageAdded           WebhookTypes = "support_ticket_message_added"
)

func NewWebhookTypesFromString(s string) (WebhookTypes, error) {
	switch s {
	case "project_conversion_status":
		return WebhookTypesProjectConversionStatus, nil
	case "chapter_conversion_status":
		return WebhookTypesChapterConversionStatus, nil
	case "post_call_transcription":
		return WebhookTypesPostCallTranscription, nil
	case "post_call_transcription_otel":
		return WebhookTypesPostCallTranscriptionOtel, nil
	case "post_call_audio":
		return WebhookTypesPostCallAudio, nil
	case "post_call_unredacted_transcription":
		return WebhookTypesPostCallUnredactedTranscription, nil
	case "post_call_unredacted_transcription_otel":
		return WebhookTypesPostCallUnredactedTranscriptionOtel, nil
	case "post_call_unredacted_audio":
		return WebhookTypesPostCallUnredactedAudio, nil
	case "post_call_analysis":
		return WebhookTypesPostCallAnalysis, nil
	case "call_initiation_failure":
		return WebhookTypesCallInitiationFailure, nil
	case "conversation_initiation_client_data":
		return WebhookTypesConversationInitiationClientData, nil
	case "answering_machine_detection":
		return WebhookTypesAnsweringMachineDetection, nil
	case "voice_removal_notice":
		return WebhookTypesVoiceRemovalNotice, nil
	case "voice_removal_notice_withdrawn":
		return WebhookTypesVoiceRemovalNoticeWithdrawn, nil
	case "voice_removed":
		return WebhookTypesVoiceRemoved, nil
	case "speech_to_text_transcription":
		return WebhookTypesSpeechToTextTranscription, nil
	case "webhook_error":
		return WebhookTypesWebhookError, nil
	case "speech_to_text_batch":
		return WebhookTypesSpeechToTextBatch, nil
	case "agent_qa_evaluation_completed":
		return WebhookTypesAgentQaEvaluationCompleted, nil
	case "convai_alerting_notification":
		return WebhookTypesConvaiAlertingNotification, nil
	case "flows_generation":
		return WebhookTypesFlowsGeneration, nil
	case "flows_template_run":
		return WebhookTypesFlowsTemplateRun, nil
	case "dubbing_project_ready":
		return WebhookTypesDubbingProjectReady, nil
	case "dubbing_project_failed":
		return WebhookTypesDubbingProjectFailed, nil
	case "dubbing_language_completed":
		return WebhookTypesDubbingLanguageCompleted, nil
	case "dubbing_language_failed":
		return WebhookTypesDubbingLanguageFailed, nil
	case "support_ticket_message_added":
		return WebhookTypesSupportTicketMessageAdded, nil
	}
	var t WebhookTypes
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookTypes) Ptr() *WebhookTypes {
	return &e
}
