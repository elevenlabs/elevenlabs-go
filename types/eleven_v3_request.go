package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the Eleven v3 TTS model.
type ElevenV3Request struct {
	// Include to send the generation's result to the workspace's configured flows webhooks once it completes or fails. The webhook payload matches the terminal response of the corresponding GET endpoint.
	Webhook *WebhookTarget `json:"webhook,omitempty" xml:"webhook,omitempty"`
	// The text to synthesize into speech.
	Text string `json:"text" xml:"text" required:"true"`
	// The ID of the voice to speak with.
	Voice string `json:"voice" xml:"voice" required:"true"`
	// The audio encoding of the output, as `codec_sampleRateHz_bitrateKbps`. `mp3_44100_192` requires the Creator tier or above.
	OutputFormat *ElevenFlashV25RequestOutputFormat `json:"output_format,omitempty" xml:"output_format,omitempty"`
	// Pronunciation dictionaries to apply to the text, in order of precedence. Up to 3.
	PronunciationDictionaryLocators []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty" maxItems:"3"`
	// The model to use for the generation.
	modelId string
	// ISO 639-1 language code to enforce on the output. Omit to detect the language from the text.
	LanguageCode *string `json:"language_code,omitempty" xml:"language_code,omitempty"`
	// Overrides for the voice's saved settings, applied to this generation only.
	VoiceSettings   *ElevenV3VoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ElevenV3Request) GetWebhook() (value *WebhookTarget) {
	if e == nil {
		return
	}
	return e.Webhook
}

func (e *ElevenV3Request) GetText() (value string) {
	if e == nil {
		return
	}
	return e.Text
}

func (e *ElevenV3Request) GetVoice() (value string) {
	if e == nil {
		return
	}
	return e.Voice
}

func (e *ElevenV3Request) GetOutputFormat() (value *ElevenFlashV25RequestOutputFormat) {
	if e == nil {
		return
	}
	return e.OutputFormat
}

func (e *ElevenV3Request) GetPronunciationDictionaryLocators() (value []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) {
	if e == nil {
		return
	}
	return e.PronunciationDictionaryLocators
}

func (e *ElevenV3Request) GetLanguageCode() (value *string) {
	if e == nil {
		return
	}
	return e.LanguageCode
}

func (e *ElevenV3Request) GetVoiceSettings() (value *ElevenV3VoiceSettings) {
	if e == nil {
		return
	}
	return e.VoiceSettings
}

func (e *ElevenV3Request) SetWebhook(webhook *WebhookTarget) {
	e.Webhook = webhook
}

func (e *ElevenV3Request) SetText(text string) {
	e.Text = text
}

func (e *ElevenV3Request) SetVoice(voice string) {
	e.Voice = voice
}

func (e *ElevenV3Request) SetOutputFormat(outputFormat *ElevenFlashV25RequestOutputFormat) {
	e.OutputFormat = outputFormat
}

func (e *ElevenV3Request) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) {
	e.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (e *ElevenV3Request) SetLanguageCode(languageCode *string) {
	e.LanguageCode = languageCode
}

func (e *ElevenV3Request) SetVoiceSettings(voiceSettings *ElevenV3VoiceSettings) {
	e.VoiceSettings = voiceSettings
}

func (e *ElevenV3Request) ModelId() string {
	return e.modelId
}

func (e *ElevenV3Request) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ElevenV3Request) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ElevenV3Request to string"
	}
	return string(jsonData)
}

func (e *ElevenV3Request) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "eleven_v3" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "eleven_v3", *literals.ModelId)
	}
	e.modelId = "eleven_v3"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "model_id")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ElevenV3Request) MarshalJSON() ([]byte, error) {
	type embed ElevenV3Request
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*e),
		ModelId: "eleven_v3",
	}
	return json.Marshal(marshaler)
}
