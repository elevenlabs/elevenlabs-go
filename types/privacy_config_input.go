package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type PrivacyConfigInput struct {
	// Whether to record the conversation
	RecordVoice *bool `json:"record_voice,omitempty" xml:"record_voice,omitempty"`
	// The number of days to retain the conversation. -1 indicates there is no retention limit
	RetentionDays *int `json:"retention_days,omitempty" xml:"retention_days,omitempty"`
	// Whether to delete the transcript and PII
	DeleteTranscriptAndPii *bool `json:"delete_transcript_and_pii,omitempty" xml:"delete_transcript_and_pii,omitempty"`
	// Whether to delete the audio
	DeleteAudio *bool `json:"delete_audio,omitempty" xml:"delete_audio,omitempty"`
	// Whether to apply the privacy settings to existing conversations
	ApplyToExistingConversations *bool `json:"apply_to_existing_conversations,omitempty" xml:"apply_to_existing_conversations,omitempty"`
	// Whether to enable zero retention mode - no PII data is stored
	ZeroRetentionMode            *bool                               `json:"zero_retention_mode,omitempty" xml:"zero_retention_mode,omitempty"`
	ConversationHistoryRedaction *ConversationHistoryRedactionConfig `json:"conversation_history_redaction,omitempty" xml:"conversation_history_redaction,omitempty"`
	extraProperties              map[string]interface{}
}

func (p *PrivacyConfigInput) GetRecordVoice() (value *bool) {
	if p == nil {
		return
	}
	return p.RecordVoice
}

func (p *PrivacyConfigInput) GetRetentionDays() (value *int) {
	if p == nil {
		return
	}
	return p.RetentionDays
}

func (p *PrivacyConfigInput) GetDeleteTranscriptAndPii() (value *bool) {
	if p == nil {
		return
	}
	return p.DeleteTranscriptAndPii
}

func (p *PrivacyConfigInput) GetDeleteAudio() (value *bool) {
	if p == nil {
		return
	}
	return p.DeleteAudio
}

func (p *PrivacyConfigInput) GetApplyToExistingConversations() (value *bool) {
	if p == nil {
		return
	}
	return p.ApplyToExistingConversations
}

func (p *PrivacyConfigInput) GetZeroRetentionMode() (value *bool) {
	if p == nil {
		return
	}
	return p.ZeroRetentionMode
}

func (p *PrivacyConfigInput) GetConversationHistoryRedaction() (value *ConversationHistoryRedactionConfig) {
	if p == nil {
		return
	}
	return p.ConversationHistoryRedaction
}

func (p *PrivacyConfigInput) SetRecordVoice(recordVoice *bool) {
	p.RecordVoice = recordVoice
}

func (p *PrivacyConfigInput) SetRetentionDays(retentionDays *int) {
	p.RetentionDays = retentionDays
}

func (p *PrivacyConfigInput) SetDeleteTranscriptAndPii(deleteTranscriptAndPii *bool) {
	p.DeleteTranscriptAndPii = deleteTranscriptAndPii
}

func (p *PrivacyConfigInput) SetDeleteAudio(deleteAudio *bool) {
	p.DeleteAudio = deleteAudio
}

func (p *PrivacyConfigInput) SetApplyToExistingConversations(applyToExistingConversations *bool) {
	p.ApplyToExistingConversations = applyToExistingConversations
}

func (p *PrivacyConfigInput) SetZeroRetentionMode(zeroRetentionMode *bool) {
	p.ZeroRetentionMode = zeroRetentionMode
}

func (p *PrivacyConfigInput) SetConversationHistoryRedaction(conversationHistoryRedaction *ConversationHistoryRedactionConfig) {
	p.ConversationHistoryRedaction = conversationHistoryRedaction
}

func (p *PrivacyConfigInput) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PrivacyConfigInput) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PrivacyConfigInput to string"
	}
	return string(jsonData)
}

func (p *PrivacyConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler PrivacyConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PrivacyConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PrivacyConfigInput) MarshalJSON() ([]byte, error) {
	type embed PrivacyConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
