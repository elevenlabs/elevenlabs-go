package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetConfigResponse struct {
	Variant    *EmbedVariant     `json:"variant,omitempty" xml:"variant,omitempty"`
	Placement  *WidgetPlacement  `json:"placement,omitempty" xml:"placement,omitempty"`
	Expandable *WidgetExpandable `json:"expandable,omitempty" xml:"expandable,omitempty"`
	// The avatar of the widget
	Avatar       *WidgetConfigResponseAvatar `json:"avatar,omitempty" xml:"avatar,omitempty"`
	FeedbackMode *WidgetFeedbackMode         `json:"feedback_mode,omitempty" xml:"feedback_mode,omitempty"`
	// Configuration for feedback collected at the end of the conversation
	EndFeedback *WidgetEndFeedbackConfig `json:"end_feedback,omitempty" xml:"end_feedback,omitempty"`
	// The background color of the widget
	BgColor *string `json:"bg_color,omitempty" xml:"bg_color,omitempty"`
	// The text color of the widget
	TextColor *string `json:"text_color,omitempty" xml:"text_color,omitempty"`
	// The button color of the widget
	BtnColor *string `json:"btn_color,omitempty" xml:"btn_color,omitempty"`
	// The button text color of the widget
	BtnTextColor *string `json:"btn_text_color,omitempty" xml:"btn_text_color,omitempty"`
	// The border color of the widget
	BorderColor *string `json:"border_color,omitempty" xml:"border_color,omitempty"`
	// The focus color of the widget
	FocusColor *string `json:"focus_color,omitempty" xml:"focus_color,omitempty"`
	// The border radius of the widget
	BorderRadius *int `json:"border_radius,omitempty" xml:"border_radius,omitempty"`
	// The button radius of the widget
	BtnRadius *int `json:"btn_radius,omitempty" xml:"btn_radius,omitempty"`
	// The action text of the widget
	ActionText *string `json:"action_text,omitempty" xml:"action_text,omitempty"`
	// The start call text of the widget
	StartCallText *string `json:"start_call_text,omitempty" xml:"start_call_text,omitempty"`
	// The end call text of the widget
	EndCallText *string `json:"end_call_text,omitempty" xml:"end_call_text,omitempty"`
	// The expand text of the widget
	ExpandText *string `json:"expand_text,omitempty" xml:"expand_text,omitempty"`
	// The text to display when the agent is listening
	ListeningText *string `json:"listening_text,omitempty" xml:"listening_text,omitempty"`
	// The text to display when the agent is speaking
	SpeakingText *string `json:"speaking_text,omitempty" xml:"speaking_text,omitempty"`
	// The text to display when sharing
	ShareablePageText *string `json:"shareable_page_text,omitempty" xml:"shareable_page_text,omitempty"`
	// Whether to show terms and conditions on the shareable page
	ShareablePageShowTerms *bool `json:"shareable_page_show_terms,omitempty" xml:"shareable_page_show_terms,omitempty"`
	// The text to display for terms and conditions
	TermsText *string `json:"terms_text,omitempty" xml:"terms_text,omitempty"`
	// The HTML to display for terms and conditions
	TermsHtml *string `json:"terms_html,omitempty" xml:"terms_html,omitempty"`
	// The key to display for terms and conditions
	TermsKey *string `json:"terms_key,omitempty" xml:"terms_key,omitempty"`
	// Whether to show the avatar when the widget is collapsed
	ShowAvatarWhenCollapsed *bool `json:"show_avatar_when_collapsed,omitempty" xml:"show_avatar_when_collapsed,omitempty"`
	// Whether to disable the banner
	DisableBanner *bool `json:"disable_banner,omitempty" xml:"disable_banner,omitempty"`
	// The override link for the widget
	OverrideLink *string `json:"override_link,omitempty" xml:"override_link,omitempty"`
	// List of allowed hostnames for clickable markdown links. Use { hostname: '*' } to allow any domain. Empty means no links are allowed.
	MarkdownLinkAllowedHosts []*AllowlistItem `json:"markdown_link_allowed_hosts,omitempty" xml:"markdown_link_allowed_hosts,omitempty"`
	// Whether to automatically include www. variants of allowed hosts
	MarkdownLinkIncludeWww *bool `json:"markdown_link_include_www,omitempty" xml:"markdown_link_include_www,omitempty"`
	// Whether to allow http:// in addition to https:// for allowed hosts
	MarkdownLinkAllowHttp *bool `json:"markdown_link_allow_http,omitempty" xml:"markdown_link_allow_http,omitempty"`
	// Whether to enable mic muting
	MicMutingEnabled *bool `json:"mic_muting_enabled,omitempty" xml:"mic_muting_enabled,omitempty"`
	// Whether the widget should show the conversation transcript as it goes on
	TranscriptEnabled *bool `json:"transcript_enabled,omitempty" xml:"transcript_enabled,omitempty"`
	// Whether the user should be able to send text messages
	TextInputEnabled *bool `json:"text_input_enabled,omitempty" xml:"text_input_enabled,omitempty"`
	// Whether to enable the conversation mode toggle in the widget
	ConversationModeToggleEnabled *bool `json:"conversation_mode_toggle_enabled,omitempty" xml:"conversation_mode_toggle_enabled,omitempty"`
	// Whether the widget should be expanded by default
	DefaultExpanded *bool `json:"default_expanded,omitempty" xml:"default_expanded,omitempty"`
	// Whether the widget should always be expanded
	AlwaysExpanded *bool `json:"always_expanded,omitempty" xml:"always_expanded,omitempty"`
	// Whether the widget can be dismissed by the user
	Dismissible *bool `json:"dismissible,omitempty" xml:"dismissible,omitempty"`
	// Whether to show agent working/done/error status during tool use
	ShowAgentStatus *bool `json:"show_agent_status,omitempty" xml:"show_agent_status,omitempty"`
	// Whether to show the conversation ID after disconnection.
	ShowConversationId *bool `json:"show_conversation_id,omitempty" xml:"show_conversation_id,omitempty"`
	// Whether to strip audio markup from messages.
	StripAudioTags *bool `json:"strip_audio_tags,omitempty" xml:"strip_audio_tags,omitempty"`
	// Theme for code block syntax highlighting. Defaults to auto-detection by the widget when not set.
	SyntaxHighlightTheme *WidgetConfigResponseSyntaxHighlightTheme `json:"syntax_highlight_theme,omitempty" xml:"syntax_highlight_theme,omitempty"`
	TextContents         *WidgetTextContents                       `json:"text_contents,omitempty" xml:"text_contents,omitempty"`
	Styles               *WidgetStyles                             `json:"styles,omitempty" xml:"styles,omitempty"`
	// Whether to show the resize button
	ShowResizeButton *bool `json:"show_resize_button,omitempty" xml:"show_resize_button,omitempty"`
	// Whether to ask the user to confirm before ending a chat or call
	EndConfirmationEnabled *bool `json:"end_confirmation_enabled,omitempty" xml:"end_confirmation_enabled,omitempty"`
	// Whether the widget opens fullscreen on mobile
	MobileFullscreen *bool `json:"mobile_fullscreen,omitempty" xml:"mobile_fullscreen,omitempty"`
	// Whether voice call transcripts are revealed as the agent speaks
	TranscriptSpeechSyncEnabled *bool    `json:"transcript_speech_sync_enabled,omitempty" xml:"transcript_speech_sync_enabled,omitempty"`
	Language                    string   `json:"language" xml:"language" required:"true"`
	SupportedLanguageOverrides  []string `json:"supported_language_overrides,omitempty" xml:"supported_language_overrides,omitempty"`
	// Language presets for the widget
	LanguagePresets map[string]*WidgetLanguagePresetResponse `json:"language_presets,omitempty"`
	// Whether the agent uses text-only mode
	TextOnly *bool `json:"text_only,omitempty" xml:"text_only,omitempty"`
	// Whether the agent can be switched to text-only mode
	SupportsTextOnly *bool   `json:"supports_text_only,omitempty" xml:"supports_text_only,omitempty"`
	FirstMessage     *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// Whether to use WebRTC for conversation connections
	UseRtc          *bool            `json:"use_rtc,omitempty" xml:"use_rtc,omitempty"`
	FileInputConfig *FileInputConfig `json:"file_input_config,omitempty" xml:"file_input_config,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WidgetConfigResponse) GetVariant() (value *EmbedVariant) {
	if w == nil {
		return
	}
	return w.Variant
}

func (w *WidgetConfigResponse) GetPlacement() (value *WidgetPlacement) {
	if w == nil {
		return
	}
	return w.Placement
}

func (w *WidgetConfigResponse) GetExpandable() (value *WidgetExpandable) {
	if w == nil {
		return
	}
	return w.Expandable
}

func (w *WidgetConfigResponse) GetAvatar() (value *WidgetConfigResponseAvatar) {
	if w == nil {
		return
	}
	return w.Avatar
}

func (w *WidgetConfigResponse) GetFeedbackMode() (value *WidgetFeedbackMode) {
	if w == nil {
		return
	}
	return w.FeedbackMode
}

func (w *WidgetConfigResponse) GetEndFeedback() (value *WidgetEndFeedbackConfig) {
	if w == nil {
		return
	}
	return w.EndFeedback
}

func (w *WidgetConfigResponse) GetBgColor() (value *string) {
	if w == nil {
		return
	}
	return w.BgColor
}

func (w *WidgetConfigResponse) GetTextColor() (value *string) {
	if w == nil {
		return
	}
	return w.TextColor
}

func (w *WidgetConfigResponse) GetBtnColor() (value *string) {
	if w == nil {
		return
	}
	return w.BtnColor
}

func (w *WidgetConfigResponse) GetBtnTextColor() (value *string) {
	if w == nil {
		return
	}
	return w.BtnTextColor
}

func (w *WidgetConfigResponse) GetBorderColor() (value *string) {
	if w == nil {
		return
	}
	return w.BorderColor
}

func (w *WidgetConfigResponse) GetFocusColor() (value *string) {
	if w == nil {
		return
	}
	return w.FocusColor
}

func (w *WidgetConfigResponse) GetBorderRadius() (value *int) {
	if w == nil {
		return
	}
	return w.BorderRadius
}

func (w *WidgetConfigResponse) GetBtnRadius() (value *int) {
	if w == nil {
		return
	}
	return w.BtnRadius
}

func (w *WidgetConfigResponse) GetActionText() (value *string) {
	if w == nil {
		return
	}
	return w.ActionText
}

func (w *WidgetConfigResponse) GetStartCallText() (value *string) {
	if w == nil {
		return
	}
	return w.StartCallText
}

func (w *WidgetConfigResponse) GetEndCallText() (value *string) {
	if w == nil {
		return
	}
	return w.EndCallText
}

func (w *WidgetConfigResponse) GetExpandText() (value *string) {
	if w == nil {
		return
	}
	return w.ExpandText
}

func (w *WidgetConfigResponse) GetListeningText() (value *string) {
	if w == nil {
		return
	}
	return w.ListeningText
}

func (w *WidgetConfigResponse) GetSpeakingText() (value *string) {
	if w == nil {
		return
	}
	return w.SpeakingText
}

func (w *WidgetConfigResponse) GetShareablePageText() (value *string) {
	if w == nil {
		return
	}
	return w.ShareablePageText
}

func (w *WidgetConfigResponse) GetShareablePageShowTerms() (value *bool) {
	if w == nil {
		return
	}
	return w.ShareablePageShowTerms
}

func (w *WidgetConfigResponse) GetTermsText() (value *string) {
	if w == nil {
		return
	}
	return w.TermsText
}

func (w *WidgetConfigResponse) GetTermsHtml() (value *string) {
	if w == nil {
		return
	}
	return w.TermsHtml
}

func (w *WidgetConfigResponse) GetTermsKey() (value *string) {
	if w == nil {
		return
	}
	return w.TermsKey
}

func (w *WidgetConfigResponse) GetShowAvatarWhenCollapsed() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowAvatarWhenCollapsed
}

func (w *WidgetConfigResponse) GetDisableBanner() (value *bool) {
	if w == nil {
		return
	}
	return w.DisableBanner
}

func (w *WidgetConfigResponse) GetOverrideLink() (value *string) {
	if w == nil {
		return
	}
	return w.OverrideLink
}

func (w *WidgetConfigResponse) GetMarkdownLinkAllowedHosts() (value []*AllowlistItem) {
	if w == nil {
		return
	}
	return w.MarkdownLinkAllowedHosts
}

func (w *WidgetConfigResponse) GetMarkdownLinkIncludeWww() (value *bool) {
	if w == nil {
		return
	}
	return w.MarkdownLinkIncludeWww
}

func (w *WidgetConfigResponse) GetMarkdownLinkAllowHttp() (value *bool) {
	if w == nil {
		return
	}
	return w.MarkdownLinkAllowHttp
}

func (w *WidgetConfigResponse) GetMicMutingEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.MicMutingEnabled
}

func (w *WidgetConfigResponse) GetTranscriptEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.TranscriptEnabled
}

func (w *WidgetConfigResponse) GetTextInputEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.TextInputEnabled
}

func (w *WidgetConfigResponse) GetConversationModeToggleEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.ConversationModeToggleEnabled
}

func (w *WidgetConfigResponse) GetDefaultExpanded() (value *bool) {
	if w == nil {
		return
	}
	return w.DefaultExpanded
}

func (w *WidgetConfigResponse) GetAlwaysExpanded() (value *bool) {
	if w == nil {
		return
	}
	return w.AlwaysExpanded
}

func (w *WidgetConfigResponse) GetDismissible() (value *bool) {
	if w == nil {
		return
	}
	return w.Dismissible
}

func (w *WidgetConfigResponse) GetShowAgentStatus() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowAgentStatus
}

func (w *WidgetConfigResponse) GetShowConversationId() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowConversationId
}

func (w *WidgetConfigResponse) GetStripAudioTags() (value *bool) {
	if w == nil {
		return
	}
	return w.StripAudioTags
}

func (w *WidgetConfigResponse) GetSyntaxHighlightTheme() (value *WidgetConfigResponseSyntaxHighlightTheme) {
	if w == nil {
		return
	}
	return w.SyntaxHighlightTheme
}

func (w *WidgetConfigResponse) GetTextContents() (value *WidgetTextContents) {
	if w == nil {
		return
	}
	return w.TextContents
}

func (w *WidgetConfigResponse) GetStyles() (value *WidgetStyles) {
	if w == nil {
		return
	}
	return w.Styles
}

func (w *WidgetConfigResponse) GetShowResizeButton() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowResizeButton
}

func (w *WidgetConfigResponse) GetEndConfirmationEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.EndConfirmationEnabled
}

func (w *WidgetConfigResponse) GetMobileFullscreen() (value *bool) {
	if w == nil {
		return
	}
	return w.MobileFullscreen
}

func (w *WidgetConfigResponse) GetTranscriptSpeechSyncEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.TranscriptSpeechSyncEnabled
}

func (w *WidgetConfigResponse) GetLanguage() (value string) {
	if w == nil {
		return
	}
	return w.Language
}

func (w *WidgetConfigResponse) GetSupportedLanguageOverrides() (value []string) {
	if w == nil {
		return
	}
	return w.SupportedLanguageOverrides
}

func (w *WidgetConfigResponse) GetLanguagePresets() (value map[string]*WidgetLanguagePresetResponse) {
	if w == nil {
		return
	}
	return w.LanguagePresets
}

func (w *WidgetConfigResponse) GetTextOnly() (value *bool) {
	if w == nil {
		return
	}
	return w.TextOnly
}

func (w *WidgetConfigResponse) GetSupportsTextOnly() (value *bool) {
	if w == nil {
		return
	}
	return w.SupportsTextOnly
}

func (w *WidgetConfigResponse) GetFirstMessage() (value *string) {
	if w == nil {
		return
	}
	return w.FirstMessage
}

func (w *WidgetConfigResponse) GetUseRtc() (value *bool) {
	if w == nil {
		return
	}
	return w.UseRtc
}

func (w *WidgetConfigResponse) GetFileInputConfig() (value *FileInputConfig) {
	if w == nil {
		return
	}
	return w.FileInputConfig
}

func (w *WidgetConfigResponse) SetVariant(variant *EmbedVariant) {
	w.Variant = variant
}

func (w *WidgetConfigResponse) SetPlacement(placement *WidgetPlacement) {
	w.Placement = placement
}

func (w *WidgetConfigResponse) SetExpandable(expandable *WidgetExpandable) {
	w.Expandable = expandable
}

func (w *WidgetConfigResponse) SetAvatar(avatar *WidgetConfigResponseAvatar) {
	w.Avatar = avatar
}

func (w *WidgetConfigResponse) SetFeedbackMode(feedbackMode *WidgetFeedbackMode) {
	w.FeedbackMode = feedbackMode
}

func (w *WidgetConfigResponse) SetEndFeedback(endFeedback *WidgetEndFeedbackConfig) {
	w.EndFeedback = endFeedback
}

func (w *WidgetConfigResponse) SetBgColor(bgColor *string) {
	w.BgColor = bgColor
}

func (w *WidgetConfigResponse) SetTextColor(textColor *string) {
	w.TextColor = textColor
}

func (w *WidgetConfigResponse) SetBtnColor(btnColor *string) {
	w.BtnColor = btnColor
}

func (w *WidgetConfigResponse) SetBtnTextColor(btnTextColor *string) {
	w.BtnTextColor = btnTextColor
}

func (w *WidgetConfigResponse) SetBorderColor(borderColor *string) {
	w.BorderColor = borderColor
}

func (w *WidgetConfigResponse) SetFocusColor(focusColor *string) {
	w.FocusColor = focusColor
}

func (w *WidgetConfigResponse) SetBorderRadius(borderRadius *int) {
	w.BorderRadius = borderRadius
}

func (w *WidgetConfigResponse) SetBtnRadius(btnRadius *int) {
	w.BtnRadius = btnRadius
}

func (w *WidgetConfigResponse) SetActionText(actionText *string) {
	w.ActionText = actionText
}

func (w *WidgetConfigResponse) SetStartCallText(startCallText *string) {
	w.StartCallText = startCallText
}

func (w *WidgetConfigResponse) SetEndCallText(endCallText *string) {
	w.EndCallText = endCallText
}

func (w *WidgetConfigResponse) SetExpandText(expandText *string) {
	w.ExpandText = expandText
}

func (w *WidgetConfigResponse) SetListeningText(listeningText *string) {
	w.ListeningText = listeningText
}

func (w *WidgetConfigResponse) SetSpeakingText(speakingText *string) {
	w.SpeakingText = speakingText
}

func (w *WidgetConfigResponse) SetShareablePageText(shareablePageText *string) {
	w.ShareablePageText = shareablePageText
}

func (w *WidgetConfigResponse) SetShareablePageShowTerms(shareablePageShowTerms *bool) {
	w.ShareablePageShowTerms = shareablePageShowTerms
}

func (w *WidgetConfigResponse) SetTermsText(termsText *string) {
	w.TermsText = termsText
}

func (w *WidgetConfigResponse) SetTermsHtml(termsHtml *string) {
	w.TermsHtml = termsHtml
}

func (w *WidgetConfigResponse) SetTermsKey(termsKey *string) {
	w.TermsKey = termsKey
}

func (w *WidgetConfigResponse) SetShowAvatarWhenCollapsed(showAvatarWhenCollapsed *bool) {
	w.ShowAvatarWhenCollapsed = showAvatarWhenCollapsed
}

func (w *WidgetConfigResponse) SetDisableBanner(disableBanner *bool) {
	w.DisableBanner = disableBanner
}

func (w *WidgetConfigResponse) SetOverrideLink(overrideLink *string) {
	w.OverrideLink = overrideLink
}

func (w *WidgetConfigResponse) SetMarkdownLinkAllowedHosts(markdownLinkAllowedHosts []*AllowlistItem) {
	w.MarkdownLinkAllowedHosts = markdownLinkAllowedHosts
}

func (w *WidgetConfigResponse) SetMarkdownLinkIncludeWww(markdownLinkIncludeWww *bool) {
	w.MarkdownLinkIncludeWww = markdownLinkIncludeWww
}

func (w *WidgetConfigResponse) SetMarkdownLinkAllowHttp(markdownLinkAllowHttp *bool) {
	w.MarkdownLinkAllowHttp = markdownLinkAllowHttp
}

func (w *WidgetConfigResponse) SetMicMutingEnabled(micMutingEnabled *bool) {
	w.MicMutingEnabled = micMutingEnabled
}

func (w *WidgetConfigResponse) SetTranscriptEnabled(transcriptEnabled *bool) {
	w.TranscriptEnabled = transcriptEnabled
}

func (w *WidgetConfigResponse) SetTextInputEnabled(textInputEnabled *bool) {
	w.TextInputEnabled = textInputEnabled
}

func (w *WidgetConfigResponse) SetConversationModeToggleEnabled(conversationModeToggleEnabled *bool) {
	w.ConversationModeToggleEnabled = conversationModeToggleEnabled
}

func (w *WidgetConfigResponse) SetDefaultExpanded(defaultExpanded *bool) {
	w.DefaultExpanded = defaultExpanded
}

func (w *WidgetConfigResponse) SetAlwaysExpanded(alwaysExpanded *bool) {
	w.AlwaysExpanded = alwaysExpanded
}

func (w *WidgetConfigResponse) SetDismissible(dismissible *bool) {
	w.Dismissible = dismissible
}

func (w *WidgetConfigResponse) SetShowAgentStatus(showAgentStatus *bool) {
	w.ShowAgentStatus = showAgentStatus
}

func (w *WidgetConfigResponse) SetShowConversationId(showConversationId *bool) {
	w.ShowConversationId = showConversationId
}

func (w *WidgetConfigResponse) SetStripAudioTags(stripAudioTags *bool) {
	w.StripAudioTags = stripAudioTags
}

func (w *WidgetConfigResponse) SetSyntaxHighlightTheme(syntaxHighlightTheme *WidgetConfigResponseSyntaxHighlightTheme) {
	w.SyntaxHighlightTheme = syntaxHighlightTheme
}

func (w *WidgetConfigResponse) SetTextContents(textContents *WidgetTextContents) {
	w.TextContents = textContents
}

func (w *WidgetConfigResponse) SetStyles(styles *WidgetStyles) {
	w.Styles = styles
}

func (w *WidgetConfigResponse) SetShowResizeButton(showResizeButton *bool) {
	w.ShowResizeButton = showResizeButton
}

func (w *WidgetConfigResponse) SetEndConfirmationEnabled(endConfirmationEnabled *bool) {
	w.EndConfirmationEnabled = endConfirmationEnabled
}

func (w *WidgetConfigResponse) SetMobileFullscreen(mobileFullscreen *bool) {
	w.MobileFullscreen = mobileFullscreen
}

func (w *WidgetConfigResponse) SetTranscriptSpeechSyncEnabled(transcriptSpeechSyncEnabled *bool) {
	w.TranscriptSpeechSyncEnabled = transcriptSpeechSyncEnabled
}

func (w *WidgetConfigResponse) SetLanguage(language string) {
	w.Language = language
}

func (w *WidgetConfigResponse) SetSupportedLanguageOverrides(supportedLanguageOverrides []string) {
	w.SupportedLanguageOverrides = supportedLanguageOverrides
}

func (w *WidgetConfigResponse) SetLanguagePresets(languagePresets map[string]*WidgetLanguagePresetResponse) {
	w.LanguagePresets = languagePresets
}

func (w *WidgetConfigResponse) SetTextOnly(textOnly *bool) {
	w.TextOnly = textOnly
}

func (w *WidgetConfigResponse) SetSupportsTextOnly(supportsTextOnly *bool) {
	w.SupportsTextOnly = supportsTextOnly
}

func (w *WidgetConfigResponse) SetFirstMessage(firstMessage *string) {
	w.FirstMessage = firstMessage
}

func (w *WidgetConfigResponse) SetUseRtc(useRtc *bool) {
	w.UseRtc = useRtc
}

func (w *WidgetConfigResponse) SetFileInputConfig(fileInputConfig *FileInputConfig) {
	w.FileInputConfig = fileInputConfig
}

func (w *WidgetConfigResponse) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetConfigResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetConfigResponse to string"
	}
	return string(jsonData)
}

func (w *WidgetConfigResponse) UnmarshalJSON(data []byte) error {
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

func (w *WidgetConfigResponse) MarshalJSON() ([]byte, error) {
	type embed WidgetConfigResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
