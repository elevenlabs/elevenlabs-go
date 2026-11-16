package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranscriptionOrderItemRequest struct {
	// The type of order item.
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// The IDs of the uploaded media files to transcribe.
	MediaIds []MediaId `json:"media_ids" xml:"media_ids" required:"true" maxItems:"50"`
	// The language code of the source media (e.g. 'en', 'es').
	SourceLanguage string `json:"source_language" xml:"source_language" required:"true"`
	// Whether to transcribe every word exactly, including filler words.
	Verbatim *bool `json:"verbatim,omitempty" xml:"verbatim,omitempty"`
	// Optional free-text instructions for the transcription team.
	Instructions    *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TranscriptionOrderItemRequest) GetKind() (value *string) {
	if t == nil {
		return
	}
	return t.Kind
}

func (t *TranscriptionOrderItemRequest) GetMediaIds() (value []MediaId) {
	if t == nil {
		return
	}
	return t.MediaIds
}

func (t *TranscriptionOrderItemRequest) GetSourceLanguage() (value string) {
	if t == nil {
		return
	}
	return t.SourceLanguage
}

func (t *TranscriptionOrderItemRequest) GetVerbatim() (value *bool) {
	if t == nil {
		return
	}
	return t.Verbatim
}

func (t *TranscriptionOrderItemRequest) GetInstructions() (value *string) {
	if t == nil {
		return
	}
	return t.Instructions
}

func (t *TranscriptionOrderItemRequest) SetKind(kind *string) {
	t.Kind = kind
}

func (t *TranscriptionOrderItemRequest) SetMediaIds(mediaIds []MediaId) {
	t.MediaIds = mediaIds
}

func (t *TranscriptionOrderItemRequest) SetSourceLanguage(sourceLanguage string) {
	t.SourceLanguage = sourceLanguage
}

func (t *TranscriptionOrderItemRequest) SetVerbatim(verbatim *bool) {
	t.Verbatim = verbatim
}

func (t *TranscriptionOrderItemRequest) SetInstructions(instructions *string) {
	t.Instructions = instructions
}

func (t *TranscriptionOrderItemRequest) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranscriptionOrderItemRequest) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranscriptionOrderItemRequest to string"
	}
	return string(jsonData)
}

func (t *TranscriptionOrderItemRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranscriptionOrderItemRequest) MarshalJSON() ([]byte, error) {
	type embed TranscriptionOrderItemRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
