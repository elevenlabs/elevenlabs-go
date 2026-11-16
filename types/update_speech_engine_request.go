package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateSpeechEngineRequest struct {
	// The speech engine ID (accepts seng_ or agent_ prefix)
	SpeechEngineId string                        `json:"-" url:"-"`
	Name           *string                       `json:"name,omitempty" xml:"name,omitempty"`
	SpeechEngine   *SpeechEngineConfig           `json:"speech_engine,omitempty" xml:"speech_engine,omitempty"`
	Asr            *AsrConversationalConfig      `json:"asr,omitempty" xml:"asr,omitempty"`
	Tts            *TtsConversationalConfigInput `json:"tts,omitempty" xml:"tts,omitempty"`
	Turn           *BaseTurnConfig               `json:"turn,omitempty" xml:"turn,omitempty"`
	Vad            *VadConfig                    `json:"vad,omitempty" xml:"vad,omitempty"`
	Conversation   *ConversationConfigInput      `json:"conversation,omitempty" xml:"conversation,omitempty"`
	Privacy        *PrivacyConfigInput           `json:"privacy,omitempty" xml:"privacy,omitempty"`
	CallLimits     *AgentCallLimits              `json:"call_limits,omitempty" xml:"call_limits,omitempty"`
	// Who may start a conversation with this speech engine
	Auth     *SpeechEngineAuthConfig `json:"auth,omitempty" xml:"auth,omitempty"`
	Language *string                 `json:"language,omitempty" xml:"language,omitempty"`
	// Time in seconds to wait for the upstream speech engine endpoint to respond before the attempt is abandoned and retried. Must be between 2 and 15 seconds.
	CascadeTimeoutSeconds *float64                                            `json:"cascade_timeout_seconds,omitempty" xml:"cascade_timeout_seconds,omitempty" min:"2" max:"15"`
	Tags                  []string                                            `json:"tags,omitempty" xml:"tags,omitempty"`
	Overrides             *SpeechEngineConversationInitiationClientDataConfig `json:"overrides,omitempty" xml:"overrides,omitempty"`
}

func (u *UpdateSpeechEngineRequest) GetSpeechEngineId() (value string) {
	if u == nil {
		return
	}
	return u.SpeechEngineId
}

func (u *UpdateSpeechEngineRequest) GetName() (value *string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateSpeechEngineRequest) GetSpeechEngine() (value *SpeechEngineConfig) {
	if u == nil {
		return
	}
	return u.SpeechEngine
}

func (u *UpdateSpeechEngineRequest) GetAsr() (value *AsrConversationalConfig) {
	if u == nil {
		return
	}
	return u.Asr
}

func (u *UpdateSpeechEngineRequest) GetTts() (value *TtsConversationalConfigInput) {
	if u == nil {
		return
	}
	return u.Tts
}

func (u *UpdateSpeechEngineRequest) GetTurn() (value *BaseTurnConfig) {
	if u == nil {
		return
	}
	return u.Turn
}

func (u *UpdateSpeechEngineRequest) GetVad() (value *VadConfig) {
	if u == nil {
		return
	}
	return u.Vad
}

func (u *UpdateSpeechEngineRequest) GetConversation() (value *ConversationConfigInput) {
	if u == nil {
		return
	}
	return u.Conversation
}

func (u *UpdateSpeechEngineRequest) GetPrivacy() (value *PrivacyConfigInput) {
	if u == nil {
		return
	}
	return u.Privacy
}

func (u *UpdateSpeechEngineRequest) GetCallLimits() (value *AgentCallLimits) {
	if u == nil {
		return
	}
	return u.CallLimits
}

func (u *UpdateSpeechEngineRequest) GetAuth() (value *SpeechEngineAuthConfig) {
	if u == nil {
		return
	}
	return u.Auth
}

func (u *UpdateSpeechEngineRequest) GetLanguage() (value *string) {
	if u == nil {
		return
	}
	return u.Language
}

func (u *UpdateSpeechEngineRequest) GetCascadeTimeoutSeconds() (value *float64) {
	if u == nil {
		return
	}
	return u.CascadeTimeoutSeconds
}

func (u *UpdateSpeechEngineRequest) GetTags() (value []string) {
	if u == nil {
		return
	}
	return u.Tags
}

func (u *UpdateSpeechEngineRequest) GetOverrides() (value *SpeechEngineConversationInitiationClientDataConfig) {
	if u == nil {
		return
	}
	return u.Overrides
}

func (u *UpdateSpeechEngineRequest) SetSpeechEngineId(speechEngineId string) {
	u.SpeechEngineId = speechEngineId
}

func (u *UpdateSpeechEngineRequest) SetName(name *string) {
	u.Name = name
}

func (u *UpdateSpeechEngineRequest) SetSpeechEngine(speechEngine *SpeechEngineConfig) {
	u.SpeechEngine = speechEngine
}

func (u *UpdateSpeechEngineRequest) SetAsr(asr *AsrConversationalConfig) {
	u.Asr = asr
}

func (u *UpdateSpeechEngineRequest) SetTts(tts *TtsConversationalConfigInput) {
	u.Tts = tts
}

func (u *UpdateSpeechEngineRequest) SetTurn(turn *BaseTurnConfig) {
	u.Turn = turn
}

func (u *UpdateSpeechEngineRequest) SetVad(vad *VadConfig) {
	u.Vad = vad
}

func (u *UpdateSpeechEngineRequest) SetConversation(conversation *ConversationConfigInput) {
	u.Conversation = conversation
}

func (u *UpdateSpeechEngineRequest) SetPrivacy(privacy *PrivacyConfigInput) {
	u.Privacy = privacy
}

func (u *UpdateSpeechEngineRequest) SetCallLimits(callLimits *AgentCallLimits) {
	u.CallLimits = callLimits
}

func (u *UpdateSpeechEngineRequest) SetAuth(auth *SpeechEngineAuthConfig) {
	u.Auth = auth
}

func (u *UpdateSpeechEngineRequest) SetLanguage(language *string) {
	u.Language = language
}

func (u *UpdateSpeechEngineRequest) SetCascadeTimeoutSeconds(cascadeTimeoutSeconds *float64) {
	u.CascadeTimeoutSeconds = cascadeTimeoutSeconds
}

func (u *UpdateSpeechEngineRequest) SetTags(tags []string) {
	u.Tags = tags
}

func (u *UpdateSpeechEngineRequest) SetOverrides(overrides *SpeechEngineConversationInitiationClientDataConfig) {
	u.Overrides = overrides
}

func (u *UpdateSpeechEngineRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateSpeechEngineRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateSpeechEngineRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateSpeechEngineRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateSpeechEngineRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateSpeechEngineRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
