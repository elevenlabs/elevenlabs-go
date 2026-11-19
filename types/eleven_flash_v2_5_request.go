package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the ElevenLabs Flash v2.5 TTS model.
type ElevenFlashV25Request struct {
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
	VoiceSettings   *ElevenFlashV25VoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ElevenFlashV25Request) GetWebhook() (value *WebhookTarget) {
	if e == nil {
		return
	}
	return e.Webhook
}

func (e *ElevenFlashV25Request) GetText() (value string) {
	if e == nil {
		return
	}
	return e.Text
}

func (e *ElevenFlashV25Request) GetVoice() (value string) {
	if e == nil {
		return
	}
	return e.Voice
}

func (e *ElevenFlashV25Request) GetOutputFormat() (value *ElevenFlashV25RequestOutputFormat) {
	if e == nil {
		return
	}
	return e.OutputFormat
}

func (e *ElevenFlashV25Request) GetPronunciationDictionaryLocators() (value []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) {
	if e == nil {
		return
	}
	return e.PronunciationDictionaryLocators
}

func (e *ElevenFlashV25Request) GetLanguageCode() (value *string) {
	if e == nil {
		return
	}
	return e.LanguageCode
}

func (e *ElevenFlashV25Request) GetVoiceSettings() (value *ElevenFlashV25VoiceSettings) {
	if e == nil {
		return
	}
	return e.VoiceSettings
}

func (e *ElevenFlashV25Request) SetWebhook(webhook *WebhookTarget) {
	e.Webhook = webhook
}

func (e *ElevenFlashV25Request) SetText(text string) {
	e.Text = text
}

func (e *ElevenFlashV25Request) SetVoice(voice string) {
	e.Voice = voice
}

func (e *ElevenFlashV25Request) SetOutputFormat(outputFormat *ElevenFlashV25RequestOutputFormat) {
	e.OutputFormat = outputFormat
}

func (e *ElevenFlashV25Request) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) {
	e.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (e *ElevenFlashV25Request) SetLanguageCode(languageCode *string) {
	e.LanguageCode = languageCode
}

func (e *ElevenFlashV25Request) SetVoiceSettings(voiceSettings *ElevenFlashV25VoiceSettings) {
	e.VoiceSettings = voiceSettings
}

func (e *ElevenFlashV25Request) ModelId() string {
	return e.modelId
}

func (e *ElevenFlashV25Request) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ElevenFlashV25Request) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ElevenFlashV25Request to string"
	}
	return string(jsonData)
}

func (e *ElevenFlashV25Request) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "eleven_flash_v2_5" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "eleven_flash_v2_5", *literals.ModelId)
	}
	e.modelId = "eleven_flash_v2_5"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "model_id")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ElevenFlashV25Request) MarshalJSON() ([]byte, error) {
	type embed ElevenFlashV25Request
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*e),
		ModelId: "eleven_flash_v2_5",
	}
	return json.Marshal(marshaler)
}
