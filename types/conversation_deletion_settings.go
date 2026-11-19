package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationDeletionSettings struct {
	DeletionTimeUnixSecs            *int  `json:"deletion_time_unix_secs,omitempty" xml:"deletion_time_unix_secs,omitempty"`
	DeletedLogsAtTimeUnixSecs       *int  `json:"deleted_logs_at_time_unix_secs,omitempty" xml:"deleted_logs_at_time_unix_secs,omitempty"`
	DeletedAudioAtTimeUnixSecs      *int  `json:"deleted_audio_at_time_unix_secs,omitempty" xml:"deleted_audio_at_time_unix_secs,omitempty"`
	DeletedTranscriptAtTimeUnixSecs *int  `json:"deleted_transcript_at_time_unix_secs,omitempty" xml:"deleted_transcript_at_time_unix_secs,omitempty"`
	DeleteTranscriptAndPii          *bool `json:"delete_transcript_and_pii,omitempty" xml:"delete_transcript_and_pii,omitempty"`
	DeleteAudio                     *bool `json:"delete_audio,omitempty" xml:"delete_audio,omitempty"`
	extraProperties                 map[string]interface{}
}

func (c *ConversationDeletionSettings) GetDeletionTimeUnixSecs() (value *int) {
	if c == nil {
		return
	}
	return c.DeletionTimeUnixSecs
}

func (c *ConversationDeletionSettings) GetDeletedLogsAtTimeUnixSecs() (value *int) {
	if c == nil {
		return
	}
	return c.DeletedLogsAtTimeUnixSecs
}

func (c *ConversationDeletionSettings) GetDeletedAudioAtTimeUnixSecs() (value *int) {
	if c == nil {
		return
	}
	return c.DeletedAudioAtTimeUnixSecs
}

func (c *ConversationDeletionSettings) GetDeletedTranscriptAtTimeUnixSecs() (value *int) {
	if c == nil {
		return
	}
	return c.DeletedTranscriptAtTimeUnixSecs
}

func (c *ConversationDeletionSettings) GetDeleteTranscriptAndPii() (value *bool) {
	if c == nil {
		return
	}
	return c.DeleteTranscriptAndPii
}

func (c *ConversationDeletionSettings) GetDeleteAudio() (value *bool) {
	if c == nil {
		return
	}
	return c.DeleteAudio
}

func (c *ConversationDeletionSettings) SetDeletionTimeUnixSecs(deletionTimeUnixSecs *int) {
	c.DeletionTimeUnixSecs = deletionTimeUnixSecs
}

func (c *ConversationDeletionSettings) SetDeletedLogsAtTimeUnixSecs(deletedLogsAtTimeUnixSecs *int) {
	c.DeletedLogsAtTimeUnixSecs = deletedLogsAtTimeUnixSecs
}

func (c *ConversationDeletionSettings) SetDeletedAudioAtTimeUnixSecs(deletedAudioAtTimeUnixSecs *int) {
	c.DeletedAudioAtTimeUnixSecs = deletedAudioAtTimeUnixSecs
}

func (c *ConversationDeletionSettings) SetDeletedTranscriptAtTimeUnixSecs(deletedTranscriptAtTimeUnixSecs *int) {
	c.DeletedTranscriptAtTimeUnixSecs = deletedTranscriptAtTimeUnixSecs
}

func (c *ConversationDeletionSettings) SetDeleteTranscriptAndPii(deleteTranscriptAndPii *bool) {
	c.DeleteTranscriptAndPii = deleteTranscriptAndPii
}

func (c *ConversationDeletionSettings) SetDeleteAudio(deleteAudio *bool) {
	c.DeleteAudio = deleteAudio
}

func (c *ConversationDeletionSettings) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationDeletionSettings) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationDeletionSettings to string"
	}
	return string(jsonData)
}

func (c *ConversationDeletionSettings) UnmarshalJSON(data []byte) error {
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

func (c *ConversationDeletionSettings) MarshalJSON() ([]byte, error) {
	type embed ConversationDeletionSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
