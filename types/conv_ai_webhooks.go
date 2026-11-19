package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConvAiWebhooks struct {
	PostCallWebhookId *string `json:"post_call_webhook_id,omitempty" xml:"post_call_webhook_id,omitempty"`
	// List of event types to send via webhook. Options: transcript, audio, call_initiation_failure, answering_machine_detection, unredacted_transcript, unredacted_audio.
	Events           []WebhookEventType       `json:"events,omitempty" xml:"events,omitempty"`
	TranscriptFormat *WebhookTranscriptFormat `json:"transcript_format,omitempty" xml:"transcript_format,omitempty"`
	// When true, JSON post-call transcription webhooks omit the turn-by-turn transcript. Analysis, metadata, and other conversation fields are still sent. Ignored for OpenTelemetry transcript format.
	ExcludeTranscript *bool `json:"exclude_transcript,omitempty" xml:"exclude_transcript,omitempty"`
	// DEPRECATED: Use 'events' field instead. Whether to send audio data with post-call webhooks for ConvAI conversations
	SendAudio       *bool `json:"send_audio,omitempty" xml:"send_audio,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConvAiWebhooks) GetPostCallWebhookId() (value *string) {
	if c == nil {
		return
	}
	return c.PostCallWebhookId
}

func (c *ConvAiWebhooks) GetEvents() (value []WebhookEventType) {
	if c == nil {
		return
	}
	return c.Events
}

func (c *ConvAiWebhooks) GetTranscriptFormat() (value *WebhookTranscriptFormat) {
	if c == nil {
		return
	}
	return c.TranscriptFormat
}

func (c *ConvAiWebhooks) GetExcludeTranscript() (value *bool) {
	if c == nil {
		return
	}
	return c.ExcludeTranscript
}

func (c *ConvAiWebhooks) GetSendAudio() (value *bool) {
	if c == nil {
		return
	}
	return c.SendAudio
}

func (c *ConvAiWebhooks) SetPostCallWebhookId(postCallWebhookId *string) {
	c.PostCallWebhookId = postCallWebhookId
}

func (c *ConvAiWebhooks) SetEvents(events []WebhookEventType) {
	c.Events = events
}

func (c *ConvAiWebhooks) SetTranscriptFormat(transcriptFormat *WebhookTranscriptFormat) {
	c.TranscriptFormat = transcriptFormat
}

func (c *ConvAiWebhooks) SetExcludeTranscript(excludeTranscript *bool) {
	c.ExcludeTranscript = excludeTranscript
}

func (c *ConvAiWebhooks) SetSendAudio(sendAudio *bool) {
	c.SendAudio = sendAudio
}

func (c *ConvAiWebhooks) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConvAiWebhooks) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConvAiWebhooks to string"
	}
	return string(jsonData)
}

func (c *ConvAiWebhooks) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConvAiWebhooks) MarshalJSON() ([]byte, error) {
	type embed ConvAiWebhooks
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
