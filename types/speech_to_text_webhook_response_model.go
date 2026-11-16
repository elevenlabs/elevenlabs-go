package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechToTextWebhookResponseModel struct {
	// The message of the webhook response.
	Message string `json:"message" xml:"message" required:"true"`
	// The request ID of the webhook response.
	RequestId string `json:"request_id" xml:"request_id" required:"true"`
	// The transcription ID of the webhook response.
	TranscriptionId *string `json:"transcription_id,omitempty" xml:"transcription_id,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechToTextWebhookResponseModel) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SpeechToTextWebhookResponseModel) GetRequestId() (value string) {
	if s == nil {
		return
	}
	return s.RequestId
}

func (s *SpeechToTextWebhookResponseModel) GetTranscriptionId() (value *string) {
	if s == nil {
		return
	}
	return s.TranscriptionId
}

func (s *SpeechToTextWebhookResponseModel) SetMessage(message string) {
	s.Message = message
}

func (s *SpeechToTextWebhookResponseModel) SetRequestId(requestId string) {
	s.RequestId = requestId
}

func (s *SpeechToTextWebhookResponseModel) SetTranscriptionId(transcriptionId *string) {
	s.TranscriptionId = transcriptionId
}

func (s *SpeechToTextWebhookResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechToTextWebhookResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechToTextWebhookResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeechToTextWebhookResponseModel) UnmarshalJSON(data []byte) error {
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

func (s *SpeechToTextWebhookResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeechToTextWebhookResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
