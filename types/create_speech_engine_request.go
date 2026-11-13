package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CreateSpeechEngineRequest struct {
	// Name of the speech engine
	Name         *string                       `json:"name,omitempty" xml:"name,omitempty"`
	SpeechEngine *SpeechEngineConfig           `json:"speech_engine" xml:"speech_engine" required:"true"`
	Asr          *AsrConversationalConfig      `json:"asr,omitempty" xml:"asr,omitempty"`
	Tts          *TtsConversationalConfigInput `json:"tts,omitempty" xml:"tts,omitempty"`
	Turn         *BaseTurnConfig               `json:"turn,omitempty" xml:"turn,omitempty"`
	Vad          *VadConfig                    `json:"vad,omitempty" xml:"vad,omitempty"`
	Conversation *ConversationConfigInput      `json:"conversation,omitempty" xml:"conversation,omitempty"`
	Privacy      *PrivacyConfigInput           `json:"privacy,omitempty" xml:"privacy,omitempty"`
	CallLimits   *AgentCallLimits              `json:"call_limits,omitempty" xml:"call_limits,omitempty"`
	Auth         *SpeechEngineAuthConfig       `json:"auth,omitempty" xml:"auth,omitempty"`
	// Language for the speech engine
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// Time in seconds to wait for the upstream speech engine endpoint to respond before the attempt is abandoned and retried. Must be between 2 and 15 seconds.
	CascadeTimeoutSeconds *float64 `json:"cascade_timeout_seconds,omitempty" xml:"cascade_timeout_seconds,omitempty" min:"2" max:"15"`
	// Tags for categorization
	Tags      []string                                            `json:"tags,omitempty" xml:"tags,omitempty"`
	Overrides *SpeechEngineConversationInitiationClientDataConfig `json:"overrides,omitempty" xml:"overrides,omitempty"`
}

func (c *CreateSpeechEngineRequest) GetName() (value *string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *CreateSpeechEngineRequest) GetSpeechEngine() (value *SpeechEngineConfig) {
	if c == nil {
		return
	}
	return c.SpeechEngine
}

func (c *CreateSpeechEngineRequest) GetAsr() (value *AsrConversationalConfig) {
	if c == nil {
		return
	}
	return c.Asr
}

func (c *CreateSpeechEngineRequest) GetTts() (value *TtsConversationalConfigInput) {
	if c == nil {
		return
	}
	return c.Tts
}

func (c *CreateSpeechEngineRequest) GetTurn() (value *BaseTurnConfig) {
	if c == nil {
		return
	}
	return c.Turn
}

func (c *CreateSpeechEngineRequest) GetVad() (value *VadConfig) {
	if c == nil {
		return
	}
	return c.Vad
}

func (c *CreateSpeechEngineRequest) GetConversation() (value *ConversationConfigInput) {
	if c == nil {
		return
	}
	return c.Conversation
}

func (c *CreateSpeechEngineRequest) GetPrivacy() (value *PrivacyConfigInput) {
	if c == nil {
		return
	}
	return c.Privacy
}

func (c *CreateSpeechEngineRequest) GetCallLimits() (value *AgentCallLimits) {
	if c == nil {
		return
	}
	return c.CallLimits
}

func (c *CreateSpeechEngineRequest) GetAuth() (value *SpeechEngineAuthConfig) {
	if c == nil {
		return
	}
	return c.Auth
}

func (c *CreateSpeechEngineRequest) GetLanguage() (value *string) {
	if c == nil {
		return
	}
	return c.Language
}

func (c *CreateSpeechEngineRequest) GetCascadeTimeoutSeconds() (value *float64) {
	if c == nil {
		return
	}
	return c.CascadeTimeoutSeconds
}

func (c *CreateSpeechEngineRequest) GetTags() (value []string) {
	if c == nil {
		return
	}
	return c.Tags
}

func (c *CreateSpeechEngineRequest) GetOverrides() (value *SpeechEngineConversationInitiationClientDataConfig) {
	if c == nil {
		return
	}
	return c.Overrides
}

func (c *CreateSpeechEngineRequest) SetName(name *string) {
	c.Name = name
}

func (c *CreateSpeechEngineRequest) SetSpeechEngine(speechEngine *SpeechEngineConfig) {
	c.SpeechEngine = speechEngine
}

func (c *CreateSpeechEngineRequest) SetAsr(asr *AsrConversationalConfig) {
	c.Asr = asr
}

func (c *CreateSpeechEngineRequest) SetTts(tts *TtsConversationalConfigInput) {
	c.Tts = tts
}

func (c *CreateSpeechEngineRequest) SetTurn(turn *BaseTurnConfig) {
	c.Turn = turn
}

func (c *CreateSpeechEngineRequest) SetVad(vad *VadConfig) {
	c.Vad = vad
}

func (c *CreateSpeechEngineRequest) SetConversation(conversation *ConversationConfigInput) {
	c.Conversation = conversation
}

func (c *CreateSpeechEngineRequest) SetPrivacy(privacy *PrivacyConfigInput) {
	c.Privacy = privacy
}

func (c *CreateSpeechEngineRequest) SetCallLimits(callLimits *AgentCallLimits) {
	c.CallLimits = callLimits
}

func (c *CreateSpeechEngineRequest) SetAuth(auth *SpeechEngineAuthConfig) {
	c.Auth = auth
}

func (c *CreateSpeechEngineRequest) SetLanguage(language *string) {
	c.Language = language
}

func (c *CreateSpeechEngineRequest) SetCascadeTimeoutSeconds(cascadeTimeoutSeconds *float64) {
	c.CascadeTimeoutSeconds = cascadeTimeoutSeconds
}

func (c *CreateSpeechEngineRequest) SetTags(tags []string) {
	c.Tags = tags
}

func (c *CreateSpeechEngineRequest) SetOverrides(overrides *SpeechEngineConversationInitiationClientDataConfig) {
	c.Overrides = overrides
}

func (c *CreateSpeechEngineRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (c CreateSpeechEngineRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateSpeechEngineRequest to string"
	}
	return string(jsonData)
}

func (c *CreateSpeechEngineRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateSpeechEngineRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateSpeechEngineRequest(tmp)
	return nil
}

func (c *CreateSpeechEngineRequest) MarshalJSON() ([]byte, error) {
	type embed CreateSpeechEngineRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
