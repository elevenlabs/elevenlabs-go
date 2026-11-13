package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WidgetConfig struct {
	Variant    *EmbedVariant     `json:"variant,omitempty" xml:"variant,omitempty"`
	Placement  *WidgetPlacement  `json:"placement,omitempty" xml:"placement,omitempty"`
	Expandable *WidgetExpandable `json:"expandable,omitempty" xml:"expandable,omitempty"`
	// The avatar of the widget
	Avatar       *WidgetConfigAvatar `json:"avatar,omitempty" xml:"avatar,omitempty"`
	FeedbackMode *WidgetFeedbackMode `json:"feedback_mode,omitempty" xml:"feedback_mode,omitempty"`
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
	SyntaxHighlightTheme *WidgetConfigSyntaxHighlightTheme `json:"syntax_highlight_theme,omitempty" xml:"syntax_highlight_theme,omitempty"`
	TextContents         *WidgetTextContents               `json:"text_contents,omitempty" xml:"text_contents,omitempty"`
	Styles               *WidgetStyles                     `json:"styles,omitempty" xml:"styles,omitempty"`
	// Whether to show the resize button
	ShowResizeButton *bool `json:"show_resize_button,omitempty" xml:"show_resize_button,omitempty"`
	// Whether to ask the user to confirm before ending a chat or call
	EndConfirmationEnabled *bool `json:"end_confirmation_enabled,omitempty" xml:"end_confirmation_enabled,omitempty"`
	// Whether the widget opens fullscreen on mobile
	MobileFullscreen *bool `json:"mobile_fullscreen,omitempty" xml:"mobile_fullscreen,omitempty"`
	// Whether voice call transcripts are revealed as the agent speaks
	TranscriptSpeechSyncEnabled *bool `json:"transcript_speech_sync_enabled,omitempty" xml:"transcript_speech_sync_enabled,omitempty"`
	// Whether to show the language selector
	LanguageSelector *bool `json:"language_selector,omitempty" xml:"language_selector,omitempty"`
	// Whether the widget can switch to text only mode
	SupportsTextOnly *bool `json:"supports_text_only,omitempty" xml:"supports_text_only,omitempty"`
	// The custom avatar path
	CustomAvatarPath *string `json:"custom_avatar_path,omitempty" xml:"custom_avatar_path,omitempty"`
	// Language presets for the widget
	LanguagePresets map[string]*WidgetLanguagePreset `json:"language_presets,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WidgetConfig) GetVariant() (value *EmbedVariant) {
	if w == nil {
		return
	}
	return w.Variant
}

func (w *WidgetConfig) GetPlacement() (value *WidgetPlacement) {
	if w == nil {
		return
	}
	return w.Placement
}

func (w *WidgetConfig) GetExpandable() (value *WidgetExpandable) {
	if w == nil {
		return
	}
	return w.Expandable
}

func (w *WidgetConfig) GetAvatar() (value *WidgetConfigAvatar) {
	if w == nil {
		return
	}
	return w.Avatar
}

func (w *WidgetConfig) GetFeedbackMode() (value *WidgetFeedbackMode) {
	if w == nil {
		return
	}
	return w.FeedbackMode
}

func (w *WidgetConfig) GetEndFeedback() (value *WidgetEndFeedbackConfig) {
	if w == nil {
		return
	}
	return w.EndFeedback
}

func (w *WidgetConfig) GetBgColor() (value *string) {
	if w == nil {
		return
	}
	return w.BgColor
}

func (w *WidgetConfig) GetTextColor() (value *string) {
	if w == nil {
		return
	}
	return w.TextColor
}

func (w *WidgetConfig) GetBtnColor() (value *string) {
	if w == nil {
		return
	}
	return w.BtnColor
}

func (w *WidgetConfig) GetBtnTextColor() (value *string) {
	if w == nil {
		return
	}
	return w.BtnTextColor
}

func (w *WidgetConfig) GetBorderColor() (value *string) {
	if w == nil {
		return
	}
	return w.BorderColor
}

func (w *WidgetConfig) GetFocusColor() (value *string) {
	if w == nil {
		return
	}
	return w.FocusColor
}

func (w *WidgetConfig) GetBorderRadius() (value *int) {
	if w == nil {
		return
	}
	return w.BorderRadius
}

func (w *WidgetConfig) GetBtnRadius() (value *int) {
	if w == nil {
		return
	}
	return w.BtnRadius
}

func (w *WidgetConfig) GetActionText() (value *string) {
	if w == nil {
		return
	}
	return w.ActionText
}

func (w *WidgetConfig) GetStartCallText() (value *string) {
	if w == nil {
		return
	}
	return w.StartCallText
}

func (w *WidgetConfig) GetEndCallText() (value *string) {
	if w == nil {
		return
	}
	return w.EndCallText
}

func (w *WidgetConfig) GetExpandText() (value *string) {
	if w == nil {
		return
	}
	return w.ExpandText
}

func (w *WidgetConfig) GetListeningText() (value *string) {
	if w == nil {
		return
	}
	return w.ListeningText
}

func (w *WidgetConfig) GetSpeakingText() (value *string) {
	if w == nil {
		return
	}
	return w.SpeakingText
}

func (w *WidgetConfig) GetShareablePageText() (value *string) {
	if w == nil {
		return
	}
	return w.ShareablePageText
}

func (w *WidgetConfig) GetShareablePageShowTerms() (value *bool) {
	if w == nil {
		return
	}
	return w.ShareablePageShowTerms
}

func (w *WidgetConfig) GetTermsText() (value *string) {
	if w == nil {
		return
	}
	return w.TermsText
}

func (w *WidgetConfig) GetTermsHtml() (value *string) {
	if w == nil {
		return
	}
	return w.TermsHtml
}

func (w *WidgetConfig) GetTermsKey() (value *string) {
	if w == nil {
		return
	}
	return w.TermsKey
}

func (w *WidgetConfig) GetShowAvatarWhenCollapsed() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowAvatarWhenCollapsed
}

func (w *WidgetConfig) GetDisableBanner() (value *bool) {
	if w == nil {
		return
	}
	return w.DisableBanner
}

func (w *WidgetConfig) GetOverrideLink() (value *string) {
	if w == nil {
		return
	}
	return w.OverrideLink
}

func (w *WidgetConfig) GetMarkdownLinkAllowedHosts() (value []*AllowlistItem) {
	if w == nil {
		return
	}
	return w.MarkdownLinkAllowedHosts
}

func (w *WidgetConfig) GetMarkdownLinkIncludeWww() (value *bool) {
	if w == nil {
		return
	}
	return w.MarkdownLinkIncludeWww
}

func (w *WidgetConfig) GetMarkdownLinkAllowHttp() (value *bool) {
	if w == nil {
		return
	}
	return w.MarkdownLinkAllowHttp
}

func (w *WidgetConfig) GetMicMutingEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.MicMutingEnabled
}

func (w *WidgetConfig) GetTranscriptEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.TranscriptEnabled
}

func (w *WidgetConfig) GetTextInputEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.TextInputEnabled
}

func (w *WidgetConfig) GetConversationModeToggleEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.ConversationModeToggleEnabled
}

func (w *WidgetConfig) GetDefaultExpanded() (value *bool) {
	if w == nil {
		return
	}
	return w.DefaultExpanded
}

func (w *WidgetConfig) GetAlwaysExpanded() (value *bool) {
	if w == nil {
		return
	}
	return w.AlwaysExpanded
}

func (w *WidgetConfig) GetDismissible() (value *bool) {
	if w == nil {
		return
	}
	return w.Dismissible
}

func (w *WidgetConfig) GetShowAgentStatus() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowAgentStatus
}

func (w *WidgetConfig) GetShowConversationId() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowConversationId
}

func (w *WidgetConfig) GetStripAudioTags() (value *bool) {
	if w == nil {
		return
	}
	return w.StripAudioTags
}

func (w *WidgetConfig) GetSyntaxHighlightTheme() (value *WidgetConfigSyntaxHighlightTheme) {
	if w == nil {
		return
	}
	return w.SyntaxHighlightTheme
}

func (w *WidgetConfig) GetTextContents() (value *WidgetTextContents) {
	if w == nil {
		return
	}
	return w.TextContents
}

func (w *WidgetConfig) GetStyles() (value *WidgetStyles) {
	if w == nil {
		return
	}
	return w.Styles
}

func (w *WidgetConfig) GetShowResizeButton() (value *bool) {
	if w == nil {
		return
	}
	return w.ShowResizeButton
}

func (w *WidgetConfig) GetEndConfirmationEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.EndConfirmationEnabled
}

func (w *WidgetConfig) GetMobileFullscreen() (value *bool) {
	if w == nil {
		return
	}
	return w.MobileFullscreen
}

func (w *WidgetConfig) GetTranscriptSpeechSyncEnabled() (value *bool) {
	if w == nil {
		return
	}
	return w.TranscriptSpeechSyncEnabled
}

func (w *WidgetConfig) GetLanguageSelector() (value *bool) {
	if w == nil {
		return
	}
	return w.LanguageSelector
}

func (w *WidgetConfig) GetSupportsTextOnly() (value *bool) {
	if w == nil {
		return
	}
	return w.SupportsTextOnly
}

func (w *WidgetConfig) GetCustomAvatarPath() (value *string) {
	if w == nil {
		return
	}
	return w.CustomAvatarPath
}

func (w *WidgetConfig) GetLanguagePresets() (value map[string]*WidgetLanguagePreset) {
	if w == nil {
		return
	}
	return w.LanguagePresets
}

func (w *WidgetConfig) SetVariant(variant *EmbedVariant) {
	w.Variant = variant
}

func (w *WidgetConfig) SetPlacement(placement *WidgetPlacement) {
	w.Placement = placement
}

func (w *WidgetConfig) SetExpandable(expandable *WidgetExpandable) {
	w.Expandable = expandable
}

func (w *WidgetConfig) SetAvatar(avatar *WidgetConfigAvatar) {
	w.Avatar = avatar
}

func (w *WidgetConfig) SetFeedbackMode(feedbackMode *WidgetFeedbackMode) {
	w.FeedbackMode = feedbackMode
}

func (w *WidgetConfig) SetEndFeedback(endFeedback *WidgetEndFeedbackConfig) {
	w.EndFeedback = endFeedback
}

func (w *WidgetConfig) SetBgColor(bgColor *string) {
	w.BgColor = bgColor
}

func (w *WidgetConfig) SetTextColor(textColor *string) {
	w.TextColor = textColor
}

func (w *WidgetConfig) SetBtnColor(btnColor *string) {
	w.BtnColor = btnColor
}

func (w *WidgetConfig) SetBtnTextColor(btnTextColor *string) {
	w.BtnTextColor = btnTextColor
}

func (w *WidgetConfig) SetBorderColor(borderColor *string) {
	w.BorderColor = borderColor
}

func (w *WidgetConfig) SetFocusColor(focusColor *string) {
	w.FocusColor = focusColor
}

func (w *WidgetConfig) SetBorderRadius(borderRadius *int) {
	w.BorderRadius = borderRadius
}

func (w *WidgetConfig) SetBtnRadius(btnRadius *int) {
	w.BtnRadius = btnRadius
}

func (w *WidgetConfig) SetActionText(actionText *string) {
	w.ActionText = actionText
}

func (w *WidgetConfig) SetStartCallText(startCallText *string) {
	w.StartCallText = startCallText
}

func (w *WidgetConfig) SetEndCallText(endCallText *string) {
	w.EndCallText = endCallText
}

func (w *WidgetConfig) SetExpandText(expandText *string) {
	w.ExpandText = expandText
}

func (w *WidgetConfig) SetListeningText(listeningText *string) {
	w.ListeningText = listeningText
}

func (w *WidgetConfig) SetSpeakingText(speakingText *string) {
	w.SpeakingText = speakingText
}

func (w *WidgetConfig) SetShareablePageText(shareablePageText *string) {
	w.ShareablePageText = shareablePageText
}

func (w *WidgetConfig) SetShareablePageShowTerms(shareablePageShowTerms *bool) {
	w.ShareablePageShowTerms = shareablePageShowTerms
}

func (w *WidgetConfig) SetTermsText(termsText *string) {
	w.TermsText = termsText
}

func (w *WidgetConfig) SetTermsHtml(termsHtml *string) {
	w.TermsHtml = termsHtml
}

func (w *WidgetConfig) SetTermsKey(termsKey *string) {
	w.TermsKey = termsKey
}

func (w *WidgetConfig) SetShowAvatarWhenCollapsed(showAvatarWhenCollapsed *bool) {
	w.ShowAvatarWhenCollapsed = showAvatarWhenCollapsed
}

func (w *WidgetConfig) SetDisableBanner(disableBanner *bool) {
	w.DisableBanner = disableBanner
}

func (w *WidgetConfig) SetOverrideLink(overrideLink *string) {
	w.OverrideLink = overrideLink
}

func (w *WidgetConfig) SetMarkdownLinkAllowedHosts(markdownLinkAllowedHosts []*AllowlistItem) {
	w.MarkdownLinkAllowedHosts = markdownLinkAllowedHosts
}

func (w *WidgetConfig) SetMarkdownLinkIncludeWww(markdownLinkIncludeWww *bool) {
	w.MarkdownLinkIncludeWww = markdownLinkIncludeWww
}

func (w *WidgetConfig) SetMarkdownLinkAllowHttp(markdownLinkAllowHttp *bool) {
	w.MarkdownLinkAllowHttp = markdownLinkAllowHttp
}

func (w *WidgetConfig) SetMicMutingEnabled(micMutingEnabled *bool) {
	w.MicMutingEnabled = micMutingEnabled
}

func (w *WidgetConfig) SetTranscriptEnabled(transcriptEnabled *bool) {
	w.TranscriptEnabled = transcriptEnabled
}

func (w *WidgetConfig) SetTextInputEnabled(textInputEnabled *bool) {
	w.TextInputEnabled = textInputEnabled
}

func (w *WidgetConfig) SetConversationModeToggleEnabled(conversationModeToggleEnabled *bool) {
	w.ConversationModeToggleEnabled = conversationModeToggleEnabled
}

func (w *WidgetConfig) SetDefaultExpanded(defaultExpanded *bool) {
	w.DefaultExpanded = defaultExpanded
}

func (w *WidgetConfig) SetAlwaysExpanded(alwaysExpanded *bool) {
	w.AlwaysExpanded = alwaysExpanded
}

func (w *WidgetConfig) SetDismissible(dismissible *bool) {
	w.Dismissible = dismissible
}

func (w *WidgetConfig) SetShowAgentStatus(showAgentStatus *bool) {
	w.ShowAgentStatus = showAgentStatus
}

func (w *WidgetConfig) SetShowConversationId(showConversationId *bool) {
	w.ShowConversationId = showConversationId
}

func (w *WidgetConfig) SetStripAudioTags(stripAudioTags *bool) {
	w.StripAudioTags = stripAudioTags
}

func (w *WidgetConfig) SetSyntaxHighlightTheme(syntaxHighlightTheme *WidgetConfigSyntaxHighlightTheme) {
	w.SyntaxHighlightTheme = syntaxHighlightTheme
}

func (w *WidgetConfig) SetTextContents(textContents *WidgetTextContents) {
	w.TextContents = textContents
}

func (w *WidgetConfig) SetStyles(styles *WidgetStyles) {
	w.Styles = styles
}

func (w *WidgetConfig) SetShowResizeButton(showResizeButton *bool) {
	w.ShowResizeButton = showResizeButton
}

func (w *WidgetConfig) SetEndConfirmationEnabled(endConfirmationEnabled *bool) {
	w.EndConfirmationEnabled = endConfirmationEnabled
}

func (w *WidgetConfig) SetMobileFullscreen(mobileFullscreen *bool) {
	w.MobileFullscreen = mobileFullscreen
}

func (w *WidgetConfig) SetTranscriptSpeechSyncEnabled(transcriptSpeechSyncEnabled *bool) {
	w.TranscriptSpeechSyncEnabled = transcriptSpeechSyncEnabled
}

func (w *WidgetConfig) SetLanguageSelector(languageSelector *bool) {
	w.LanguageSelector = languageSelector
}

func (w *WidgetConfig) SetSupportsTextOnly(supportsTextOnly *bool) {
	w.SupportsTextOnly = supportsTextOnly
}

func (w *WidgetConfig) SetCustomAvatarPath(customAvatarPath *string) {
	w.CustomAvatarPath = customAvatarPath
}

func (w *WidgetConfig) SetLanguagePresets(languagePresets map[string]*WidgetLanguagePreset) {
	w.LanguagePresets = languagePresets
}

func (w *WidgetConfig) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WidgetConfig) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WidgetConfig to string"
	}
	return string(jsonData)
}

func (w *WidgetConfig) UnmarshalJSON(data []byte) error {
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

func (w *WidgetConfig) MarshalJSON() ([]byte, error) {
	type embed WidgetConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
