package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request body for the ElevenLabs Multilingual v2 TTS model.
type ElevenMultilingualV2Request struct {
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
	// Overrides for the voice's saved settings, applied to this generation only.
	VoiceSettings   *TtsVoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	extraProperties map[string]interface{}
}

func (e *ElevenMultilingualV2Request) GetWebhook() (value *WebhookTarget) {
	if e == nil {
		return
	}
	return e.Webhook
}

func (e *ElevenMultilingualV2Request) GetText() (value string) {
	if e == nil {
		return
	}
	return e.Text
}

func (e *ElevenMultilingualV2Request) GetVoice() (value string) {
	if e == nil {
		return
	}
	return e.Voice
}

func (e *ElevenMultilingualV2Request) GetOutputFormat() (value *ElevenFlashV25RequestOutputFormat) {
	if e == nil {
		return
	}
	return e.OutputFormat
}

func (e *ElevenMultilingualV2Request) GetPronunciationDictionaryLocators() (value []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) {
	if e == nil {
		return
	}
	return e.PronunciationDictionaryLocators
}

func (e *ElevenMultilingualV2Request) GetVoiceSettings() (value *TtsVoiceSettings) {
	if e == nil {
		return
	}
	return e.VoiceSettings
}

func (e *ElevenMultilingualV2Request) SetWebhook(webhook *WebhookTarget) {
	e.Webhook = webhook
}

func (e *ElevenMultilingualV2Request) SetText(text string) {
	e.Text = text
}

func (e *ElevenMultilingualV2Request) SetVoice(voice string) {
	e.Voice = voice
}

func (e *ElevenMultilingualV2Request) SetOutputFormat(outputFormat *ElevenFlashV25RequestOutputFormat) {
	e.OutputFormat = outputFormat
}

func (e *ElevenMultilingualV2Request) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*ElevenFlashV25RequestPronunciationDictionaryLocatorsItem) {
	e.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (e *ElevenMultilingualV2Request) SetVoiceSettings(voiceSettings *TtsVoiceSettings) {
	e.VoiceSettings = voiceSettings
}

func (e *ElevenMultilingualV2Request) ModelId() string {
	return e.modelId
}

func (e *ElevenMultilingualV2Request) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e ElevenMultilingualV2Request) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ElevenMultilingualV2Request to string"
	}
	return string(jsonData)
}

func (e *ElevenMultilingualV2Request) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	var literals struct {
		ModelId *string `json:"model_id"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.ModelId != nil && *literals.ModelId != "eleven_multilingual_v2" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "eleven_multilingual_v2", *literals.ModelId)
	}
	e.modelId = "eleven_multilingual_v2"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "model_id")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *ElevenMultilingualV2Request) MarshalJSON() ([]byte, error) {
	type embed ElevenMultilingualV2Request
	var marshaler = struct {
		embed
		ModelId string `json:"model_id"`
	}{
		embed:   embed(*e),
		ModelId: "eleven_multilingual_v2",
	}
	return json.Marshal(marshaler)
}
