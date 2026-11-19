package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetTextContents struct {
	// Call to action displayed inside the compact and full variants.
	MainLabel *string `json:"main_label,omitempty" xml:"main_label,omitempty"`
	// Text and ARIA label for the start call button.
	StartCall *string `json:"start_call,omitempty" xml:"start_call,omitempty"`
	// Text and ARIA label for the start chat button (text only)
	StartChat *string `json:"start_chat,omitempty" xml:"start_chat,omitempty"`
	// Text and ARIA label for the new call button. Displayed when the caller already finished at least one call in order ot start the next one.
	NewCall *string `json:"new_call,omitempty" xml:"new_call,omitempty"`
	// Text and ARIA label for the end call button.
	EndCall *string `json:"end_call,omitempty" xml:"end_call,omitempty"`
	// ARIA label for the mute microphone button.
	MuteMicrophone *string `json:"mute_microphone,omitempty" xml:"mute_microphone,omitempty"`
	// ARIA label for the change language dropdown.
	ChangeLanguage *string `json:"change_language,omitempty" xml:"change_language,omitempty"`
	// ARIA label for the collapse button.
	Collapse *string `json:"collapse,omitempty" xml:"collapse,omitempty"`
	// ARIA label for the expand button.
	Expand *string `json:"expand,omitempty" xml:"expand,omitempty"`
	// Text displayed when the user copies a value using the copy button.
	Copied *string `json:"copied,omitempty" xml:"copied,omitempty"`
	// Text and ARIA label for the accept terms button.
	AcceptTerms *string `json:"accept_terms,omitempty" xml:"accept_terms,omitempty"`
	// Text and ARIA label for the cancel terms button.
	DismissTerms *string `json:"dismiss_terms,omitempty" xml:"dismiss_terms,omitempty"`
	// Status displayed when the agent is listening.
	ListeningStatus *string `json:"listening_status,omitempty" xml:"listening_status,omitempty"`
	// Status displayed when the agent is speaking.
	SpeakingStatus *string `json:"speaking_status,omitempty" xml:"speaking_status,omitempty"`
	// Status displayed when the agent is connecting.
	ConnectingStatus *string `json:"connecting_status,omitempty" xml:"connecting_status,omitempty"`
	// Status displayed when the agent is chatting (text only)
	ChattingStatus *string `json:"chatting_status,omitempty" xml:"chatting_status,omitempty"`
	// Status displayed while waiting in the queue for an available agent.
	QueueWaitingStatus *string `json:"queue_waiting_status,omitempty" xml:"queue_waiting_status,omitempty"`
	// Short status displayed while waiting for an available agent.
	QueueWaitingStatusShort *string `json:"queue_waiting_status_short,omitempty" xml:"queue_waiting_status_short,omitempty"`
	// ARIA label for the text message input.
	InputLabel *string `json:"input_label,omitempty" xml:"input_label,omitempty"`
	// Placeholder text for the text message input.
	InputPlaceholder *string `json:"input_placeholder,omitempty" xml:"input_placeholder,omitempty"`
	// Placeholder text for the text message input (text only)
	InputPlaceholderTextOnly *string `json:"input_placeholder_text_only,omitempty" xml:"input_placeholder_text_only,omitempty"`
	// Placeholder text for the text message input when starting a new conversation (text only)
	InputPlaceholderNewConversation *string `json:"input_placeholder_new_conversation,omitempty" xml:"input_placeholder_new_conversation,omitempty"`
	// Information message displayed when the user ends the conversation.
	UserEndedConversation *string `json:"user_ended_conversation,omitempty" xml:"user_ended_conversation,omitempty"`
	// Information message displayed when the agent ends the conversation.
	AgentEndedConversation *string `json:"agent_ended_conversation,omitempty" xml:"agent_ended_conversation,omitempty"`
	// Text label used next to the conversation ID.
	ConversationId *string `json:"conversation_id,omitempty" xml:"conversation_id,omitempty"`
	// Text label used when an error occurs.
	ErrorOccurred *string `json:"error_occurred,omitempty" xml:"error_occurred,omitempty"`
	// Error message displayed when the queue wait times out.
	QueueTimedOut *string `json:"queue_timed_out,omitempty" xml:"queue_timed_out,omitempty"`
	// Text and ARIA label used for the copy ID button.
	CopyId *string `json:"copy_id,omitempty" xml:"copy_id,omitempty"`
	// Text displayed to prompt the user for feedback.
	InitiateFeedback *string `json:"initiate_feedback,omitempty" xml:"initiate_feedback,omitempty"`
	// Text displayed to request additional feedback details.
	RequestFollowUpFeedback *string `json:"request_follow_up_feedback,omitempty" xml:"request_follow_up_feedback,omitempty"`
	// Text displayed to thank the user for providing feedback.
	ThanksForFeedback *string `json:"thanks_for_feedback,omitempty" xml:"thanks_for_feedback,omitempty"`
	// Additional text displayed explaining the value of user feedback.
	ThanksForFeedbackDetails *string `json:"thanks_for_feedback_details,omitempty" xml:"thanks_for_feedback_details,omitempty"`
	// Placeholder text for the follow-up feedback input field.
	FollowUpFeedbackPlaceholder *string `json:"follow_up_feedback_placeholder,omitempty" xml:"follow_up_feedback_placeholder,omitempty"`
	// Text and ARIA label for the submit button.
	Submit *string `json:"submit,omitempty" xml:"submit,omitempty"`
	// Text and ARIA label for the go back button.
	GoBack *string `json:"go_back,omitempty" xml:"go_back,omitempty"`
	// Text and ARIA label for the send message button.
	SendMessage *string `json:"send_message,omitempty" xml:"send_message,omitempty"`
	// Text and ARIA label for the switch to text mode button.
	TextMode *string `json:"text_mode,omitempty" xml:"text_mode,omitempty"`
	// Text and ARIA label for the switch to voice mode button.
	VoiceMode *string `json:"voice_mode,omitempty" xml:"voice_mode,omitempty"`
	// Toast notification displayed when switching to text mode.
	SwitchedToTextMode *string `json:"switched_to_text_mode,omitempty" xml:"switched_to_text_mode,omitempty"`
	// Toast notification displayed when switching to voice mode.
	SwitchedToVoiceMode *string `json:"switched_to_voice_mode,omitempty" xml:"switched_to_voice_mode,omitempty"`
	// Text and ARIA label for the copy button.
	Copy *string `json:"copy,omitempty" xml:"copy,omitempty"`
	// Text and ARIA label for the download button.
	Download *string `json:"download,omitempty" xml:"download,omitempty"`
	// Text and ARIA label for the wrap toggle button.
	Wrap *string `json:"wrap,omitempty" xml:"wrap,omitempty"`
	// Status text displayed when the agent is processing a tool call.
	AgentWorking *string `json:"agent_working,omitempty" xml:"agent_working,omitempty"`
	// Status text displayed when the agent finishes processing a tool call.
	AgentDone *string `json:"agent_done,omitempty" xml:"agent_done,omitempty"`
	// Status text displayed when the agent encounters an error during a tool call.
	AgentError *string `json:"agent_error,omitempty" xml:"agent_error,omitempty"`
	// Text and ARIA label for the attach file button.
	AttachFile *string `json:"attach_file,omitempty" xml:"attach_file,omitempty"`
	// ARIA label for the remove file button.
	RemoveFile *string `json:"remove_file,omitempty" xml:"remove_file,omitempty"`
	// Error message displayed when a file fails to upload.
	FileUploadError *string `json:"file_upload_error,omitempty" xml:"file_upload_error,omitempty"`
	// Error message displayed when an unsupported file type is selected. Followed by the list of accepted types.
	FileTypeUnsupported *string `json:"file_type_unsupported,omitempty" xml:"file_type_unsupported,omitempty"`
	// Error message displayed when a file exceeds the maximum size limit.
	FileTooLarge *string `json:"file_too_large,omitempty" xml:"file_too_large,omitempty"`
	// Error message displayed when the maximum number of files for a conversation is reached.
	FileLimitReached *string `json:"file_limit_reached,omitempty" xml:"file_limit_reached,omitempty"`
	// Status text displayed while the agent is typing.
	TypingIndicator *string `json:"typing_indicator,omitempty" xml:"typing_indicator,omitempty"`
	// Fallback message displayed when rich content cannot be rendered.
	RichContentUnavailable *string `json:"rich_content_unavailable,omitempty" xml:"rich_content_unavailable,omitempty"`
	extraProperties        map[string]interface{}
}

func (w *WidgetTextContents) GetMainLabel() (value *string) {
	if w == nil {
		return
	}
	return w.MainLabel
}

func (w *WidgetTextContents) GetStartCall() (value *string) {
	if w == nil {
		return
	}
	return w.StartCall
}

func (w *WidgetTextContents) GetStartChat() (value *string) {
	if w == nil {
		return
	}
	return w.StartChat
}

func (w *WidgetTextContents) GetNewCall() (value *string) {
	if w == nil {
		return
	}
	return w.NewCall
}

func (w *WidgetTextContents) GetEndCall() (value *string) {
	if w == nil {
		return
	}
	return w.EndCall
}

func (w *WidgetTextContents) GetMuteMicrophone() (value *string) {
	if w == nil {
		return
	}
	return w.MuteMicrophone
}

func (w *WidgetTextContents) GetChangeLanguage() (value *string) {
	if w == nil {
		return
	}
	return w.ChangeLanguage
}

func (w *WidgetTextContents) GetCollapse() (value *string) {
	if w == nil {
		return
	}
	return w.Collapse
}

func (w *WidgetTextContents) GetExpand() (value *string) {
	if w == nil {
		return
	}
	return w.Expand
}

func (w *WidgetTextContents) GetCopied() (value *string) {
	if w == nil {
		return
	}
	return w.Copied
}

func (w *WidgetTextContents) GetAcceptTerms() (value *string) {
	if w == nil {
		return
	}
	return w.AcceptTerms
}

func (w *WidgetTextContents) GetDismissTerms() (value *string) {
	if w == nil {
		return
	}
	return w.DismissTerms
}

func (w *WidgetTextContents) GetListeningStatus() (value *string) {
	if w == nil {
		return
	}
	return w.ListeningStatus
}

func (w *WidgetTextContents) GetSpeakingStatus() (value *string) {
	if w == nil {
		return
	}
	return w.SpeakingStatus
}

func (w *WidgetTextContents) GetConnectingStatus() (value *string) {
	if w == nil {
		return
	}
	return w.ConnectingStatus
}

func (w *WidgetTextContents) GetChattingStatus() (value *string) {
	if w == nil {
		return
	}
	return w.ChattingStatus
}

func (w *WidgetTextContents) GetQueueWaitingStatus() (value *string) {
	if w == nil {
		return
	}
	return w.QueueWaitingStatus
}

func (w *WidgetTextContents) GetQueueWaitingStatusShort() (value *string) {
	if w == nil {
		return
	}
	return w.QueueWaitingStatusShort
}

func (w *WidgetTextContents) GetInputLabel() (value *string) {
	if w == nil {
		return
	}
	return w.InputLabel
}

func (w *WidgetTextContents) GetInputPlaceholder() (value *string) {
	if w == nil {
		return
	}
	return w.InputPlaceholder
}

func (w *WidgetTextContents) GetInputPlaceholderTextOnly() (value *string) {
	if w == nil {
		return
	}
	return w.InputPlaceholderTextOnly
}

func (w *WidgetTextContents) GetInputPlaceholderNewConversation() (value *string) {
	if w == nil {
		return
	}
	return w.InputPlaceholderNewConversation
}

func (w *WidgetTextContents) GetUserEndedConversation() (value *string) {
	if w == nil {
		return
	}
	return w.UserEndedConversation
}

func (w *WidgetTextContents) GetAgentEndedConversation() (value *string) {
	if w == nil {
		return
	}
	return w.AgentEndedConversation
}

func (w *WidgetTextContents) GetConversationId() (value *string) {
	if w == nil {
		return
	}
	return w.ConversationId
}

func (w *WidgetTextContents) GetErrorOccurred() (value *string) {
	if w == nil {
		return
	}
	return w.ErrorOccurred
}

func (w *WidgetTextContents) GetQueueTimedOut() (value *string) {
	if w == nil {
		return
	}
	return w.QueueTimedOut
}

func (w *WidgetTextContents) GetCopyId() (value *string) {
	if w == nil {
		return
	}
	return w.CopyId
}

func (w *WidgetTextContents) GetInitiateFeedback() (value *string) {
	if w == nil {
		return
	}
	return w.InitiateFeedback
}

func (w *WidgetTextContents) GetRequestFollowUpFeedback() (value *string) {
	if w == nil {
		return
	}
	return w.RequestFollowUpFeedback
}

func (w *WidgetTextContents) GetThanksForFeedback() (value *string) {
	if w == nil {
		return
	}
	return w.ThanksForFeedback
}

func (w *WidgetTextContents) GetThanksForFeedbackDetails() (value *string) {
	if w == nil {
		return
	}
	return w.ThanksForFeedbackDetails
}

func (w *WidgetTextContents) GetFollowUpFeedbackPlaceholder() (value *string) {
	if w == nil {
		return
	}
	return w.FollowUpFeedbackPlaceholder
}

func (w *WidgetTextContents) GetSubmit() (value *string) {
	if w == nil {
		return
	}
	return w.Submit
}

func (w *WidgetTextContents) GetGoBack() (value *string) {
	if w == nil {
		return
	}
	return w.GoBack
}

func (w *WidgetTextContents) GetSendMessage() (value *string) {
	if w == nil {
		return
	}
	return w.SendMessage
}

func (w *WidgetTextContents) GetTextMode() (value *string) {
	if w == nil {
		return
	}
	return w.TextMode
}

func (w *WidgetTextContents) GetVoiceMode() (value *string) {
	if w == nil {
		return
	}
	return w.VoiceMode
}

func (w *WidgetTextContents) GetSwitchedToTextMode() (value *string) {
	if w == nil {
		return
	}
	return w.SwitchedToTextMode
}

func (w *WidgetTextContents) GetSwitchedToVoiceMode() (value *string) {
	if w == nil {
		return
	}
	return w.SwitchedToVoiceMode
}

func (w *WidgetTextContents) GetCopy() (value *string) {
	if w == nil {
		return
	}
	return w.Copy
}

func (w *WidgetTextContents) GetDownload() (value *string) {
	if w == nil {
		return
	}
	return w.Download
}

func (w *WidgetTextContents) GetWrap() (value *string) {
	if w == nil {
		return
	}
	return w.Wrap
}

func (w *WidgetTextContents) GetAgentWorking() (value *string) {
	if w == nil {
		return
	}
	return w.AgentWorking
}

func (w *WidgetTextContents) GetAgentDone() (value *string) {
	if w == nil {
		return
	}
	return w.AgentDone
}

func (w *WidgetTextContents) GetAgentError() (value *string) {
	if w == nil {
		return
	}
	return w.AgentError
}

func (w *WidgetTextContents) GetAttachFile() (value *string) {
	if w == nil {
		return
	}
	return w.AttachFile
}

func (w *WidgetTextContents) GetRemoveFile() (value *string) {
	if w == nil {
		return
	}
	return w.RemoveFile
}

func (w *WidgetTextContents) GetFileUploadError() (value *string) {
	if w == nil {
		return
	}
	return w.FileUploadError
}

func (w *WidgetTextContents) GetFileTypeUnsupported() (value *string) {
	if w == nil {
		return
	}
	return w.FileTypeUnsupported
}

func (w *WidgetTextContents) GetFileTooLarge() (value *string) {
	if w == nil {
		return
	}
	return w.FileTooLarge
}

func (w *WidgetTextContents) GetFileLimitReached() (value *string) {
	if w == nil {
		return
	}
	return w.FileLimitReached
}

func (w *WidgetTextContents) GetTypingIndicator() (value *string) {
	if w == nil {
		return
	}
	return w.TypingIndicator
}

func (w *WidgetTextContents) GetRichContentUnavailable() (value *string) {
	if w == nil {
		return
	}
	return w.RichContentUnavailable
}

func (w *WidgetTextContents) SetMainLabel(mainLabel *string) {
	w.MainLabel = mainLabel
}

func (w *WidgetTextContents) SetStartCall(startCall *string) {
	w.StartCall = startCall
}

func (w *WidgetTextContents) SetStartChat(startChat *string) {
	w.StartChat = startChat
}

func (w *WidgetTextContents) SetNewCall(newCall *string) {
	w.NewCall = newCall
}

func (w *WidgetTextContents) SetEndCall(endCall *string) {
	w.EndCall = endCall
}

func (w *WidgetTextContents) SetMuteMicrophone(muteMicrophone *string) {
	w.MuteMicrophone = muteMicrophone
}

func (w *WidgetTextContents) SetChangeLanguage(changeLanguage *string) {
	w.ChangeLanguage = changeLanguage
}

func (w *WidgetTextContents) SetCollapse(collapse *string) {
	w.Collapse = collapse
}

func (w *WidgetTextContents) SetExpand(expand *string) {
	w.Expand = expand
}

func (w *WidgetTextContents) SetCopied(copied *string) {
	w.Copied = copied
}

func (w *WidgetTextContents) SetAcceptTerms(acceptTerms *string) {
	w.AcceptTerms = acceptTerms
}

func (w *WidgetTextContents) SetDismissTerms(dismissTerms *string) {
	w.DismissTerms = dismissTerms
}

func (w *WidgetTextContents) SetListeningStatus(listeningStatus *string) {
	w.ListeningStatus = listeningStatus
}

func (w *WidgetTextContents) SetSpeakingStatus(speakingStatus *string) {
	w.SpeakingStatus = speakingStatus
}

func (w *WidgetTextContents) SetConnectingStatus(connectingStatus *string) {
	w.ConnectingStatus = connectingStatus
}

func (w *WidgetTextContents) SetChattingStatus(chattingStatus *string) {
	w.ChattingStatus = chattingStatus
}

func (w *WidgetTextContents) SetQueueWaitingStatus(queueWaitingStatus *string) {
	w.QueueWaitingStatus = queueWaitingStatus
}

func (w *WidgetTextContents) SetQueueWaitingStatusShort(queueWaitingStatusShort *string) {
	w.QueueWaitingStatusShort = queueWaitingStatusShort
}

func (w *WidgetTextContents) SetInputLabel(inputLabel *string) {
	w.InputLabel = inputLabel
}

func (w *WidgetTextContents) SetInputPlaceholder(inputPlaceholder *string) {
	w.InputPlaceholder = inputPlaceholder
}

func (w *WidgetTextContents) SetInputPlaceholderTextOnly(inputPlaceholderTextOnly *string) {
	w.InputPlaceholderTextOnly = inputPlaceholderTextOnly
}

func (w *WidgetTextContents) SetInputPlaceholderNewConversation(inputPlaceholderNewConversation *string) {
	w.InputPlaceholderNewConversation = inputPlaceholderNewConversation
}

func (w *WidgetTextContents) SetUserEndedConversation(userEndedConversation *string) {
	w.UserEndedConversation = userEndedConversation
}

func (w *WidgetTextContents) SetAgentEndedConversation(agentEndedConversation *string) {
	w.AgentEndedConversation = agentEndedConversation
}

func (w *WidgetTextContents) SetConversationId(conversationId *string) {
	w.ConversationId = conversationId
}

func (w *WidgetTextContents) SetErrorOccurred(errorOccurred *string) {
	w.ErrorOccurred = errorOccurred
}

func (w *WidgetTextContents) SetQueueTimedOut(queueTimedOut *string) {
	w.QueueTimedOut = queueTimedOut
}

func (w *WidgetTextContents) SetCopyId(copyId *string) {
	w.CopyId = copyId
}

func (w *WidgetTextContents) SetInitiateFeedback(initiateFeedback *string) {
	w.InitiateFeedback = initiateFeedback
}

func (w *WidgetTextContents) SetRequestFollowUpFeedback(requestFollowUpFeedback *string) {
	w.RequestFollowUpFeedback = requestFollowUpFeedback
}

func (w *WidgetTextContents) SetThanksForFeedback(thanksForFeedback *string) {
	w.ThanksForFeedback = thanksForFeedback
}

func (w *WidgetTextContents) SetThanksForFeedbackDetails(thanksForFeedbackDetails *string) {
	w.ThanksForFeedbackDetails = thanksForFeedbackDetails
}

func (w *WidgetTextContents) SetFollowUpFeedbackPlaceholder(followUpFeedbackPlaceholder *string) {
	w.FollowUpFeedbackPlaceholder = followUpFeedbackPlaceholder
}

func (w *WidgetTextContents) SetSubmit(submit *string) {
	w.Submit = submit
}

func (w *WidgetTextContents) SetGoBack(goBack *string) {
	w.GoBack = goBack
}

func (w *WidgetTextContents) SetSendMessage(sendMessage *string) {
	w.SendMessage = sendMessage
}

func (w *WidgetTextContents) SetTextMode(textMode *string) {
	w.TextMode = textMode
}

func (w *WidgetTextContents) SetVoiceMode(voiceMode *string) {
	w.VoiceMode = voiceMode
}

func (w *WidgetTextContents) SetSwitchedToTextMode(switchedToTextMode *string) {
	w.SwitchedToTextMode = switchedToTextMode
}

func (w *WidgetTextContents) SetSwitchedToVoiceMode(switchedToVoiceMode *string) {
	w.SwitchedToVoiceMode = switchedToVoiceMode
}

func (w *WidgetTextContents) SetCopy(copy *string) {
	w.Copy = copy
}

func (w *WidgetTextContents) SetDownload(download *string) {
	w.Download = download
}

func (w *WidgetTextContents) SetWrap(wrap *string) {
	w.Wrap = wrap
}

func (w *WidgetTextContents) SetAgentWorking(agentWorking *string) {
	w.AgentWorking = agentWorking
}

func (w *WidgetTextContents) SetAgentDone(agentDone *string) {
	w.AgentDone = agentDone
}

func (w *WidgetTextContents) SetAgentError(agentError *string) {
	w.AgentError = agentError
}

func (w *WidgetTextContents) SetAttachFile(attachFile *string) {
	w.AttachFile = attachFile
}

func (w *WidgetTextContents) SetRemoveFile(removeFile *string) {
	w.RemoveFile = removeFile
}

func (w *WidgetTextContents) SetFileUploadError(fileUploadError *string) {
	w.FileUploadError = fileUploadError
}

func (w *WidgetTextContents) SetFileTypeUnsupported(fileTypeUnsupported *string) {
	w.FileTypeUnsupported = fileTypeUnsupported
}

func (w *WidgetTextContents) SetFileTooLarge(fileTooLarge *string) {
	w.FileTooLarge = fileTooLarge
}

func (w *WidgetTextContents) SetFileLimitReached(fileLimitReached *string) {
	w.FileLimitReached = fileLimitReached
}

func (w *WidgetTextContents) SetTypingIndicator(typingIndicator *string) {
	w.TypingIndicator = typingIndicator
}

func (w *WidgetTextContents) SetRichContentUnavailable(richContentUnavailable *string) {
	w.RichContentUnavailable = richContentUnavailable
}

func (w *WidgetTextContents) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetTextContents) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetTextContents to string"
	}
	return string(jsonData)
}

func (w *WidgetTextContents) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WidgetTextContents) MarshalJSON() ([]byte, error) {
	type embed WidgetTextContents
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
