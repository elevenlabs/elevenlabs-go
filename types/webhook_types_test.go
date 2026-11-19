package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWebhookTypesVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant WebhookTypes
		wire    string
	}
	cases := []enumCase{
		{variant: WebhookTypesProjectConversionStatus, wire: "\"project_conversion_status\""},
		{variant: WebhookTypesChapterConversionStatus, wire: "\"chapter_conversion_status\""},
		{variant: WebhookTypesPostCallTranscription, wire: "\"post_call_transcription\""},
		{variant: WebhookTypesPostCallTranscriptionOtel, wire: "\"post_call_transcription_otel\""},
		{variant: WebhookTypesPostCallAudio, wire: "\"post_call_audio\""},
		{variant: WebhookTypesPostCallUnredactedTranscription, wire: "\"post_call_unredacted_transcription\""},
		{variant: WebhookTypesPostCallUnredactedTranscriptionOtel, wire: "\"post_call_unredacted_transcription_otel\""},
		{variant: WebhookTypesPostCallUnredactedAudio, wire: "\"post_call_unredacted_audio\""},
		{variant: WebhookTypesPostCallAnalysis, wire: "\"post_call_analysis\""},
		{variant: WebhookTypesCallInitiationFailure, wire: "\"call_initiation_failure\""},
		{variant: WebhookTypesConversationInitiationClientData, wire: "\"conversation_initiation_client_data\""},
		{variant: WebhookTypesAnsweringMachineDetection, wire: "\"answering_machine_detection\""},
		{variant: WebhookTypesVoiceRemovalNotice, wire: "\"voice_removal_notice\""},
		{variant: WebhookTypesVoiceRemovalNoticeWithdrawn, wire: "\"voice_removal_notice_withdrawn\""},
		{variant: WebhookTypesVoiceRemoved, wire: "\"voice_removed\""},
		{variant: WebhookTypesSpeechToTextTranscription, wire: "\"speech_to_text_transcription\""},
		{variant: WebhookTypesWebhookError, wire: "\"webhook_error\""},
		{variant: WebhookTypesSpeechToTextBatch, wire: "\"speech_to_text_batch\""},
		{variant: WebhookTypesAgentQaEvaluationCompleted, wire: "\"agent_qa_evaluation_completed\""},
		{variant: WebhookTypesConvaiAlertingNotification, wire: "\"convai_alerting_notification\""},
		{variant: WebhookTypesFlowsGeneration, wire: "\"flows_generation\""},
		{variant: WebhookTypesFlowsTemplateRun, wire: "\"flows_template_run\""},
		{variant: WebhookTypesDubbingProjectReady, wire: "\"dubbing_project_ready\""},
		{variant: WebhookTypesDubbingProjectFailed, wire: "\"dubbing_project_failed\""},
		{variant: WebhookTypesDubbingLanguageCompleted, wire: "\"dubbing_language_completed\""},
		{variant: WebhookTypesDubbingLanguageFailed, wire: "\"dubbing_language_failed\""},
		{variant: WebhookTypesSupportTicketMessageAdded, wire: "\"support_ticket_message_added\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded WebhookTypes
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
