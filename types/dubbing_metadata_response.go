package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type DubbingMetadataResponse struct {
	// The ID of the dubbing project.
	DubbingId string `json:"dubbing_id" xml:"dubbing_id" required:"true"`
	// The name of the dubbing project.
	Name string `json:"name" xml:"name" required:"true"`
	// The state this dub is in.
	Status string `json:"status" xml:"status" required:"true"`
	// Once dubbing has completed, the ISO-639-1 code of the original media's source language.
	SourceLanguage *string `json:"source_language" xml:"source_language" required:"true"`
	// The ISO-639-1 code of the languages this media has been dubbed into.
	TargetLanguages []string `json:"target_languages" xml:"target_languages" required:"true"`
	// Whether this dubbing project is editable in Dubbing Studio.
	Editable *bool `json:"editable,omitempty" xml:"editable,omitempty"`
	// Timestamp this dub was created.
	CreatedAt time.Time `json:"created_at" xml:"created_at" required:"true"`
	// Metadata, such as the length in seconds and content type, of the dubbed content.
	MediaMetadata *DubbingMediaMetadata `json:"media_metadata,omitempty" xml:"media_metadata,omitempty"`
	// Error message indicate, if this dub has failed, what happened.
	Error           *string `json:"error,omitempty" xml:"error,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingMetadataResponse) GetDubbingId() (value string) {
	if d == nil {
		return
	}
	return d.DubbingId
}

func (d *DubbingMetadataResponse) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DubbingMetadataResponse) GetStatus() (value string) {
	if d == nil {
		return
	}
	return d.Status
}

func (d *DubbingMetadataResponse) GetSourceLanguage() (value *string) {
	if d == nil {
		return
	}
	return d.SourceLanguage
}

func (d *DubbingMetadataResponse) GetTargetLanguages() (value []string) {
	if d == nil {
		return
	}
	return d.TargetLanguages
}

func (d *DubbingMetadataResponse) GetEditable() (value *bool) {
	if d == nil {
		return
	}
	return d.Editable
}

func (d *DubbingMetadataResponse) GetCreatedAt() (value time.Time) {
	if d == nil {
		return
	}
	return d.CreatedAt
}

func (d *DubbingMetadataResponse) GetMediaMetadata() (value *DubbingMediaMetadata) {
	if d == nil {
		return
	}
	return d.MediaMetadata
}

func (d *DubbingMetadataResponse) GetError() (value *string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingMetadataResponse) SetDubbingId(dubbingId string) {
	d.DubbingId = dubbingId
}

func (d *DubbingMetadataResponse) SetName(name string) {
	d.Name = name
}

func (d *DubbingMetadataResponse) SetStatus(status string) {
	d.Status = status
}

func (d *DubbingMetadataResponse) SetSourceLanguage(sourceLanguage *string) {
	d.SourceLanguage = sourceLanguage
}

func (d *DubbingMetadataResponse) SetTargetLanguages(targetLanguages []string) {
	d.TargetLanguages = targetLanguages
}

func (d *DubbingMetadataResponse) SetEditable(editable *bool) {
	d.Editable = editable
}

func (d *DubbingMetadataResponse) SetCreatedAt(createdAt time.Time) {
	d.CreatedAt = createdAt
}

func (d *DubbingMetadataResponse) SetMediaMetadata(mediaMetadata *DubbingMediaMetadata) {
	d.MediaMetadata = mediaMetadata
}

func (d *DubbingMetadataResponse) SetError(error *string) {
	d.Error = error
}

func (d *DubbingMetadataResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingMetadataResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingMetadataResponse to string"
	}
	return string(jsonData)
}

func (d *DubbingMetadataResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingMetadataResponse) MarshalJSON() ([]byte, error) {
	type embed DubbingMetadataResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
