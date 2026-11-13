package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechEngineResponse struct {
	// The speech engine resource ID
	SpeechEngineId string `json:"speech_engine_id" xml:"speech_engine_id" required:"true"`
	// Human-readable name for the speech engine
	Name         string                         `json:"name" xml:"name" required:"true"`
	SpeechEngine *SpeechEngineConfig            `json:"speech_engine" xml:"speech_engine" required:"true"`
	Asr          *AsrConversationalConfig       `json:"asr" xml:"asr" required:"true"`
	Tts          *TtsConversationalConfigOutput `json:"tts" xml:"tts" required:"true"`
	Turn         *BaseTurnConfig                `json:"turn" xml:"turn" required:"true"`
	Vad          *VadConfig                     `json:"vad" xml:"vad" required:"true"`
	Conversation *ConversationConfigOutput      `json:"conversation" xml:"conversation" required:"true"`
	Privacy      *PrivacyConfigOutput           `json:"privacy" xml:"privacy" required:"true"`
	CallLimits   *AgentCallLimits               `json:"call_limits" xml:"call_limits" required:"true"`
	Auth         *SpeechEngineAuthConfig        `json:"auth" xml:"auth" required:"true"`
	// ISO language code used by the speech engine (e.g. 'en')
	Language string `json:"language" xml:"language" required:"true"`
	// Time in seconds to wait for the upstream speech engine endpoint to respond before the attempt is abandoned and retried. Must be between 2 and 15 seconds.
	CascadeTimeoutSeconds float64 `json:"cascade_timeout_seconds" xml:"cascade_timeout_seconds" required:"true"`
	// Arbitrary tags for categorization and filtering
	Tags      []string                                            `json:"tags" xml:"tags" required:"true"`
	Overrides *SpeechEngineConversationInitiationClientDataConfig `json:"overrides" xml:"overrides" required:"true"`
	Metadata  *AgentMetadataDbModel                               `json:"metadata" xml:"metadata" required:"true"`
	// The access information of the speech engine for the user
	AccessInfo      *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechEngineResponse) GetSpeechEngineId() (value string) {
	if s == nil {
		return
	}
	return s.SpeechEngineId
}

func (s *SpeechEngineResponse) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SpeechEngineResponse) GetSpeechEngine() (value *SpeechEngineConfig) {
	if s == nil {
		return
	}
	return s.SpeechEngine
}

func (s *SpeechEngineResponse) GetAsr() (value *AsrConversationalConfig) {
	if s == nil {
		return
	}
	return s.Asr
}

func (s *SpeechEngineResponse) GetTts() (value *TtsConversationalConfigOutput) {
	if s == nil {
		return
	}
	return s.Tts
}

func (s *SpeechEngineResponse) GetTurn() (value *BaseTurnConfig) {
	if s == nil {
		return
	}
	return s.Turn
}

func (s *SpeechEngineResponse) GetVad() (value *VadConfig) {
	if s == nil {
		return
	}
	return s.Vad
}

func (s *SpeechEngineResponse) GetConversation() (value *ConversationConfigOutput) {
	if s == nil {
		return
	}
	return s.Conversation
}

func (s *SpeechEngineResponse) GetPrivacy() (value *PrivacyConfigOutput) {
	if s == nil {
		return
	}
	return s.Privacy
}

func (s *SpeechEngineResponse) GetCallLimits() (value *AgentCallLimits) {
	if s == nil {
		return
	}
	return s.CallLimits
}

func (s *SpeechEngineResponse) GetAuth() (value *SpeechEngineAuthConfig) {
	if s == nil {
		return
	}
	return s.Auth
}

func (s *SpeechEngineResponse) GetLanguage() (value string) {
	if s == nil {
		return
	}
	return s.Language
}

func (s *SpeechEngineResponse) GetCascadeTimeoutSeconds() (value float64) {
	if s == nil {
		return
	}
	return s.CascadeTimeoutSeconds
}

func (s *SpeechEngineResponse) GetTags() (value []string) {
	if s == nil {
		return
	}
	return s.Tags
}

func (s *SpeechEngineResponse) GetOverrides() (value *SpeechEngineConversationInitiationClientDataConfig) {
	if s == nil {
		return
	}
	return s.Overrides
}

func (s *SpeechEngineResponse) GetMetadata() (value *AgentMetadataDbModel) {
	if s == nil {
		return
	}
	return s.Metadata
}

func (s *SpeechEngineResponse) GetAccessInfo() (value *ResourceAccessInfo) {
	if s == nil {
		return
	}
	return s.AccessInfo
}

func (s *SpeechEngineResponse) SetSpeechEngineId(speechEngineId string) {
	s.SpeechEngineId = speechEngineId
}

func (s *SpeechEngineResponse) SetName(name string) {
	s.Name = name
}

func (s *SpeechEngineResponse) SetSpeechEngine(speechEngine *SpeechEngineConfig) {
	s.SpeechEngine = speechEngine
}

func (s *SpeechEngineResponse) SetAsr(asr *AsrConversationalConfig) {
	s.Asr = asr
}

func (s *SpeechEngineResponse) SetTts(tts *TtsConversationalConfigOutput) {
	s.Tts = tts
}

func (s *SpeechEngineResponse) SetTurn(turn *BaseTurnConfig) {
	s.Turn = turn
}

func (s *SpeechEngineResponse) SetVad(vad *VadConfig) {
	s.Vad = vad
}

func (s *SpeechEngineResponse) SetConversation(conversation *ConversationConfigOutput) {
	s.Conversation = conversation
}

func (s *SpeechEngineResponse) SetPrivacy(privacy *PrivacyConfigOutput) {
	s.Privacy = privacy
}

func (s *SpeechEngineResponse) SetCallLimits(callLimits *AgentCallLimits) {
	s.CallLimits = callLimits
}

func (s *SpeechEngineResponse) SetAuth(auth *SpeechEngineAuthConfig) {
	s.Auth = auth
}

func (s *SpeechEngineResponse) SetLanguage(language string) {
	s.Language = language
}

func (s *SpeechEngineResponse) SetCascadeTimeoutSeconds(cascadeTimeoutSeconds float64) {
	s.CascadeTimeoutSeconds = cascadeTimeoutSeconds
}

func (s *SpeechEngineResponse) SetTags(tags []string) {
	s.Tags = tags
}

func (s *SpeechEngineResponse) SetOverrides(overrides *SpeechEngineConversationInitiationClientDataConfig) {
	s.Overrides = overrides
}

func (s *SpeechEngineResponse) SetMetadata(metadata *AgentMetadataDbModel) {
	s.Metadata = metadata
}

func (s *SpeechEngineResponse) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	s.AccessInfo = accessInfo
}

func (s *SpeechEngineResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechEngineResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechEngineResponse to string"
	}
	return string(jsonData)
}

func (s *SpeechEngineResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechEngineResponse) MarshalJSON() ([]byte, error) {
	type embed SpeechEngineResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
